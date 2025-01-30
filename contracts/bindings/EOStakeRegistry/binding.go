// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractEOStakeRegistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IStakeRegistryTypesStakeUpdate is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryTypesStakeUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	Stake                 *big.Int
}

// IStakeRegistryTypesStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryTypesStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// EOStakeRegistryMetaData contains all meta data concerning the EOStakeRegistry contract.
var EOStakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slashingRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_WEIGHING_FUNCTION_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WEIGHTING_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategies\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentStake\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentTotalStake\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestStakeUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRegistryTypes.StakeUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeAtBlockNumber\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeHistory\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StakeUpdate[]\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRegistryTypes.StakeUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeUpdateIndexAtBlockNumber\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeAtBlockNumberFromIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStakeUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRegistryTypes.StakeUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeSlashableStakeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorSetQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumStakeForQuorum\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyStrategyParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"strategyIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"newMultipliers\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategies\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"indicesToRemove\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainManager\",\"inputs\":[{\"name\":\"newChainManager\",\"type\":\"address\",\"internalType\":\"contractIEOChainManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumStakeForQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashableStakeLookahead\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_lookAheadBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashableStakeLookAheadPerQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeTypePerQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIStakeRegistryTypes.StakeType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategiesPerQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyParamsByIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyParamsLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateOperatorStake\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weightOfOperatorForQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LookAheadPeriodChanged\",\"inputs\":[{\"name\":\"oldLookAheadBlocks\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newLookAheadBlocks\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumStakeForQuorumUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorStakeUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"stake\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumCreated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeTypeSet\",\"inputs\":[{\"name\":\"newStakeType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIStakeRegistryTypes.StakeType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyMultiplierUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BelowMinimumStakeRequirement\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyStakeHistory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputDuplicateStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputMultiplierZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySlashingRegistryCoordinator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySlashingRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]}]",
}

// EOStakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EOStakeRegistryMetaData.ABI instead.
var EOStakeRegistryABI = EOStakeRegistryMetaData.ABI

// EOStakeRegistry is an auto generated Go binding around an Ethereum contract.
type EOStakeRegistry struct {
	EOStakeRegistryCaller     // Read-only binding to the contract
	EOStakeRegistryTransactor // Write-only binding to the contract
	EOStakeRegistryFilterer   // Log filterer for contract events
}

// EOStakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EOStakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOStakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EOStakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOStakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EOStakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOStakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EOStakeRegistrySession struct {
	Contract     *EOStakeRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EOStakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EOStakeRegistryCallerSession struct {
	Contract *EOStakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EOStakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EOStakeRegistryTransactorSession struct {
	Contract     *EOStakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EOStakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EOStakeRegistryRaw struct {
	Contract *EOStakeRegistry // Generic contract binding to access the raw methods on
}

// EOStakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EOStakeRegistryCallerRaw struct {
	Contract *EOStakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EOStakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EOStakeRegistryTransactorRaw struct {
	Contract *EOStakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEOStakeRegistry creates a new instance of EOStakeRegistry, bound to a specific deployed contract.
func NewEOStakeRegistry(address common.Address, backend bind.ContractBackend) (*EOStakeRegistry, error) {
	contract, err := bindEOStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistry{EOStakeRegistryCaller: EOStakeRegistryCaller{contract: contract}, EOStakeRegistryTransactor: EOStakeRegistryTransactor{contract: contract}, EOStakeRegistryFilterer: EOStakeRegistryFilterer{contract: contract}}, nil
}

// NewEOStakeRegistryCaller creates a new read-only instance of EOStakeRegistry, bound to a specific deployed contract.
func NewEOStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*EOStakeRegistryCaller, error) {
	contract, err := bindEOStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryCaller{contract: contract}, nil
}

// NewEOStakeRegistryTransactor creates a new write-only instance of EOStakeRegistry, bound to a specific deployed contract.
func NewEOStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EOStakeRegistryTransactor, error) {
	contract, err := bindEOStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryTransactor{contract: contract}, nil
}

// NewEOStakeRegistryFilterer creates a new log filterer instance of EOStakeRegistry, bound to a specific deployed contract.
func NewEOStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EOStakeRegistryFilterer, error) {
	contract, err := bindEOStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryFilterer{contract: contract}, nil
}

// bindEOStakeRegistry binds a generic wrapper to an already deployed contract.
func bindEOStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EOStakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EOStakeRegistry *EOStakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOStakeRegistry.Contract.EOStakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EOStakeRegistry *EOStakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.EOStakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EOStakeRegistry *EOStakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.EOStakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EOStakeRegistry *EOStakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOStakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EOStakeRegistry *EOStakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EOStakeRegistry *EOStakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistryCaller) MAXWEIGHINGFUNCTIONLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "MAX_WEIGHING_FUNCTION_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistrySession) MAXWEIGHINGFUNCTIONLENGTH() (uint8, error) {
	return _EOStakeRegistry.Contract.MAXWEIGHINGFUNCTIONLENGTH(&_EOStakeRegistry.CallOpts)
}

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) MAXWEIGHINGFUNCTIONLENGTH() (uint8, error) {
	return _EOStakeRegistry.Contract.MAXWEIGHINGFUNCTIONLENGTH(&_EOStakeRegistry.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCaller) WEIGHTINGDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "WEIGHTING_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistrySession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _EOStakeRegistry.Contract.WEIGHTINGDIVISOR(&_EOStakeRegistry.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _EOStakeRegistry.Contract.WEIGHTINGDIVISOR(&_EOStakeRegistry.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) AllocationManager() (common.Address, error) {
	return _EOStakeRegistry.Contract.AllocationManager(&_EOStakeRegistry.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) AllocationManager() (common.Address, error) {
	return _EOStakeRegistry.Contract.AllocationManager(&_EOStakeRegistry.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) AvsDirectory() (common.Address, error) {
	return _EOStakeRegistry.Contract.AvsDirectory(&_EOStakeRegistry.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) AvsDirectory() (common.Address, error) {
	return _EOStakeRegistry.Contract.AvsDirectory(&_EOStakeRegistry.CallOpts)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) ChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "chainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) ChainManager() (common.Address, error) {
	return _EOStakeRegistry.Contract.ChainManager(&_EOStakeRegistry.CallOpts)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) ChainManager() (common.Address, error) {
	return _EOStakeRegistry.Contract.ChainManager(&_EOStakeRegistry.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) Delegation() (common.Address, error) {
	return _EOStakeRegistry.Contract.Delegation(&_EOStakeRegistry.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) Delegation() (common.Address, error) {
	return _EOStakeRegistry.Contract.Delegation(&_EOStakeRegistry.CallOpts)
}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetCurrentStake(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getCurrentStake", operatorId, quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) GetCurrentStake(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetCurrentStake(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetCurrentStake(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetCurrentStake(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetCurrentTotalStake(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getCurrentTotalStake", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) GetCurrentTotalStake(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetCurrentTotalStake(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetCurrentTotalStake(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetCurrentTotalStake(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCaller) GetLatestStakeUpdate(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (IStakeRegistryTypesStakeUpdate, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getLatestStakeUpdate", operatorId, quorumNumber)

	if err != nil {
		return *new(IStakeRegistryTypesStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryTypesStakeUpdate)).(*IStakeRegistryTypesStakeUpdate)

	return out0, err

}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistrySession) GetLatestStakeUpdate(operatorId [32]byte, quorumNumber uint8) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetLatestStakeUpdate(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetLatestStakeUpdate(operatorId [32]byte, quorumNumber uint8) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetLatestStakeUpdate(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeAtBlockNumber(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeAtBlockNumber", operatorId, quorumNumber, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeAtBlockNumber(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeAtBlockNumber(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeAtBlockNumberAndIndex", quorumNumber, blockNumber, operatorId, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeAtBlockNumberAndIndex(&_EOStakeRegistry.CallOpts, quorumNumber, blockNumber, operatorId, index)
}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeAtBlockNumberAndIndex(&_EOStakeRegistry.CallOpts, quorumNumber, blockNumber, operatorId, index)
}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeHistory(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryTypesStakeUpdate, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeHistory", operatorId, quorumNumber)

	if err != nil {
		return *new([]IStakeRegistryTypesStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakeRegistryTypesStakeUpdate)).(*[]IStakeRegistryTypesStakeUpdate)

	return out0, err

}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeHistory(operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetStakeHistory(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeHistory(operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetStakeHistory(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeHistoryLength(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeHistoryLength", operatorId, quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeHistoryLength(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeHistoryLength(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeHistoryLength(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetStakeHistoryLength(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeUpdateAtIndex", quorumNumber, operatorId, index)

	if err != nil {
		return *new(IStakeRegistryTypesStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryTypesStakeUpdate)).(*IStakeRegistryTypesStakeUpdate)

	return out0, err

}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeUpdateAtIndex(quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetStakeUpdateAtIndex(&_EOStakeRegistry.CallOpts, quorumNumber, operatorId, index)
}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeUpdateAtIndex(quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetStakeUpdateAtIndex(&_EOStakeRegistry.CallOpts, quorumNumber, operatorId, index)
}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetStakeUpdateIndexAtBlockNumber(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getStakeUpdateIndexAtBlockNumber", operatorId, quorumNumber, blockNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistrySession) GetStakeUpdateIndexAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _EOStakeRegistry.Contract.GetStakeUpdateIndexAtBlockNumber(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetStakeUpdateIndexAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _EOStakeRegistry.Contract.GetStakeUpdateIndexAtBlockNumber(&_EOStakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetTotalStakeAtBlockNumberFromIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getTotalStakeAtBlockNumberFromIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) GetTotalStakeAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeAtBlockNumberFromIndex(&_EOStakeRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetTotalStakeAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeAtBlockNumberFromIndex(&_EOStakeRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCaller) GetTotalStakeHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getTotalStakeHistoryLength", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistrySession) GetTotalStakeHistoryLength(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeHistoryLength(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetTotalStakeHistoryLength(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeHistoryLength(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_EOStakeRegistry *EOStakeRegistryCaller) GetTotalStakeIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getTotalStakeIndicesAtBlockNumber", blockNumber, quorumNumbers)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_EOStakeRegistry *EOStakeRegistrySession) GetTotalStakeIndicesAtBlockNumber(blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeIndicesAtBlockNumber(&_EOStakeRegistry.CallOpts, blockNumber, quorumNumbers)
}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetTotalStakeIndicesAtBlockNumber(blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeIndicesAtBlockNumber(&_EOStakeRegistry.CallOpts, blockNumber, quorumNumbers)
}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCaller) GetTotalStakeUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "getTotalStakeUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IStakeRegistryTypesStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryTypesStakeUpdate)).(*IStakeRegistryTypesStakeUpdate)

	return out0, err

}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistrySession) GetTotalStakeUpdateAtIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeUpdateAtIndex(&_EOStakeRegistry.CallOpts, quorumNumber, index)
}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_EOStakeRegistry *EOStakeRegistryCallerSession) GetTotalStakeUpdateAtIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStakeUpdate, error) {
	return _EOStakeRegistry.Contract.GetTotalStakeUpdateAtIndex(&_EOStakeRegistry.CallOpts, quorumNumber, index)
}

// IsOperatorSetQuorum is a free data retrieval call binding the contract method 0x9f8aff26.
//
// Solidity: function isOperatorSetQuorum(uint8 quorumNumber) view returns(bool)
func (_EOStakeRegistry *EOStakeRegistryCaller) IsOperatorSetQuorum(opts *bind.CallOpts, quorumNumber uint8) (bool, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "isOperatorSetQuorum", quorumNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetQuorum is a free data retrieval call binding the contract method 0x9f8aff26.
//
// Solidity: function isOperatorSetQuorum(uint8 quorumNumber) view returns(bool)
func (_EOStakeRegistry *EOStakeRegistrySession) IsOperatorSetQuorum(quorumNumber uint8) (bool, error) {
	return _EOStakeRegistry.Contract.IsOperatorSetQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// IsOperatorSetQuorum is a free data retrieval call binding the contract method 0x9f8aff26.
//
// Solidity: function isOperatorSetQuorum(uint8 quorumNumber) view returns(bool)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) IsOperatorSetQuorum(quorumNumber uint8) (bool, error) {
	return _EOStakeRegistry.Contract.IsOperatorSetQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) MinimumStakeForQuorum(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "minimumStakeForQuorum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) MinimumStakeForQuorum(arg0 uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.MinimumStakeForQuorum(&_EOStakeRegistry.CallOpts, arg0)
}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) MinimumStakeForQuorum(arg0 uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.MinimumStakeForQuorum(&_EOStakeRegistry.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _EOStakeRegistry.Contract.RegistryCoordinator(&_EOStakeRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _EOStakeRegistry.Contract.RegistryCoordinator(&_EOStakeRegistry.CallOpts)
}

// SlashableStakeLookAheadPerQuorum is a free data retrieval call binding the contract method 0x9ab4d6ff.
//
// Solidity: function slashableStakeLookAheadPerQuorum(uint8 quorumNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistryCaller) SlashableStakeLookAheadPerQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "slashableStakeLookAheadPerQuorum", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlashableStakeLookAheadPerQuorum is a free data retrieval call binding the contract method 0x9ab4d6ff.
//
// Solidity: function slashableStakeLookAheadPerQuorum(uint8 quorumNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistrySession) SlashableStakeLookAheadPerQuorum(quorumNumber uint8) (uint32, error) {
	return _EOStakeRegistry.Contract.SlashableStakeLookAheadPerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// SlashableStakeLookAheadPerQuorum is a free data retrieval call binding the contract method 0x9ab4d6ff.
//
// Solidity: function slashableStakeLookAheadPerQuorum(uint8 quorumNumber) view returns(uint32)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) SlashableStakeLookAheadPerQuorum(quorumNumber uint8) (uint32, error) {
	return _EOStakeRegistry.Contract.SlashableStakeLookAheadPerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// StakeTypePerQuorum is a free data retrieval call binding the contract method 0x697fbd93.
//
// Solidity: function stakeTypePerQuorum(uint8 quorumNumber) view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistryCaller) StakeTypePerQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint8, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "stakeTypePerQuorum", quorumNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StakeTypePerQuorum is a free data retrieval call binding the contract method 0x697fbd93.
//
// Solidity: function stakeTypePerQuorum(uint8 quorumNumber) view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistrySession) StakeTypePerQuorum(quorumNumber uint8) (uint8, error) {
	return _EOStakeRegistry.Contract.StakeTypePerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// StakeTypePerQuorum is a free data retrieval call binding the contract method 0x697fbd93.
//
// Solidity: function stakeTypePerQuorum(uint8 quorumNumber) view returns(uint8)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) StakeTypePerQuorum(quorumNumber uint8) (uint8, error) {
	return _EOStakeRegistry.Contract.StakeTypePerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 quorumNumber, uint256 ) view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCaller) StrategiesPerQuorum(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "strategiesPerQuorum", quorumNumber, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 quorumNumber, uint256 ) view returns(address)
func (_EOStakeRegistry *EOStakeRegistrySession) StrategiesPerQuorum(quorumNumber uint8, arg1 *big.Int) (common.Address, error) {
	return _EOStakeRegistry.Contract.StrategiesPerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber, arg1)
}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 quorumNumber, uint256 ) view returns(address)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) StrategiesPerQuorum(quorumNumber uint8, arg1 *big.Int) (common.Address, error) {
	return _EOStakeRegistry.Contract.StrategiesPerQuorum(&_EOStakeRegistry.CallOpts, quorumNumber, arg1)
}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 quorumNumber, uint256 ) view returns(address strategy, uint96 multiplier)
func (_EOStakeRegistry *EOStakeRegistryCaller) StrategyParams(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "strategyParams", quorumNumber, arg1)

	outstruct := new(struct {
		Strategy   common.Address
		Multiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Multiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 quorumNumber, uint256 ) view returns(address strategy, uint96 multiplier)
func (_EOStakeRegistry *EOStakeRegistrySession) StrategyParams(quorumNumber uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _EOStakeRegistry.Contract.StrategyParams(&_EOStakeRegistry.CallOpts, quorumNumber, arg1)
}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 quorumNumber, uint256 ) view returns(address strategy, uint96 multiplier)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) StrategyParams(quorumNumber uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _EOStakeRegistry.Contract.StrategyParams(&_EOStakeRegistry.CallOpts, quorumNumber, arg1)
}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_EOStakeRegistry *EOStakeRegistryCaller) StrategyParamsByIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStrategyParams, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "strategyParamsByIndex", quorumNumber, index)

	if err != nil {
		return *new(IStakeRegistryTypesStrategyParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryTypesStrategyParams)).(*IStakeRegistryTypesStrategyParams)

	return out0, err

}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_EOStakeRegistry *EOStakeRegistrySession) StrategyParamsByIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStrategyParams, error) {
	return _EOStakeRegistry.Contract.StrategyParamsByIndex(&_EOStakeRegistry.CallOpts, quorumNumber, index)
}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_EOStakeRegistry *EOStakeRegistryCallerSession) StrategyParamsByIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryTypesStrategyParams, error) {
	return _EOStakeRegistry.Contract.StrategyParamsByIndex(&_EOStakeRegistry.CallOpts, quorumNumber, index)
}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCaller) StrategyParamsLength(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "strategyParamsLength", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistrySession) StrategyParamsLength(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.StrategyParamsLength(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) StrategyParamsLength(quorumNumber uint8) (*big.Int, error) {
	return _EOStakeRegistry.Contract.StrategyParamsLength(&_EOStakeRegistry.CallOpts, quorumNumber)
}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCaller) WeightOfOperatorForQuorum(opts *bind.CallOpts, quorumNumber uint8, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EOStakeRegistry.contract.Call(opts, &out, "weightOfOperatorForQuorum", quorumNumber, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistrySession) WeightOfOperatorForQuorum(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _EOStakeRegistry.Contract.WeightOfOperatorForQuorum(&_EOStakeRegistry.CallOpts, quorumNumber, operator)
}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_EOStakeRegistry *EOStakeRegistryCallerSession) WeightOfOperatorForQuorum(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _EOStakeRegistry.Contract.WeightOfOperatorForQuorum(&_EOStakeRegistry.CallOpts, quorumNumber, operator)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) AddStrategies(opts *bind.TransactOpts, quorumNumber uint8, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "addStrategies", quorumNumber, _strategyParams)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) AddStrategies(quorumNumber uint8, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.AddStrategies(&_EOStakeRegistry.TransactOpts, quorumNumber, _strategyParams)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) AddStrategies(quorumNumber uint8, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.AddStrategies(&_EOStakeRegistry.TransactOpts, quorumNumber, _strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "deregisterOperator", operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.DeregisterOperator(&_EOStakeRegistry.TransactOpts, operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.DeregisterOperator(&_EOStakeRegistry.TransactOpts, operatorId, quorumNumbers)
}

// InitializeDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x75d4173a.
//
// Solidity: function initializeDelegatedStakeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) InitializeDelegatedStakeQuorum(opts *bind.TransactOpts, quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "initializeDelegatedStakeQuorum", quorumNumber, minimumStake, _strategyParams)
}

// InitializeDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x75d4173a.
//
// Solidity: function initializeDelegatedStakeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) InitializeDelegatedStakeQuorum(quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.InitializeDelegatedStakeQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake, _strategyParams)
}

// InitializeDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x75d4173a.
//
// Solidity: function initializeDelegatedStakeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) InitializeDelegatedStakeQuorum(quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.InitializeDelegatedStakeQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake, _strategyParams)
}

// InitializeSlashableStakeQuorum is a paid mutator transaction binding the contract method 0xcc5a7c20.
//
// Solidity: function initializeSlashableStakeQuorum(uint8 quorumNumber, uint96 minimumStake, uint32 lookAheadPeriod, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) InitializeSlashableStakeQuorum(opts *bind.TransactOpts, quorumNumber uint8, minimumStake *big.Int, lookAheadPeriod uint32, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "initializeSlashableStakeQuorum", quorumNumber, minimumStake, lookAheadPeriod, _strategyParams)
}

// InitializeSlashableStakeQuorum is a paid mutator transaction binding the contract method 0xcc5a7c20.
//
// Solidity: function initializeSlashableStakeQuorum(uint8 quorumNumber, uint96 minimumStake, uint32 lookAheadPeriod, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) InitializeSlashableStakeQuorum(quorumNumber uint8, minimumStake *big.Int, lookAheadPeriod uint32, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.InitializeSlashableStakeQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake, lookAheadPeriod, _strategyParams)
}

// InitializeSlashableStakeQuorum is a paid mutator transaction binding the contract method 0xcc5a7c20.
//
// Solidity: function initializeSlashableStakeQuorum(uint8 quorumNumber, uint96 minimumStake, uint32 lookAheadPeriod, (address,uint96)[] _strategyParams) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) InitializeSlashableStakeQuorum(quorumNumber uint8, minimumStake *big.Int, lookAheadPeriod uint32, _strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.InitializeSlashableStakeQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake, lookAheadPeriod, _strategyParams)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) ModifyStrategyParams(opts *bind.TransactOpts, quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "modifyStrategyParams", quorumNumber, strategyIndices, newMultipliers)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) ModifyStrategyParams(quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.ModifyStrategyParams(&_EOStakeRegistry.TransactOpts, quorumNumber, strategyIndices, newMultipliers)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) ModifyStrategyParams(quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.ModifyStrategyParams(&_EOStakeRegistry.TransactOpts, quorumNumber, strategyIndices, newMultipliers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_EOStakeRegistry *EOStakeRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "registerOperator", operator, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_EOStakeRegistry *EOStakeRegistrySession) RegisterOperator(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.RegisterOperator(&_EOStakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) RegisterOperator(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.RegisterOperator(&_EOStakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) RemoveStrategies(opts *bind.TransactOpts, quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "removeStrategies", quorumNumber, indicesToRemove)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) RemoveStrategies(quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.RemoveStrategies(&_EOStakeRegistry.TransactOpts, quorumNumber, indicesToRemove)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) RemoveStrategies(quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.RemoveStrategies(&_EOStakeRegistry.TransactOpts, quorumNumber, indicesToRemove)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) SetChainManager(opts *bind.TransactOpts, newChainManager common.Address) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "setChainManager", newChainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) SetChainManager(newChainManager common.Address) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetChainManager(&_EOStakeRegistry.TransactOpts, newChainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address newChainManager) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) SetChainManager(newChainManager common.Address) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetChainManager(&_EOStakeRegistry.TransactOpts, newChainManager)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) SetMinimumStakeForQuorum(opts *bind.TransactOpts, quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "setMinimumStakeForQuorum", quorumNumber, minimumStake)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) SetMinimumStakeForQuorum(quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetMinimumStakeForQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) SetMinimumStakeForQuorum(quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetMinimumStakeForQuorum(&_EOStakeRegistry.TransactOpts, quorumNumber, minimumStake)
}

// SetSlashableStakeLookahead is a paid mutator transaction binding the contract method 0xe086adb3.
//
// Solidity: function setSlashableStakeLookahead(uint8 quorumNumber, uint32 _lookAheadBlocks) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactor) SetSlashableStakeLookahead(opts *bind.TransactOpts, quorumNumber uint8, _lookAheadBlocks uint32) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "setSlashableStakeLookahead", quorumNumber, _lookAheadBlocks)
}

// SetSlashableStakeLookahead is a paid mutator transaction binding the contract method 0xe086adb3.
//
// Solidity: function setSlashableStakeLookahead(uint8 quorumNumber, uint32 _lookAheadBlocks) returns()
func (_EOStakeRegistry *EOStakeRegistrySession) SetSlashableStakeLookahead(quorumNumber uint8, _lookAheadBlocks uint32) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetSlashableStakeLookahead(&_EOStakeRegistry.TransactOpts, quorumNumber, _lookAheadBlocks)
}

// SetSlashableStakeLookahead is a paid mutator transaction binding the contract method 0xe086adb3.
//
// Solidity: function setSlashableStakeLookahead(uint8 quorumNumber, uint32 _lookAheadBlocks) returns()
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) SetSlashableStakeLookahead(quorumNumber uint8, _lookAheadBlocks uint32) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.SetSlashableStakeLookahead(&_EOStakeRegistry.TransactOpts, quorumNumber, _lookAheadBlocks)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_EOStakeRegistry *EOStakeRegistryTransactor) UpdateOperatorStake(opts *bind.TransactOpts, operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.contract.Transact(opts, "updateOperatorStake", operator, operatorId, quorumNumbers)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_EOStakeRegistry *EOStakeRegistrySession) UpdateOperatorStake(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.UpdateOperatorStake(&_EOStakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_EOStakeRegistry *EOStakeRegistryTransactorSession) UpdateOperatorStake(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _EOStakeRegistry.Contract.UpdateOperatorStake(&_EOStakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// EOStakeRegistryLookAheadPeriodChangedIterator is returned from FilterLookAheadPeriodChanged and is used to iterate over the raw logs and unpacked data for LookAheadPeriodChanged events raised by the EOStakeRegistry contract.
type EOStakeRegistryLookAheadPeriodChangedIterator struct {
	Event *EOStakeRegistryLookAheadPeriodChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryLookAheadPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryLookAheadPeriodChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryLookAheadPeriodChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryLookAheadPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryLookAheadPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryLookAheadPeriodChanged represents a LookAheadPeriodChanged event raised by the EOStakeRegistry contract.
type EOStakeRegistryLookAheadPeriodChanged struct {
	OldLookAheadBlocks uint32
	NewLookAheadBlocks uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLookAheadPeriodChanged is a free log retrieval operation binding the contract event 0x28d7358b79f02d21b8b7e17aefc4185a64308aa37406fa5befc05b91932c39c7.
//
// Solidity: event LookAheadPeriodChanged(uint32 oldLookAheadBlocks, uint32 newLookAheadBlocks)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterLookAheadPeriodChanged(opts *bind.FilterOpts) (*EOStakeRegistryLookAheadPeriodChangedIterator, error) {

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "LookAheadPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryLookAheadPeriodChangedIterator{contract: _EOStakeRegistry.contract, event: "LookAheadPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchLookAheadPeriodChanged is a free log subscription operation binding the contract event 0x28d7358b79f02d21b8b7e17aefc4185a64308aa37406fa5befc05b91932c39c7.
//
// Solidity: event LookAheadPeriodChanged(uint32 oldLookAheadBlocks, uint32 newLookAheadBlocks)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchLookAheadPeriodChanged(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryLookAheadPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "LookAheadPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryLookAheadPeriodChanged)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "LookAheadPeriodChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLookAheadPeriodChanged is a log parse operation binding the contract event 0x28d7358b79f02d21b8b7e17aefc4185a64308aa37406fa5befc05b91932c39c7.
//
// Solidity: event LookAheadPeriodChanged(uint32 oldLookAheadBlocks, uint32 newLookAheadBlocks)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseLookAheadPeriodChanged(log types.Log) (*EOStakeRegistryLookAheadPeriodChanged, error) {
	event := new(EOStakeRegistryLookAheadPeriodChanged)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "LookAheadPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryMinimumStakeForQuorumUpdatedIterator is returned from FilterMinimumStakeForQuorumUpdated and is used to iterate over the raw logs and unpacked data for MinimumStakeForQuorumUpdated events raised by the EOStakeRegistry contract.
type EOStakeRegistryMinimumStakeForQuorumUpdatedIterator struct {
	Event *EOStakeRegistryMinimumStakeForQuorumUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryMinimumStakeForQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryMinimumStakeForQuorumUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryMinimumStakeForQuorumUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryMinimumStakeForQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryMinimumStakeForQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryMinimumStakeForQuorumUpdated represents a MinimumStakeForQuorumUpdated event raised by the EOStakeRegistry contract.
type EOStakeRegistryMinimumStakeForQuorumUpdated struct {
	QuorumNumber uint8
	MinimumStake *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinimumStakeForQuorumUpdated is a free log retrieval operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterMinimumStakeForQuorumUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*EOStakeRegistryMinimumStakeForQuorumUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "MinimumStakeForQuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryMinimumStakeForQuorumUpdatedIterator{contract: _EOStakeRegistry.contract, event: "MinimumStakeForQuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumStakeForQuorumUpdated is a free log subscription operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchMinimumStakeForQuorumUpdated(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryMinimumStakeForQuorumUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "MinimumStakeForQuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryMinimumStakeForQuorumUpdated)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "MinimumStakeForQuorumUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinimumStakeForQuorumUpdated is a log parse operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseMinimumStakeForQuorumUpdated(log types.Log) (*EOStakeRegistryMinimumStakeForQuorumUpdated, error) {
	event := new(EOStakeRegistryMinimumStakeForQuorumUpdated)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "MinimumStakeForQuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryOperatorStakeUpdateIterator is returned from FilterOperatorStakeUpdate and is used to iterate over the raw logs and unpacked data for OperatorStakeUpdate events raised by the EOStakeRegistry contract.
type EOStakeRegistryOperatorStakeUpdateIterator struct {
	Event *EOStakeRegistryOperatorStakeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryOperatorStakeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryOperatorStakeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryOperatorStakeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryOperatorStakeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryOperatorStakeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryOperatorStakeUpdate represents a OperatorStakeUpdate event raised by the EOStakeRegistry contract.
type EOStakeRegistryOperatorStakeUpdate struct {
	OperatorId   [32]byte
	QuorumNumber uint8
	Stake        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorStakeUpdate is a free log retrieval operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterOperatorStakeUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*EOStakeRegistryOperatorStakeUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "OperatorStakeUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryOperatorStakeUpdateIterator{contract: _EOStakeRegistry.contract, event: "OperatorStakeUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorStakeUpdate is a free log subscription operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchOperatorStakeUpdate(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryOperatorStakeUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "OperatorStakeUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryOperatorStakeUpdate)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "OperatorStakeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorStakeUpdate is a log parse operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseOperatorStakeUpdate(log types.Log) (*EOStakeRegistryOperatorStakeUpdate, error) {
	event := new(EOStakeRegistryOperatorStakeUpdate)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "OperatorStakeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryQuorumCreatedIterator is returned from FilterQuorumCreated and is used to iterate over the raw logs and unpacked data for QuorumCreated events raised by the EOStakeRegistry contract.
type EOStakeRegistryQuorumCreatedIterator struct {
	Event *EOStakeRegistryQuorumCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryQuorumCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryQuorumCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryQuorumCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryQuorumCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryQuorumCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryQuorumCreated represents a QuorumCreated event raised by the EOStakeRegistry contract.
type EOStakeRegistryQuorumCreated struct {
	QuorumNumber uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumCreated is a free log retrieval operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterQuorumCreated(opts *bind.FilterOpts, quorumNumber []uint8) (*EOStakeRegistryQuorumCreatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryQuorumCreatedIterator{contract: _EOStakeRegistry.contract, event: "QuorumCreated", logs: logs, sub: sub}, nil
}

// WatchQuorumCreated is a free log subscription operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchQuorumCreated(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryQuorumCreated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryQuorumCreated)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumCreated is a log parse operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseQuorumCreated(log types.Log) (*EOStakeRegistryQuorumCreated, error) {
	event := new(EOStakeRegistryQuorumCreated)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryStakeTypeSetIterator is returned from FilterStakeTypeSet and is used to iterate over the raw logs and unpacked data for StakeTypeSet events raised by the EOStakeRegistry contract.
type EOStakeRegistryStakeTypeSetIterator struct {
	Event *EOStakeRegistryStakeTypeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryStakeTypeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryStakeTypeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryStakeTypeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryStakeTypeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryStakeTypeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryStakeTypeSet represents a StakeTypeSet event raised by the EOStakeRegistry contract.
type EOStakeRegistryStakeTypeSet struct {
	NewStakeType uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeTypeSet is a free log retrieval operation binding the contract event 0x7c112e863ccf007862e2c9e25819c933fedbc9350a6443423b4a8599c2e8a52d.
//
// Solidity: event StakeTypeSet(uint8 newStakeType)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterStakeTypeSet(opts *bind.FilterOpts) (*EOStakeRegistryStakeTypeSetIterator, error) {

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "StakeTypeSet")
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryStakeTypeSetIterator{contract: _EOStakeRegistry.contract, event: "StakeTypeSet", logs: logs, sub: sub}, nil
}

// WatchStakeTypeSet is a free log subscription operation binding the contract event 0x7c112e863ccf007862e2c9e25819c933fedbc9350a6443423b4a8599c2e8a52d.
//
// Solidity: event StakeTypeSet(uint8 newStakeType)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchStakeTypeSet(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryStakeTypeSet) (event.Subscription, error) {

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "StakeTypeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryStakeTypeSet)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "StakeTypeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeTypeSet is a log parse operation binding the contract event 0x7c112e863ccf007862e2c9e25819c933fedbc9350a6443423b4a8599c2e8a52d.
//
// Solidity: event StakeTypeSet(uint8 newStakeType)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseStakeTypeSet(log types.Log) (*EOStakeRegistryStakeTypeSet, error) {
	event := new(EOStakeRegistryStakeTypeSet)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "StakeTypeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryStrategyAddedToQuorumIterator is returned from FilterStrategyAddedToQuorum and is used to iterate over the raw logs and unpacked data for StrategyAddedToQuorum events raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyAddedToQuorumIterator struct {
	Event *EOStakeRegistryStrategyAddedToQuorum // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryStrategyAddedToQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryStrategyAddedToQuorum)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryStrategyAddedToQuorum)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryStrategyAddedToQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryStrategyAddedToQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryStrategyAddedToQuorum represents a StrategyAddedToQuorum event raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyAddedToQuorum struct {
	QuorumNumber uint8
	Strategy     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQuorum is a free log retrieval operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterStrategyAddedToQuorum(opts *bind.FilterOpts, quorumNumber []uint8) (*EOStakeRegistryStrategyAddedToQuorumIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "StrategyAddedToQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryStrategyAddedToQuorumIterator{contract: _EOStakeRegistry.contract, event: "StrategyAddedToQuorum", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQuorum is a free log subscription operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchStrategyAddedToQuorum(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryStrategyAddedToQuorum, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "StrategyAddedToQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryStrategyAddedToQuorum)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyAddedToQuorum", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAddedToQuorum is a log parse operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseStrategyAddedToQuorum(log types.Log) (*EOStakeRegistryStrategyAddedToQuorum, error) {
	event := new(EOStakeRegistryStrategyAddedToQuorum)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyAddedToQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryStrategyMultiplierUpdatedIterator is returned from FilterStrategyMultiplierUpdated and is used to iterate over the raw logs and unpacked data for StrategyMultiplierUpdated events raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyMultiplierUpdatedIterator struct {
	Event *EOStakeRegistryStrategyMultiplierUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryStrategyMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryStrategyMultiplierUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryStrategyMultiplierUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryStrategyMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryStrategyMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryStrategyMultiplierUpdated represents a StrategyMultiplierUpdated event raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyMultiplierUpdated struct {
	QuorumNumber uint8
	Strategy     common.Address
	Multiplier   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyMultiplierUpdated is a free log retrieval operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterStrategyMultiplierUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*EOStakeRegistryStrategyMultiplierUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "StrategyMultiplierUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryStrategyMultiplierUpdatedIterator{contract: _EOStakeRegistry.contract, event: "StrategyMultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyMultiplierUpdated is a free log subscription operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchStrategyMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryStrategyMultiplierUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "StrategyMultiplierUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryStrategyMultiplierUpdated)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyMultiplierUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyMultiplierUpdated is a log parse operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseStrategyMultiplierUpdated(log types.Log) (*EOStakeRegistryStrategyMultiplierUpdated, error) {
	event := new(EOStakeRegistryStrategyMultiplierUpdated)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyMultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOStakeRegistryStrategyRemovedFromQuorumIterator is returned from FilterStrategyRemovedFromQuorum and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQuorum events raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyRemovedFromQuorumIterator struct {
	Event *EOStakeRegistryStrategyRemovedFromQuorum // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EOStakeRegistryStrategyRemovedFromQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOStakeRegistryStrategyRemovedFromQuorum)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EOStakeRegistryStrategyRemovedFromQuorum)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EOStakeRegistryStrategyRemovedFromQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOStakeRegistryStrategyRemovedFromQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOStakeRegistryStrategyRemovedFromQuorum represents a StrategyRemovedFromQuorum event raised by the EOStakeRegistry contract.
type EOStakeRegistryStrategyRemovedFromQuorum struct {
	QuorumNumber uint8
	Strategy     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQuorum is a free log retrieval operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) FilterStrategyRemovedFromQuorum(opts *bind.FilterOpts, quorumNumber []uint8) (*EOStakeRegistryStrategyRemovedFromQuorumIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.FilterLogs(opts, "StrategyRemovedFromQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &EOStakeRegistryStrategyRemovedFromQuorumIterator{contract: _EOStakeRegistry.contract, event: "StrategyRemovedFromQuorum", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQuorum is a free log subscription operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) WatchStrategyRemovedFromQuorum(opts *bind.WatchOpts, sink chan<- *EOStakeRegistryStrategyRemovedFromQuorum, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _EOStakeRegistry.contract.WatchLogs(opts, "StrategyRemovedFromQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOStakeRegistryStrategyRemovedFromQuorum)
				if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyRemovedFromQuorum", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRemovedFromQuorum is a log parse operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_EOStakeRegistry *EOStakeRegistryFilterer) ParseStrategyRemovedFromQuorum(log types.Log) (*EOStakeRegistryStrategyRemovedFromQuorum, error) {
	event := new(EOStakeRegistryStrategyRemovedFromQuorum)
	if err := _EOStakeRegistry.contract.UnpackLog(event, "StrategyRemovedFromQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
