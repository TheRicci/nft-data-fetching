// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brazuera

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

// BrazueraMetaData contains all meta data concerning the Brazuera contract.
var BrazueraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_initBaseURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"_transferSend\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"_checkContractOnERC721Received\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"_transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractIsFreeMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getAddressLastMintedTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getNftEtherValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQtdAvailableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"isAffliliated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnlyWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recommendedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_indicationType\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"endUser\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"setAffiliateProgramPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setAllowAffiliateProgram\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setFreeMintPerAddressLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setIsFreeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMintAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftEtherValue\",\"type\":\"uint256\"}],\"name\":\"setNftEtherValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setOnlyWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recommendedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_indicationType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"split\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseExtension\",\"type\":\"string\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BrazueraABI is the input ABI used to generate the binding from.
// Deprecated: Use BrazueraMetaData.ABI instead.
var BrazueraABI = BrazueraMetaData.ABI

// Brazuera is an auto generated Go binding around an Ethereum contract.
type Brazuera struct {
	BrazueraCaller     // Read-only binding to the contract
	BrazueraTransactor // Write-only binding to the contract
	BrazueraFilterer   // Log filterer for contract events
}

// BrazueraCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrazueraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrazueraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrazueraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrazueraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrazueraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrazueraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrazueraSession struct {
	Contract     *Brazuera         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrazueraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrazueraCallerSession struct {
	Contract *BrazueraCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BrazueraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrazueraTransactorSession struct {
	Contract     *BrazueraTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BrazueraRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrazueraRaw struct {
	Contract *Brazuera // Generic contract binding to access the raw methods on
}

// BrazueraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrazueraCallerRaw struct {
	Contract *BrazueraCaller // Generic read-only contract binding to access the raw methods on
}

// BrazueraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrazueraTransactorRaw struct {
	Contract *BrazueraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrazuera creates a new instance of Brazuera, bound to a specific deployed contract.
func NewBrazuera(address common.Address, backend bind.ContractBackend) (*Brazuera, error) {
	contract, err := bindBrazuera(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brazuera{BrazueraCaller: BrazueraCaller{contract: contract}, BrazueraTransactor: BrazueraTransactor{contract: contract}, BrazueraFilterer: BrazueraFilterer{contract: contract}}, nil
}

// NewBrazueraCaller creates a new read-only instance of Brazuera, bound to a specific deployed contract.
func NewBrazueraCaller(address common.Address, caller bind.ContractCaller) (*BrazueraCaller, error) {
	contract, err := bindBrazuera(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrazueraCaller{contract: contract}, nil
}

// NewBrazueraTransactor creates a new write-only instance of Brazuera, bound to a specific deployed contract.
func NewBrazueraTransactor(address common.Address, transactor bind.ContractTransactor) (*BrazueraTransactor, error) {
	contract, err := bindBrazuera(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrazueraTransactor{contract: contract}, nil
}

// NewBrazueraFilterer creates a new log filterer instance of Brazuera, bound to a specific deployed contract.
func NewBrazueraFilterer(address common.Address, filterer bind.ContractFilterer) (*BrazueraFilterer, error) {
	contract, err := bindBrazuera(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrazueraFilterer{contract: contract}, nil
}

// bindBrazuera binds a generic wrapper to an already deployed contract.
func bindBrazuera(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrazueraABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brazuera *BrazueraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brazuera.Contract.BrazueraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brazuera *BrazueraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brazuera.Contract.BrazueraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brazuera *BrazueraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brazuera.Contract.BrazueraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brazuera *BrazueraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brazuera.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brazuera *BrazueraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brazuera.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brazuera *BrazueraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brazuera.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Brazuera *BrazueraCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Brazuera *BrazueraSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Brazuera.Contract.BalanceOf(&_Brazuera.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Brazuera *BrazueraCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Brazuera.Contract.BalanceOf(&_Brazuera.CallOpts, owner)
}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Brazuera *BrazueraCaller) ContractIsFreeMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "contractIsFreeMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Brazuera *BrazueraSession) ContractIsFreeMint() (bool, error) {
	return _Brazuera.Contract.ContractIsFreeMint(&_Brazuera.CallOpts)
}

// ContractIsFreeMint is a free data retrieval call binding the contract method 0x4fc9f84a.
//
// Solidity: function contractIsFreeMint() view returns(bool)
func (_Brazuera *BrazueraCallerSession) ContractIsFreeMint() (bool, error) {
	return _Brazuera.Contract.ContractIsFreeMint(&_Brazuera.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brazuera *BrazueraCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brazuera *BrazueraSession) ContractURI() (string, error) {
	return _Brazuera.Contract.ContractURI(&_Brazuera.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Brazuera *BrazueraCallerSession) ContractURI() (string, error) {
	return _Brazuera.Contract.ContractURI(&_Brazuera.CallOpts)
}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Brazuera *BrazueraCaller) GetAddressLastMintedTokenId(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getAddressLastMintedTokenId", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Brazuera *BrazueraSession) GetAddressLastMintedTokenId(wallet common.Address) (*big.Int, error) {
	return _Brazuera.Contract.GetAddressLastMintedTokenId(&_Brazuera.CallOpts, wallet)
}

// GetAddressLastMintedTokenId is a free data retrieval call binding the contract method 0x6c21fb00.
//
// Solidity: function getAddressLastMintedTokenId(address wallet) view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetAddressLastMintedTokenId(wallet common.Address) (*big.Int, error) {
	return _Brazuera.Contract.GetAddressLastMintedTokenId(&_Brazuera.CallOpts, wallet)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Brazuera.Contract.GetApproved(&_Brazuera.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Brazuera.Contract.GetApproved(&_Brazuera.CallOpts, tokenId)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brazuera *BrazueraCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brazuera *BrazueraSession) GetBalance() (*big.Int, error) {
	return _Brazuera.Contract.GetBalance(&_Brazuera.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetBalance() (*big.Int, error) {
	return _Brazuera.Contract.GetBalance(&_Brazuera.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brazuera *BrazueraCaller) GetBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brazuera *BrazueraSession) GetBaseURI() (string, error) {
	return _Brazuera.Contract.GetBaseURI(&_Brazuera.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_Brazuera *BrazueraCallerSession) GetBaseURI() (string, error) {
	return _Brazuera.Contract.GetBaseURI(&_Brazuera.CallOpts)
}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Brazuera *BrazueraCaller) GetMaxMintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getMaxMintAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Brazuera *BrazueraSession) GetMaxMintAmount() (*big.Int, error) {
	return _Brazuera.Contract.GetMaxMintAmount(&_Brazuera.CallOpts)
}

// GetMaxMintAmount is a free data retrieval call binding the contract method 0xd6e7b8e4.
//
// Solidity: function getMaxMintAmount() view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetMaxMintAmount() (*big.Int, error) {
	return _Brazuera.Contract.GetMaxMintAmount(&_Brazuera.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brazuera *BrazueraCaller) GetMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brazuera *BrazueraSession) GetMaxSupply() (*big.Int, error) {
	return _Brazuera.Contract.GetMaxSupply(&_Brazuera.CallOpts)
}

// GetMaxSupply is a free data retrieval call binding the contract method 0x4c0f38c2.
//
// Solidity: function getMaxSupply() view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetMaxSupply() (*big.Int, error) {
	return _Brazuera.Contract.GetMaxSupply(&_Brazuera.CallOpts)
}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraCaller) GetNFTURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getNFTURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraSession) GetNFTURI(tokenId *big.Int) (string, error) {
	return _Brazuera.Contract.GetNFTURI(&_Brazuera.CallOpts, tokenId)
}

// GetNFTURI is a free data retrieval call binding the contract method 0x00621fda.
//
// Solidity: function getNFTURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraCallerSession) GetNFTURI(tokenId *big.Int) (string, error) {
	return _Brazuera.Contract.GetNFTURI(&_Brazuera.CallOpts, tokenId)
}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Brazuera *BrazueraCaller) GetNftEtherValue(opts *bind.CallOpts, wallet common.Address, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getNftEtherValue", wallet, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Brazuera *BrazueraSession) GetNftEtherValue(wallet common.Address, proof [][32]byte) (*big.Int, error) {
	return _Brazuera.Contract.GetNftEtherValue(&_Brazuera.CallOpts, wallet, proof)
}

// GetNftEtherValue is a free data retrieval call binding the contract method 0x77dbd17f.
//
// Solidity: function getNftEtherValue(address wallet, bytes32[] proof) view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetNftEtherValue(wallet common.Address, proof [][32]byte) (*big.Int, error) {
	return _Brazuera.Contract.GetNftEtherValue(&_Brazuera.CallOpts, wallet, proof)
}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Brazuera *BrazueraCaller) GetQtdAvailableTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "getQtdAvailableTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Brazuera *BrazueraSession) GetQtdAvailableTokens() (*big.Int, error) {
	return _Brazuera.Contract.GetQtdAvailableTokens(&_Brazuera.CallOpts)
}

// GetQtdAvailableTokens is a free data retrieval call binding the contract method 0xb4ed16b3.
//
// Solidity: function getQtdAvailableTokens() view returns(uint256)
func (_Brazuera *BrazueraCallerSession) GetQtdAvailableTokens() (*big.Int, error) {
	return _Brazuera.Contract.GetQtdAvailableTokens(&_Brazuera.CallOpts)
}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Brazuera *BrazueraCaller) IsAffliliated(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isAffliliated", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Brazuera *BrazueraSession) IsAffliliated(wallet common.Address) (bool, error) {
	return _Brazuera.Contract.IsAffliliated(&_Brazuera.CallOpts, wallet)
}

// IsAffliliated is a free data retrieval call binding the contract method 0x9b0e665d.
//
// Solidity: function isAffliliated(address wallet) view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsAffliliated(wallet common.Address) (bool, error) {
	return _Brazuera.Contract.IsAffliliated(&_Brazuera.CallOpts, wallet)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Brazuera *BrazueraCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Brazuera *BrazueraSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Brazuera.Contract.IsApprovedForAll(&_Brazuera.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Brazuera.Contract.IsApprovedForAll(&_Brazuera.CallOpts, owner, operator)
}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Brazuera *BrazueraCaller) IsOnlyWhitelist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isOnlyWhitelist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Brazuera *BrazueraSession) IsOnlyWhitelist() (bool, error) {
	return _Brazuera.Contract.IsOnlyWhitelist(&_Brazuera.CallOpts)
}

// IsOnlyWhitelist is a free data retrieval call binding the contract method 0x6f50c93c.
//
// Solidity: function isOnlyWhitelist() view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsOnlyWhitelist() (bool, error) {
	return _Brazuera.Contract.IsOnlyWhitelist(&_Brazuera.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brazuera *BrazueraCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brazuera *BrazueraSession) IsPaused() (bool, error) {
	return _Brazuera.Contract.IsPaused(&_Brazuera.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsPaused() (bool, error) {
	return _Brazuera.Contract.IsPaused(&_Brazuera.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Brazuera *BrazueraCaller) IsValid(opts *bind.CallOpts, proof [][32]byte, leaf [32]byte) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isValid", proof, leaf)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Brazuera *BrazueraSession) IsValid(proof [][32]byte, leaf [32]byte) (bool, error) {
	return _Brazuera.Contract.IsValid(&_Brazuera.CallOpts, proof, leaf)
}

// IsValid is a free data retrieval call binding the contract method 0xb8a20ed0.
//
// Solidity: function isValid(bytes32[] proof, bytes32 leaf) view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsValid(proof [][32]byte, leaf [32]byte) (bool, error) {
	return _Brazuera.Contract.IsValid(&_Brazuera.CallOpts, proof, leaf)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Brazuera *BrazueraCaller) IsWhitelisted(opts *bind.CallOpts, _user common.Address, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "isWhitelisted", _user, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Brazuera *BrazueraSession) IsWhitelisted(_user common.Address, proof [][32]byte) (bool, error) {
	return _Brazuera.Contract.IsWhitelisted(&_Brazuera.CallOpts, _user, proof)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x5a23dd99.
//
// Solidity: function isWhitelisted(address _user, bytes32[] proof) view returns(bool)
func (_Brazuera *BrazueraCallerSession) IsWhitelisted(_user common.Address, proof [][32]byte) (bool, error) {
	return _Brazuera.Contract.IsWhitelisted(&_Brazuera.CallOpts, _user, proof)
}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Brazuera *BrazueraCaller) MerkleProof(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "merkleProof", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Brazuera *BrazueraSession) MerkleProof(arg0 *big.Int) ([32]byte, error) {
	return _Brazuera.Contract.MerkleProof(&_Brazuera.CallOpts, arg0)
}

// MerkleProof is a free data retrieval call binding the contract method 0xfa9a309d.
//
// Solidity: function merkleProof(uint256 ) view returns(bytes32)
func (_Brazuera *BrazueraCallerSession) MerkleProof(arg0 *big.Int) ([32]byte, error) {
	return _Brazuera.Contract.MerkleProof(&_Brazuera.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brazuera *BrazueraCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brazuera *BrazueraSession) Name() (string, error) {
	return _Brazuera.Contract.Name(&_Brazuera.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brazuera *BrazueraCallerSession) Name() (string, error) {
	return _Brazuera.Contract.Name(&_Brazuera.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brazuera *BrazueraCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brazuera *BrazueraSession) Owner() (common.Address, error) {
	return _Brazuera.Contract.Owner(&_Brazuera.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brazuera *BrazueraCallerSession) Owner() (common.Address, error) {
	return _Brazuera.Contract.Owner(&_Brazuera.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Brazuera.Contract.OwnerOf(&_Brazuera.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Brazuera *BrazueraCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Brazuera.Contract.OwnerOf(&_Brazuera.CallOpts, tokenId)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Brazuera *BrazueraCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Brazuera *BrazueraSession) Root() ([32]byte, error) {
	return _Brazuera.Contract.Root(&_Brazuera.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Brazuera *BrazueraCallerSession) Root() ([32]byte, error) {
	return _Brazuera.Contract.Root(&_Brazuera.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brazuera *BrazueraCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brazuera *BrazueraSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Brazuera.Contract.SupportsInterface(&_Brazuera.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Brazuera *BrazueraCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Brazuera.Contract.SupportsInterface(&_Brazuera.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brazuera *BrazueraCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brazuera *BrazueraSession) Symbol() (string, error) {
	return _Brazuera.Contract.Symbol(&_Brazuera.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brazuera *BrazueraCallerSession) Symbol() (string, error) {
	return _Brazuera.Contract.Symbol(&_Brazuera.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Brazuera *BrazueraCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int, baseExtension string) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "tokenURI", tokenId, baseExtension)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Brazuera *BrazueraSession) TokenURI(tokenId *big.Int, baseExtension string) (string, error) {
	return _Brazuera.Contract.TokenURI(&_Brazuera.CallOpts, tokenId, baseExtension)
}

// TokenURI is a free data retrieval call binding the contract method 0x41dcf454.
//
// Solidity: function tokenURI(uint256 tokenId, string baseExtension) view returns(string)
func (_Brazuera *BrazueraCallerSession) TokenURI(tokenId *big.Int, baseExtension string) (string, error) {
	return _Brazuera.Contract.TokenURI(&_Brazuera.CallOpts, tokenId, baseExtension)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraCaller) TokenURI0(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "tokenURI0", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraSession) TokenURI0(tokenId *big.Int) (string, error) {
	return _Brazuera.Contract.TokenURI0(&_Brazuera.CallOpts, tokenId)
}

// TokenURI0 is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Brazuera *BrazueraCallerSession) TokenURI0(tokenId *big.Int) (string, error) {
	return _Brazuera.Contract.TokenURI0(&_Brazuera.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brazuera *BrazueraCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brazuera *BrazueraSession) TotalSupply() (*big.Int, error) {
	return _Brazuera.Contract.TotalSupply(&_Brazuera.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brazuera *BrazueraCallerSession) TotalSupply() (*big.Int, error) {
	return _Brazuera.Contract.TotalSupply(&_Brazuera.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Brazuera *BrazueraCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Brazuera *BrazueraSession) WalletOfOwner(_owner common.Address) (*big.Int, error) {
	return _Brazuera.Contract.WalletOfOwner(&_Brazuera.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256)
func (_Brazuera *BrazueraCallerSession) WalletOfOwner(_owner common.Address) (*big.Int, error) {
	return _Brazuera.Contract.WalletOfOwner(&_Brazuera.CallOpts, _owner)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brazuera *BrazueraCaller) WhitelistClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "whitelistClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brazuera *BrazueraSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Brazuera.Contract.WhitelistClaimed(&_Brazuera.CallOpts, arg0)
}

// WhitelistClaimed is a free data retrieval call binding the contract method 0xdb4bec44.
//
// Solidity: function whitelistClaimed(address ) view returns(bool)
func (_Brazuera *BrazueraCallerSession) WhitelistClaimed(arg0 common.Address) (bool, error) {
	return _Brazuera.Contract.WhitelistClaimed(&_Brazuera.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brazuera *BrazueraCaller) WhitelistedAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Brazuera.contract.Call(opts, &out, "whitelistedAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brazuera *BrazueraSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Brazuera.Contract.WhitelistedAddresses(&_Brazuera.CallOpts, arg0)
}

// WhitelistedAddresses is a free data retrieval call binding the contract method 0xba4e5c49.
//
// Solidity: function whitelistedAddresses(uint256 ) view returns(address)
func (_Brazuera *BrazueraCallerSession) WhitelistedAddresses(arg0 *big.Int) (common.Address, error) {
	return _Brazuera.Contract.WhitelistedAddresses(&_Brazuera.CallOpts, arg0)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Brazuera *BrazueraTransactor) CheckContractOnERC721Received(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "_checkContractOnERC721Received", from, to, tokenId, _data)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Brazuera *BrazueraSession) CheckContractOnERC721Received(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.Contract.CheckContractOnERC721Received(&_Brazuera.TransactOpts, from, to, tokenId, _data)
}

// CheckContractOnERC721Received is a paid mutator transaction binding the contract method 0xd88343e2.
//
// Solidity: function _checkContractOnERC721Received(address from, address to, uint256 tokenId, bytes _data) returns(bool)
func (_Brazuera *BrazueraTransactorSession) CheckContractOnERC721Received(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.Contract.CheckContractOnERC721Received(&_Brazuera.TransactOpts, from, to, tokenId, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "_transfer", from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraSession) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.Transfer(&_Brazuera.TransactOpts, from, to, tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0x30e0789e.
//
// Solidity: function _transfer(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactorSession) Transfer(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.Transfer(&_Brazuera.TransactOpts, from, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.Approve(&_Brazuera.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.Approve(&_Brazuera.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraTransactor) Mint(opts *bind.TransactOpts, _mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "mint", _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraSession) Mint(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.Contract.Mint(&_Brazuera.TransactOpts, _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x9293237d.
//
// Solidity: function mint(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, address endUser, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraTransactorSession) Mint(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, endUser common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.Contract.Mint(&_Brazuera.TransactOpts, _mintAmount, _recommendedBy, _indicationType, endUser, proof)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brazuera *BrazueraTransactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brazuera *BrazueraSession) Pause(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.Pause(&_Brazuera.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Brazuera *BrazueraTransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.Pause(&_Brazuera.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brazuera *BrazueraTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brazuera *BrazueraSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brazuera.Contract.RenounceOwnership(&_Brazuera.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brazuera *BrazueraTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brazuera.Contract.RenounceOwnership(&_Brazuera.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SafeTransferFrom(&_Brazuera.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SafeTransferFrom(&_Brazuera.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Brazuera *BrazueraTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Brazuera *BrazueraSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.Contract.SafeTransferFrom0(&_Brazuera.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Brazuera *BrazueraTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Brazuera.Contract.SafeTransferFrom0(&_Brazuera.TransactOpts, from, to, tokenId, _data)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Brazuera *BrazueraTransactor) SetAffiliate(opts *bind.TransactOpts, manager common.Address, state bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setAffiliate", manager, state)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Brazuera *BrazueraSession) SetAffiliate(manager common.Address, state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAffiliate(&_Brazuera.TransactOpts, manager, state)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(address manager, bool state) returns()
func (_Brazuera *BrazueraTransactorSession) SetAffiliate(manager common.Address, state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAffiliate(&_Brazuera.TransactOpts, manager, state)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Brazuera *BrazueraTransactor) SetAffiliateProgramPercentage(opts *bind.TransactOpts, percentage *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setAffiliateProgramPercentage", percentage)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Brazuera *BrazueraSession) SetAffiliateProgramPercentage(percentage *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAffiliateProgramPercentage(&_Brazuera.TransactOpts, percentage)
}

// SetAffiliateProgramPercentage is a paid mutator transaction binding the contract method 0x474c68ec.
//
// Solidity: function setAffiliateProgramPercentage(uint256 percentage) returns()
func (_Brazuera *BrazueraTransactorSession) SetAffiliateProgramPercentage(percentage *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAffiliateProgramPercentage(&_Brazuera.TransactOpts, percentage)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Brazuera *BrazueraTransactor) SetAllowAffiliateProgram(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setAllowAffiliateProgram", _state)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Brazuera *BrazueraSession) SetAllowAffiliateProgram(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAllowAffiliateProgram(&_Brazuera.TransactOpts, _state)
}

// SetAllowAffiliateProgram is a paid mutator transaction binding the contract method 0x94cf4b6a.
//
// Solidity: function setAllowAffiliateProgram(bool _state) returns()
func (_Brazuera *BrazueraTransactorSession) SetAllowAffiliateProgram(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetAllowAffiliateProgram(&_Brazuera.TransactOpts, _state)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brazuera *BrazueraTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brazuera *BrazueraSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetApprovalForAll(&_Brazuera.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Brazuera *BrazueraTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetApprovalForAll(&_Brazuera.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brazuera *BrazueraTransactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brazuera *BrazueraSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Brazuera.Contract.SetBaseURI(&_Brazuera.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Brazuera *BrazueraTransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Brazuera.Contract.SetBaseURI(&_Brazuera.TransactOpts, _newBaseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brazuera *BrazueraTransactor) SetContractURI(opts *bind.TransactOpts, contractURI_ string) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setContractURI", contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brazuera *BrazueraSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Brazuera.Contract.SetContractURI(&_Brazuera.TransactOpts, contractURI_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string contractURI_) returns()
func (_Brazuera *BrazueraTransactorSession) SetContractURI(contractURI_ string) (*types.Transaction, error) {
	return _Brazuera.Contract.SetContractURI(&_Brazuera.TransactOpts, contractURI_)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Brazuera *BrazueraTransactor) SetFreeMintPerAddressLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setFreeMintPerAddressLimit", _limit)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Brazuera *BrazueraSession) SetFreeMintPerAddressLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetFreeMintPerAddressLimit(&_Brazuera.TransactOpts, _limit)
}

// SetFreeMintPerAddressLimit is a paid mutator transaction binding the contract method 0x08181816.
//
// Solidity: function setFreeMintPerAddressLimit(uint256 _limit) returns()
func (_Brazuera *BrazueraTransactorSession) SetFreeMintPerAddressLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetFreeMintPerAddressLimit(&_Brazuera.TransactOpts, _limit)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Brazuera *BrazueraTransactor) SetIsFreeMint(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setIsFreeMint", state)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Brazuera *BrazueraSession) SetIsFreeMint(state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetIsFreeMint(&_Brazuera.TransactOpts, state)
}

// SetIsFreeMint is a paid mutator transaction binding the contract method 0xef5d87d7.
//
// Solidity: function setIsFreeMint(bool state) returns()
func (_Brazuera *BrazueraTransactorSession) SetIsFreeMint(state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetIsFreeMint(&_Brazuera.TransactOpts, state)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Brazuera *BrazueraTransactor) SetMaxMintAmount(opts *bind.TransactOpts, _maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setMaxMintAmount", _maxMintAmount)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Brazuera *BrazueraSession) SetMaxMintAmount(_maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetMaxMintAmount(&_Brazuera.TransactOpts, _maxMintAmount)
}

// SetMaxMintAmount is a paid mutator transaction binding the contract method 0x088a4ed0.
//
// Solidity: function setMaxMintAmount(uint256 _maxMintAmount) returns()
func (_Brazuera *BrazueraTransactorSession) SetMaxMintAmount(_maxMintAmount *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetMaxMintAmount(&_Brazuera.TransactOpts, _maxMintAmount)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Brazuera *BrazueraTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Brazuera *BrazueraSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetMaxSupply(&_Brazuera.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Brazuera *BrazueraTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetMaxSupply(&_Brazuera.TransactOpts, _maxSupply)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Brazuera *BrazueraTransactor) SetNftEtherValue(opts *bind.TransactOpts, nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setNftEtherValue", nftEtherValue)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Brazuera *BrazueraSession) SetNftEtherValue(nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetNftEtherValue(&_Brazuera.TransactOpts, nftEtherValue)
}

// SetNftEtherValue is a paid mutator transaction binding the contract method 0x6e17f797.
//
// Solidity: function setNftEtherValue(uint256 nftEtherValue) returns()
func (_Brazuera *BrazueraTransactorSession) SetNftEtherValue(nftEtherValue *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.SetNftEtherValue(&_Brazuera.TransactOpts, nftEtherValue)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Brazuera *BrazueraTransactor) SetOnlyWhitelisted(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "setOnlyWhitelisted", _state)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Brazuera *BrazueraSession) SetOnlyWhitelisted(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetOnlyWhitelisted(&_Brazuera.TransactOpts, _state)
}

// SetOnlyWhitelisted is a paid mutator transaction binding the contract method 0x3c952764.
//
// Solidity: function setOnlyWhitelisted(bool _state) returns()
func (_Brazuera *BrazueraTransactorSession) SetOnlyWhitelisted(_state bool) (*types.Transaction, error) {
	return _Brazuera.Contract.SetOnlyWhitelisted(&_Brazuera.TransactOpts, _state)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraTransactor) Split(opts *bind.TransactOpts, _mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "split", _mintAmount, _recommendedBy, _indicationType, proof)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraSession) Split(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.Contract.Split(&_Brazuera.TransactOpts, _mintAmount, _recommendedBy, _indicationType, proof)
}

// Split is a paid mutator transaction binding the contract method 0x8c584664.
//
// Solidity: function split(uint256 _mintAmount, address _recommendedBy, uint256 _indicationType, bytes32[] proof) payable returns()
func (_Brazuera *BrazueraTransactorSession) Split(_mintAmount *big.Int, _recommendedBy common.Address, _indicationType *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Brazuera.Contract.Split(&_Brazuera.TransactOpts, _mintAmount, _recommendedBy, _indicationType, proof)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.TransferFrom(&_Brazuera.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Brazuera *BrazueraTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Brazuera.Contract.TransferFrom(&_Brazuera.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brazuera *BrazueraTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brazuera *BrazueraSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brazuera.Contract.TransferOwnership(&_Brazuera.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brazuera *BrazueraTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brazuera.Contract.TransferOwnership(&_Brazuera.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brazuera *BrazueraTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brazuera.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brazuera *BrazueraSession) Withdraw() (*types.Transaction, error) {
	return _Brazuera.Contract.Withdraw(&_Brazuera.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Brazuera *BrazueraTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Brazuera.Contract.Withdraw(&_Brazuera.TransactOpts)
}

// BrazueraApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Brazuera contract.
type BrazueraApprovalIterator struct {
	Event *BrazueraApproval // Event containing the contract specifics and raw log

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
func (it *BrazueraApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrazueraApproval)
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
		it.Event = new(BrazueraApproval)
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
func (it *BrazueraApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrazueraApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrazueraApproval represents a Approval event raised by the Brazuera contract.
type BrazueraApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BrazueraApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Brazuera.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BrazueraApprovalIterator{contract: _Brazuera.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BrazueraApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Brazuera.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrazueraApproval)
				if err := _Brazuera.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) ParseApproval(log types.Log) (*BrazueraApproval, error) {
	event := new(BrazueraApproval)
	if err := _Brazuera.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrazueraApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Brazuera contract.
type BrazueraApprovalForAllIterator struct {
	Event *BrazueraApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BrazueraApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrazueraApprovalForAll)
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
		it.Event = new(BrazueraApprovalForAll)
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
func (it *BrazueraApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrazueraApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrazueraApprovalForAll represents a ApprovalForAll event raised by the Brazuera contract.
type BrazueraApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Brazuera *BrazueraFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BrazueraApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Brazuera.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BrazueraApprovalForAllIterator{contract: _Brazuera.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Brazuera *BrazueraFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BrazueraApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Brazuera.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrazueraApprovalForAll)
				if err := _Brazuera.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Brazuera *BrazueraFilterer) ParseApprovalForAll(log types.Log) (*BrazueraApprovalForAll, error) {
	event := new(BrazueraApprovalForAll)
	if err := _Brazuera.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrazueraOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Brazuera contract.
type BrazueraOwnershipTransferredIterator struct {
	Event *BrazueraOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrazueraOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrazueraOwnershipTransferred)
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
		it.Event = new(BrazueraOwnershipTransferred)
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
func (it *BrazueraOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrazueraOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrazueraOwnershipTransferred represents a OwnershipTransferred event raised by the Brazuera contract.
type BrazueraOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brazuera *BrazueraFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrazueraOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brazuera.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrazueraOwnershipTransferredIterator{contract: _Brazuera.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brazuera *BrazueraFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrazueraOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brazuera.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrazueraOwnershipTransferred)
				if err := _Brazuera.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Brazuera *BrazueraFilterer) ParseOwnershipTransferred(log types.Log) (*BrazueraOwnershipTransferred, error) {
	event := new(BrazueraOwnershipTransferred)
	if err := _Brazuera.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrazueraTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Brazuera contract.
type BrazueraTransferIterator struct {
	Event *BrazueraTransfer // Event containing the contract specifics and raw log

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
func (it *BrazueraTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrazueraTransfer)
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
		it.Event = new(BrazueraTransfer)
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
func (it *BrazueraTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrazueraTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrazueraTransfer represents a Transfer event raised by the Brazuera contract.
type BrazueraTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BrazueraTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Brazuera.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BrazueraTransferIterator{contract: _Brazuera.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BrazueraTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Brazuera.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrazueraTransfer)
				if err := _Brazuera.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Brazuera *BrazueraFilterer) ParseTransfer(log types.Log) (*BrazueraTransfer, error) {
	event := new(BrazueraTransfer)
	if err := _Brazuera.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrazueraTransferSendIterator is returned from FilterTransferSend and is used to iterate over the raw logs and unpacked data for TransferSend events raised by the Brazuera contract.
type BrazueraTransferSendIterator struct {
	Event *BrazueraTransferSend // Event containing the contract specifics and raw log

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
func (it *BrazueraTransferSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrazueraTransferSend)
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
		it.Event = new(BrazueraTransferSend)
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
func (it *BrazueraTransferSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrazueraTransferSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrazueraTransferSend represents a TransferSend event raised by the Brazuera contract.
type BrazueraTransferSend struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferSend is a free log retrieval operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Brazuera *BrazueraFilterer) FilterTransferSend(opts *bind.FilterOpts) (*BrazueraTransferSendIterator, error) {

	logs, sub, err := _Brazuera.contract.FilterLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return &BrazueraTransferSendIterator{contract: _Brazuera.contract, event: "_transferSend", logs: logs, sub: sub}, nil
}

// WatchTransferSend is a free log subscription operation binding the contract event 0xf70a742b6ec5de639ce8e8d0df10a44a88c323ac660a09f1fc0c190e830aea74.
//
// Solidity: event _transferSend(address _from, address _to, uint256 _amount)
func (_Brazuera *BrazueraFilterer) WatchTransferSend(opts *bind.WatchOpts, sink chan<- *BrazueraTransferSend) (event.Subscription, error) {

	logs, sub, err := _Brazuera.contract.WatchLogs(opts, "_transferSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrazueraTransferSend)
				if err := _Brazuera.contract.UnpackLog(event, "_transferSend", log); err != nil {
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
func (_Brazuera *BrazueraFilterer) ParseTransferSend(log types.Log) (*BrazueraTransferSend, error) {
	event := new(BrazueraTransferSend)
	if err := _Brazuera.contract.UnpackLog(event, "_transferSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
