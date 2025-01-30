package cmd

import (
	"context"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	eoconfig "github.com/eoracle/eoracle-operator-cli/contracts/bindings/EOConfig"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"math/big"
)

func NewResetConfigurationCommand() *cli.Command {
	return &cli.Command{
		Name:        "reset-configuration",
		Description: "Reset configuration in eOracle chain",
		Before:      setProfile,
		Action:      runResetConfiguration,
		Flags: []cli.Flag{
			ProfileFlag,
			EthRPCFlag,
			EOChainRPCFlag,
			EOConfigAddressFlag,
			PassphraseFlag,
			KeyStorePathFlag,
		},
	}
}

func runResetConfiguration(c *cli.Context) error {
	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("passphrase and keystore-path are required")
	}

	ecdsaOperatorPair, ecdsaAliasPair, err := getOperatorAndAliasKeys(c)
	if err != nil {
		return err
	}

	operatorAddress := crypto.PubkeyToAddress(ecdsaOperatorPair.PublicKey)
	operatorAliasAddress := crypto.PubkeyToAddress(ecdsaAliasPair.PublicKey)

	eoChainEthClient, err := createEthClient(c.String(EOChainRPCFlag.Name))
	if err != nil {
		return err
	}

	contractEOConfig, err := getEOConfigContract(eoChainEthClient)
	if err != nil {
		return err
	}

	operatorAlias, err := contractEOConfig.OperatorToAlias(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if err != nil || operatorAlias == (gethcommon.Address{}) {
		utils.Fatalf("Failed to get operator %v alias %v", operatorAddress, err)
	}

	logger.Info("operator details",
		"operator address", operatorAddress.Hex(),
		"alias address", operatorAliasAddress.Hex(),
		"operator is EOA", true,
	)

	if operatorAlias != operatorAliasAddress {
		utils.Fatalf("Operator (%v) alias (%v) does not match the expected alias (%v), please contact eOracle support",
			operatorAddress.Hex(), operatorAlias.Hex(), operatorAliasAddress.Hex())
	}

	balance, err := eoChainEthClient.BalanceAt(context.Background(), operatorAddress, nil)
	if err != nil {
		utils.Fatalf("Error while getting the operator (%v) balance %v on eoChain", operatorAddress, err)
	}

	if balance.Cmp(big.NewInt(500000000000000000)) > 0 {
		eochainChainIDBigInt, err := eoChainEthClient.ChainID(context.Background())
		if err != nil {
			utils.Fatalf("Error getting chainId (%v): %v", c.String(EOChainRPCFlag.Name), err)
		}

		returnBalance := balance.Sub(balance, big.NewInt(500000000000000000))
		signerV2, signerAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaOperatorPair}, eochainChainIDBigInt)
		if err != nil {
			utils.Fatalf("Error creating the signer function for operator %v on eoChain (%v) %v",
				crypto.PubkeyToAddress(ecdsaOperatorPair.PublicKey), c.String(EOChainRPCFlag.Name), err)
		}

		txSender, err := wallet.NewPrivateKeyWallet(eoChainEthClient, signerV2, signerAddr, logger)
		if err != nil {
			utils.Fatalf("Error creating the transaction sender for operator %v on eoChain (%v) %v",
				crypto.PubkeyToAddress(ecdsaOperatorPair.PublicKey), c.String(EOChainRPCFlag.Name), err)
		}
		txMgr := txmgr.NewSimpleTxManager(txSender, eoChainEthClient, logger, signerAddr)
		txOpts, err := txMgr.GetNoSendTxOpts()
		if err != nil {
			utils.Fatalf("Error generating transaction for resetting eochain gas balance of operator %v on eoChain (%v) %v",
				operatorAddress.Hex(), c.String(EOChainRPCFlag.Name), err)
		}
		txOpts.Value = returnBalance

		contractEOConfigRaw := eoconfig.EOConfigRaw{Contract: contractEOConfig}
		tx, err := contractEOConfigRaw.Transfer(txOpts)
		if err != nil {
			utils.Fatalf("Error reseting the operator %v balance on eochain (%v) %v", operatorAddress.Hex(), c.String(EOChainRPCFlag.Name), err)
		}

		ctx := context.Background()
		receipt, err := txMgr.Send(ctx, tx, true)
		if err != nil {
			utils.Fatalf("Error sending the reset balance transaction of operator %v on eochain (%v) %v",
				operatorAddress.Hex(), c.String(EOChainRPCFlag.Name), err)
		}

		if receipt.Status != 1 {
			utils.Fatalf("The transaction %v to reset the operator %v balance on eochain (%v) reverted",
				receipt.TxHash.Hex(), operatorAddress.Hex(), c.String(EOChainRPCFlag.Name))
		}

		balance, err = eoChainEthClient.BalanceAt(context.Background(), operatorAddress, nil)
		if err != nil {
			utils.Fatalf("Error while getting the operator (%v) balance %v on eoChain", operatorAddress, err)
		}
	}
	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(1e18)))
	logger.Info("Operator balance", "operator address", operatorAddress.Hex(), "balance", balanceInEth.String())

	return nil
}

func getEOConfigContract(ethClient *ethclient.Client) (*eoconfig.EOConfig, error) {
	return eoconfig.NewEOConfig(profile.EOConfigAddress, ethClient)
}
