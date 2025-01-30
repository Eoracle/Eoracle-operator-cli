package cmd

import (
	"context"
	"crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	eoconfig "github.com/eoracle/eoracle-operator-cli/contracts/bindings/EOConfig"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func NewDeclareAliasCommand() *cli.Command {
	return &cli.Command{
		Name:        "declare-alias",
		Description: "Declare the alias in the eochain",
		Before:      setProfile,
		Action:      runDeclareAlias,
		Flags: []cli.Flag{
			ProfileFlag,
			PassphraseFlag,
			KeyStorePathFlag,
			EOChainRPCFlag,
			EOConfigAddressFlag,
		},
	}
}

func runDeclareAlias(c *cli.Context) error {
	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("passphrase and keystore-path are required")
	}

	ethEcdsaPair, aliasEcdsaPair, err := getOperatorAndAliasKeys(c)
	if err != nil {
		return err
	}

	txMgr, contractEOConfig, err := getTxMgrForEOChain(ethEcdsaPair)
	if err != nil {
		return err
	}

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		utils.Fatalf("Error creating transaction object %v", err)
	}

	tx, err := contractEOConfig.DeclareAlias(
		noSendTxOpts,
		crypto.PubkeyToAddress(aliasEcdsaPair.PublicKey),
	)
	if err != nil {
		utils.Fatalf("Failed to create EOConfig.declareAlias transaction %v", err)
	}

	ctx := context.Background()
	receipt, err := txMgr.Send(ctx, tx, true)
	if err != nil {
		utils.Fatalf("declareAlias transaction failed %s", err)
	}

	if receipt.Status != 1 {
		utils.Fatalf(
			"declareAlias transaction %v for operator %v on Ethereum mainnet/Holesky (%v) reverted",
			receipt.TxHash.Hex(),
			crypto.PubkeyToAddress(ethEcdsaPair.PublicKey),
			profile.EOChainRPCEndpoint,
		)
	}

	logger.Info(
		"successfully declared an alias in the eochain",
		"Ethereum address", crypto.PubkeyToAddress(ethEcdsaPair.PublicKey),
		"eochain address", crypto.PubkeyToAddress(aliasEcdsaPair.PublicKey),
		"tx hash", receipt.TxHash.Hex(),
	)
	return nil
}

func getTxMgrForEOChain(ethEcdsaPair *ecdsa.PrivateKey) (*txmgr.SimpleTxManager, *eoconfig.EOConfig, error) {
	ethClient, err := createEthClient(profile.EOChainRPCEndpoint)
	if err != nil {
		return nil, nil, err
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		utils.Fatalf("cannot get chainId (%v): %v", profile.EOChainRPCEndpoint, err)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ethEcdsaPair}, chainIDBigInt)
	if err != nil {
		utils.Fatalf("Error creating the signer function for %v %v", crypto.PubkeyToAddress(ethEcdsaPair.PublicKey), err)
	}
	txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		utils.Fatalf("Failed to create transaction sender for declaring alias of operator %v on eoChain (%v) %v",
			crypto.PubkeyToAddress(ethEcdsaPair.PublicKey), profile.EOChainRPCEndpoint, err)
	}

	contractEOConfig, err := eoconfig.NewEOConfig(profile.EOConfigAddress, ethClient)
	if err != nil {
		utils.Fatalf("Failed to bind the eoconfig contract %v", err)
	}

	return txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerAddr), contractEOConfig, nil
}
