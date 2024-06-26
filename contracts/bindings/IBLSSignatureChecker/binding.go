// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIBLSSignatureChecker

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ContractIBLSSignatureCheckerMetaData contains all meta data concerning the ContractIBLSSignatureChecker contract.
var ContractIBLSSignatureCheckerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEORegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// ContractIBLSSignatureCheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIBLSSignatureCheckerMetaData.ABI instead.
var ContractIBLSSignatureCheckerABI = ContractIBLSSignatureCheckerMetaData.ABI

// ContractIBLSSignatureCheckerMethods is an auto generated interface around an Ethereum contract.
type ContractIBLSSignatureCheckerMethods interface {
	ContractIBLSSignatureCheckerCalls
	ContractIBLSSignatureCheckerTransacts
	ContractIBLSSignatureCheckerFilters
}

// ContractIBLSSignatureCheckerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIBLSSignatureCheckerCalls interface {
	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractIBLSSignatureCheckerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIBLSSignatureCheckerTransacts interface {
}

// ContractIBLSSignatureCheckerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIBLSSignatureCheckerFilters interface {
	FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator, error)
	WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate) (event.Subscription, error)
	ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate, error)
}

// ContractIBLSSignatureChecker is an auto generated Go binding around an Ethereum contract.
type ContractIBLSSignatureChecker struct {
	ContractIBLSSignatureCheckerCaller     // Read-only binding to the contract
	ContractIBLSSignatureCheckerTransactor // Write-only binding to the contract
	ContractIBLSSignatureCheckerFilterer   // Log filterer for contract events
}

// ContractIBLSSignatureChecker implements the ContractIBLSSignatureCheckerMethods interface.
var _ ContractIBLSSignatureCheckerMethods = (*ContractIBLSSignatureChecker)(nil)

// ContractIBLSSignatureCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIBLSSignatureCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSSignatureCheckerCaller implements the ContractIBLSSignatureCheckerCalls interface.
var _ ContractIBLSSignatureCheckerCalls = (*ContractIBLSSignatureCheckerCaller)(nil)

// ContractIBLSSignatureCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIBLSSignatureCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSSignatureCheckerTransactor implements the ContractIBLSSignatureCheckerTransacts interface.
var _ ContractIBLSSignatureCheckerTransacts = (*ContractIBLSSignatureCheckerTransactor)(nil)

// ContractIBLSSignatureCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIBLSSignatureCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSSignatureCheckerFilterer implements the ContractIBLSSignatureCheckerFilters interface.
var _ ContractIBLSSignatureCheckerFilters = (*ContractIBLSSignatureCheckerFilterer)(nil)

// ContractIBLSSignatureCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIBLSSignatureCheckerSession struct {
	Contract     *ContractIBLSSignatureChecker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractIBLSSignatureCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIBLSSignatureCheckerCallerSession struct {
	Contract *ContractIBLSSignatureCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ContractIBLSSignatureCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIBLSSignatureCheckerTransactorSession struct {
	Contract     *ContractIBLSSignatureCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ContractIBLSSignatureCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIBLSSignatureCheckerRaw struct {
	Contract *ContractIBLSSignatureChecker // Generic contract binding to access the raw methods on
}

// ContractIBLSSignatureCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIBLSSignatureCheckerCallerRaw struct {
	Contract *ContractIBLSSignatureCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIBLSSignatureCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIBLSSignatureCheckerTransactorRaw struct {
	Contract *ContractIBLSSignatureCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIBLSSignatureChecker creates a new instance of ContractIBLSSignatureChecker, bound to a specific deployed contract.
func NewContractIBLSSignatureChecker(address common.Address, backend bind.ContractBackend) (*ContractIBLSSignatureChecker, error) {
	contract, err := bindContractIBLSSignatureChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSSignatureChecker{ContractIBLSSignatureCheckerCaller: ContractIBLSSignatureCheckerCaller{contract: contract}, ContractIBLSSignatureCheckerTransactor: ContractIBLSSignatureCheckerTransactor{contract: contract}, ContractIBLSSignatureCheckerFilterer: ContractIBLSSignatureCheckerFilterer{contract: contract}}, nil
}

// NewContractIBLSSignatureCheckerCaller creates a new read-only instance of ContractIBLSSignatureChecker, bound to a specific deployed contract.
func NewContractIBLSSignatureCheckerCaller(address common.Address, caller bind.ContractCaller) (*ContractIBLSSignatureCheckerCaller, error) {
	contract, err := bindContractIBLSSignatureChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSSignatureCheckerCaller{contract: contract}, nil
}

// NewContractIBLSSignatureCheckerTransactor creates a new write-only instance of ContractIBLSSignatureChecker, bound to a specific deployed contract.
func NewContractIBLSSignatureCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIBLSSignatureCheckerTransactor, error) {
	contract, err := bindContractIBLSSignatureChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSSignatureCheckerTransactor{contract: contract}, nil
}

// NewContractIBLSSignatureCheckerFilterer creates a new log filterer instance of ContractIBLSSignatureChecker, bound to a specific deployed contract.
func NewContractIBLSSignatureCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIBLSSignatureCheckerFilterer, error) {
	contract, err := bindContractIBLSSignatureChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSSignatureCheckerFilterer{contract: contract}, nil
}

// bindContractIBLSSignatureChecker binds a generic wrapper to an already deployed contract.
func bindContractIBLSSignatureChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIBLSSignatureCheckerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIBLSSignatureChecker.Contract.ContractIBLSSignatureCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIBLSSignatureChecker.Contract.ContractIBLSSignatureCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIBLSSignatureChecker.Contract.ContractIBLSSignatureCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIBLSSignatureChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIBLSSignatureChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIBLSSignatureChecker.Contract.contract.Transact(opts, method, params...)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSSignatureChecker.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.BlsApkRegistry(&_ContractIBLSSignatureChecker.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.BlsApkRegistry(&_ContractIBLSSignatureChecker.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIBLSSignatureChecker.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIBLSSignatureChecker.Contract.CheckSignatures(&_ContractIBLSSignatureChecker.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns((uint96[],uint96[]), bytes32)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIBLSSignatureChecker.Contract.CheckSignatures(&_ContractIBLSSignatureChecker.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSSignatureChecker.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerSession) Delegation() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.Delegation(&_ContractIBLSSignatureChecker.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerSession) Delegation() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.Delegation(&_ContractIBLSSignatureChecker.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSSignatureChecker.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.RegistryCoordinator(&_ContractIBLSSignatureChecker.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.RegistryCoordinator(&_ContractIBLSSignatureChecker.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSSignatureChecker.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerSession) StakeRegistry() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.StakeRegistry(&_ContractIBLSSignatureChecker.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIBLSSignatureChecker.Contract.StakeRegistry(&_ContractIBLSSignatureChecker.CallOpts)
}

// ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIBLSSignatureChecker contract.
type ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate)
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
func (it *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIBLSSignatureChecker contract.
type ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIBLSSignatureChecker.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIBLSSignatureCheckerStaleStakesForbiddenUpdateIterator{contract: _ContractIBLSSignatureChecker.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIBLSSignatureChecker.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate)
				if err := _ContractIBLSSignatureChecker.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIBLSSignatureChecker *ContractIBLSSignatureCheckerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIBLSSignatureCheckerStaleStakesForbiddenUpdate)
	if err := _ContractIBLSSignatureChecker.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
