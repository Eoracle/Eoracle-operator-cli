// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIBLSApkRegistry

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

// IBLSApkRegistryTypesApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryTypesPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IBLSApkRegistryMetaData contains all meta data concerning the IBLSApkRegistry contract.
var IBLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPubkeyG2\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrRegisterOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"NewG2PubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BLSPubkeyAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberBeforeFirstUpdate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberNotLatest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberTooRecent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"G2PubkeyAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignatureOrPrivateKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPubKey\",\"inputs\":[]}]",
}

// IBLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IBLSApkRegistryMetaData.ABI instead.
var IBLSApkRegistryABI = IBLSApkRegistryMetaData.ABI

// IBLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type IBLSApkRegistry struct {
	IBLSApkRegistryCaller     // Read-only binding to the contract
	IBLSApkRegistryTransactor // Write-only binding to the contract
	IBLSApkRegistryFilterer   // Log filterer for contract events
}

// IBLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBLSApkRegistrySession struct {
	Contract     *IBLSApkRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBLSApkRegistryCallerSession struct {
	Contract *IBLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IBLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBLSApkRegistryTransactorSession struct {
	Contract     *IBLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IBLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBLSApkRegistryRaw struct {
	Contract *IBLSApkRegistry // Generic contract binding to access the raw methods on
}

// IBLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBLSApkRegistryCallerRaw struct {
	Contract *IBLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IBLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBLSApkRegistryTransactorRaw struct {
	Contract *IBLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBLSApkRegistry creates a new instance of IBLSApkRegistry, bound to a specific deployed contract.
func NewIBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*IBLSApkRegistry, error) {
	contract, err := bindIBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistry{IBLSApkRegistryCaller: IBLSApkRegistryCaller{contract: contract}, IBLSApkRegistryTransactor: IBLSApkRegistryTransactor{contract: contract}, IBLSApkRegistryFilterer: IBLSApkRegistryFilterer{contract: contract}}, nil
}

// NewIBLSApkRegistryCaller creates a new read-only instance of IBLSApkRegistry, bound to a specific deployed contract.
func NewIBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*IBLSApkRegistryCaller, error) {
	contract, err := bindIBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryCaller{contract: contract}, nil
}

// NewIBLSApkRegistryTransactor creates a new write-only instance of IBLSApkRegistry, bound to a specific deployed contract.
func NewIBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IBLSApkRegistryTransactor, error) {
	contract, err := bindIBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryTransactor{contract: contract}, nil
}

// NewIBLSApkRegistryFilterer creates a new log filterer instance of IBLSApkRegistry, bound to a specific deployed contract.
func NewIBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IBLSApkRegistryFilterer, error) {
	contract, err := bindIBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryFilterer{contract: contract}, nil
}

// bindIBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindIBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLSApkRegistry *IBLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLSApkRegistry.Contract.IBLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLSApkRegistry *IBLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.IBLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLSApkRegistry *IBLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.IBLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLSApkRegistry *IBLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLSApkRegistry *IBLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLSApkRegistry *IBLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 index) view returns(bytes24, uint32, uint32)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) ([24]byte, uint32, uint32, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "apkHistory", quorumNumber, index)

	if err != nil {
		return *new([24]byte), *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return out0, out1, out2, err

}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 index) view returns(bytes24, uint32, uint32)
func (_IBLSApkRegistry *IBLSApkRegistrySession) ApkHistory(quorumNumber uint8, index *big.Int) ([24]byte, uint32, uint32, error) {
	return _IBLSApkRegistry.Contract.ApkHistory(&_IBLSApkRegistry.CallOpts, quorumNumber, index)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 index) view returns(bytes24, uint32, uint32)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) ApkHistory(quorumNumber uint8, index *big.Int) ([24]byte, uint32, uint32, error) {
	return _IBLSApkRegistry.Contract.ApkHistory(&_IBLSApkRegistry.CallOpts, quorumNumber, index)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "currentApk", quorumNumber)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistrySession) CurrentApk(quorumNumber uint8) (*big.Int, *big.Int, error) {
	return _IBLSApkRegistry.Contract.CurrentApk(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) CurrentApk(quorumNumber uint8) (*big.Int, *big.Int, error) {
	return _IBLSApkRegistry.Contract.CurrentApk(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _IBLSApkRegistry.Contract.GetApk(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _IBLSApkRegistry.Contract.GetApk(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _IBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_IBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _IBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_IBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _IBLSApkRegistry.Contract.GetApkHistoryLength(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _IBLSApkRegistry.Contract.GetApkHistoryLength(&_IBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _IBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_IBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _IBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_IBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryTypesApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryTypesApkUpdate)).(*IBLSApkRegistryTypesApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _IBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_IBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _IBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_IBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _IBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_IBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _IBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_IBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _IBLSApkRegistry.Contract.GetOperatorId(&_IBLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _IBLSApkRegistry.Contract.GetOperatorId(&_IBLSApkRegistry.CallOpts, operator)
}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetOperatorPubkeyG2(opts *bind.CallOpts, operator common.Address) (BN254G2Point, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getOperatorPubkeyG2", operator)

	if err != nil {
		return *new(BN254G2Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G2Point)).(*BN254G2Point)

	return out0, err

}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetOperatorPubkeyG2(operator common.Address) (BN254G2Point, error) {
	return _IBLSApkRegistry.Contract.GetOperatorPubkeyG2(&_IBLSApkRegistry.CallOpts, operator)
}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetOperatorPubkeyG2(operator common.Address) (BN254G2Point, error) {
	return _IBLSApkRegistry.Contract.GetOperatorPubkeyG2(&_IBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _IBLSApkRegistry.Contract.GetRegisteredPubkey(&_IBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _IBLSApkRegistry.Contract.GetRegisteredPubkey(&_IBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, operator common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", operator)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistrySession) OperatorToPubkey(operator common.Address) (*big.Int, *big.Int, error) {
	return _IBLSApkRegistry.Contract.OperatorToPubkey(&_IBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256, uint256)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) OperatorToPubkey(operator common.Address) (*big.Int, *big.Int, error) {
	return _IBLSApkRegistry.Contract.OperatorToPubkey(&_IBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistrySession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _IBLSApkRegistry.Contract.OperatorToPubkeyHash(&_IBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _IBLSApkRegistry.Contract.OperatorToPubkeyHash(&_IBLSApkRegistry.CallOpts, operator)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistrySession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _IBLSApkRegistry.Contract.PubkeyHashToOperator(&_IBLSApkRegistry.CallOpts, pubkeyHash)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _IBLSApkRegistry.Contract.PubkeyHashToOperator(&_IBLSApkRegistry.CallOpts, pubkeyHash)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_IBLSApkRegistry *IBLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_IBLSApkRegistry *IBLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _IBLSApkRegistry.Contract.RegistryCoordinator(&_IBLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_IBLSApkRegistry *IBLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _IBLSApkRegistry.Contract.RegistryCoordinator(&_IBLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.DeregisterOperator(&_IBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.DeregisterOperator(&_IBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// GetOrRegisterOperatorId is a paid mutator transaction binding the contract method 0x03c5a6b6.
//
// Solidity: function getOrRegisterOperatorId(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryTransactor) GetOrRegisterOperatorId(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.contract.Transact(opts, "getOrRegisterOperatorId", operator, params, pubkeyRegistrationMessageHash)
}

// GetOrRegisterOperatorId is a paid mutator transaction binding the contract method 0x03c5a6b6.
//
// Solidity: function getOrRegisterOperatorId(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistrySession) GetOrRegisterOperatorId(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.GetOrRegisterOperatorId(&_IBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// GetOrRegisterOperatorId is a paid mutator transaction binding the contract method 0x03c5a6b6.
//
// Solidity: function getOrRegisterOperatorId(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryTransactorSession) GetOrRegisterOperatorId(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.GetOrRegisterOperatorId(&_IBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _IBLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_IBLSApkRegistry *IBLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.InitializeQuorum(&_IBLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.InitializeQuorum(&_IBLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.RegisterBLSPublicKey(&_IBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_IBLSApkRegistry *IBLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.RegisterBLSPublicKey(&_IBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.RegisterOperator(&_IBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_IBLSApkRegistry *IBLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _IBLSApkRegistry.Contract.RegisterOperator(&_IBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// IBLSApkRegistryNewG2PubkeyRegistrationIterator is returned from FilterNewG2PubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewG2PubkeyRegistration events raised by the IBLSApkRegistry contract.
type IBLSApkRegistryNewG2PubkeyRegistrationIterator struct {
	Event *IBLSApkRegistryNewG2PubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *IBLSApkRegistryNewG2PubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBLSApkRegistryNewG2PubkeyRegistration)
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
		it.Event = new(IBLSApkRegistryNewG2PubkeyRegistration)
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
func (it *IBLSApkRegistryNewG2PubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBLSApkRegistryNewG2PubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBLSApkRegistryNewG2PubkeyRegistration represents a NewG2PubkeyRegistration event raised by the IBLSApkRegistry contract.
type IBLSApkRegistryNewG2PubkeyRegistration struct {
	Operator common.Address
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewG2PubkeyRegistration is a free log retrieval operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) FilterNewG2PubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*IBLSApkRegistryNewG2PubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBLSApkRegistry.contract.FilterLogs(opts, "NewG2PubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryNewG2PubkeyRegistrationIterator{contract: _IBLSApkRegistry.contract, event: "NewG2PubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewG2PubkeyRegistration is a free log subscription operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) WatchNewG2PubkeyRegistration(opts *bind.WatchOpts, sink chan<- *IBLSApkRegistryNewG2PubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBLSApkRegistry.contract.WatchLogs(opts, "NewG2PubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBLSApkRegistryNewG2PubkeyRegistration)
				if err := _IBLSApkRegistry.contract.UnpackLog(event, "NewG2PubkeyRegistration", log); err != nil {
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

// ParseNewG2PubkeyRegistration is a log parse operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) ParseNewG2PubkeyRegistration(log types.Log) (*IBLSApkRegistryNewG2PubkeyRegistration, error) {
	event := new(IBLSApkRegistryNewG2PubkeyRegistration)
	if err := _IBLSApkRegistry.contract.UnpackLog(event, "NewG2PubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the IBLSApkRegistry contract.
type IBLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *IBLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *IBLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(IBLSApkRegistryNewPubkeyRegistration)
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
func (it *IBLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the IBLSApkRegistry contract.
type IBLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*IBLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryNewPubkeyRegistrationIterator{contract: _IBLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *IBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBLSApkRegistryNewPubkeyRegistration)
				if err := _IBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*IBLSApkRegistryNewPubkeyRegistration, error) {
	event := new(IBLSApkRegistryNewPubkeyRegistration)
	if err := _IBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the IBLSApkRegistry contract.
type IBLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *IBLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *IBLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(IBLSApkRegistryOperatorAddedToQuorums)
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
func (it *IBLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the IBLSApkRegistry contract.
type IBLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*IBLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _IBLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryOperatorAddedToQuorumsIterator{contract: _IBLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *IBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _IBLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBLSApkRegistryOperatorAddedToQuorums)
				if err := _IBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*IBLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(IBLSApkRegistryOperatorAddedToQuorums)
	if err := _IBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the IBLSApkRegistry contract.
type IBLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *IBLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *IBLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(IBLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *IBLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the IBLSApkRegistry contract.
type IBLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*IBLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _IBLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &IBLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _IBLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *IBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _IBLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBLSApkRegistryOperatorRemovedFromQuorums)
				if err := _IBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_IBLSApkRegistry *IBLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*IBLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(IBLSApkRegistryOperatorRemovedFromQuorums)
	if err := _IBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
