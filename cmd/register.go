package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	regcoord "github.com/eodata/operator-cli/contracts/bindings/EORegistryCoordinator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func NewRegisterCommand() *cli.Command {
	return &cli.Command{
		Name:        "register",
		Description: "Register the operator",
		Before:      setProfile,
		Action:      runRegister,
		Flags: []cli.Flag{
			ProfileFlag,
			EthRPCFlag,
			EOChainRPCFlag,
			RegistryCoordinatorFlag,
			EOConfigAddressFlag,
			PassphraseFlag,
			KeyStorePathFlag,
			SaltFlag,
			ExpiryFlag,
			EcdsaPrivateKeyFlag,
			BlsPrivateKeyFlag,
			ChainValidatorG1PointSignatureFlag,
			ValidatorRoleFlag,
			QuorumNumberFlag,
		},
	}
}

func runRegister(c *cli.Context) error {
	if (!c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name)) && (!c.IsSet(EcdsaPrivateKeyFlag.Name) || !c.IsSet(BlsPrivateKeyFlag.Name)) {
		utils.Fatalf("either passphrase and keystore-path or ecdsa-private-key and bls-private-key are required")
	}

	ecdsaPair, blsKeyPair, err := getKeys(c)
	if err != nil {
		return err
	}

	ethClient, err := createEthClient(profile.EthRPCEndpoint)
	if err != nil {
		return err
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		utils.Fatalf("Error getting chainId (%v): %v", profile.EthRPCEndpoint, err)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPair}, chainIDBigInt)
	if err != nil {
		utils.Fatalf("Error creating the register transaction signer for operator %v on Ethereum mainnet/holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey), profile.EthRPCEndpoint, err,
		)
	}

	saltBytes, err := getSaltBytes(c)
	if err != nil {
		return err
	}

	expiry, err := getExpiry(c)
	if err != nil {
		return err
	}

	avsClient, err := buildAVSClient(ethClient)
	if err != nil {
		utils.Fatalf("Error creating AVS client %v", err)
	}

	g1HashedMsgToSign, err := avsClient.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, signerAddr)
	if err != nil {
		utils.Fatalf("Error getting PubkeyRegistrationMessageHash from registryCoordinator contract %v", err)
	}

	signedMsg := convertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(convertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := convertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := convertToBN254G2Point(blsKeyPair.GetPubKeyG2())

	pubkeyRegParams := regcoord.IBLSApkRegistryTypesPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	msgToSign, err := avsClient.elReader.CalculateOperatorAVSRegistrationDigestHash(
		context.Background(),
		signerAddr,
		avsClient.serviceManagerAddr,
		saltBytes,
		expiry,
	)
	if err != nil {
		utils.Fatalf(
			"Error generating message to sign by operator %v on Ethereum mainnet/Holeskey (%v) using CalculateOperatorAVSRegistrationDigestHash %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	operatorSignature, err := crypto.Sign(msgToSign[:], ecdsaPair)
	if err != nil {
		utils.Fatalf(
			"Error signing the message using CalculateOperatorAVSRegistrationDigestHash for operator %v on Ethereum mainnet/Holeskey (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      saltBytes,
		Expiry:    expiry,
	}

	txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		utils.Fatalf(
			"Error creating the register transaction sender for operator %v on Ethereum mainnet/Holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	address, err := txSender.SenderAddress(context.Background())
	if err != nil {
		return err
	}

	addr := address.String()
	logger.Info("sender address", "address", addr)

	txMgr := txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerAddr)
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		utils.Fatalf("error creating transaction object %v", err)
	}

	noSendTxOpts.GasLimit = 2_000_000

	tx, err := avsClient.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		[]byte{0},
		"0.0.0.0:0",
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		utils.Fatalf(
			"Error creating the register transaction for operator %v on Ethereum mainnet/Holeskey (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	ctx := context.Background()
	receipt, err := txMgr.Send(ctx, tx, true)
	if err != nil {
		utils.Fatalf(
			"register transaction for operator %v on Ethereum mainnet/Holeskey (%v) failed %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile,
			err,
		)
	}
	if receipt.Status != 1 {
		utils.Fatalf(
			"register transaction %v for operator %v on Ethereum mainnet/Holeskey (%v) reverted",
			receipt.TxHash.Hex(),
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
		)
	}

	logger.Info("successfully registered to eoracle AVS", "address", signerAddr, "tx hash", receipt.TxHash.Hex())

	return nil
}

func convertBn254GethToGnark(input regcoord.BN254G1Point) *bn254.G1Affine {
	return eigensdkbls.NewG1Point(input.X, input.Y).G1Affine
}

func convertToBN254G2Point(input *eigensdkbls.G2Point) regcoord.BN254G2Point {
	output := regcoord.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
func convertToBN254G1Point(input *eigensdkbls.G1Point) regcoord.BN254G1Point {
	output := regcoord.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func getExpiry(c *cli.Context) (*big.Int, error) {
	expiry, ok := big.NewInt(0).SetString(c.String(ExpiryFlag.Name), 10)
	if !ok {
		utils.Fatalf("Invalid expiry")
	}
	return expiry, nil
}

func getKeys(c *cli.Context) (*ecdsa.PrivateKey, *eigensdkbls.KeyPair, error) {
	var ecdsaPair *ecdsa.PrivateKey
	var blsKeyPair *eigensdkbls.KeyPair
	var err error

	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		ecdsaPair, err = crypto.HexToECDSA(c.String(EcdsaPrivateKeyFlag.Name))
		if err != nil {
			utils.Fatalf("Invalid EDCSA private key %v", err)
		}
		blsKeyPair, err = eigensdkbls.NewKeyPairFromString(c.String(BlsPrivateKeyFlag.Name))
		if err != nil {
			utils.Fatalf("Invalid BLS private key %v", err)
		}
	} else {
		ecdsaPair, err = eigensdkecdsa.ReadKey(
			filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaEncryptedWallet.json"),
			c.String(PassphraseFlag.Name),
		)
		if err != nil {
			utils.Fatalf("Failed to read ecdsaEncryptedWallet.json file %v", err)
		}
		blsKeyPair, err = eigensdkbls.ReadPrivateKeyFromFile(
			filepath.Join(c.String(KeyStorePathFlag.Name), "blsEncryptedWallet.json"),
			c.String(PassphraseFlag.Name),
		)
		if err != nil {
			utils.Fatalf("Failed to read blsEncryptedWallet.json file %v", err)
		}
	}
	return ecdsaPair, blsKeyPair, nil
}

func getSaltBytes(c *cli.Context) ([32]byte, error) {
	var saltBytes [32]byte
	inputBytes, err := hex.DecodeString(strings.TrimPrefix(c.String(SaltFlag.Name), "0x"))
	if err != nil {
		utils.Fatalf("Invalid salt %v", err)
	}
	copy(saltBytes[:], inputBytes)
	return saltBytes, nil
}

func getChainValidatorG1PointSignature(c *cli.Context) (regcoord.BN254G1Point, error) {
	chainValidatorG1PointSignature := convertToBN254G1Point(eigensdkbls.NewG1Point(big.NewInt(0), big.NewInt(0)))
	if c.String(ValidatorRoleFlag.Name) != "DATA_VALIDATOR" {
		if len(c.StringSlice(ChainValidatorG1PointSignatureFlag.Name)) != 2 {
			utils.Fatalf("chain-validator-g1-point-signature is required or has too many values")
		}
		x, ok := new(big.Int).SetString(c.StringSlice(ChainValidatorG1PointSignatureFlag.Name)[0][2:], 16)
		if !ok {
			utils.Fatalf("Invalid chain-validator-g1-point-signature (x)")
		}
		y, ok := new(big.Int).SetString(c.StringSlice(ChainValidatorG1PointSignatureFlag.Name)[1][2:], 16)
		if !ok {
			utils.Fatalf("Invalid chain-validator-g1-point-signature (y)")
		}
		chainValidatorG1PointSignature = convertToBN254G1Point(eigensdkbls.NewG1Point(x, y))
	}
	return chainValidatorG1PointSignature, nil
}
