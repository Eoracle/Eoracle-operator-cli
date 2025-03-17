package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"path/filepath"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	smbase "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	eigensdkbls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	regcoord "github.com/eoracle/eoracle-operator-cli/contracts/bindings/EORegistryCoordinator"
	stakeregistry "github.com/eoracle/eoracle-operator-cli/contracts/bindings/EOStakeRegistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

type avsClient struct {
	registryCoordinatorAddr gethcommon.Address
	serviceManagerAddr      gethcommon.Address
	delegationManagerAddr   gethcommon.Address
	avsDirectoryAddr        gethcommon.Address
	registryCoordinator     *regcoord.EORegistryCoordinator
	serviceManager          *smbase.ContractServiceManagerBase
	stakeRegistry           *stakeregistry.EOStakeRegistry
	elReader                *elcontracts.ChainReader
}

func buildAVSClient(
	ethClient *ethclient.Client,
) (*avsClient, error) {

	avsClient := &avsClient{
		registryCoordinatorAddr: profile.RegistryCoordinatorAddress,
	}

	registryCoordinator, err := regcoord.NewEORegistryCoordinator(
		profile.RegistryCoordinatorAddress,
		ethClient,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create RegistryCoordinator contract %v", err)
	}
	avsClient.registryCoordinator = registryCoordinator

	serviceManagerAddr, err := registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get ServiceManager from registryCoordinator contract %v", err)
	}
	avsClient.serviceManagerAddr = serviceManagerAddr

	serviceManager, err := smbase.NewContractServiceManagerBase(serviceManagerAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create serviceManager contract %v", err)
	}
	avsClient.serviceManager = serviceManager

	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get stakeRegistryAddr %v", err)
	}
	stakeRegistry, err := stakeregistry.NewEOStakeRegistry(stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create stakeRegistry contract %v", err)
	}
	avsClient.stakeRegistry = stakeRegistry

	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get delegationManagerAddr %v", err)
	}
	avsClient.delegationManagerAddr = delegationManagerAddr

	avsDirectoryAddr, err := serviceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get avsDirectoryAddr %v", err)
	}
	avsClient.avsDirectoryAddr = avsDirectoryAddr

	elConfig := elcontracts.Config{
		RewardsCoordinatorAddress: 	 gethcommon.Address{},
		PermissionControllerAddress: gethcommon.Address{},
		DontUseAllocationManager: 	 true,
	}
	elReader, _, err := elcontracts.BuildReadClients(elConfig, ethClient, logger, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ELChainReader %v", err)
	}
	avsClient.elReader = elReader

	return avsClient, nil
}

func createEthClient(rpcEndpoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		utils.Fatalf("failed to create Eth client %v %v", rpcEndpoint, err)
	}
	return ethclient.NewClient(rpcClient), nil
}

func getECDSAPrivateKey(c *cli.Context) (*ecdsa.PrivateKey, error) {
	if !c.IsSet(PassphraseFlag.Name) {
		utils.Fatalf("passphrase is required")
	}
	if !c.IsSet(KeyStorePathFlag.Name) {
		utils.Fatalf("keystore-path is required")
	}
	if c.IsSet(EcdsaPrivateKeyFlag.Name) {
		return crypto.HexToECDSA(c.String(EcdsaPrivateKeyFlag.Name))
	}
	return eigensdkecdsa.ReadKey(
		filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaEncryptedWallet.json"),
		c.String(PassphraseFlag.Name),
	)
}

func getBLSPrivateKey(c *cli.Context) (*eigensdkbls.KeyPair, error) {
	return eigensdkbls.NewKeyPairFromString(c.String(BlsPrivateKeyFlag.Name))
}

func getOperatorAndAliasKeys(c *cli.Context) (*ecdsa.PrivateKey, *ecdsa.PrivateKey, error) {
	ecdsaOperatorPair, err := eigensdkecdsa.ReadKey(filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaEncryptedWallet.json"), c.String(PassphraseFlag.Name))
	if err != nil {
		utils.Fatalf("Failed to read ecdsaEncryptedWallet.json file %v", err)
	}

	ecdsaAliasPair, err := eigensdkecdsa.ReadKey(filepath.Join(c.String(KeyStorePathFlag.Name), "ecdsaAliasedEncryptedWallet.json"), c.String(PassphraseFlag.Name))
	if err != nil {
		utils.Fatalf("Failed to read ecdsaAliasedEncryptedWallet.json file %v", err)
	}

	return ecdsaOperatorPair, ecdsaAliasPair, nil
}
