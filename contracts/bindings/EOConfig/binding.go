// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eoconfig

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

// EOConfigMetaData contains all meta data concerning the EOConfig contract.
var EOConfigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addSource\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addSources\",\"inputs\":[{\"name\":\"names\",\"type\":\"bytes24[]\",\"internalType\":\"bytes24[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addSymbol\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addSymbols\",\"inputs\":[{\"name\":\"names\",\"type\":\"bytes8[]\",\"internalType\":\"bytes8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEOAggregator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aliasToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignAlias\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorAlias\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeAlias\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newAlias\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"declareAlias\",\"inputs\":[{\"name\":\"operatorAlias\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evmSourceConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceSymbolId\",\"inputs\":[{\"name\":\"source\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"symbol\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceSymbols\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSymbolSources\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isIdAllowed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"notify\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onStateReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToAlias\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToRegistrationData\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"packPair\",\"inputs\":[{\"name\":\"source\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"symbol\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"removeSource\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSymbol\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainManager\",\"inputs\":[{\"name\":\"_chainManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEvmSourceConfig\",\"inputs\":[{\"name\":\"_evmSourceConfig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSourceSymbolIds\",\"inputs\":[{\"name\":\"source\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"symbols\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"names\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sourceById\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sourceId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sourceIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sourceIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sourceSymbolOverrides\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subscribe\",\"inputs\":[{\"name\":\"listener\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbolById\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"id\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbolId\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbolIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbolIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unsubscribe\",\"inputs\":[{\"name\":\"listener\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSource\",\"inputs\":[{\"name\":\"_sourceId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"_symbols\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSymbol\",\"inputs\":[{\"name\":\"_symbolId\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"_sources\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnConfigChange\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSourceAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes24\",\"indexed\":false,\"internalType\":\"bytes24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSourceRemoved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSourceUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes24\",\"indexed\":false,\"internalType\":\"bytes24\"},{\"name\":\"symbols\",\"type\":\"uint16[]\",\"indexed\":false,\"internalType\":\"uint16[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSymbolAdded\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes8\",\"indexed\":false,\"internalType\":\"bytes8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSymbolRemoved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSymbolUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes8\",\"indexed\":false,\"internalType\":\"bytes8\"},{\"name\":\"sources\",\"type\":\"uint16[]\",\"indexed\":false,\"internalType\":\"uint16[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// EOConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use EOConfigMetaData.ABI instead.
var EOConfigABI = EOConfigMetaData.ABI

// EOConfig is an auto generated Go binding around an Ethereum contract.
type EOConfig struct {
	EOConfigCaller     // Read-only binding to the contract
	EOConfigTransactor // Write-only binding to the contract
	EOConfigFilterer   // Log filterer for contract events
}

// EOConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type EOConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EOConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EOConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EOConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EOConfigSession struct {
	Contract     *EOConfig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EOConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EOConfigCallerSession struct {
	Contract *EOConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EOConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EOConfigTransactorSession struct {
	Contract     *EOConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EOConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type EOConfigRaw struct {
	Contract *EOConfig // Generic contract binding to access the raw methods on
}

// EOConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EOConfigCallerRaw struct {
	Contract *EOConfigCaller // Generic read-only contract binding to access the raw methods on
}

// EOConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EOConfigTransactorRaw struct {
	Contract *EOConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEOConfig creates a new instance of EOConfig, bound to a specific deployed contract.
func NewEOConfig(address common.Address, backend bind.ContractBackend) (*EOConfig, error) {
	contract, err := bindEOConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EOConfig{EOConfigCaller: EOConfigCaller{contract: contract}, EOConfigTransactor: EOConfigTransactor{contract: contract}, EOConfigFilterer: EOConfigFilterer{contract: contract}}, nil
}

// NewEOConfigCaller creates a new read-only instance of EOConfig, bound to a specific deployed contract.
func NewEOConfigCaller(address common.Address, caller bind.ContractCaller) (*EOConfigCaller, error) {
	contract, err := bindEOConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EOConfigCaller{contract: contract}, nil
}

// NewEOConfigTransactor creates a new write-only instance of EOConfig, bound to a specific deployed contract.
func NewEOConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*EOConfigTransactor, error) {
	contract, err := bindEOConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EOConfigTransactor{contract: contract}, nil
}

// NewEOConfigFilterer creates a new log filterer instance of EOConfig, bound to a specific deployed contract.
func NewEOConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*EOConfigFilterer, error) {
	contract, err := bindEOConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EOConfigFilterer{contract: contract}, nil
}

// bindEOConfig binds a generic wrapper to an already deployed contract.
func bindEOConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EOConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EOConfig *EOConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOConfig.Contract.EOConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EOConfig *EOConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.Contract.EOConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EOConfig *EOConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOConfig.Contract.EOConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EOConfig *EOConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EOConfig *EOConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EOConfig *EOConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOConfig.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EOConfig *EOConfigCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EOConfig *EOConfigSession) Aggregator() (common.Address, error) {
	return _EOConfig.Contract.Aggregator(&_EOConfig.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EOConfig *EOConfigCallerSession) Aggregator() (common.Address, error) {
	return _EOConfig.Contract.Aggregator(&_EOConfig.CallOpts)
}

// AliasToOperator is a free data retrieval call binding the contract method 0xbfa65472.
//
// Solidity: function aliasToOperator(address ) view returns(address)
func (_EOConfig *EOConfigCaller) AliasToOperator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "aliasToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AliasToOperator is a free data retrieval call binding the contract method 0xbfa65472.
//
// Solidity: function aliasToOperator(address ) view returns(address)
func (_EOConfig *EOConfigSession) AliasToOperator(arg0 common.Address) (common.Address, error) {
	return _EOConfig.Contract.AliasToOperator(&_EOConfig.CallOpts, arg0)
}

// AliasToOperator is a free data retrieval call binding the contract method 0xbfa65472.
//
// Solidity: function aliasToOperator(address ) view returns(address)
func (_EOConfig *EOConfigCallerSession) AliasToOperator(arg0 common.Address) (common.Address, error) {
	return _EOConfig.Contract.AliasToOperator(&_EOConfig.CallOpts, arg0)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOConfig *EOConfigCaller) ChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "chainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOConfig *EOConfigSession) ChainManager() (common.Address, error) {
	return _EOConfig.Contract.ChainManager(&_EOConfig.CallOpts)
}

// ChainManager is a free data retrieval call binding the contract method 0x5d824812.
//
// Solidity: function chainManager() view returns(address)
func (_EOConfig *EOConfigCallerSession) ChainManager() (common.Address, error) {
	return _EOConfig.Contract.ChainManager(&_EOConfig.CallOpts)
}

// EvmSourceConfig is a free data retrieval call binding the contract method 0x980690cf.
//
// Solidity: function evmSourceConfig() view returns(address)
func (_EOConfig *EOConfigCaller) EvmSourceConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "evmSourceConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvmSourceConfig is a free data retrieval call binding the contract method 0x980690cf.
//
// Solidity: function evmSourceConfig() view returns(address)
func (_EOConfig *EOConfigSession) EvmSourceConfig() (common.Address, error) {
	return _EOConfig.Contract.EvmSourceConfig(&_EOConfig.CallOpts)
}

// EvmSourceConfig is a free data retrieval call binding the contract method 0x980690cf.
//
// Solidity: function evmSourceConfig() view returns(address)
func (_EOConfig *EOConfigCallerSession) EvmSourceConfig() (common.Address, error) {
	return _EOConfig.Contract.EvmSourceConfig(&_EOConfig.CallOpts)
}

// GetSourceSymbolId is a free data retrieval call binding the contract method 0xce2fc469.
//
// Solidity: function getSourceSymbolId(uint16 source, uint16 symbol) view returns(string)
func (_EOConfig *EOConfigCaller) GetSourceSymbolId(opts *bind.CallOpts, source uint16, symbol uint16) (string, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "getSourceSymbolId", source, symbol)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSourceSymbolId is a free data retrieval call binding the contract method 0xce2fc469.
//
// Solidity: function getSourceSymbolId(uint16 source, uint16 symbol) view returns(string)
func (_EOConfig *EOConfigSession) GetSourceSymbolId(source uint16, symbol uint16) (string, error) {
	return _EOConfig.Contract.GetSourceSymbolId(&_EOConfig.CallOpts, source, symbol)
}

// GetSourceSymbolId is a free data retrieval call binding the contract method 0xce2fc469.
//
// Solidity: function getSourceSymbolId(uint16 source, uint16 symbol) view returns(string)
func (_EOConfig *EOConfigCallerSession) GetSourceSymbolId(source uint16, symbol uint16) (string, error) {
	return _EOConfig.Contract.GetSourceSymbolId(&_EOConfig.CallOpts, source, symbol)
}

// GetSourceSymbols is a free data retrieval call binding the contract method 0x3ae23dfb.
//
// Solidity: function getSourceSymbols(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigCaller) GetSourceSymbols(opts *bind.CallOpts, id uint16) ([]uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "getSourceSymbols", id)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetSourceSymbols is a free data retrieval call binding the contract method 0x3ae23dfb.
//
// Solidity: function getSourceSymbols(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigSession) GetSourceSymbols(id uint16) ([]uint16, error) {
	return _EOConfig.Contract.GetSourceSymbols(&_EOConfig.CallOpts, id)
}

// GetSourceSymbols is a free data retrieval call binding the contract method 0x3ae23dfb.
//
// Solidity: function getSourceSymbols(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigCallerSession) GetSourceSymbols(id uint16) ([]uint16, error) {
	return _EOConfig.Contract.GetSourceSymbols(&_EOConfig.CallOpts, id)
}

// GetSymbolSources is a free data retrieval call binding the contract method 0x14832bc8.
//
// Solidity: function getSymbolSources(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigCaller) GetSymbolSources(opts *bind.CallOpts, id uint16) ([]uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "getSymbolSources", id)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetSymbolSources is a free data retrieval call binding the contract method 0x14832bc8.
//
// Solidity: function getSymbolSources(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigSession) GetSymbolSources(id uint16) ([]uint16, error) {
	return _EOConfig.Contract.GetSymbolSources(&_EOConfig.CallOpts, id)
}

// GetSymbolSources is a free data retrieval call binding the contract method 0x14832bc8.
//
// Solidity: function getSymbolSources(uint16 id) view returns(uint16[])
func (_EOConfig *EOConfigCallerSession) GetSymbolSources(id uint16) ([]uint16, error) {
	return _EOConfig.Contract.GetSymbolSources(&_EOConfig.CallOpts, id)
}

// IsIdAllowed is a free data retrieval call binding the contract method 0x04495eb1.
//
// Solidity: function isIdAllowed(bytes4 ) view returns(bool)
func (_EOConfig *EOConfigCaller) IsIdAllowed(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "isIdAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdAllowed is a free data retrieval call binding the contract method 0x04495eb1.
//
// Solidity: function isIdAllowed(bytes4 ) view returns(bool)
func (_EOConfig *EOConfigSession) IsIdAllowed(arg0 [4]byte) (bool, error) {
	return _EOConfig.Contract.IsIdAllowed(&_EOConfig.CallOpts, arg0)
}

// IsIdAllowed is a free data retrieval call binding the contract method 0x04495eb1.
//
// Solidity: function isIdAllowed(bytes4 ) view returns(bool)
func (_EOConfig *EOConfigCallerSession) IsIdAllowed(arg0 [4]byte) (bool, error) {
	return _EOConfig.Contract.IsIdAllowed(&_EOConfig.CallOpts, arg0)
}

// OperatorToAlias is a free data retrieval call binding the contract method 0x92c94ab5.
//
// Solidity: function operatorToAlias(address ) view returns(address)
func (_EOConfig *EOConfigCaller) OperatorToAlias(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "operatorToAlias", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorToAlias is a free data retrieval call binding the contract method 0x92c94ab5.
//
// Solidity: function operatorToAlias(address ) view returns(address)
func (_EOConfig *EOConfigSession) OperatorToAlias(arg0 common.Address) (common.Address, error) {
	return _EOConfig.Contract.OperatorToAlias(&_EOConfig.CallOpts, arg0)
}

// OperatorToAlias is a free data retrieval call binding the contract method 0x92c94ab5.
//
// Solidity: function operatorToAlias(address ) view returns(address)
func (_EOConfig *EOConfigCallerSession) OperatorToAlias(arg0 common.Address) (common.Address, error) {
	return _EOConfig.Contract.OperatorToAlias(&_EOConfig.CallOpts, arg0)
}

// OperatorToRegistrationData is a free data retrieval call binding the contract method 0x4a9910fa.
//
// Solidity: function operatorToRegistrationData(address ) view returns(bytes)
func (_EOConfig *EOConfigCaller) OperatorToRegistrationData(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "operatorToRegistrationData", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// OperatorToRegistrationData is a free data retrieval call binding the contract method 0x4a9910fa.
//
// Solidity: function operatorToRegistrationData(address ) view returns(bytes)
func (_EOConfig *EOConfigSession) OperatorToRegistrationData(arg0 common.Address) ([]byte, error) {
	return _EOConfig.Contract.OperatorToRegistrationData(&_EOConfig.CallOpts, arg0)
}

// OperatorToRegistrationData is a free data retrieval call binding the contract method 0x4a9910fa.
//
// Solidity: function operatorToRegistrationData(address ) view returns(bytes)
func (_EOConfig *EOConfigCallerSession) OperatorToRegistrationData(arg0 common.Address) ([]byte, error) {
	return _EOConfig.Contract.OperatorToRegistrationData(&_EOConfig.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EOConfig *EOConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EOConfig *EOConfigSession) Owner() (common.Address, error) {
	return _EOConfig.Contract.Owner(&_EOConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EOConfig *EOConfigCallerSession) Owner() (common.Address, error) {
	return _EOConfig.Contract.Owner(&_EOConfig.CallOpts)
}

// PackPair is a free data retrieval call binding the contract method 0xc18cb7b0.
//
// Solidity: function packPair(uint16 source, uint16 symbol) pure returns(bytes4 id)
func (_EOConfig *EOConfigCaller) PackPair(opts *bind.CallOpts, source uint16, symbol uint16) ([4]byte, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "packPair", source, symbol)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// PackPair is a free data retrieval call binding the contract method 0xc18cb7b0.
//
// Solidity: function packPair(uint16 source, uint16 symbol) pure returns(bytes4 id)
func (_EOConfig *EOConfigSession) PackPair(source uint16, symbol uint16) ([4]byte, error) {
	return _EOConfig.Contract.PackPair(&_EOConfig.CallOpts, source, symbol)
}

// PackPair is a free data retrieval call binding the contract method 0xc18cb7b0.
//
// Solidity: function packPair(uint16 source, uint16 symbol) pure returns(bytes4 id)
func (_EOConfig *EOConfigCallerSession) PackPair(source uint16, symbol uint16) ([4]byte, error) {
	return _EOConfig.Contract.PackPair(&_EOConfig.CallOpts, source, symbol)
}

// SourceById is a free data retrieval call binding the contract method 0xbf81551e.
//
// Solidity: function sourceById(uint16 ) view returns(bytes24 name, uint16 id)
func (_EOConfig *EOConfigCaller) SourceById(opts *bind.CallOpts, arg0 uint16) (struct {
	Name [24]byte
	Id   uint16
}, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "sourceById", arg0)

	outstruct := new(struct {
		Name [24]byte
		Id   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.Id = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// SourceById is a free data retrieval call binding the contract method 0xbf81551e.
//
// Solidity: function sourceById(uint16 ) view returns(bytes24 name, uint16 id)
func (_EOConfig *EOConfigSession) SourceById(arg0 uint16) (struct {
	Name [24]byte
	Id   uint16
}, error) {
	return _EOConfig.Contract.SourceById(&_EOConfig.CallOpts, arg0)
}

// SourceById is a free data retrieval call binding the contract method 0xbf81551e.
//
// Solidity: function sourceById(uint16 ) view returns(bytes24 name, uint16 id)
func (_EOConfig *EOConfigCallerSession) SourceById(arg0 uint16) (struct {
	Name [24]byte
	Id   uint16
}, error) {
	return _EOConfig.Contract.SourceById(&_EOConfig.CallOpts, arg0)
}

// SourceId is a free data retrieval call binding the contract method 0x88dea97d.
//
// Solidity: function sourceId(bytes24 ) view returns(uint16)
func (_EOConfig *EOConfigCaller) SourceId(opts *bind.CallOpts, arg0 [24]byte) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "sourceId", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SourceId is a free data retrieval call binding the contract method 0x88dea97d.
//
// Solidity: function sourceId(bytes24 ) view returns(uint16)
func (_EOConfig *EOConfigSession) SourceId(arg0 [24]byte) (uint16, error) {
	return _EOConfig.Contract.SourceId(&_EOConfig.CallOpts, arg0)
}

// SourceId is a free data retrieval call binding the contract method 0x88dea97d.
//
// Solidity: function sourceId(bytes24 ) view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SourceId(arg0 [24]byte) (uint16, error) {
	return _EOConfig.Contract.SourceId(&_EOConfig.CallOpts, arg0)
}

// SourceIds is a free data retrieval call binding the contract method 0x6d34c28b.
//
// Solidity: function sourceIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigCaller) SourceIds(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "sourceIds", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SourceIds is a free data retrieval call binding the contract method 0x6d34c28b.
//
// Solidity: function sourceIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigSession) SourceIds(arg0 *big.Int) (uint16, error) {
	return _EOConfig.Contract.SourceIds(&_EOConfig.CallOpts, arg0)
}

// SourceIds is a free data retrieval call binding the contract method 0x6d34c28b.
//
// Solidity: function sourceIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SourceIds(arg0 *big.Int) (uint16, error) {
	return _EOConfig.Contract.SourceIds(&_EOConfig.CallOpts, arg0)
}

// SourceIdsLength is a free data retrieval call binding the contract method 0x009c03fb.
//
// Solidity: function sourceIdsLength() view returns(uint16)
func (_EOConfig *EOConfigCaller) SourceIdsLength(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "sourceIdsLength")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SourceIdsLength is a free data retrieval call binding the contract method 0x009c03fb.
//
// Solidity: function sourceIdsLength() view returns(uint16)
func (_EOConfig *EOConfigSession) SourceIdsLength() (uint16, error) {
	return _EOConfig.Contract.SourceIdsLength(&_EOConfig.CallOpts)
}

// SourceIdsLength is a free data retrieval call binding the contract method 0x009c03fb.
//
// Solidity: function sourceIdsLength() view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SourceIdsLength() (uint16, error) {
	return _EOConfig.Contract.SourceIdsLength(&_EOConfig.CallOpts)
}

// SourceSymbolOverrides is a free data retrieval call binding the contract method 0x98332e91.
//
// Solidity: function sourceSymbolOverrides(uint16 , uint16 ) view returns(string)
func (_EOConfig *EOConfigCaller) SourceSymbolOverrides(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (string, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "sourceSymbolOverrides", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SourceSymbolOverrides is a free data retrieval call binding the contract method 0x98332e91.
//
// Solidity: function sourceSymbolOverrides(uint16 , uint16 ) view returns(string)
func (_EOConfig *EOConfigSession) SourceSymbolOverrides(arg0 uint16, arg1 uint16) (string, error) {
	return _EOConfig.Contract.SourceSymbolOverrides(&_EOConfig.CallOpts, arg0, arg1)
}

// SourceSymbolOverrides is a free data retrieval call binding the contract method 0x98332e91.
//
// Solidity: function sourceSymbolOverrides(uint16 , uint16 ) view returns(string)
func (_EOConfig *EOConfigCallerSession) SourceSymbolOverrides(arg0 uint16, arg1 uint16) (string, error) {
	return _EOConfig.Contract.SourceSymbolOverrides(&_EOConfig.CallOpts, arg0, arg1)
}

// SymbolById is a free data retrieval call binding the contract method 0xa0a79a7c.
//
// Solidity: function symbolById(uint16 ) view returns(bytes8 name, uint16 id)
func (_EOConfig *EOConfigCaller) SymbolById(opts *bind.CallOpts, arg0 uint16) (struct {
	Name [8]byte
	Id   uint16
}, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "symbolById", arg0)

	outstruct := new(struct {
		Name [8]byte
		Id   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)
	outstruct.Id = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// SymbolById is a free data retrieval call binding the contract method 0xa0a79a7c.
//
// Solidity: function symbolById(uint16 ) view returns(bytes8 name, uint16 id)
func (_EOConfig *EOConfigSession) SymbolById(arg0 uint16) (struct {
	Name [8]byte
	Id   uint16
}, error) {
	return _EOConfig.Contract.SymbolById(&_EOConfig.CallOpts, arg0)
}

// SymbolById is a free data retrieval call binding the contract method 0xa0a79a7c.
//
// Solidity: function symbolById(uint16 ) view returns(bytes8 name, uint16 id)
func (_EOConfig *EOConfigCallerSession) SymbolById(arg0 uint16) (struct {
	Name [8]byte
	Id   uint16
}, error) {
	return _EOConfig.Contract.SymbolById(&_EOConfig.CallOpts, arg0)
}

// SymbolId is a free data retrieval call binding the contract method 0x2d75981c.
//
// Solidity: function symbolId(bytes8 ) view returns(uint16)
func (_EOConfig *EOConfigCaller) SymbolId(opts *bind.CallOpts, arg0 [8]byte) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "symbolId", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SymbolId is a free data retrieval call binding the contract method 0x2d75981c.
//
// Solidity: function symbolId(bytes8 ) view returns(uint16)
func (_EOConfig *EOConfigSession) SymbolId(arg0 [8]byte) (uint16, error) {
	return _EOConfig.Contract.SymbolId(&_EOConfig.CallOpts, arg0)
}

// SymbolId is a free data retrieval call binding the contract method 0x2d75981c.
//
// Solidity: function symbolId(bytes8 ) view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SymbolId(arg0 [8]byte) (uint16, error) {
	return _EOConfig.Contract.SymbolId(&_EOConfig.CallOpts, arg0)
}

// SymbolIds is a free data retrieval call binding the contract method 0xedc24d69.
//
// Solidity: function symbolIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigCaller) SymbolIds(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "symbolIds", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SymbolIds is a free data retrieval call binding the contract method 0xedc24d69.
//
// Solidity: function symbolIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigSession) SymbolIds(arg0 *big.Int) (uint16, error) {
	return _EOConfig.Contract.SymbolIds(&_EOConfig.CallOpts, arg0)
}

// SymbolIds is a free data retrieval call binding the contract method 0xedc24d69.
//
// Solidity: function symbolIds(uint256 ) view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SymbolIds(arg0 *big.Int) (uint16, error) {
	return _EOConfig.Contract.SymbolIds(&_EOConfig.CallOpts, arg0)
}

// SymbolIdsLength is a free data retrieval call binding the contract method 0x20a332a4.
//
// Solidity: function symbolIdsLength() view returns(uint16)
func (_EOConfig *EOConfigCaller) SymbolIdsLength(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EOConfig.contract.Call(opts, &out, "symbolIdsLength")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SymbolIdsLength is a free data retrieval call binding the contract method 0x20a332a4.
//
// Solidity: function symbolIdsLength() view returns(uint16)
func (_EOConfig *EOConfigSession) SymbolIdsLength() (uint16, error) {
	return _EOConfig.Contract.SymbolIdsLength(&_EOConfig.CallOpts)
}

// SymbolIdsLength is a free data retrieval call binding the contract method 0x20a332a4.
//
// Solidity: function symbolIdsLength() view returns(uint16)
func (_EOConfig *EOConfigCallerSession) SymbolIdsLength() (uint16, error) {
	return _EOConfig.Contract.SymbolIdsLength(&_EOConfig.CallOpts)
}

// AddSource is a paid mutator transaction binding the contract method 0x13996105.
//
// Solidity: function addSource(bytes24 name) returns()
func (_EOConfig *EOConfigTransactor) AddSource(opts *bind.TransactOpts, name [24]byte) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "addSource", name)
}

// AddSource is a paid mutator transaction binding the contract method 0x13996105.
//
// Solidity: function addSource(bytes24 name) returns()
func (_EOConfig *EOConfigSession) AddSource(name [24]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSource(&_EOConfig.TransactOpts, name)
}

// AddSource is a paid mutator transaction binding the contract method 0x13996105.
//
// Solidity: function addSource(bytes24 name) returns()
func (_EOConfig *EOConfigTransactorSession) AddSource(name [24]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSource(&_EOConfig.TransactOpts, name)
}

// AddSources is a paid mutator transaction binding the contract method 0xb2ee3c95.
//
// Solidity: function addSources(bytes24[] names) returns()
func (_EOConfig *EOConfigTransactor) AddSources(opts *bind.TransactOpts, names [][24]byte) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "addSources", names)
}

// AddSources is a paid mutator transaction binding the contract method 0xb2ee3c95.
//
// Solidity: function addSources(bytes24[] names) returns()
func (_EOConfig *EOConfigSession) AddSources(names [][24]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSources(&_EOConfig.TransactOpts, names)
}

// AddSources is a paid mutator transaction binding the contract method 0xb2ee3c95.
//
// Solidity: function addSources(bytes24[] names) returns()
func (_EOConfig *EOConfigTransactorSession) AddSources(names [][24]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSources(&_EOConfig.TransactOpts, names)
}

// AddSymbol is a paid mutator transaction binding the contract method 0x08059117.
//
// Solidity: function addSymbol(bytes8 name) returns()
func (_EOConfig *EOConfigTransactor) AddSymbol(opts *bind.TransactOpts, name [8]byte) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "addSymbol", name)
}

// AddSymbol is a paid mutator transaction binding the contract method 0x08059117.
//
// Solidity: function addSymbol(bytes8 name) returns()
func (_EOConfig *EOConfigSession) AddSymbol(name [8]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSymbol(&_EOConfig.TransactOpts, name)
}

// AddSymbol is a paid mutator transaction binding the contract method 0x08059117.
//
// Solidity: function addSymbol(bytes8 name) returns()
func (_EOConfig *EOConfigTransactorSession) AddSymbol(name [8]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSymbol(&_EOConfig.TransactOpts, name)
}

// AddSymbols is a paid mutator transaction binding the contract method 0xba868392.
//
// Solidity: function addSymbols(bytes8[] names) returns()
func (_EOConfig *EOConfigTransactor) AddSymbols(opts *bind.TransactOpts, names [][8]byte) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "addSymbols", names)
}

// AddSymbols is a paid mutator transaction binding the contract method 0xba868392.
//
// Solidity: function addSymbols(bytes8[] names) returns()
func (_EOConfig *EOConfigSession) AddSymbols(names [][8]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSymbols(&_EOConfig.TransactOpts, names)
}

// AddSymbols is a paid mutator transaction binding the contract method 0xba868392.
//
// Solidity: function addSymbols(bytes8[] names) returns()
func (_EOConfig *EOConfigTransactorSession) AddSymbols(names [][8]byte) (*types.Transaction, error) {
	return _EOConfig.Contract.AddSymbols(&_EOConfig.TransactOpts, names)
}

// AssignAlias is a paid mutator transaction binding the contract method 0x2e0ec10c.
//
// Solidity: function assignAlias(address operator, address operatorAlias) returns()
func (_EOConfig *EOConfigTransactor) AssignAlias(opts *bind.TransactOpts, operator common.Address, operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "assignAlias", operator, operatorAlias)
}

// AssignAlias is a paid mutator transaction binding the contract method 0x2e0ec10c.
//
// Solidity: function assignAlias(address operator, address operatorAlias) returns()
func (_EOConfig *EOConfigSession) AssignAlias(operator common.Address, operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.AssignAlias(&_EOConfig.TransactOpts, operator, operatorAlias)
}

// AssignAlias is a paid mutator transaction binding the contract method 0x2e0ec10c.
//
// Solidity: function assignAlias(address operator, address operatorAlias) returns()
func (_EOConfig *EOConfigTransactorSession) AssignAlias(operator common.Address, operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.AssignAlias(&_EOConfig.TransactOpts, operator, operatorAlias)
}

// ChangeAlias is a paid mutator transaction binding the contract method 0x5c7c8b9a.
//
// Solidity: function changeAlias(address operator, address newAlias) returns()
func (_EOConfig *EOConfigTransactor) ChangeAlias(opts *bind.TransactOpts, operator common.Address, newAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "changeAlias", operator, newAlias)
}

// ChangeAlias is a paid mutator transaction binding the contract method 0x5c7c8b9a.
//
// Solidity: function changeAlias(address operator, address newAlias) returns()
func (_EOConfig *EOConfigSession) ChangeAlias(operator common.Address, newAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.ChangeAlias(&_EOConfig.TransactOpts, operator, newAlias)
}

// ChangeAlias is a paid mutator transaction binding the contract method 0x5c7c8b9a.
//
// Solidity: function changeAlias(address operator, address newAlias) returns()
func (_EOConfig *EOConfigTransactorSession) ChangeAlias(operator common.Address, newAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.ChangeAlias(&_EOConfig.TransactOpts, operator, newAlias)
}

// DeclareAlias is a paid mutator transaction binding the contract method 0xf405566d.
//
// Solidity: function declareAlias(address operatorAlias) returns()
func (_EOConfig *EOConfigTransactor) DeclareAlias(opts *bind.TransactOpts, operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "declareAlias", operatorAlias)
}

// DeclareAlias is a paid mutator transaction binding the contract method 0xf405566d.
//
// Solidity: function declareAlias(address operatorAlias) returns()
func (_EOConfig *EOConfigSession) DeclareAlias(operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.DeclareAlias(&_EOConfig.TransactOpts, operatorAlias)
}

// DeclareAlias is a paid mutator transaction binding the contract method 0xf405566d.
//
// Solidity: function declareAlias(address operatorAlias) returns()
func (_EOConfig *EOConfigTransactorSession) DeclareAlias(operatorAlias common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.DeclareAlias(&_EOConfig.TransactOpts, operatorAlias)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EOConfig *EOConfigTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EOConfig *EOConfigSession) Initialize() (*types.Transaction, error) {
	return _EOConfig.Contract.Initialize(&_EOConfig.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_EOConfig *EOConfigTransactorSession) Initialize() (*types.Transaction, error) {
	return _EOConfig.Contract.Initialize(&_EOConfig.TransactOpts)
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_EOConfig *EOConfigTransactor) Notify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "notify")
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_EOConfig *EOConfigSession) Notify() (*types.Transaction, error) {
	return _EOConfig.Contract.Notify(&_EOConfig.TransactOpts)
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_EOConfig *EOConfigTransactorSession) Notify() (*types.Transaction, error) {
	return _EOConfig.Contract.Notify(&_EOConfig.TransactOpts)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_EOConfig *EOConfigTransactor) OnStateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "onStateReceive", arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_EOConfig *EOConfigSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _EOConfig.Contract.OnStateReceive(&_EOConfig.TransactOpts, arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_EOConfig *EOConfigTransactorSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _EOConfig.Contract.OnStateReceive(&_EOConfig.TransactOpts, arg0, sender, data)
}

// RemoveSource is a paid mutator transaction binding the contract method 0x7672db8e.
//
// Solidity: function removeSource(uint16 id) returns()
func (_EOConfig *EOConfigTransactor) RemoveSource(opts *bind.TransactOpts, id uint16) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "removeSource", id)
}

// RemoveSource is a paid mutator transaction binding the contract method 0x7672db8e.
//
// Solidity: function removeSource(uint16 id) returns()
func (_EOConfig *EOConfigSession) RemoveSource(id uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.RemoveSource(&_EOConfig.TransactOpts, id)
}

// RemoveSource is a paid mutator transaction binding the contract method 0x7672db8e.
//
// Solidity: function removeSource(uint16 id) returns()
func (_EOConfig *EOConfigTransactorSession) RemoveSource(id uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.RemoveSource(&_EOConfig.TransactOpts, id)
}

// RemoveSymbol is a paid mutator transaction binding the contract method 0x5aa264e4.
//
// Solidity: function removeSymbol(uint16 id) returns()
func (_EOConfig *EOConfigTransactor) RemoveSymbol(opts *bind.TransactOpts, id uint16) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "removeSymbol", id)
}

// RemoveSymbol is a paid mutator transaction binding the contract method 0x5aa264e4.
//
// Solidity: function removeSymbol(uint16 id) returns()
func (_EOConfig *EOConfigSession) RemoveSymbol(id uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.RemoveSymbol(&_EOConfig.TransactOpts, id)
}

// RemoveSymbol is a paid mutator transaction binding the contract method 0x5aa264e4.
//
// Solidity: function removeSymbol(uint16 id) returns()
func (_EOConfig *EOConfigTransactorSession) RemoveSymbol(id uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.RemoveSymbol(&_EOConfig.TransactOpts, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EOConfig *EOConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EOConfig *EOConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _EOConfig.Contract.RenounceOwnership(&_EOConfig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EOConfig *EOConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EOConfig.Contract.RenounceOwnership(&_EOConfig.TransactOpts)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_EOConfig *EOConfigTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_EOConfig *EOConfigSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetAggregator(&_EOConfig.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_EOConfig *EOConfigTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetAggregator(&_EOConfig.TransactOpts, _aggregator)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address _chainManager) returns()
func (_EOConfig *EOConfigTransactor) SetChainManager(opts *bind.TransactOpts, _chainManager common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "setChainManager", _chainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address _chainManager) returns()
func (_EOConfig *EOConfigSession) SetChainManager(_chainManager common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetChainManager(&_EOConfig.TransactOpts, _chainManager)
}

// SetChainManager is a paid mutator transaction binding the contract method 0xd70cf0ed.
//
// Solidity: function setChainManager(address _chainManager) returns()
func (_EOConfig *EOConfigTransactorSession) SetChainManager(_chainManager common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetChainManager(&_EOConfig.TransactOpts, _chainManager)
}

// SetEvmSourceConfig is a paid mutator transaction binding the contract method 0x4571adf1.
//
// Solidity: function setEvmSourceConfig(address _evmSourceConfig) returns()
func (_EOConfig *EOConfigTransactor) SetEvmSourceConfig(opts *bind.TransactOpts, _evmSourceConfig common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "setEvmSourceConfig", _evmSourceConfig)
}

// SetEvmSourceConfig is a paid mutator transaction binding the contract method 0x4571adf1.
//
// Solidity: function setEvmSourceConfig(address _evmSourceConfig) returns()
func (_EOConfig *EOConfigSession) SetEvmSourceConfig(_evmSourceConfig common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetEvmSourceConfig(&_EOConfig.TransactOpts, _evmSourceConfig)
}

// SetEvmSourceConfig is a paid mutator transaction binding the contract method 0x4571adf1.
//
// Solidity: function setEvmSourceConfig(address _evmSourceConfig) returns()
func (_EOConfig *EOConfigTransactorSession) SetEvmSourceConfig(_evmSourceConfig common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.SetEvmSourceConfig(&_EOConfig.TransactOpts, _evmSourceConfig)
}

// SetSourceSymbolIds is a paid mutator transaction binding the contract method 0xe000b072.
//
// Solidity: function setSourceSymbolIds(uint16 source, uint16[] symbols, string[] names) returns()
func (_EOConfig *EOConfigTransactor) SetSourceSymbolIds(opts *bind.TransactOpts, source uint16, symbols []uint16, names []string) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "setSourceSymbolIds", source, symbols, names)
}

// SetSourceSymbolIds is a paid mutator transaction binding the contract method 0xe000b072.
//
// Solidity: function setSourceSymbolIds(uint16 source, uint16[] symbols, string[] names) returns()
func (_EOConfig *EOConfigSession) SetSourceSymbolIds(source uint16, symbols []uint16, names []string) (*types.Transaction, error) {
	return _EOConfig.Contract.SetSourceSymbolIds(&_EOConfig.TransactOpts, source, symbols, names)
}

// SetSourceSymbolIds is a paid mutator transaction binding the contract method 0xe000b072.
//
// Solidity: function setSourceSymbolIds(uint16 source, uint16[] symbols, string[] names) returns()
func (_EOConfig *EOConfigTransactorSession) SetSourceSymbolIds(source uint16, symbols []uint16, names []string) (*types.Transaction, error) {
	return _EOConfig.Contract.SetSourceSymbolIds(&_EOConfig.TransactOpts, source, symbols, names)
}

// Subscribe is a paid mutator transaction binding the contract method 0x41a7726a.
//
// Solidity: function subscribe(address listener) returns()
func (_EOConfig *EOConfigTransactor) Subscribe(opts *bind.TransactOpts, listener common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "subscribe", listener)
}

// Subscribe is a paid mutator transaction binding the contract method 0x41a7726a.
//
// Solidity: function subscribe(address listener) returns()
func (_EOConfig *EOConfigSession) Subscribe(listener common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.Subscribe(&_EOConfig.TransactOpts, listener)
}

// Subscribe is a paid mutator transaction binding the contract method 0x41a7726a.
//
// Solidity: function subscribe(address listener) returns()
func (_EOConfig *EOConfigTransactorSession) Subscribe(listener common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.Subscribe(&_EOConfig.TransactOpts, listener)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EOConfig *EOConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EOConfig *EOConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.TransferOwnership(&_EOConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EOConfig *EOConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.TransferOwnership(&_EOConfig.TransactOpts, newOwner)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address listener) returns()
func (_EOConfig *EOConfigTransactor) Unsubscribe(opts *bind.TransactOpts, listener common.Address) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "unsubscribe", listener)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address listener) returns()
func (_EOConfig *EOConfigSession) Unsubscribe(listener common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.Unsubscribe(&_EOConfig.TransactOpts, listener)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x7262561c.
//
// Solidity: function unsubscribe(address listener) returns()
func (_EOConfig *EOConfigTransactorSession) Unsubscribe(listener common.Address) (*types.Transaction, error) {
	return _EOConfig.Contract.Unsubscribe(&_EOConfig.TransactOpts, listener)
}

// UpdateSource is a paid mutator transaction binding the contract method 0x88c54990.
//
// Solidity: function updateSource(uint16 _sourceId, bytes24 name, uint16[] _symbols) returns()
func (_EOConfig *EOConfigTransactor) UpdateSource(opts *bind.TransactOpts, _sourceId uint16, name [24]byte, _symbols []uint16) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "updateSource", _sourceId, name, _symbols)
}

// UpdateSource is a paid mutator transaction binding the contract method 0x88c54990.
//
// Solidity: function updateSource(uint16 _sourceId, bytes24 name, uint16[] _symbols) returns()
func (_EOConfig *EOConfigSession) UpdateSource(_sourceId uint16, name [24]byte, _symbols []uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.UpdateSource(&_EOConfig.TransactOpts, _sourceId, name, _symbols)
}

// UpdateSource is a paid mutator transaction binding the contract method 0x88c54990.
//
// Solidity: function updateSource(uint16 _sourceId, bytes24 name, uint16[] _symbols) returns()
func (_EOConfig *EOConfigTransactorSession) UpdateSource(_sourceId uint16, name [24]byte, _symbols []uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.UpdateSource(&_EOConfig.TransactOpts, _sourceId, name, _symbols)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x9483a58d.
//
// Solidity: function updateSymbol(uint16 _symbolId, bytes8 name, uint16[] _sources) returns()
func (_EOConfig *EOConfigTransactor) UpdateSymbol(opts *bind.TransactOpts, _symbolId uint16, name [8]byte, _sources []uint16) (*types.Transaction, error) {
	return _EOConfig.contract.Transact(opts, "updateSymbol", _symbolId, name, _sources)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x9483a58d.
//
// Solidity: function updateSymbol(uint16 _symbolId, bytes8 name, uint16[] _sources) returns()
func (_EOConfig *EOConfigSession) UpdateSymbol(_symbolId uint16, name [8]byte, _sources []uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.UpdateSymbol(&_EOConfig.TransactOpts, _symbolId, name, _sources)
}

// UpdateSymbol is a paid mutator transaction binding the contract method 0x9483a58d.
//
// Solidity: function updateSymbol(uint16 _symbolId, bytes8 name, uint16[] _sources) returns()
func (_EOConfig *EOConfigTransactorSession) UpdateSymbol(_symbolId uint16, name [8]byte, _sources []uint16) (*types.Transaction, error) {
	return _EOConfig.Contract.UpdateSymbol(&_EOConfig.TransactOpts, _symbolId, name, _sources)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EOConfig *EOConfigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOConfig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EOConfig *EOConfigSession) Receive() (*types.Transaction, error) {
	return _EOConfig.Contract.Receive(&_EOConfig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EOConfig *EOConfigTransactorSession) Receive() (*types.Transaction, error) {
	return _EOConfig.Contract.Receive(&_EOConfig.TransactOpts)
}

// EOConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EOConfig contract.
type EOConfigInitializedIterator struct {
	Event *EOConfigInitialized // Event containing the contract specifics and raw log

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
func (it *EOConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigInitialized)
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
		it.Event = new(EOConfigInitialized)
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
func (it *EOConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigInitialized represents a Initialized event raised by the EOConfig contract.
type EOConfigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EOConfig *EOConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*EOConfigInitializedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EOConfigInitializedIterator{contract: _EOConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EOConfig *EOConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EOConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigInitialized)
				if err := _EOConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EOConfig *EOConfigFilterer) ParseInitialized(log types.Log) (*EOConfigInitialized, error) {
	event := new(EOConfigInitialized)
	if err := _EOConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnConfigChangeIterator is returned from FilterOnConfigChange and is used to iterate over the raw logs and unpacked data for OnConfigChange events raised by the EOConfig contract.
type EOConfigOnConfigChangeIterator struct {
	Event *EOConfigOnConfigChange // Event containing the contract specifics and raw log

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
func (it *EOConfigOnConfigChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnConfigChange)
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
		it.Event = new(EOConfigOnConfigChange)
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
func (it *EOConfigOnConfigChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnConfigChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnConfigChange represents a OnConfigChange event raised by the EOConfig contract.
type EOConfigOnConfigChange struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnConfigChange is a free log retrieval operation binding the contract event 0x662f307c99b27e8c90dbce053f13da64171deb8c531b3ae9c9446ab010da7fdc.
//
// Solidity: event OnConfigChange()
func (_EOConfig *EOConfigFilterer) FilterOnConfigChange(opts *bind.FilterOpts) (*EOConfigOnConfigChangeIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnConfigChange")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnConfigChangeIterator{contract: _EOConfig.contract, event: "OnConfigChange", logs: logs, sub: sub}, nil
}

// WatchOnConfigChange is a free log subscription operation binding the contract event 0x662f307c99b27e8c90dbce053f13da64171deb8c531b3ae9c9446ab010da7fdc.
//
// Solidity: event OnConfigChange()
func (_EOConfig *EOConfigFilterer) WatchOnConfigChange(opts *bind.WatchOpts, sink chan<- *EOConfigOnConfigChange) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnConfigChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnConfigChange)
				if err := _EOConfig.contract.UnpackLog(event, "OnConfigChange", log); err != nil {
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

// ParseOnConfigChange is a log parse operation binding the contract event 0x662f307c99b27e8c90dbce053f13da64171deb8c531b3ae9c9446ab010da7fdc.
//
// Solidity: event OnConfigChange()
func (_EOConfig *EOConfigFilterer) ParseOnConfigChange(log types.Log) (*EOConfigOnConfigChange, error) {
	event := new(EOConfigOnConfigChange)
	if err := _EOConfig.contract.UnpackLog(event, "OnConfigChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSourceAddedIterator is returned from FilterOnSourceAdded and is used to iterate over the raw logs and unpacked data for OnSourceAdded events raised by the EOConfig contract.
type EOConfigOnSourceAddedIterator struct {
	Event *EOConfigOnSourceAdded // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSourceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSourceAdded)
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
		it.Event = new(EOConfigOnSourceAdded)
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
func (it *EOConfigOnSourceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSourceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSourceAdded represents a OnSourceAdded event raised by the EOConfig contract.
type EOConfigOnSourceAdded struct {
	Id   uint16
	Name [24]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSourceAdded is a free log retrieval operation binding the contract event 0x209da2b4b51d7df154f686f6190e38bf6eb987bf19f48df080b9a14111762d88.
//
// Solidity: event OnSourceAdded(uint16 id, bytes24 name)
func (_EOConfig *EOConfigFilterer) FilterOnSourceAdded(opts *bind.FilterOpts) (*EOConfigOnSourceAddedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSourceAdded")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSourceAddedIterator{contract: _EOConfig.contract, event: "OnSourceAdded", logs: logs, sub: sub}, nil
}

// WatchOnSourceAdded is a free log subscription operation binding the contract event 0x209da2b4b51d7df154f686f6190e38bf6eb987bf19f48df080b9a14111762d88.
//
// Solidity: event OnSourceAdded(uint16 id, bytes24 name)
func (_EOConfig *EOConfigFilterer) WatchOnSourceAdded(opts *bind.WatchOpts, sink chan<- *EOConfigOnSourceAdded) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSourceAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSourceAdded)
				if err := _EOConfig.contract.UnpackLog(event, "OnSourceAdded", log); err != nil {
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

// ParseOnSourceAdded is a log parse operation binding the contract event 0x209da2b4b51d7df154f686f6190e38bf6eb987bf19f48df080b9a14111762d88.
//
// Solidity: event OnSourceAdded(uint16 id, bytes24 name)
func (_EOConfig *EOConfigFilterer) ParseOnSourceAdded(log types.Log) (*EOConfigOnSourceAdded, error) {
	event := new(EOConfigOnSourceAdded)
	if err := _EOConfig.contract.UnpackLog(event, "OnSourceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSourceRemovedIterator is returned from FilterOnSourceRemoved and is used to iterate over the raw logs and unpacked data for OnSourceRemoved events raised by the EOConfig contract.
type EOConfigOnSourceRemovedIterator struct {
	Event *EOConfigOnSourceRemoved // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSourceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSourceRemoved)
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
		it.Event = new(EOConfigOnSourceRemoved)
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
func (it *EOConfigOnSourceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSourceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSourceRemoved represents a OnSourceRemoved event raised by the EOConfig contract.
type EOConfigOnSourceRemoved struct {
	Id  uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnSourceRemoved is a free log retrieval operation binding the contract event 0xce77a2b83741e02bd444c230b0f3ab40ea28201581cd73d3fe76809d16cca361.
//
// Solidity: event OnSourceRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) FilterOnSourceRemoved(opts *bind.FilterOpts) (*EOConfigOnSourceRemovedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSourceRemoved")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSourceRemovedIterator{contract: _EOConfig.contract, event: "OnSourceRemoved", logs: logs, sub: sub}, nil
}

// WatchOnSourceRemoved is a free log subscription operation binding the contract event 0xce77a2b83741e02bd444c230b0f3ab40ea28201581cd73d3fe76809d16cca361.
//
// Solidity: event OnSourceRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) WatchOnSourceRemoved(opts *bind.WatchOpts, sink chan<- *EOConfigOnSourceRemoved) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSourceRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSourceRemoved)
				if err := _EOConfig.contract.UnpackLog(event, "OnSourceRemoved", log); err != nil {
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

// ParseOnSourceRemoved is a log parse operation binding the contract event 0xce77a2b83741e02bd444c230b0f3ab40ea28201581cd73d3fe76809d16cca361.
//
// Solidity: event OnSourceRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) ParseOnSourceRemoved(log types.Log) (*EOConfigOnSourceRemoved, error) {
	event := new(EOConfigOnSourceRemoved)
	if err := _EOConfig.contract.UnpackLog(event, "OnSourceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSourceUpdatedIterator is returned from FilterOnSourceUpdated and is used to iterate over the raw logs and unpacked data for OnSourceUpdated events raised by the EOConfig contract.
type EOConfigOnSourceUpdatedIterator struct {
	Event *EOConfigOnSourceUpdated // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSourceUpdated)
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
		it.Event = new(EOConfigOnSourceUpdated)
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
func (it *EOConfigOnSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSourceUpdated represents a OnSourceUpdated event raised by the EOConfig contract.
type EOConfigOnSourceUpdated struct {
	Id      uint16
	Name    [24]byte
	Symbols []uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnSourceUpdated is a free log retrieval operation binding the contract event 0x051fa95d237354f8475b41634acc198ae4da9758585b7ab07f2f6839f71492a2.
//
// Solidity: event OnSourceUpdated(uint16 id, bytes24 name, uint16[] symbols)
func (_EOConfig *EOConfigFilterer) FilterOnSourceUpdated(opts *bind.FilterOpts) (*EOConfigOnSourceUpdatedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSourceUpdated")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSourceUpdatedIterator{contract: _EOConfig.contract, event: "OnSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchOnSourceUpdated is a free log subscription operation binding the contract event 0x051fa95d237354f8475b41634acc198ae4da9758585b7ab07f2f6839f71492a2.
//
// Solidity: event OnSourceUpdated(uint16 id, bytes24 name, uint16[] symbols)
func (_EOConfig *EOConfigFilterer) WatchOnSourceUpdated(opts *bind.WatchOpts, sink chan<- *EOConfigOnSourceUpdated) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSourceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSourceUpdated)
				if err := _EOConfig.contract.UnpackLog(event, "OnSourceUpdated", log); err != nil {
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

// ParseOnSourceUpdated is a log parse operation binding the contract event 0x051fa95d237354f8475b41634acc198ae4da9758585b7ab07f2f6839f71492a2.
//
// Solidity: event OnSourceUpdated(uint16 id, bytes24 name, uint16[] symbols)
func (_EOConfig *EOConfigFilterer) ParseOnSourceUpdated(log types.Log) (*EOConfigOnSourceUpdated, error) {
	event := new(EOConfigOnSourceUpdated)
	if err := _EOConfig.contract.UnpackLog(event, "OnSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSymbolAddedIterator is returned from FilterOnSymbolAdded and is used to iterate over the raw logs and unpacked data for OnSymbolAdded events raised by the EOConfig contract.
type EOConfigOnSymbolAddedIterator struct {
	Event *EOConfigOnSymbolAdded // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSymbolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSymbolAdded)
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
		it.Event = new(EOConfigOnSymbolAdded)
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
func (it *EOConfigOnSymbolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSymbolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSymbolAdded represents a OnSymbolAdded event raised by the EOConfig contract.
type EOConfigOnSymbolAdded struct {
	Id   uint16
	Name [8]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSymbolAdded is a free log retrieval operation binding the contract event 0x17c9d771a4294ad678b91f00d84d6c838ce1f073b54b67bf7070320331667e99.
//
// Solidity: event OnSymbolAdded(uint16 id, bytes8 name)
func (_EOConfig *EOConfigFilterer) FilterOnSymbolAdded(opts *bind.FilterOpts) (*EOConfigOnSymbolAddedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSymbolAdded")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSymbolAddedIterator{contract: _EOConfig.contract, event: "OnSymbolAdded", logs: logs, sub: sub}, nil
}

// WatchOnSymbolAdded is a free log subscription operation binding the contract event 0x17c9d771a4294ad678b91f00d84d6c838ce1f073b54b67bf7070320331667e99.
//
// Solidity: event OnSymbolAdded(uint16 id, bytes8 name)
func (_EOConfig *EOConfigFilterer) WatchOnSymbolAdded(opts *bind.WatchOpts, sink chan<- *EOConfigOnSymbolAdded) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSymbolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSymbolAdded)
				if err := _EOConfig.contract.UnpackLog(event, "OnSymbolAdded", log); err != nil {
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

// ParseOnSymbolAdded is a log parse operation binding the contract event 0x17c9d771a4294ad678b91f00d84d6c838ce1f073b54b67bf7070320331667e99.
//
// Solidity: event OnSymbolAdded(uint16 id, bytes8 name)
func (_EOConfig *EOConfigFilterer) ParseOnSymbolAdded(log types.Log) (*EOConfigOnSymbolAdded, error) {
	event := new(EOConfigOnSymbolAdded)
	if err := _EOConfig.contract.UnpackLog(event, "OnSymbolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSymbolRemovedIterator is returned from FilterOnSymbolRemoved and is used to iterate over the raw logs and unpacked data for OnSymbolRemoved events raised by the EOConfig contract.
type EOConfigOnSymbolRemovedIterator struct {
	Event *EOConfigOnSymbolRemoved // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSymbolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSymbolRemoved)
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
		it.Event = new(EOConfigOnSymbolRemoved)
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
func (it *EOConfigOnSymbolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSymbolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSymbolRemoved represents a OnSymbolRemoved event raised by the EOConfig contract.
type EOConfigOnSymbolRemoved struct {
	Id  uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnSymbolRemoved is a free log retrieval operation binding the contract event 0xaf1d0220552f1ffa37378a7942dffc37d3f6e51062806798e46e650aa95ae693.
//
// Solidity: event OnSymbolRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) FilterOnSymbolRemoved(opts *bind.FilterOpts) (*EOConfigOnSymbolRemovedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSymbolRemoved")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSymbolRemovedIterator{contract: _EOConfig.contract, event: "OnSymbolRemoved", logs: logs, sub: sub}, nil
}

// WatchOnSymbolRemoved is a free log subscription operation binding the contract event 0xaf1d0220552f1ffa37378a7942dffc37d3f6e51062806798e46e650aa95ae693.
//
// Solidity: event OnSymbolRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) WatchOnSymbolRemoved(opts *bind.WatchOpts, sink chan<- *EOConfigOnSymbolRemoved) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSymbolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSymbolRemoved)
				if err := _EOConfig.contract.UnpackLog(event, "OnSymbolRemoved", log); err != nil {
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

// ParseOnSymbolRemoved is a log parse operation binding the contract event 0xaf1d0220552f1ffa37378a7942dffc37d3f6e51062806798e46e650aa95ae693.
//
// Solidity: event OnSymbolRemoved(uint16 id)
func (_EOConfig *EOConfigFilterer) ParseOnSymbolRemoved(log types.Log) (*EOConfigOnSymbolRemoved, error) {
	event := new(EOConfigOnSymbolRemoved)
	if err := _EOConfig.contract.UnpackLog(event, "OnSymbolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOnSymbolUpdatedIterator is returned from FilterOnSymbolUpdated and is used to iterate over the raw logs and unpacked data for OnSymbolUpdated events raised by the EOConfig contract.
type EOConfigOnSymbolUpdatedIterator struct {
	Event *EOConfigOnSymbolUpdated // Event containing the contract specifics and raw log

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
func (it *EOConfigOnSymbolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOnSymbolUpdated)
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
		it.Event = new(EOConfigOnSymbolUpdated)
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
func (it *EOConfigOnSymbolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOnSymbolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOnSymbolUpdated represents a OnSymbolUpdated event raised by the EOConfig contract.
type EOConfigOnSymbolUpdated struct {
	Id      uint16
	Name    [8]byte
	Sources []uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnSymbolUpdated is a free log retrieval operation binding the contract event 0x05aefdf7098c008b0f258f7f5aceb8ed4e93f8faedba69832f087a0c3a09d1df.
//
// Solidity: event OnSymbolUpdated(uint16 id, bytes8 name, uint16[] sources)
func (_EOConfig *EOConfigFilterer) FilterOnSymbolUpdated(opts *bind.FilterOpts) (*EOConfigOnSymbolUpdatedIterator, error) {

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OnSymbolUpdated")
	if err != nil {
		return nil, err
	}
	return &EOConfigOnSymbolUpdatedIterator{contract: _EOConfig.contract, event: "OnSymbolUpdated", logs: logs, sub: sub}, nil
}

// WatchOnSymbolUpdated is a free log subscription operation binding the contract event 0x05aefdf7098c008b0f258f7f5aceb8ed4e93f8faedba69832f087a0c3a09d1df.
//
// Solidity: event OnSymbolUpdated(uint16 id, bytes8 name, uint16[] sources)
func (_EOConfig *EOConfigFilterer) WatchOnSymbolUpdated(opts *bind.WatchOpts, sink chan<- *EOConfigOnSymbolUpdated) (event.Subscription, error) {

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OnSymbolUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOnSymbolUpdated)
				if err := _EOConfig.contract.UnpackLog(event, "OnSymbolUpdated", log); err != nil {
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

// ParseOnSymbolUpdated is a log parse operation binding the contract event 0x05aefdf7098c008b0f258f7f5aceb8ed4e93f8faedba69832f087a0c3a09d1df.
//
// Solidity: event OnSymbolUpdated(uint16 id, bytes8 name, uint16[] sources)
func (_EOConfig *EOConfigFilterer) ParseOnSymbolUpdated(log types.Log) (*EOConfigOnSymbolUpdated, error) {
	event := new(EOConfigOnSymbolUpdated)
	if err := _EOConfig.contract.UnpackLog(event, "OnSymbolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EOConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EOConfig contract.
type EOConfigOwnershipTransferredIterator struct {
	Event *EOConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EOConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EOConfigOwnershipTransferred)
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
		it.Event = new(EOConfigOwnershipTransferred)
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
func (it *EOConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EOConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EOConfigOwnershipTransferred represents a OwnershipTransferred event raised by the EOConfig contract.
type EOConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EOConfig *EOConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EOConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EOConfig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EOConfigOwnershipTransferredIterator{contract: _EOConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EOConfig *EOConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EOConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EOConfig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EOConfigOwnershipTransferred)
				if err := _EOConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EOConfig *EOConfigFilterer) ParseOwnershipTransferred(log types.Log) (*EOConfigOwnershipTransferred, error) {
	event := new(EOConfigOwnershipTransferred)
	if err := _EOConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
