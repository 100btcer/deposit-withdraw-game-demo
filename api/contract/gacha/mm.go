// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gacha

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

// MMMetaData contains all meta data concerning the MM contract.
var MMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticketAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockTimestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimal\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofs\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getDepositByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLpBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"sysTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllLp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MMABI is the input ABI used to generate the binding from.
// Deprecated: Use MMMetaData.ABI instead.
var MMABI = MMMetaData.ABI

// MM is an auto generated Go binding around an Ethereum contract.
type MM struct {
	MMCaller     // Read-only binding to the contract
	MMTransactor // Write-only binding to the contract
	MMFilterer   // Log filterer for contract events
}

// MMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MMSession struct {
	Contract     *MM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MMCallerSession struct {
	Contract *MMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MMTransactorSession struct {
	Contract     *MMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MMRaw struct {
	Contract *MM // Generic contract binding to access the raw methods on
}

// MMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MMCallerRaw struct {
	Contract *MMCaller // Generic read-only contract binding to access the raw methods on
}

// MMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MMTransactorRaw struct {
	Contract *MMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMM creates a new instance of MM, bound to a specific deployed contract.
func NewMM(address common.Address, backend bind.ContractBackend) (*MM, error) {
	contract, err := bindMM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MM{MMCaller: MMCaller{contract: contract}, MMTransactor: MMTransactor{contract: contract}, MMFilterer: MMFilterer{contract: contract}}, nil
}

// NewMMCaller creates a new read-only instance of MM, bound to a specific deployed contract.
func NewMMCaller(address common.Address, caller bind.ContractCaller) (*MMCaller, error) {
	contract, err := bindMM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MMCaller{contract: contract}, nil
}

// NewMMTransactor creates a new write-only instance of MM, bound to a specific deployed contract.
func NewMMTransactor(address common.Address, transactor bind.ContractTransactor) (*MMTransactor, error) {
	contract, err := bindMM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MMTransactor{contract: contract}, nil
}

// NewMMFilterer creates a new log filterer instance of MM, bound to a specific deployed contract.
func NewMMFilterer(address common.Address, filterer bind.ContractFilterer) (*MMFilterer, error) {
	contract, err := bindMM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MMFilterer{contract: contract}, nil
}

// bindMM binds a generic wrapper to an already deployed contract.
func bindMM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MM *MMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MM.Contract.MMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MM *MMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MM.Contract.MMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MM *MMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MM.Contract.MMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MM *MMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MM *MMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MM *MMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MM.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _token) view returns(uint256)
func (_MM *MMCaller) BalanceOf(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "balanceOf", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _token) view returns(uint256)
func (_MM *MMSession) BalanceOf(_token common.Address) (*big.Int, error) {
	return _MM.Contract.BalanceOf(&_MM.CallOpts, _token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _token) view returns(uint256)
func (_MM *MMCallerSession) BalanceOf(_token common.Address) (*big.Int, error) {
	return _MM.Contract.BalanceOf(&_MM.CallOpts, _token)
}

// GetDepositByPoolId is a free data retrieval call binding the contract method 0xedb4e5ac.
//
// Solidity: function getDepositByPoolId(uint256 _pid) view returns(uint256, uint256, uint256)
func (_MM *MMCaller) GetDepositByPoolId(opts *bind.CallOpts, _pid *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "getDepositByPoolId", _pid)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetDepositByPoolId is a free data retrieval call binding the contract method 0xedb4e5ac.
//
// Solidity: function getDepositByPoolId(uint256 _pid) view returns(uint256, uint256, uint256)
func (_MM *MMSession) GetDepositByPoolId(_pid *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _MM.Contract.GetDepositByPoolId(&_MM.CallOpts, _pid)
}

// GetDepositByPoolId is a free data retrieval call binding the contract method 0xedb4e5ac.
//
// Solidity: function getDepositByPoolId(uint256 _pid) view returns(uint256, uint256, uint256)
func (_MM *MMCallerSession) GetDepositByPoolId(_pid *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _MM.Contract.GetDepositByPoolId(&_MM.CallOpts, _pid)
}

// GetLpBalance is a free data retrieval call binding the contract method 0x8421403d.
//
// Solidity: function getLpBalance() view returns(uint256)
func (_MM *MMCaller) GetLpBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "getLpBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpBalance is a free data retrieval call binding the contract method 0x8421403d.
//
// Solidity: function getLpBalance() view returns(uint256)
func (_MM *MMSession) GetLpBalance() (*big.Int, error) {
	return _MM.Contract.GetLpBalance(&_MM.CallOpts)
}

// GetLpBalance is a free data retrieval call binding the contract method 0x8421403d.
//
// Solidity: function getLpBalance() view returns(uint256)
func (_MM *MMCallerSession) GetLpBalance() (*big.Int, error) {
	return _MM.Contract.GetLpBalance(&_MM.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xd12acf73.
//
// Solidity: function isClaimed(address , address , uint256 ) view returns(bool)
func (_MM *MMCaller) IsClaimed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "isClaimed", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xd12acf73.
//
// Solidity: function isClaimed(address , address , uint256 ) view returns(bool)
func (_MM *MMSession) IsClaimed(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _MM.Contract.IsClaimed(&_MM.CallOpts, arg0, arg1, arg2)
}

// IsClaimed is a free data retrieval call binding the contract method 0xd12acf73.
//
// Solidity: function isClaimed(address , address , uint256 ) view returns(bool)
func (_MM *MMCallerSession) IsClaimed(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _MM.Contract.IsClaimed(&_MM.CallOpts, arg0, arg1, arg2)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MM *MMCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MM *MMSession) MerkleRoot() ([32]byte, error) {
	return _MM.Contract.MerkleRoot(&_MM.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MM *MMCallerSession) MerkleRoot() ([32]byte, error) {
	return _MM.Contract.MerkleRoot(&_MM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MM *MMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MM *MMSession) Owner() (common.Address, error) {
	return _MM.Contract.Owner(&_MM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MM *MMCallerSession) Owner() (common.Address, error) {
	return _MM.Contract.Owner(&_MM.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0xb7c7c4e0.
//
// Solidity: function addPool(uint256 _pid, uint256 lockDay, uint256 rewardRatio, uint256 decimal) returns()
func (_MM *MMTransactor) AddPool(opts *bind.TransactOpts, _pid *big.Int, lockDay *big.Int, rewardRatio *big.Int, decimal *big.Int) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "addPool", _pid, lockDay, rewardRatio, decimal)
}

// AddPool is a paid mutator transaction binding the contract method 0xb7c7c4e0.
//
// Solidity: function addPool(uint256 _pid, uint256 lockDay, uint256 rewardRatio, uint256 decimal) returns()
func (_MM *MMSession) AddPool(_pid *big.Int, lockDay *big.Int, rewardRatio *big.Int, decimal *big.Int) (*types.Transaction, error) {
	return _MM.Contract.AddPool(&_MM.TransactOpts, _pid, lockDay, rewardRatio, decimal)
}

// AddPool is a paid mutator transaction binding the contract method 0xb7c7c4e0.
//
// Solidity: function addPool(uint256 _pid, uint256 lockDay, uint256 rewardRatio, uint256 decimal) returns()
func (_MM *MMTransactorSession) AddPool(_pid *big.Int, lockDay *big.Int, rewardRatio *big.Int, decimal *big.Int) (*types.Transaction, error) {
	return _MM.Contract.AddPool(&_MM.TransactOpts, _pid, lockDay, rewardRatio, decimal)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _token, uint256 _amount, uint256 _id, bytes32[] _proofs) returns()
func (_MM *MMTransactor) Claim(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _id *big.Int, _proofs [][32]byte) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "claim", _token, _amount, _id, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _token, uint256 _amount, uint256 _id, bytes32[] _proofs) returns()
func (_MM *MMSession) Claim(_token common.Address, _amount *big.Int, _id *big.Int, _proofs [][32]byte) (*types.Transaction, error) {
	return _MM.Contract.Claim(&_MM.TransactOpts, _token, _amount, _id, _proofs)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _token, uint256 _amount, uint256 _id, bytes32[] _proofs) returns()
func (_MM *MMTransactorSession) Claim(_token common.Address, _amount *big.Int, _id *big.Int, _proofs [][32]byte) (*types.Transaction, error) {
	return _MM.Contract.Claim(&_MM.TransactOpts, _token, _amount, _id, _proofs)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MM *MMTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MM *MMSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MM.Contract.Deposit(&_MM.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MM *MMTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MM.Contract.Deposit(&_MM.TransactOpts, _pid, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MM *MMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MM *MMSession) RenounceOwnership() (*types.Transaction, error) {
	return _MM.Contract.RenounceOwnership(&_MM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MM *MMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MM.Contract.RenounceOwnership(&_MM.TransactOpts)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_MM *MMTransactor) SetMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "setMerkleRoot", _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_MM *MMSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _MM.Contract.SetMerkleRoot(&_MM.TransactOpts, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_MM *MMTransactorSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _MM.Contract.SetMerkleRoot(&_MM.TransactOpts, _merkleRoot)
}

// SysTransfer is a paid mutator transaction binding the contract method 0xb3f48f49.
//
// Solidity: function sysTransfer(address _token, uint256 _amount, uint256 _id) returns()
func (_MM *MMTransactor) SysTransfer(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "sysTransfer", _token, _amount, _id)
}

// SysTransfer is a paid mutator transaction binding the contract method 0xb3f48f49.
//
// Solidity: function sysTransfer(address _token, uint256 _amount, uint256 _id) returns()
func (_MM *MMSession) SysTransfer(_token common.Address, _amount *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _MM.Contract.SysTransfer(&_MM.TransactOpts, _token, _amount, _id)
}

// SysTransfer is a paid mutator transaction binding the contract method 0xb3f48f49.
//
// Solidity: function sysTransfer(address _token, uint256 _amount, uint256 _id) returns()
func (_MM *MMTransactorSession) SysTransfer(_token common.Address, _amount *big.Int, _id *big.Int) (*types.Transaction, error) {
	return _MM.Contract.SysTransfer(&_MM.TransactOpts, _token, _amount, _id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MM *MMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MM *MMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MM.Contract.TransferOwnership(&_MM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MM *MMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MM.Contract.TransferOwnership(&_MM.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MM *MMTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MM *MMSession) Withdraw() (*types.Transaction, error) {
	return _MM.Contract.Withdraw(&_MM.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MM *MMTransactorSession) Withdraw() (*types.Transaction, error) {
	return _MM.Contract.Withdraw(&_MM.TransactOpts)
}

// WithdrawAllLp is a paid mutator transaction binding the contract method 0x8c538ed4.
//
// Solidity: function withdrawAllLp() returns()
func (_MM *MMTransactor) WithdrawAllLp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "withdrawAllLp")
}

// WithdrawAllLp is a paid mutator transaction binding the contract method 0x8c538ed4.
//
// Solidity: function withdrawAllLp() returns()
func (_MM *MMSession) WithdrawAllLp() (*types.Transaction, error) {
	return _MM.Contract.WithdrawAllLp(&_MM.TransactOpts)
}

// WithdrawAllLp is a paid mutator transaction binding the contract method 0x8c538ed4.
//
// Solidity: function withdrawAllLp() returns()
func (_MM *MMTransactorSession) WithdrawAllLp() (*types.Transaction, error) {
	return _MM.Contract.WithdrawAllLp(&_MM.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_MM *MMTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _MM.contract.Transact(opts, "withdrawToken", _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_MM *MMSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _MM.Contract.WithdrawToken(&_MM.TransactOpts, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_MM *MMTransactorSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _MM.Contract.WithdrawToken(&_MM.TransactOpts, _token)
}

// MMDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MM contract.
type MMDepositIterator struct {
	Event *MMDeposit // Event containing the contract specifics and raw log

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
func (it *MMDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MMDeposit)
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
		it.Event = new(MMDeposit)
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
func (it *MMDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MMDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MMDeposit represents a Deposit event raised by the MM contract.
type MMDeposit struct {
	User          common.Address
	Pid           *big.Int
	DepositAmount *big.Int
	TicketAmount  *big.Int
	LockTimestamp *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 depositAmount, uint256 ticketAmount, uint256 lockTimestamp)
func (_MM *MMFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MMDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MM.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MMDepositIterator{contract: _MM.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 depositAmount, uint256 ticketAmount, uint256 lockTimestamp)
func (_MM *MMFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MMDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MM.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MMDeposit)
				if err := _MM.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 depositAmount, uint256 ticketAmount, uint256 lockTimestamp)
func (_MM *MMFilterer) ParseDeposit(log types.Log) (*MMDeposit, error) {
	event := new(MMDeposit)
	if err := _MM.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MM contract.
type MMOwnershipTransferredIterator struct {
	Event *MMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MMOwnershipTransferred)
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
		it.Event = new(MMOwnershipTransferred)
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
func (it *MMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MMOwnershipTransferred represents a OwnershipTransferred event raised by the MM contract.
type MMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MM *MMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MMOwnershipTransferredIterator{contract: _MM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MM *MMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MMOwnershipTransferred)
				if err := _MM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MM *MMFilterer) ParseOwnershipTransferred(log types.Log) (*MMOwnershipTransferred, error) {
	event := new(MMOwnershipTransferred)
	if err := _MM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MMWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MM contract.
type MMWithdrawIterator struct {
	Event *MMWithdraw // Event containing the contract specifics and raw log

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
func (it *MMWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MMWithdraw)
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
		it.Event = new(MMWithdraw)
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
func (it *MMWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MMWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MMWithdraw represents a Withdraw event raised by the MM contract.
type MMWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_MM *MMFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*MMWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MM.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &MMWithdrawIterator{contract: _MM.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_MM *MMFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MMWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MM.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MMWithdraw)
				if err := _MM.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_MM *MMFilterer) ParseWithdraw(log types.Log) (*MMWithdraw, error) {
	event := new(MMWithdraw)
	if err := _MM.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
