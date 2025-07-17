package cmd

import (
	"context"
	"fmt"
	"math/big"

	eoconfig "github.com/eodata/operator-cli/contracts/bindings/EOConfig"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

func NewPrintStatusCommand() *cli.Command {
	return &cli.Command{
		Name:        "print-status",
		Description: "Print the operator status",
		Before:      setProfile,
		Action:      runPrintStatus,
		Flags: []cli.Flag{
			ProfileFlag,
			EthRPCFlag,
			EOChainRPCFlag,
			EOConfigAddressFlag,
			RegistryCoordinatorFlag,
			PassphraseFlag,
			KeyStorePathFlag,
			QuorumNumberFlag,
		},
	}
}

func runPrintStatus(c *cli.Context) error {
	if !c.IsSet(PassphraseFlag.Name) || !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("passphrase and keystore-path are required")
	}

	ecdsaOperatorPair, ecdsaAliasPair, err := getOperatorAndAliasKeys(c)
	if err != nil {
		return err
	}

	ethClient, err := createEthClient(profile.EthRPCEndpoint)
	if err != nil {
		return err
	}

	avsClient, err := buildAVSClient(ethClient)
	if err != nil {
		utils.Fatalf("Failed to create AVS client %v", err)
	}

	operatorAddress := crypto.PubkeyToAddress(ecdsaOperatorPair.PublicKey)
	operatorAliasAddress := crypto.PubkeyToAddress(ecdsaAliasPair.PublicKey)

	operatorIsEOA, err := isEOA(ethClient, operatorAddress)
	if err != nil {
		return err
	}

	eoChainEthClient, err := createEthClient(profile.EOChainRPCEndpoint)
	if err != nil {
		return err
	}

	contractEOConfig, err := eoconfig.NewEOConfig(profile.EOConfigAddress, eoChainEthClient)
	if err != nil {
		utils.Fatalf("Failed to load EOConfig contract %v", err)
	}

	operatorAlias, err := contractEOConfig.OperatorToAlias(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if err != nil || operatorAlias == (gethcommon.Address{}) {
		utils.Fatalf("Failed to get operator %v alias %v", operatorAddress, err)
	}

	logger.Info(
		"operator details",
		"operator address", operatorAddress,
		"alias address", operatorAliasAddress,
		"operator is EOA", operatorIsEOA,
	)

	if operatorAlias != operatorAliasAddress {
		utils.Fatalf(
			"Operator (%v) alias (%v) does not match the expected alias (%v)",
			operatorAddress,
			operatorAlias,
			operatorAliasAddress,
		)
	}

	id, err := avsClient.registryCoordinator.GetOperatorId(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if err != nil || id == [32]byte{} {
		utils.Fatalf("Error while GetOperatorId %v", err)
	}

	status, err := avsClient.registryCoordinator.GetOperatorStatus(&bind.CallOpts{Context: context.Background()}, operatorAddress)
	if err != nil {
		utils.Fatalf("Error while GetOperatorStatus %v", err)
	}

	switch status {
	case 0:
		logger.Info("Operator Status", "status", "NEVER REGISTERED")
	case 1:
		logger.Info("Operator Status", "status", "REGISTERED")
	case 2:
		logger.Info("Operator Status", "status", "DEREGISTERED")
	default:
		utils.Fatalf("Unknown operator status %v", status)
	}

	stake, err := avsClient.stakeRegistry.GetLatestStakeUpdate(
		&bind.CallOpts{Context: context.Background()},
		id,
		uint8(c.Int(QuorumNumberFlag.Name)),
	)
	if err != nil {
		utils.Fatalf("Error while GetLatestStakeUpdate %v", err)
	}

	logger.Info("Operator stake update", "stake", stake.Stake, "block number", stake.UpdateBlockNumber)

	printBalance(eoChainEthClient, operatorAddress, "operator address")
	printBalance(eoChainEthClient, operatorAliasAddress, "operator alias address")

	return nil
}

func isEOA(ethClient *ethclient.Client, address gethcommon.Address) (bool, error) {
	code, err := ethClient.CodeAt(context.Background(), address, nil)
	if err != nil {
		utils.Fatalf("Error checking if address is EOA %v", err)
	}
	return len(code) == 0, nil
}

func printBalance(ethClient *ethclient.Client, address gethcommon.Address, label string) {
	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		utils.Fatalf("Error while getting the %s balance %v", label, err)
	}
	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(1e18)))
	logger.Info(fmt.Sprintf("%s balance", label), "address", address.Hex(), "balance", balanceInEth.String())
}
