package operatorcli

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v2"

	"github.com/consensys/gnark-crypto/ecc/bn254"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	smbase "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

	regcoord "github.com/Eoracle/core-go/contracts/bindings/EORegistryCoordinator"
	stakeregistry "github.com/Eoracle/core-go/contracts/bindings/EOStakeRegistry"

	eoconfig "github.com/Eoracle/core-go/contracts/bindings/EOConfig"
	"github.com/Eoracle/core-go/internal/flag"
)

type avsClient struct {
	registryCoordinatorAddr gethcommon.Address
	serviceManagerAddr      gethcommon.Address
	delegationManagerAddr   gethcommon.Address
	avsDirectoryAddr        gethcommon.Address
	registryCoordinator     *regcoord.ContractEORegistryCoordinator
	serviceManager          *smbase.ContractServiceManagerBase
	stakeRegistry           *stakeregistry.ContractEOStakeRegistry
	elReader                elcontracts.ELReader
}

func RunEncrypt(c *cli.Context) error {
	passphrase := c.String(flag.PassphraseFlag.Name)
	if passphrase == "" {
		return cli.Exit("passphrase is required", 1)
	}

	keyStorePath := c.String(flag.KeyStorePathFlag.Name)
	if keyStorePath == "" {
		return cli.Exit("keystore-path is required", 1)
	}

	// Encrypt the ecdsa private key and save it to a file
	ecdsaPair, err := crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
	}

	if err = eigensdkecdsa.WriteKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), ecdsaPair, passphrase); err != nil {
		return cli.Exit(fmt.Sprintf("Error writing the ecdsaEncryptedWallet.json file %v", err), 1)
	}
	fmt.Println("ecdsa address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "saved")

	// Encrypt the bls private key and save it to a file
	blsKeyPair, err := eigensdkbls.NewKeyPairFromString(c.String(flag.BlsPrivateKeyFlag.Name))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Invalid BLS private key %v", err), 1)
	}

	if err = blsKeyPair.SaveToFile(filepath.Join(keyStorePath, "blsEncryptedWallet.json"), passphrase); err != nil {
		return cli.Exit(fmt.Sprintf("Error writing the blsEncryptedWallet.json file %v", err), 1)
	}
	fmt.Println("bls address G1, G2 ", blsKeyPair.GetPubKeyG1().String(), ", ", blsKeyPair.GetPubKeyG2().String(), "saved")

	return nil
}

func RunDecrypt(c *cli.Context) error {
	passphrase := c.String(flag.PassphraseFlag.Name)
	if passphrase == "" {
		return cli.Exit("passphrase is required", 1)
	}

	keyStorePath := c.String(flag.KeyStorePathFlag.Name)
	if keyStorePath == "" {
		return cli.Exit("keystore-path is required", 1)
	}

	ecdsaPair, err := eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), passphrase)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error reading the ecdsaEncryptedWallet.json file %v", err), 1)
	}
	fmt.Println("ecdsa address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "private key", hex.EncodeToString(ecdsaPair.D.Bytes()))

	blsKeyPair, err := eigensdkbls.ReadPrivateKeyFromFile(filepath.Join(keyStorePath, "blsEncryptedWallet.json"), passphrase)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Error reading the blsEncryptedWallet.json file %v", err), 1)
	}
	fmt.Println("bls address G1, G2 ", blsKeyPair.GetPubKeyG1().String(), ", ", blsKeyPair.GetPubKeyG2().String(), "private key", blsKeyPair.PrivKey.String())

	ecdsaEOChainPair, err := eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaAliasedEncryptedWallet.json"), passphrase)
	if err != nil {
		if err == os.ErrNotExist {
			fmt.Println("EOChain alias was not set in the system")
			return nil
		}
		return cli.Exit(fmt.Sprintf("Error reading the ecdsaAliasedEncryptedWallet.json file %v", err), 1)
	}
	fmt.Println("EOChain ecdsa address ", crypto.PubkeyToAddress(ecdsaEOChainPair.PublicKey), "private key", hex.EncodeToString(ecdsaEOChainPair.D.Bytes()))

	return nil
}

func RunRegister(c *cli.Context) error {
	passphrase := c.String(flag.PassphraseFlag.Name)
	keyStorePath := c.String(flag.KeyStorePathFlag.Name)

	var ecdsaPair *ecdsa.PrivateKey
	var blsKeyPair *eigensdkbls.KeyPair
	var err error

	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating logger %v", err), 1)
	}

	if passphrase == "" || keyStorePath == "" {
		if c.String(flag.EcdsaPrivateKeyFlag.Name) == "" || c.String(flag.BlsPrivateKeyFlag.Name) == "" {
			return cli.Exit("either passphrase and keystore-path or ecdsa-private-key and bls-private-key are required", 1)
		}
		ecdsaPair, err = crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
		if err != nil {
			return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
		}
		blsKeyPair, err = eigensdkbls.NewKeyPairFromString(c.String(flag.BlsPrivateKeyFlag.Name))
		if err != nil {
			return cli.Exit(fmt.Sprintf("Invalid BLS private key %v", err), 1)
		}
	} else {
		ecdsaPair, err = eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), passphrase)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to read ecdsaEncryptedWallet.json file %v", err), 1)
		}
		blsKeyPair, err = eigensdkbls.ReadPrivateKeyFromFile(filepath.Join(keyStorePath, "blsEncryptedWallet.json"), passphrase)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to read blsEncryptedWallet.json file %v", err), 1)
		}
	}

	if c.String(flag.EthRPCFlag.Name) == "" {
		return cli.Exit("eth-rpc is required", 1)
	}

	ethClient, err := eth.NewClient(c.String(flag.EthRPCFlag.Name))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create Eth client %v %v", c.String(flag.EthRPCFlag.Name), err), 1)
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		return cli.Exit(fmt.Sprintf("cannot get chainId: %v", err), 1)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPair}, chainIDBigInt)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating the signer function for %v %v", crypto.PubkeyToAddress(ecdsaPair.PublicKey), err), 1)
	}

	var saltBytes [32]byte
	inputBytes, err := hex.DecodeString(strings.TrimPrefix(c.String(flag.SaltFlag.Name), "0x"))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Invalid salt %v", err), 1)
	}
	copy(saltBytes[:], inputBytes)

	expiry, ok := big.NewInt(0).SetString(c.String(flag.ExpiryFlag.Name), 10)
	if !ok {
		return cli.Exit("Invalid expiry", 1)
	}

	chainValidatorG1PointSignature := ConvertToBN254G1Point(eigensdkbls.NewG1Point(big.NewInt(0), big.NewInt(0)))

	if c.String(flag.ValidatorRoleFlag.Name) != "DATA_VALIDATOR" {
		if len(c.StringSlice(flag.ChainValidatorG1PointSignatureFlag.Name)) != 2 {
			return cli.Exit("chain-validator-g1-point-signature is required or has too many values", 1)
		}
		x, ok := new(big.Int).SetString(c.StringSlice(flag.ChainValidatorG1PointSignatureFlag.Name)[0][2:], 16)
		if !ok {
			return cli.Exit("Invalid chain-validator-g1-point-signature", 1)
		}
		y, ok := new(big.Int).SetString(c.StringSlice(flag.ChainValidatorG1PointSignatureFlag.Name)[1][2:], 16)
		if !ok {
			return cli.Exit("Invalid chain-validator-g1-point-signature", 1)
		}
		chainValidatorG1PointSignature = ConvertToBN254G1Point(eigensdkbls.NewG1Point(x, y))
	}

	if c.String(flag.RegistryCoordinatorFlag.Name) == "" {
		return cli.Exit("registry-coordinator is required", 1)
	}
	registryCoordinatorAddr := gethcommon.HexToAddress(c.String(flag.RegistryCoordinatorFlag.Name))
	avsClient, err := buildAVSClient(registryCoordinatorAddr, ethClient, logger)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create AVS client %v", err), 1)
	}

	g1HashedMsgToSign, err := avsClient.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, signerAddr)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to get PubkeyRegistrationMessageHash from registryCoordinator contract %v", err), 1)
	}
	signedMsg := ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)

	G1pubkeyBN254 := ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IEOBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		ChainValidatorSignature:     chainValidatorG1PointSignature,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// Params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := avsClient.elReader.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{},
		signerAddr,
		avsClient.serviceManagerAddr,
		saltBytes,
		expiry,
	)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to CalculateOperatorAVSRegistrationDigestHash %v", err), 1)
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], ecdsaPair)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to sign %v", err), 1)
	}

	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      saltBytes,
		Expiry:    expiry,
	}

	txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create transaction sender %v", err), 1)
	}
	txMgr := txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerV2, signerAddr)

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating transaction object %v", err), 1)
	}

	tx, err := avsClient.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		[]byte{0},
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create RegisterOperator transaction %v", err), 1)
	}

	ctx := context.Background()
	receipt, err := txMgr.Send(ctx, tx)
	if err != nil {
		return cli.Exit(fmt.Sprintf("register transaction failed %s", err), 1)
	}
	if receipt.Status != 1 {
		return cli.Exit(fmt.Sprintf("register transaction %v reverted", receipt.TxHash.Hex()), 1)
	}

	logger.Info("succesfully registered to eoracle AVS", "address", signerAddr, "tx hash", receipt.TxHash.Hex())

	return nil
}

func RunDeregister(c *cli.Context) error {
	passphrase := c.String(flag.PassphraseFlag.Name)
	keyStorePath := c.String(flag.KeyStorePathFlag.Name)

	var ecdsaPair *ecdsa.PrivateKey
	var err error

	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating logger %v", err), 1)
	}

	if passphrase == "" || keyStorePath == "" {
		if c.String(flag.EcdsaPrivateKeyFlag.Name) == "" {
			return cli.Exit("either passphrase and keystore-path or ecdsa-private-key are required", 1)
		}
		ecdsaPair, err = crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
		if err != nil {
			return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
		}
	} else {
		ecdsaPair, err = eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), passphrase)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to read ecdsaEncryptedWallet.json file %v", err), 1)
		}
	}

	if c.String(flag.EthRPCFlag.Name) == "" {
		return cli.Exit("eth-rpc is required", 1)
	}

	ethClient, err := eth.NewClient(c.String(flag.EthRPCFlag.Name))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create Eth client %v %v", c.String(flag.EthRPCFlag.Name), err), 1)
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		return cli.Exit(fmt.Sprintf("cannot get chainId: %v", err), 1)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPair}, chainIDBigInt)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating the signer function for %v %v", crypto.PubkeyToAddress(ecdsaPair.PublicKey), err), 1)
	}

	if c.String(flag.RegistryCoordinatorFlag.Name) == "" {
		return cli.Exit("registry-coordinator is required", 1)
	}
	registryCoordinatorAddr := gethcommon.HexToAddress(c.String(flag.RegistryCoordinatorFlag.Name))
	avsClient, err := buildAVSClient(registryCoordinatorAddr, ethClient, logger)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create AVS client %v", err), 1)
	}

	txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create transaction sender %v", err), 1)
	}
	txMgr := txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerV2, signerAddr)

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()

	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating transaction object %v", err), 1)
	}
	tx, err := avsClient.registryCoordinator.DeregisterOperator(
		noSendTxOpts,
		[]byte{0},
	)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create DeregisterOperator transaction %v", err), 1)
	}
	ctx := context.Background()
	receipt, err := txMgr.Send(ctx, tx)
	if err != nil {
		return cli.Exit(fmt.Sprintf("deregister transaction failed %s", err), 1)
	}
	if receipt.Status != 1 {
		return cli.Exit(fmt.Sprintf("deregister transaction %v reverted", receipt.TxHash.Hex()), 1)
	}
	logger.Info("DeregisterOperator", "gas", receipt.GasUsed, "txHash", receipt.TxHash.Hex())

	logger.Info("succesfully deregistered from eoracle AVS", "address", signerAddr, "tx hash", receipt.TxHash.Hex())

	return nil
}

func RunPrintStatus(c *cli.Context) error {
	passphrase := c.String(flag.PassphraseFlag.Name)
	keyStorePath := c.String(flag.KeyStorePathFlag.Name)

	var ecdsaPair *ecdsa.PrivateKey
	var err error

	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating logger %v", err), 1)
	}

	if passphrase == "" || keyStorePath == "" {
		if c.String(flag.EcdsaPrivateKeyFlag.Name) == "" {
			return cli.Exit("either passphrase and keystore-path or ecdsa-private-key are required", 1)
		}
		ecdsaPair, err = crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
		if err != nil {
			return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
		}
	} else {
		ecdsaPair, err = eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), passphrase)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to read ecdsaEncryptedWallet.json file %v", err), 1)
		}
	}
	operatorAddress := crypto.PubkeyToAddress(ecdsaPair.PublicKey)

	if c.String(flag.EthRPCFlag.Name) == "" {
		return cli.Exit("eth-rpc is required", 1)
	}

	ethClient, err := eth.NewClient(c.String(flag.EthRPCFlag.Name))
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create Eth client %v %v", c.String(flag.EthRPCFlag.Name), err), 1)
	}

	if c.String(flag.RegistryCoordinatorFlag.Name) == "" {
		return cli.Exit("registry-coordinator is required", 1)
	}
	registryCoordinatorAddr := gethcommon.HexToAddress(c.String(flag.RegistryCoordinatorFlag.Name))

	avsClient, err := buildAVSClient(registryCoordinatorAddr, ethClient, logger)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to create AVS client %v", err), 1)
	}

	var zeroID [32]byte
	id, err := avsClient.registryCoordinator.GetOperatorId(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if (err != nil) || (id == zeroID) {
		cli.Exit(fmt.Sprintf("Error while GetOperatorId %v", err), 1)
	}

	status, err := avsClient.registryCoordinator.GetOperatorStatus(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if err != nil {
		cli.Exit(fmt.Sprintf("Error while GetOperatorStatus %v", err), 1)
	}

	switch status {
	case 0:
		logger.Info("Operator Status", "status", "NEVER REGISTERED")
	case 1:
		logger.Info("Operator Status", "status", "REGISTERED")
	case 2:
		logger.Info("Operator Status", "status", "DEREGISTERED")
	default:
		cli.Exit(fmt.Sprintf("Unknown operator status %v", status), 1)
	}

	stake, err := avsClient.stakeRegistry.GetLatestStakeUpdate(&bind.CallOpts{Context: context.Background()}, id, uint8(c.Int(flag.QuorumNumberFlag.Name)))
	if err != nil {
		cli.Exit(fmt.Sprintf("Error while GetLatestStakeUpdate %v", err), 1)
	}
	logger.Info("Operator stake update", "stake", stake.Stake, "block number", stake.UpdateBlockNumber)

	return nil
}

func RunGenerateBLSKey(c *cli.Context) error {
	keyPair, err := eigensdkbls.GenRandomBlsKeys()
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to generate BLS key pair %v", err), 1)
	}
	fmt.Println("BLS private key", keyPair.PrivKey.String())
	fmt.Println("BLS public key G1", keyPair.GetPubKeyG1().String())
	fmt.Println("BLS public key G2", keyPair.GetPubKeyG2().String())
	return nil
}

func RunEOChainSetAlias(c *cli.Context) error {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		return cli.Exit(fmt.Sprintf("error creating logger %v", err), 1)
	}

	passphrase := c.String(flag.PassphraseFlag.Name)
	if passphrase == "" {
		return cli.Exit("passphrase is required", 1)
	}

	keyStorePath := c.String(flag.KeyStorePathFlag.Name)
	if keyStorePath == "" {
		return cli.Exit("keystore-path is required", 1)
	}

	EthEcdsaPair, err := eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaEncryptedWallet.json"), passphrase)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Failed to read ecdsaEncryptedWallet.json file %v", err), 1)
	}

	// The following summarizes the logic for setting the alias in the eochain
	// An alias exists | specified as argument | override flag 	| expected behavior 
	//   yes           |   no        		   |    no     		|    use existing
	//   yes		   |   no        		   |    yes    		|    use existing
	//   yes           |   yes       		   |    no     		|    error
	//   yes           |   yes       		   |    yes    		|    use the cli
	//   no            |   no        		   |    no     		|    generate new
	//   no            |   no        		   |    yes    		|    generate new
	//   no            |   yes       		   |    no     		|    use the cli
	//   no            |   yes       		   |    yes    		|    use the cli 

	aliasEcdsaKeyUpdated := false
	var ecdsaPair *ecdsa.PrivateKey
	ecdsaPair, err = eigensdkecdsa.ReadKey(filepath.Join(keyStorePath, "ecdsaAliasedEncryptedWallet.json"), passphrase)
	if err != nil {
		// there was an error reading the alias key (either the file doesn't exist or it is corrupted), consider as if the alias doesn't exist
		if c.String(flag.EcdsaPrivateKeyFlag.Name) != "" {
			// Use the private key passed in the command line
			ecdsaPair, err = crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
			if err != nil {
				return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
			}
			aliasEcdsaKeyUpdated = true
		} else {
			// Generate a new key
			ecdsaPair, err = crypto.GenerateKey()
			if err != nil {
				return cli.Exit(fmt.Sprintf("Failed to generate ECDSA key %v", err), 1)
			}
			aliasEcdsaKeyUpdated = true
			fmt.Println("a new alias ecdsa was generated address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "private key", hex.EncodeToString(ecdsaPair.D.Bytes()))
		}
	} else {
		if c.String(flag.EcdsaPrivateKeyFlag.Name) != "" {
			if !c.Bool(flag.OverrideFlag.Name) {
				return cli.Exit("The alias key already exists, cannot override", 1)
			} 
			// Use the private key passed in the command line
			ecdsaPair, err = crypto.HexToECDSA(c.String(flag.EcdsaPrivateKeyFlag.Name))
			if err != nil {
				return cli.Exit(fmt.Sprintf("Invalid EDCSA private key %v", err), 1)
			}
			aliasEcdsaKeyUpdated = true
		}
	}

	// Save the private key to a file
	if err = eigensdkecdsa.WriteKey(filepath.Join(keyStorePath, "ecdsaAliasedEncryptedWallet.json"), ecdsaPair, passphrase); err != nil {
		return cli.Exit(fmt.Sprintf("Error writing the ecdsaAliasedEncryptedWallet.json file %v", err), 1)
	}
	if aliasEcdsaKeyUpdated {
		fmt.Println("alias ecdsa address ", crypto.PubkeyToAddress(ecdsaPair.PublicKey), "saved")
	}

	if !c.Bool(flag.EncryptOnlyFlag.Name) {
		// Update the alias in the eochain
		if c.String(flag.EOChainEthRPCFlag.Name) == "" {
			return cli.Exit("eochain-eth-rpc is required", 1)
		}

		if c.String(flag.EOConfigAddressFlag.Name) == "" {
			return cli.Exit("eoconfig is required", 1)
		}
		eoConfigAddr := gethcommon.HexToAddress(c.String(flag.EOConfigAddressFlag.Name))

		ethClient, err := eth.NewClient(c.String(flag.EOChainEthRPCFlag.Name))
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create Eth client %v %v", c.String(flag.EOChainEthRPCFlag.Name), err), 1)
		}

		chainIDBigInt, err := ethClient.ChainID(context.Background())
		if err != nil {
			return cli.Exit(fmt.Sprintf("cannot get chainId: %v", err), 1)
		}

		contractEOConfig, err := eoconfig.NewContractEOConfig(eoConfigAddr, ethClient)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to bind the eoconfig contract %v", err), 1)
		}

		signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: EthEcdsaPair}, chainIDBigInt)
		if err != nil {
			return cli.Exit(fmt.Sprintf("error creating the signer function for %v %v", crypto.PubkeyToAddress(EthEcdsaPair.PublicKey), err), 1)
		}

		txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create transaction sender %v", err), 1)
		}
		txMgr := txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerV2, signerAddr)

		noSendTxOpts, err := txMgr.GetNoSendTxOpts()
		if err != nil {
			return cli.Exit(fmt.Sprintf("error creating transaction object %v", err), 1)
		}

		tx, err := contractEOConfig.DeclareAlias(
			noSendTxOpts,
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
		)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Failed to create EOConfig.declareAlias transaction %v", err), 1)
		}

		ctx := context.Background()
		receipt, err := txMgr.Send(ctx, tx)
		if err != nil {
			return cli.Exit(fmt.Sprintf("declareAlias transaction failed %s", err), 1)
		}
		if receipt.Status != 1 {
			return cli.Exit(fmt.Sprintf("declareAlias transaction %v reverted", receipt.TxHash.Hex()), 1)
		}

		logger.Info("succesfully declare the alias in the eochain", "Ethereum address", signerAddr, "eochain address", crypto.PubkeyToAddress(EthEcdsaPair.PublicKey), "tx hash", receipt.TxHash.Hex())
	}
	return nil
}

func buildAVSClient(
	registryCoordinatorAddr gethcommon.Address,
	ethClient eth.Client,
	logger logging.Logger,
) (*avsClient, error) {

	avsClient := &avsClient{
		registryCoordinatorAddr: registryCoordinatorAddr,
	}

	registryCoordinator, err := regcoord.NewContractEORegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to create RegistryCoordinator contract %v", err))
	}
	avsClient.registryCoordinator = registryCoordinator

	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to get ServiceManager from registryCoordinator contract %v", err))
	}
	avsClient.serviceManagerAddr = serviceManagerAddr

	serviceManager, err := smbase.NewContractServiceManagerBase(serviceManagerAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to create serviceManager contract %v", err))
	}
	avsClient.serviceManager = serviceManager

	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to get stakeRegistryAddr %v", err))
	}
	stakeRegistry, err := stakeregistry.NewContractEOStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to create stakeRegistry contract %v", err))
	}
	avsClient.stakeRegistry = stakeRegistry

	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to get delegationManagerAddr %v", err))
	}
	avsClient.delegationManagerAddr = delegationManagerAddr

	avsDirectoryAddr, err := serviceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to get avsDirectoryAddr %v", err))
	}
	avsClient.avsDirectoryAddr = avsDirectoryAddr

	elReader, err := elcontracts.BuildELChainReader(delegationManagerAddr, avsDirectoryAddr, ethClient, logger)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to create ELChainReader %v", err))
	}
	avsClient.elReader = elReader

	return avsClient, nil
}

func ConvertBn254GethToGnark(input regcoord.BN254G1Point) *bn254.G1Affine {
	return eigensdkbls.NewG1Point(input.X, input.Y).G1Affine
}

func ConvertToBN254G1Point(input *eigensdkbls.G1Point) regcoord.BN254G1Point {
	output := regcoord.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *eigensdkbls.G2Point) regcoord.BN254G2Point {
	output := regcoord.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
