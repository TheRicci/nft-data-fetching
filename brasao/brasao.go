// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brasao

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
)

// BrasaoMetaData contains all meta data concerning the Brasao contract.
var BrasaoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_transferSend\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"endUser\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"endUser\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"claimBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupplyMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BrasaoABI is the input ABI used to generate the binding from.
// Deprecated: Use BrasaoMetaData.ABI instead.
var BrasaoABI = BrasaoMetaData.ABI

// Brasao is an auto generated Go binding around an Ethereum contract.
type Brasao struct {
	BrasaoCaller     // Read-only binding to the contract
	BrasaoTransactor // Write-only binding to the contract
	BrasaoFilterer   // Log filterer for contract events
}

// BrasaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrasaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrasaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrasaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrasaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrasaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrasaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrasaoSession struct {
	Contract     *Brasao           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrasaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrasaoCallerSession struct {
	Contract *BrasaoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrasaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrasaoTransactorSession struct {
	Contract     *BrasaoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrasaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrasaoRaw struct {
	Contract *Brasao // Generic contract binding to access the raw methods on
}

// BrasaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrasaoCallerRaw struct {
	Contract *BrasaoCaller // Generic read-only contract binding to access the raw methods on
}

// BrasaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrasaoTransactorRaw struct {
	Contract *BrasaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrasao creates a new instance of Brasao, bound to a specific deployed contract.
func NewBrasao(address common.Address, backend bind.ContractBackend) (*Brasao, error) {
	contract, err := bindBrasao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brasao{BrasaoCaller: BrasaoCaller{contract: contract}, BrasaoTransactor: BrasaoTransactor{contract: contract}, BrasaoFilterer: BrasaoFilterer{contract: contract}}, nil
}

// NewBrasaoCaller creates a new read-only instance of Brasao, bound to a specific deployed contract.
func NewBrasaoCaller(address common.Address, caller bind.ContractCaller) (*BrasaoCaller, error) {
	contract, err := bindBrasao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrasaoCaller{contract: contract}, nil
}

// NewBrasaoTransactor creates a new write-only instance of Brasao, bound to a specific deployed contract.
func NewBrasaoTransactor(address common.Address, transactor bind.ContractTransactor) (*BrasaoTransactor, error) {
	contract, err := bindBrasao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrasaoTransactor{contract: contract}, nil
}

// NewBrasaoFilterer creates a new log filterer instance of Brasao, bound to a specific deployed contract.
func NewBrasaoFilterer(address common.Address, filterer bind.ContractFilterer) (*BrasaoFilterer, error) {
	contract, err := bindBrasao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrasaoFilterer{contract: contract}, nil
}

// bindBrasao binds a generic wrapper to an already deployed contract.
func bindBrasao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrasaoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brasao *BrasaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brasao.Contract.BrasaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brasao *BrasaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brasao.Contract.BrasaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brasao *BrasaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brasao.Contract.BrasaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brasao *BrasaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brasao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brasao *BrasaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brasao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brasao *BrasaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brasao.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Brasao *BrasaoCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Brasao *BrasaoSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Brasao.Contract.BalanceOf(&_Brasao.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Brasao *BrasaoCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Brasao.Contract.BalanceOf(&_Brasao.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Brasao *BrasaoCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Brasao *BrasaoSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Brasao.Contract.BalanceOfBatch(&_Brasao.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Brasao *BrasaoCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Brasao.Contract.BalanceOfBatch(&_Brasao.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brasao *BrasaoCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brasao *BrasaoSession) ContractURI() (string, error) {
	return _Brasao.Contract.ContractURI(&_Brasao.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brasao *BrasaoCallerSession) ContractURI() (string, error) {
	return _Brasao.Contract.ContractURI(&_Brasao.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brasao *BrasaoCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brasao *BrasaoSession) GetBalance() (*big.Int, error) {
	return _Brasao.Contract.GetBalance(&_Brasao.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brasao *BrasaoCallerSession) GetBalance() (*big.Int, error) {
	return _Brasao.Contract.GetBalance(&_Brasao.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brasao *BrasaoCaller) GetBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "getBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brasao *BrasaoSession) GetBaseURI() (string, error) {
	return _Brasao.Contract.GetBaseURI(&_Brasao.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brasao *BrasaoCallerSession) GetBaseURI() (string, error) {
	return _Brasao.Contract.GetBaseURI(&_Brasao.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brasao *BrasaoCaller) GetMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "getMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brasao *BrasaoSession) GetMaxSupply() (*big.Int, error) {
	return _Brasao.Contract.GetMaxSupply(&_Brasao.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brasao *BrasaoCallerSession) GetMaxSupply() (*big.Int, error) {
	return _Brasao.Contract.GetMaxSupply(&_Brasao.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Brasao *BrasaoCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Brasao *BrasaoSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Brasao.Contract.IsApprovedForAll(&_Brasao.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Brasao *BrasaoCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Brasao.Contract.IsApprovedForAll(&_Brasao.CallOpts, account, operator)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brasao *BrasaoCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brasao *BrasaoSession) IsPaused() (bool, error) {
	return _Brasao.Contract.IsPaused(&_Brasao.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brasao *BrasaoCallerSession) IsPaused() (bool, error) {
	return _Brasao.Contract.IsPaused(&_Brasao.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brasao *BrasaoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brasao *BrasaoSession) Name() (string, error) {
	return _Brasao.Contract.Name(&_Brasao.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brasao *BrasaoCallerSession) Name() (string, error) {
	return _Brasao.Contract.Name(&_Brasao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brasao *BrasaoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brasao *BrasaoSession) Owner() (common.Address, error) {
	return _Brasao.Contract.Owner(&_Brasao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brasao *BrasaoCallerSession) Owner() (common.Address, error) {
	return _Brasao.Contract.Owner(&_Brasao.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brasao *BrasaoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brasao *BrasaoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Brasao.Contract.SupportsInterface(&_Brasao.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brasao *BrasaoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Brasao.Contract.SupportsInterface(&_Brasao.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brasao *BrasaoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brasao *BrasaoSession) Symbol() (string, error) {
	return _Brasao.Contract.Symbol(&_Brasao.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brasao *BrasaoCallerSession) Symbol() (string, error) {
	return _Brasao.Contract.Symbol(&_Brasao.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brasao *BrasaoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brasao *BrasaoSession) TotalSupply() (*big.Int, error) {
	return _Brasao.Contract.TotalSupply(&_Brasao.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brasao *BrasaoCallerSession) TotalSupply() (*big.Int, error) {
	return _Brasao.Contract.TotalSupply(&_Brasao.CallOpts)
}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Brasao *BrasaoCaller) TotalSupplyMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "totalSupplyMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Brasao *BrasaoSession) TotalSupplyMint() (*big.Int, error) {
	return _Brasao.Contract.TotalSupplyMint(&_Brasao.CallOpts)
}

// TotalSupplyMint is a free data retrieval call binding the contract method 0xf5b0778b.
//
// Solidity: function totalSupplyMint() view returns(uint256)
func (_Brasao *BrasaoCallerSession) TotalSupplyMint() (*big.Int, error) {
	return _Brasao.Contract.TotalSupplyMint(&_Brasao.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Brasao *BrasaoCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Brasao *BrasaoSession) Uri(arg0 *big.Int) (string, error) {
	return _Brasao.Contract.Uri(&_Brasao.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Brasao *BrasaoCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Brasao.Contract.Uri(&_Brasao.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brasao *BrasaoCaller) WhitelistClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "whitelistClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brasao *BrasaoSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Brasao.Contract.WhitelistClaimed(&_Brasao.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brasao *BrasaoCallerSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Brasao.Contract.WhitelistClaimed(&_Brasao.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brasao *BrasaoCaller) WhitelistedAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Brasao.contract.Call(opts, &out, "whitelistedAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brasao *BrasaoSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Brasao.Contract.WhitelistedAddresses(&_Brasao.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brasao *BrasaoCallerSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Brasao.Contract.WhitelistedAddresses(&_Brasao.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Brasao *BrasaoTransactor) Claim(opts *bind.TransactOpts, _mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "claim", _mintAmount, _number, endUser)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Brasao *BrasaoSession) Claim(_mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Brasao.Contract.Claim(&_Brasao.TransactOpts, _mintAmount, _number, endUser)
}

// Claim is a paid mutator transaction binding the contract method 0xac44ff31.
//
// Solidity: function claim(uint256 _mintAmount, uint256 _number, address endUser) returns()
func (_Brasao *BrasaoTransactorSession) Claim(_mintAmount *big.Int, _number *big.Int, endUser common.Address) (*types.Transaction, error) {
	return _Brasao.Contract.Claim(&_Brasao.TransactOpts, _mintAmount, _number, endUser)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Brasao *BrasaoTransactor) ClaimBatch(opts *bind.TransactOpts, endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "claimBatch", endUser, amount)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Brasao *BrasaoSession) ClaimBatch(endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Brasao.Contract.ClaimBatch(&_Brasao.TransactOpts, endUser, amount)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0xd7c2dfc8.
//
// Solidity: function claimBatch(address[] endUser, uint256[] amount) returns()
func (_Brasao *BrasaoTransactorSession) ClaimBatch(endUser []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Brasao.Contract.ClaimBatch(&_Brasao.TransactOpts, endUser, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brasao *BrasaoTransactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brasao *BrasaoSession) Pause(_state bool) (*types.Transaction, error) {
	return _Brasao.Contract.Pause(&_Brasao.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brasao *BrasaoTransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Brasao.Contract.Pause(&_Brasao.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brasao *BrasaoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brasao *BrasaoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brasao.Contract.RenounceOwnership(&_Brasao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brasao *BrasaoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brasao.Contract.RenounceOwnership(&_Brasao.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Brasao *BrasaoTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Brasao *BrasaoSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.Contract.SafeBatchTransferFrom(&_Brasao.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Brasao *BrasaoTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.Contract.SafeBatchTransferFrom(&_Brasao.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Brasao *BrasaoTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Brasao *BrasaoSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.Contract.SafeTransferFrom(&_Brasao.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Brasao *BrasaoTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Brasao.Contract.SafeTransferFrom(&_Brasao.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brasao *BrasaoTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brasao *BrasaoSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brasao.Contract.SetApprovalForAll(&_Brasao.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brasao *BrasaoTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brasao.Contract.SetApprovalForAll(&_Brasao.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brasao *BrasaoTransactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brasao *BrasaoSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Brasao.Contract.SetBaseURI(&_Brasao.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brasao *BrasaoTransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Brasao.Contract.SetBaseURI(&_Brasao.TransactOpts, _newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brasao *BrasaoTransactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brasao *BrasaoSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Brasao.Contract.SetContractURI(&_Brasao.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brasao *BrasaoTransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Brasao.Contract.SetContractURI(&_Brasao.TransactOpts, contractURI_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brasao *BrasaoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brasao *BrasaoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brasao.Contract.TransferOwnership(&_Brasao.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brasao *BrasaoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brasao.Contract.TransferOwnership(&_Brasao.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brasao *BrasaoTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brasao.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brasao *BrasaoSession) Withdraw() (*types.Transaction, error) {
	return _Brasao.Contract.Withdraw(&_Brasao.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brasao *BrasaoTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Brasao.Contract.Withdraw(&_Brasao.TransactOpts)
}

// BrasaoApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Brasao contract.
type BrasaoApprovalForAllIterator struct {
	Event *BrasaoApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BrasaoApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoApprovalForAll)
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
		it.Event = new(BrasaoApprovalForAll)
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
func (it *BrasaoApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoApprovalForAll represents a ApprovalForAll event raised by the Brasao contract.
type BrasaoApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Brasao *BrasaoFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*BrasaoApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BrasaoApprovalForAllIterator{contract: _Brasao.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Brasao *BrasaoFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BrasaoApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoApprovalForAll)
				if err := _Brasao.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Brasao *BrasaoFilterer) ParseApprovalForAll(log types.Log) (*BrasaoApprovalForAll, error) {
	event := new(BrasaoApprovalForAll)
	if err := _Brasao.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrasaoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Brasao contract.
type BrasaoOwnershipTransferredIterator struct {
	Event *BrasaoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrasaoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoOwnershipTransferred)
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
		it.Event = new(BrasaoOwnershipTransferred)
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
func (it *BrasaoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoOwnershipTransferred represents a OwnershipTransferred event raised by the Brasao contract.
type BrasaoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brasao *BrasaoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrasaoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrasaoOwnershipTransferredIterator{contract: _Brasao.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brasao *BrasaoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrasaoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoOwnershipTransferred)
				if err := _Brasao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Brasao *BrasaoFilterer) ParseOwnershipTransferred(log types.Log) (*BrasaoOwnershipTransferred, error) {
	event := new(BrasaoOwnershipTransferred)
	if err := _Brasao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrasaoTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Brasao contract.
type BrasaoTransferBatchIterator struct {
	Event *BrasaoTransferBatch // Event containing the contract specifics and raw log

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
func (it *BrasaoTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoTransferBatch)
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
		it.Event = new(BrasaoTransferBatch)
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
func (it *BrasaoTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoTransferBatch represents a TransferBatch event raised by the Brasao contract.
type BrasaoTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Brasao *BrasaoFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BrasaoTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrasaoTransferBatchIterator{contract: _Brasao.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Brasao *BrasaoFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *BrasaoTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoTransferBatch)
				if err := _Brasao.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Brasao *BrasaoFilterer) ParseTransferBatch(log types.Log) (*BrasaoTransferBatch, error) {
	event := new(BrasaoTransferBatch)
	if err := _Brasao.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrasaoTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Brasao contract.
type BrasaoTransferSingleIterator struct {
	Event *BrasaoTransferSingle // Event containing the contract specifics and raw log

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
func (it *BrasaoTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoTransferSingle)
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
		it.Event = new(BrasaoTransferSingle)
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
func (it *BrasaoTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoTransferSingle represents a TransferSingle event raised by the Brasao contract.
type BrasaoTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Brasao *BrasaoFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BrasaoTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrasaoTransferSingleIterator{contract: _Brasao.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Brasao *BrasaoFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *BrasaoTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoTransferSingle)
				if err := _Brasao.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Brasao *BrasaoFilterer) ParseTransferSingle(log types.Log) (*BrasaoTransferSingle, error) {
	event := new(BrasaoTransferSingle)
	if err := _Brasao.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrasaoURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Brasao contract.
type BrasaoURIIterator struct {
	Event *BrasaoURI // Event containing the contract specifics and raw log

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
func (it *BrasaoURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoURI)
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
		it.Event = new(BrasaoURI)
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
func (it *BrasaoURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoURI represents a URI event raised by the Brasao contract.
type BrasaoURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Brasao *BrasaoFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*BrasaoURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &BrasaoURIIterator{contract: _Brasao.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Brasao *BrasaoFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *BrasaoURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoURI)
				if err := _Brasao.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Brasao *BrasaoFilterer) ParseURI(log types.Log) (*BrasaoURI, error) {
	event := new(BrasaoURI)
	if err := _Brasao.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrasaoTransferSendIterator is returned from FilterTransferSend and is used to iterate over the raw logs and unpacked data for TransferSend events raised by the Brasao contract.
type BrasaoTransferSendIterator struct {
	Event *BrasaoTransferSend // Event containing the contract specifics and raw log

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
func (it *BrasaoTransferSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrasaoTransferSend)
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
		it.Event = new(BrasaoTransferSend)
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
func (it *BrasaoTransferSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrasaoTransferSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrasaoTransferSend represents a TransferSend event raised by the Brasao contract.
type BrasaoTransferSend struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSend is a free log retrieval operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Brasao *BrasaoFilterer) FilterTransferSend(opts *bind.FilterOpts) (*BrasaoTransferSendIterator, error) {

	logs, sub, err := _Brasao.contract.FilterLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return &BrasaoTransferSendIterator{contract: _Brasao.contract, event: "_transferSend", logs: logs, sub: sub}, nil
}

// WatchTransferSend is a free log subscription operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Brasao *BrasaoFilterer) WatchTransferSend(opts *bind.WatchOpts, sink chan<- *BrasaoTransferSend) (event.Subscription, error) {

	logs, sub, err := _Brasao.contract.WatchLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrasaoTransferSend)
				if err := _Brasao.contract.UnpackLog(event, "_transferSend", log); err != nil {
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

// ParseTransferSend is a log parse operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Brasao *BrasaoFilterer) ParseTransferSend(log types.Log) (*BrasaoTransferSend, error) {
	event := new(BrasaoTransferSend)
	if err := _Brasao.contract.UnpackLog(event, "_transferSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
