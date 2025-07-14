package cmd

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func NewDeregisterCommand() *cli.Command {
	return &cli.Command{
		Name:        "deregister",
		Description: "Deregister the operator",
		Before:      setProfile,
		Action:      runDeregister,
		Flags: []cli.Flag{
			ProfileFlag,
			EthRPCFlag,
			RegistryCoordinatorFlag,
			PassphraseFlag,
			KeyStorePathFlag,
			EcdsaPrivateKeyFlag,
			QuorumNumberFlag,
		},
	}
}

func runDeregister(c *cli.Context) error {
	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("passphrase and keystore-path are required")
	}

	ecdsaPair, err := getECDSAPrivateKey(c)
	if err != nil {
		return err
	}

	ethClient, err := createEthClient(profile.EthRPCEndpoint)
	if err != nil {
		return err
	}

	chainIDBigInt, err := ethClient.ChainID(context.Background())
	if err != nil {
		utils.Fatalf("cannot get chainId (%v): %v", profile.EthRPCEndpoint, err)
	}

	avsClient, err := buildAVSClient(ethClient)
	if err != nil {
		utils.Fatalf("Error creating AVS client %v", err)
	}

	signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPair}, chainIDBigInt)
	if err != nil {
		utils.Fatalf(
			"Error creating the deregister transaction signer for operator %v on Ethereum mainnet/Holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	txSender, err := wallet.NewPrivateKeyWallet(ethClient, signerV2, signerAddr, logger)
	if err != nil {
		utils.Fatalf(
			"Error creating the deregister transaction sender for operator %v on Ethereum mainnet/Holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	txMgr := txmgr.NewSimpleTxManager(txSender, ethClient, logger, signerAddr)
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		utils.Fatalf("Error creating transaction object %v", err)
	}

	tx, err := avsClient.registryCoordinator.DeregisterOperator(
		noSendTxOpts,
		crypto.PubkeyToAddress(ecdsaPair.PublicKey),
		common.Address{},
		[]uint32{0},
	)
	if err != nil {
		utils.Fatalf(
			"Error creating the deregister transaction for operator %v on Ethereum mainnet/Holesky (%v) %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}

	ctx := context.Background()
	receipt, err := txMgr.Send(ctx, tx, true)
	if err != nil {
		utils.Fatalf(
			"deregister transaction for operator %v on Ethereum mainnet/Holesky (%v) failed %v",
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
			err,
		)
	}
	if receipt.Status != 1 {
		utils.Fatalf(
			"deregister transaction %v for operator %v on Ethereum mainnet/Holesky (%v) reverted",
			receipt.TxHash.Hex(),
			crypto.PubkeyToAddress(ecdsaPair.PublicKey),
			profile.EthRPCEndpoint,
		)
	}

	logger.Info(
		"successfully deregistered from eoracle AVS",
		"address", signerAddr,
		"tx hash", receipt.TxHash.Hex(),
	)

	return nil
}
