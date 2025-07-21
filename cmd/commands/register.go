package commands

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/eodata/operator-cli/cmd/flags"
	allocationmanager "github.com/eodata/operator-cli/contracts/bindings/AllocationManager"
	regcoord "github.com/eodata/operator-cli/contracts/bindings/EORegistryCoordinator"
	iblsapkregistry "github.com/eodata/operator-cli/contracts/bindings/IBLSApkRegistry"
	"github.com/eodata/operator-cli/internal/eigen"
	"github.com/eodata/operator-cli/internal/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v3"
)

func NewRegisterCommand(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:        "register",
		Description: "Register the operator",
		Version:     version,
		HideVersion: true,
		Action:      runRegister,
		Flags: []cli.Flag{
			flags.ProfileFlag,
			flags.EthRPCFlag,
			flags.EOChainRPCFlag,
			flags.OperatorAddressFlag,
			flags.RegistryCoordinatorFlag,
			flags.PassphraseFlag,
			flags.EcdsaPassphraseFlag,
			flags.BlsPassphraseFlag,
			flags.KeyStorePathFlag,
			flags.SaltFlag,
			flags.ExpiryFlag,
			flags.OperatorSetIDsFlag,
		},
	}
}

type operatorSetData struct {
	RegistrationType int32
	Socket           string
	Params           iblsapkregistry.IBLSApkRegistryTypesPubkeyRegistrationParams
	OperatorAliases  []common.Address
}

func runRegister(ctx context.Context, c *cli.Command) error {
	logger := flags.Logger

	err := flags.SetProfile(ctx, c)
	if err != nil {
		return fmt.Errorf("error setting profile: %v", err)
	}

	var ecdsaPassphrase string
	var blsPassphrase string

	if c.IsSet(flags.PassphraseFlag.Name) {
		ecdsaPassphrase = c.String(flags.PassphraseFlag.Name)
		blsPassphrase = c.String(flags.PassphraseFlag.Name)
	}

	if c.IsSet(flags.EcdsaPassphraseFlag.Name) {
		if ecdsaPassphrase != "" {
			return fmt.Errorf("either common passphrase or ecdsa passphrases should be set")
		}
		ecdsaPassphrase = c.String(flags.EcdsaPassphraseFlag.Name)
	}

	if ecdsaPassphrase == "" {
		return fmt.Errorf("either common passphrase or ecdsa passphrases should be set")
	}

	ecdsaPair, err := keystore.GetECDSAPrivateKey(ecdsaPassphrase, c.String(flags.KeyStorePathFlag.Name), "")
	if err != nil {
		return fmt.Errorf("error getting ECDSA private key %v", err)
	}

	if c.IsSet(flags.BlsPassphraseFlag.Name) {
		if blsPassphrase != "" {
			return fmt.Errorf("either common passphrase or bls passphrases should be set")
		}
		blsPassphrase = c.String(flags.BlsPassphraseFlag.Name)
	}

	if blsPassphrase == "" {
		return fmt.Errorf("either common passphrase or bls passphrases should be set")
	}

	blsKeyPair, err := keystore.GetBLSPrivateKey(blsPassphrase, c.String(flags.KeyStorePathFlag.Name))
	if err != nil {
		return fmt.Errorf("error getting BLS private key %v", err)
	}

	if !common.IsHexAddress(c.String(flags.RegistryCoordinatorFlag.Name)) {
		return fmt.Errorf("registry-coordinator-address must be a valid Ethereum address")
	}
	registryCoordinatorAddress := common.HexToAddress(c.String(flags.RegistryCoordinatorFlag.Name))
	if registryCoordinatorAddress == (common.Address{}) {
		return fmt.Errorf("registry-coordinator-address must be a valid Ethereum address")
	}

	if !common.IsHexAddress(c.String(flags.OperatorAddressFlag.Name)) {
		return fmt.Errorf("operator-address must be a valid Ethereum address")
	}
	operatorAddress := common.HexToAddress(c.String(flags.OperatorAddressFlag.Name))
	if operatorAddress == (common.Address{}) {
		return fmt.Errorf("operator-address must be a valid Ethereum address")
	}

	// parse operator-set-ids comma separated int:string tuple. The first item is integer of operator set id.
	// The second item is the alias address for that operator set.
	operatorSetIDStrings := c.StringSlice(flags.OperatorSetIDsFlag.Name)
	operatorSetIds := []uint32{}
	operatorSetAliases := []common.Address{}

	for _, idString := range operatorSetIDStrings {
		parts := strings.Split(idString, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid operator-set-id format: %s", idString)
		}

		id, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid operator-set-id integer: %s", parts[0])
		}

		if !common.IsHexAddress(parts[1]) {
			return fmt.Errorf("invalid operator-set-id alias: %s", parts[1])
		}

		alias := common.HexToAddress(parts[1])

		operatorSetIds = append(operatorSetIds, uint32(id))
		operatorSetAliases = append(operatorSetAliases, alias)
	}

	ethClient, err := eigen.CreateEthClient(c.String(flags.EthRPCFlag.Name))
	if err != nil {
		return fmt.Errorf("error creating Eth client %v", err)
	}

	avsClient, err := eigen.BuildAVSClient(ethClient, registryCoordinatorAddress)
	if err != nil {
		return fmt.Errorf("error creating AVS client %v", err)
	}

	g1HashedMsgToSign, err := avsClient.RegistryCoordinator.PubkeyRegistrationMessageHash(
		&bind.CallOpts{},
		operatorAddress,
	)
	if err != nil {
		return fmt.Errorf("error getting PubkeyRegistrationMessageHash from registryCoordinator contract %v", err)
	}

	signedMsg := convertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(convertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := convertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := convertToBN254G2Point(blsKeyPair.GetPubKeyG2())

	operatorSetData := operatorSetData{
		RegistrationType: 1,
		Socket:           "",
		Params: iblsapkregistry.IBLSApkRegistryTypesPubkeyRegistrationParams{
			PubkeyRegistrationSignature: signedMsg,
			PubkeyG1:                    G1pubkeyBN254,
			PubkeyG2:                    G2pubkeyBN254,
		},
		OperatorAliases: operatorSetAliases,
	}

	const operatorSetDataABI = `[{
    "type": "tuple",
    "components": [
        {"type": "int32", "name": "registrationType"},
        {"type": "string", "name": "socket"},
        {"type": "tuple", "name": "params", "components": [
            {"type": "tuple", "name": "pubkeyRegistrationSignature", "components": [
                {"type": "uint256", "name": "x"},
                {"type": "uint256", "name": "y"}
            ]},
            {"type": "tuple", "name": "pubkeyG1", "components": [
                {"type": "uint256", "name": "x"},
                {"type": "uint256", "name": "y"}
            ]},
            {"type": "tuple", "name": "pubkeyG2", "components": [
                {"type": "uint256[2]", "name": "x"},
                {"type": "uint256[2]", "name": "y"}
            ]}
        ]},
        {"type": "address[]", "name": "operatorAliases"}
    ]
}]`

	parsedABI, err := abi.JSON(strings.NewReader(operatorSetDataABI))
	if err != nil {
		return fmt.Errorf("error parsing ABI: %v", err)
	}

	encodedData, err := parsedABI.Pack("",
		operatorSetData.RegistrationType,
		operatorSetData.Socket,
		[]interface{}{
			[]interface{}{
				operatorSetData.Params.PubkeyRegistrationSignature.X,
				operatorSetData.Params.PubkeyRegistrationSignature.Y,
			},
			[]interface{}{
				operatorSetData.Params.PubkeyG1.X,
				operatorSetData.Params.PubkeyG1.Y,
			},
			[]interface{}{
				operatorSetData.Params.PubkeyG2.X,
				operatorSetData.Params.PubkeyG2.Y,
			},
		},
		operatorSetData.OperatorAliases,
	)
	if err != nil {
		return fmt.Errorf("error encoding operator set data: %v", err)
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("error getting chainId (%v): %v", c.String(flags.EthRPCFlag.Name), err)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPair}, chainIDBigInt)
	if err != nil {
		return fmt.Errorf("error creating the signer object of %v on Ethereum mainnet/holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey), c.String(flags.EthRPCFlag.Name), err,
		)
	}

	wallet, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		return fmt.Errorf("error creating the wallet object of %v on Ethereum mainnet/holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey), c.String(flags.EthRPCFlag.Name), err,
		)
	}

	txMgr := txmgr.NewSimpleTxManager(
		wallet,
		ethClient,
		logger,
		signerAddr,
	)

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		return fmt.Errorf("error creating a noopSender parameters on Ethereum mainnet/Holesky (%v) %v",
			c.String(flags.EthRPCFlag.Name), err,
		)
	}

	registerParams := allocationmanager.IAllocationManagerTypesRegisterParams{
		Avs:            registryCoordinatorAddress,
		OperatorSetIds: operatorSetIds,
		Data:           encodedData,
	}

	tx, err := avsClient.AllocationManager.RegisterForOperatorSets(
		noSendTxOpts,
		operatorAddress,
		registerParams,
	)
	if err != nil {
		return fmt.Errorf(
			"error building the registerForOperatorSets transaction of operator %v on Ethereum mainnet/Holesky (%v) %v",
			operatorAddress.Hex(),
			c.String(flags.EthRPCFlag.Name),
			err,
		)
	}

	receipt, err := txMgr.Send(ctx, tx, true)
	if err != nil {
		return fmt.Errorf(
			"error sending the registerForOperatorSets transaction of operator %v on Ethereum mainnet/Holesky (%v) %v",
			operatorAddress.Hex(),
			c.String(flags.EthRPCFlag.Name),
			err,
		)
	}
	if receipt.Status != 1 {
		return fmt.Errorf(
			"registerForOperatorSets transaction %v of operator %v on Ethereum mainnet/Holeskey (%v) reverted",
			receipt.TxHash.Hex(),
			operatorAddress.Hex(),
			c.String(flags.EthRPCFlag.Name),
		)
	}

	logger.Info(
		"successfully registered to eoracle AVS",
		"operator", operatorAddress.Hex(),
		"tx hash", receipt.TxHash.Hex(),
		"sender", signerAddr.Hex(),
	)

	return nil
}

func convertBn254GethToGnark(input regcoord.BN254G1Point) *bn254.G1Affine {
	return eigensdkbls.NewG1Point(input.X, input.Y).G1Affine
}

func convertToBN254G2Point(input *eigensdkbls.G2Point) iblsapkregistry.BN254G2Point {
	output := iblsapkregistry.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
func convertToBN254G1Point(input *eigensdkbls.G1Point) iblsapkregistry.BN254G1Point {
	output := iblsapkregistry.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}
