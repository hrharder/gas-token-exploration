// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GasToken2ABI is the input ABI used to generate the binding from.
const GasToken2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeFromUpTo\",\"outputs\":[{\"name\":\"freed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"freeUpTo\",\"outputs\":[{\"name\":\"freed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// GasToken2 is an auto generated Go binding around an Ethereum contract.
type GasToken2 struct {
	GasToken2Caller     // Read-only binding to the contract
	GasToken2Transactor // Write-only binding to the contract
	GasToken2Filterer   // Log filterer for contract events
}

// GasToken2Caller is an auto generated read-only Go binding around an Ethereum contract.
type GasToken2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasToken2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GasToken2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasToken2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasToken2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasToken2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasToken2Session struct {
	Contract     *GasToken2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasToken2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasToken2CallerSession struct {
	Contract *GasToken2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GasToken2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasToken2TransactorSession struct {
	Contract     *GasToken2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GasToken2Raw is an auto generated low-level Go binding around an Ethereum contract.
type GasToken2Raw struct {
	Contract *GasToken2 // Generic contract binding to access the raw methods on
}

// GasToken2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasToken2CallerRaw struct {
	Contract *GasToken2Caller // Generic read-only contract binding to access the raw methods on
}

// GasToken2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasToken2TransactorRaw struct {
	Contract *GasToken2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGasToken2 creates a new instance of GasToken2, bound to a specific deployed contract.
func NewGasToken2(address common.Address, backend bind.ContractBackend) (*GasToken2, error) {
	contract, err := bindGasToken2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasToken2{GasToken2Caller: GasToken2Caller{contract: contract}, GasToken2Transactor: GasToken2Transactor{contract: contract}, GasToken2Filterer: GasToken2Filterer{contract: contract}}, nil
}

// NewGasToken2Caller creates a new read-only instance of GasToken2, bound to a specific deployed contract.
func NewGasToken2Caller(address common.Address, caller bind.ContractCaller) (*GasToken2Caller, error) {
	contract, err := bindGasToken2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasToken2Caller{contract: contract}, nil
}

// NewGasToken2Transactor creates a new write-only instance of GasToken2, bound to a specific deployed contract.
func NewGasToken2Transactor(address common.Address, transactor bind.ContractTransactor) (*GasToken2Transactor, error) {
	contract, err := bindGasToken2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasToken2Transactor{contract: contract}, nil
}

// NewGasToken2Filterer creates a new log filterer instance of GasToken2, bound to a specific deployed contract.
func NewGasToken2Filterer(address common.Address, filterer bind.ContractFilterer) (*GasToken2Filterer, error) {
	contract, err := bindGasToken2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasToken2Filterer{contract: contract}, nil
}

// bindGasToken2 binds a generic wrapper to an already deployed contract.
func bindGasToken2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasToken2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken2 *GasToken2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken2.Contract.GasToken2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken2 *GasToken2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken2.Contract.GasToken2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken2 *GasToken2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken2.Contract.GasToken2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasToken2 *GasToken2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasToken2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasToken2 *GasToken2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasToken2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasToken2 *GasToken2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasToken2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_GasToken2 *GasToken2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_GasToken2 *GasToken2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GasToken2.Contract.Allowance(&_GasToken2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_GasToken2 *GasToken2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GasToken2.Contract.Allowance(&_GasToken2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_GasToken2 *GasToken2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_GasToken2 *GasToken2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GasToken2.Contract.BalanceOf(&_GasToken2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_GasToken2 *GasToken2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GasToken2.Contract.BalanceOf(&_GasToken2.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GasToken2 *GasToken2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GasToken2 *GasToken2Session) Decimals() (uint8, error) {
	return _GasToken2.Contract.Decimals(&_GasToken2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_GasToken2 *GasToken2CallerSession) Decimals() (uint8, error) {
	return _GasToken2.Contract.Decimals(&_GasToken2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GasToken2 *GasToken2Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GasToken2 *GasToken2Session) Name() (string, error) {
	return _GasToken2.Contract.Name(&_GasToken2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_GasToken2 *GasToken2CallerSession) Name() (string, error) {
	return _GasToken2.Contract.Name(&_GasToken2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GasToken2 *GasToken2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GasToken2 *GasToken2Session) Symbol() (string, error) {
	return _GasToken2.Contract.Symbol(&_GasToken2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_GasToken2 *GasToken2CallerSession) Symbol() (string, error) {
	return _GasToken2.Contract.Symbol(&_GasToken2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 supply)
func (_GasToken2 *GasToken2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasToken2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 supply)
func (_GasToken2 *GasToken2Session) TotalSupply() (*big.Int, error) {
	return _GasToken2.Contract.TotalSupply(&_GasToken2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 supply)
func (_GasToken2 *GasToken2CallerSession) TotalSupply() (*big.Int, error) {
	return _GasToken2.Contract.TotalSupply(&_GasToken2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Approve(&_GasToken2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Approve(&_GasToken2.TransactOpts, spender, value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Transactor) Free(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "free", value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Session) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Free(&_GasToken2.TransactOpts, value)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 value) returns(bool success)
func (_GasToken2 *GasToken2TransactorSession) Free(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Free(&_GasToken2.TransactOpts, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Transactor) FreeFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "freeFrom", from, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Session) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeFrom(&_GasToken2.TransactOpts, from, value)
}

// FreeFrom is a paid mutator transaction binding the contract method 0x5f2e2b45.
//
// Solidity: function freeFrom(address from, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2TransactorSession) FreeFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeFrom(&_GasToken2.TransactOpts, from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2Transactor) FreeFromUpTo(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "freeFromUpTo", from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2Session) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeFromUpTo(&_GasToken2.TransactOpts, from, value)
}

// FreeFromUpTo is a paid mutator transaction binding the contract method 0x079d229f.
//
// Solidity: function freeFromUpTo(address from, uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2TransactorSession) FreeFromUpTo(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeFromUpTo(&_GasToken2.TransactOpts, from, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2Transactor) FreeUpTo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "freeUpTo", value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2Session) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeUpTo(&_GasToken2.TransactOpts, value)
}

// FreeUpTo is a paid mutator transaction binding the contract method 0x6366b936.
//
// Solidity: function freeUpTo(uint256 value) returns(uint256 freed)
func (_GasToken2 *GasToken2TransactorSession) FreeUpTo(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.FreeUpTo(&_GasToken2.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken2 *GasToken2Transactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken2 *GasToken2Session) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Mint(&_GasToken2.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_GasToken2 *GasToken2TransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Mint(&_GasToken2.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Transfer(&_GasToken2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.Transfer(&_GasToken2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.TransferFrom(&_GasToken2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_GasToken2 *GasToken2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _GasToken2.Contract.TransferFrom(&_GasToken2.TransactOpts, from, to, value)
}

// GasToken2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GasToken2 contract.
type GasToken2ApprovalIterator struct {
	Event *GasToken2Approval // Event containing the contract specifics and raw log

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
func (it *GasToken2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasToken2Approval)
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
		it.Event = new(GasToken2Approval)
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
func (it *GasToken2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasToken2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasToken2Approval represents a Approval event raised by the GasToken2 contract.
type GasToken2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GasToken2 *GasToken2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GasToken2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GasToken2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GasToken2ApprovalIterator{contract: _GasToken2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GasToken2 *GasToken2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GasToken2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GasToken2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasToken2Approval)
				if err := _GasToken2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// GasToken2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GasToken2 contract.
type GasToken2TransferIterator struct {
	Event *GasToken2Transfer // Event containing the contract specifics and raw log

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
func (it *GasToken2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasToken2Transfer)
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
		it.Event = new(GasToken2Transfer)
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
func (it *GasToken2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasToken2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasToken2Transfer represents a Transfer event raised by the GasToken2 contract.
type GasToken2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GasToken2 *GasToken2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GasToken2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasToken2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GasToken2TransferIterator{contract: _GasToken2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GasToken2 *GasToken2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GasToken2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasToken2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasToken2Transfer)
				if err := _GasToken2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
