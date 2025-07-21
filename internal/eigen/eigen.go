package eigen

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	smbase "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	"github.com/Layr-Labs/eigensdk-go/logging"
	allocationmanager "github.com/eodata/operator-cli/contracts/bindings/AllocationManager"
	regcoord "github.com/eodata/operator-cli/contracts/bindings/EORegistryCoordinator"
	stakeregistry "github.com/eodata/operator-cli/contracts/bindings/EOStakeRegistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type AvsClient struct {
	registryCoordinatorAddr gethcommon.Address
	allocationManagerAddr   gethcommon.Address
	serviceManagerAddr      gethcommon.Address
	delegationManagerAddr   gethcommon.Address
	stakeRegistryAddr       gethcommon.Address
	avsDirectoryAddr        gethcommon.Address
	RegistryCoordinator     *regcoord.EORegistryCoordinator
	AllocationManager       *allocationmanager.AllocationManager
	ServiceManager          *smbase.ContractServiceManagerBase
	StakeRegistry           *stakeregistry.EOStakeRegistry
	ElReader                *elcontracts.ChainReader
}

func BuildAVSClient(
	ethClient *ethclient.Client,
	registryCoordinatorAddress gethcommon.Address,
) (*AvsClient, error) {

	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger %v", err)
	}

	avsClient := &AvsClient{
		registryCoordinatorAddr: registryCoordinatorAddress,
	}

	registryCoordinator, err := regcoord.NewEORegistryCoordinator(
		avsClient.registryCoordinatorAddr,
		ethClient,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RegistryCoordinator contract %v", err)
	}
	avsClient.RegistryCoordinator = registryCoordinator

	avsClient.allocationManagerAddr, err = registryCoordinator.AllocationManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get AllocationManager from registryCoordinator contract %v", err)
	}

	avsClient.AllocationManager, err = allocationmanager.NewAllocationManager(avsClient.allocationManagerAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create AllocationManager contract %v", err)
	}

	avsClient.serviceManagerAddr, err = registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get ServiceManager from registryCoordinator contract %v", err)
	}

	avsClient.serviceManagerAddr, err = registryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to create serviceManager contract %v", err)
	}

	avsClient.ServiceManager, err = smbase.NewContractServiceManagerBase(avsClient.serviceManagerAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create ServiceManager contract %v", err)
	}

	avsClient.stakeRegistryAddr, err = registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get stakeRegistryAddr %v", err)
	}
	avsClient.StakeRegistry, err = stakeregistry.NewEOStakeRegistry(avsClient.stakeRegistryAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create stakeRegistry contract %v", err)
	}

	avsClient.delegationManagerAddr, err = avsClient.StakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get delegationManagerAddr %v", err)
	}

	avsClient.avsDirectoryAddr, err = avsClient.ServiceManager.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get avsDirectoryAddr %v", err)
	}

	elConfig := elcontracts.Config{
		AvsDirectoryAddress:         avsClient.avsDirectoryAddr,
		RewardsCoordinatorAddress:   gethcommon.Address{},
		PermissionControllerAddress: gethcommon.Address{},
		DontUseAllocationManager:    true,
	}
	elReader, _, err := elcontracts.BuildReadClients(elConfig, ethClient, logger, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ELChainReader %v", err)
	}
	avsClient.ElReader = elReader

	return avsClient, nil
}

func CreateEthClient(rpcEndpoint string) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create Eth client %v %v", rpcEndpoint, err)
	}
	return ethclient.NewClient(rpcClient), nil
}
