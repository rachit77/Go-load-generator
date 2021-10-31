// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package loadbot

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

// LoadbotMetaData contains all meta data concerning the Loadbot contract.
var LoadbotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"name\":\"InflationFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatePeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"name\":\"InflationParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"TransferComment\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inflationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inflationFactorUpdatePeriod\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"initialBalanceAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"initialBalanceValues\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"exchangeIdentifier\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatePeriod\",\"type\":\"uint256\"}],\"name\":\"setInflationParameters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mintForLoadBot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"findBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"transferWithComment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInflationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExchangeRegistryId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"units\",\"type\":\"uint256\"}],\"name\":\"unitsToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"debitGasFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gatewayFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tipTxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gatewayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTxFee\",\"type\":\"uint256\"}],\"name\":\"creditGasFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LoadbotABI is the input ABI used to generate the binding from.
// Deprecated: Use LoadbotMetaData.ABI instead.
var LoadbotABI = LoadbotMetaData.ABI

// Loadbot is an auto generated Go binding around an Ethereum contract.
type Loadbot struct {
	LoadbotCaller     // Read-only binding to the contract
	LoadbotTransactor // Write-only binding to the contract
	LoadbotFilterer   // Log filterer for contract events
}

// LoadbotCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoadbotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadbotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoadbotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadbotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoadbotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadbotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoadbotSession struct {
	Contract     *Loadbot          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoadbotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoadbotCallerSession struct {
	Contract *LoadbotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LoadbotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoadbotTransactorSession struct {
	Contract     *LoadbotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LoadbotRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoadbotRaw struct {
	Contract *Loadbot // Generic contract binding to access the raw methods on
}

// LoadbotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoadbotCallerRaw struct {
	Contract *LoadbotCaller // Generic read-only contract binding to access the raw methods on
}

// LoadbotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoadbotTransactorRaw struct {
	Contract *LoadbotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoadbot creates a new instance of Loadbot, bound to a specific deployed contract.
func NewLoadbot(address common.Address, backend bind.ContractBackend) (*Loadbot, error) {
	contract, err := bindLoadbot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Loadbot{LoadbotCaller: LoadbotCaller{contract: contract}, LoadbotTransactor: LoadbotTransactor{contract: contract}, LoadbotFilterer: LoadbotFilterer{contract: contract}}, nil
}

// NewLoadbotCaller creates a new read-only instance of Loadbot, bound to a specific deployed contract.
func NewLoadbotCaller(address common.Address, caller bind.ContractCaller) (*LoadbotCaller, error) {
	contract, err := bindLoadbot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoadbotCaller{contract: contract}, nil
}

// NewLoadbotTransactor creates a new write-only instance of Loadbot, bound to a specific deployed contract.
func NewLoadbotTransactor(address common.Address, transactor bind.ContractTransactor) (*LoadbotTransactor, error) {
	contract, err := bindLoadbot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoadbotTransactor{contract: contract}, nil
}

// NewLoadbotFilterer creates a new log filterer instance of Loadbot, bound to a specific deployed contract.
func NewLoadbotFilterer(address common.Address, filterer bind.ContractFilterer) (*LoadbotFilterer, error) {
	contract, err := bindLoadbot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoadbotFilterer{contract: contract}, nil
}

// bindLoadbot binds a generic wrapper to an already deployed contract.
func bindLoadbot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoadbotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loadbot *LoadbotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loadbot.Contract.LoadbotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loadbot *LoadbotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loadbot.Contract.LoadbotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loadbot *LoadbotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loadbot.Contract.LoadbotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loadbot *LoadbotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loadbot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loadbot *LoadbotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loadbot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loadbot *LoadbotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loadbot.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) view returns(uint256)
func (_Loadbot *LoadbotCaller) Allowance(opts *bind.CallOpts, accountOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "allowance", accountOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) view returns(uint256)
func (_Loadbot *LoadbotSession) Allowance(accountOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Loadbot.Contract.Allowance(&_Loadbot.CallOpts, accountOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) Allowance(accountOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Loadbot.Contract.Allowance(&_Loadbot.CallOpts, accountOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) view returns(uint256)
func (_Loadbot *LoadbotCaller) BalanceOf(opts *bind.CallOpts, accountOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "balanceOf", accountOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) view returns(uint256)
func (_Loadbot *LoadbotSession) BalanceOf(accountOwner common.Address) (*big.Int, error) {
	return _Loadbot.Contract.BalanceOf(&_Loadbot.CallOpts, accountOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) BalanceOf(accountOwner common.Address) (*big.Int, error) {
	return _Loadbot.Contract.BalanceOf(&_Loadbot.CallOpts, accountOwner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Loadbot *LoadbotCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Loadbot *LoadbotSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Loadbot.Contract.Balances(&_Loadbot.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Loadbot.Contract.Balances(&_Loadbot.CallOpts, arg0)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Loadbot *LoadbotCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Loadbot *LoadbotSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Loadbot.Contract.CheckProofOfPossession(&_Loadbot.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Loadbot *LoadbotCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Loadbot.Contract.CheckProofOfPossession(&_Loadbot.CallOpts, sender, blsKey, blsPop)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Loadbot *LoadbotCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Loadbot *LoadbotSession) Decimals() (uint8, error) {
	return _Loadbot.Contract.Decimals(&_Loadbot.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Loadbot *LoadbotCallerSession) Decimals() (uint8, error) {
	return _Loadbot.Contract.Decimals(&_Loadbot.CallOpts)
}

// FindBalance is a free data retrieval call binding the contract method 0xcfba4fe0.
//
// Solidity: function findBalance(address to) view returns(uint256)
func (_Loadbot *LoadbotCaller) FindBalance(opts *bind.CallOpts, to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "findBalance", to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindBalance is a free data retrieval call binding the contract method 0xcfba4fe0.
//
// Solidity: function findBalance(address to) view returns(uint256)
func (_Loadbot *LoadbotSession) FindBalance(to common.Address) (*big.Int, error) {
	return _Loadbot.Contract.FindBalance(&_Loadbot.CallOpts, to)
}

// FindBalance is a free data retrieval call binding the contract method 0xcfba4fe0.
//
// Solidity: function findBalance(address to) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) FindBalance(to common.Address) (*big.Int, error) {
	return _Loadbot.Contract.FindBalance(&_Loadbot.CallOpts, to)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Loadbot *LoadbotCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Loadbot *LoadbotSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Loadbot.Contract.FractionMulExp(&_Loadbot.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Loadbot *LoadbotCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Loadbot.Contract.FractionMulExp(&_Loadbot.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Loadbot *LoadbotCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Loadbot *LoadbotSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Loadbot.Contract.GetBlockNumberFromHeader(&_Loadbot.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Loadbot.Contract.GetBlockNumberFromHeader(&_Loadbot.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Loadbot *LoadbotCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Loadbot *LoadbotSession) GetEpochNumber() (*big.Int, error) {
	return _Loadbot.Contract.GetEpochNumber(&_Loadbot.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Loadbot *LoadbotCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Loadbot.Contract.GetEpochNumber(&_Loadbot.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.GetEpochNumberOfBlock(&_Loadbot.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.GetEpochNumberOfBlock(&_Loadbot.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Loadbot *LoadbotCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Loadbot *LoadbotSession) GetEpochSize() (*big.Int, error) {
	return _Loadbot.Contract.GetEpochSize(&_Loadbot.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Loadbot *LoadbotCallerSession) GetEpochSize() (*big.Int, error) {
	return _Loadbot.Contract.GetEpochSize(&_Loadbot.CallOpts)
}

// GetExchangeRegistryId is a free data retrieval call binding the contract method 0x40a12f64.
//
// Solidity: function getExchangeRegistryId() view returns(bytes32)
func (_Loadbot *LoadbotCaller) GetExchangeRegistryId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getExchangeRegistryId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetExchangeRegistryId is a free data retrieval call binding the contract method 0x40a12f64.
//
// Solidity: function getExchangeRegistryId() view returns(bytes32)
func (_Loadbot *LoadbotSession) GetExchangeRegistryId() ([32]byte, error) {
	return _Loadbot.Contract.GetExchangeRegistryId(&_Loadbot.CallOpts)
}

// GetExchangeRegistryId is a free data retrieval call binding the contract method 0x40a12f64.
//
// Solidity: function getExchangeRegistryId() view returns(bytes32)
func (_Loadbot *LoadbotCallerSession) GetExchangeRegistryId() ([32]byte, error) {
	return _Loadbot.Contract.GetExchangeRegistryId(&_Loadbot.CallOpts)
}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotCaller) GetInflationParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getInflationParameters")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotSession) GetInflationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Loadbot.Contract.GetInflationParameters(&_Loadbot.CallOpts)
}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotCallerSession) GetInflationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Loadbot.Contract.GetInflationParameters(&_Loadbot.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Loadbot *LoadbotCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Loadbot *LoadbotSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Loadbot.Contract.GetParentSealBitmap(&_Loadbot.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Loadbot *LoadbotCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Loadbot.Contract.GetParentSealBitmap(&_Loadbot.CallOpts, blockNumber)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Loadbot.Contract.GetVerifiedSealBitmapFromHeader(&_Loadbot.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Loadbot.Contract.GetVerifiedSealBitmapFromHeader(&_Loadbot.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "getVersionNumber")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Loadbot.Contract.GetVersionNumber(&_Loadbot.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Loadbot *LoadbotCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Loadbot.Contract.GetVersionNumber(&_Loadbot.CallOpts)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotSession) HashHeader(header []byte) ([32]byte, error) {
	return _Loadbot.Contract.HashHeader(&_Loadbot.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Loadbot *LoadbotCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Loadbot.Contract.HashHeader(&_Loadbot.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Loadbot *LoadbotCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Loadbot *LoadbotSession) Initialized() (bool, error) {
	return _Loadbot.Contract.Initialized(&_Loadbot.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Loadbot *LoadbotCallerSession) Initialized() (bool, error) {
	return _Loadbot.Contract.Initialized(&_Loadbot.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Loadbot *LoadbotCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Loadbot *LoadbotSession) IsOwner() (bool, error) {
	return _Loadbot.Contract.IsOwner(&_Loadbot.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Loadbot *LoadbotCallerSession) IsOwner() (bool, error) {
	return _Loadbot.Contract.IsOwner(&_Loadbot.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.MinQuorumSize(&_Loadbot.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.MinQuorumSize(&_Loadbot.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Loadbot.Contract.MinQuorumSizeInCurrentSet(&_Loadbot.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Loadbot.Contract.MinQuorumSizeInCurrentSet(&_Loadbot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Loadbot *LoadbotCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Loadbot *LoadbotSession) Name() (string, error) {
	return _Loadbot.Contract.Name(&_Loadbot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Loadbot *LoadbotCallerSession) Name() (string, error) {
	return _Loadbot.Contract.Name(&_Loadbot.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Loadbot.Contract.NumberValidatorsInCurrentSet(&_Loadbot.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Loadbot *LoadbotCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Loadbot.Contract.NumberValidatorsInCurrentSet(&_Loadbot.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.NumberValidatorsInSet(&_Loadbot.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.NumberValidatorsInSet(&_Loadbot.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Loadbot *LoadbotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Loadbot *LoadbotSession) Owner() (common.Address, error) {
	return _Loadbot.Contract.Owner(&_Loadbot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Loadbot *LoadbotCallerSession) Owner() (common.Address, error) {
	return _Loadbot.Contract.Owner(&_Loadbot.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Loadbot *LoadbotCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Loadbot *LoadbotSession) Registry() (common.Address, error) {
	return _Loadbot.Contract.Registry(&_Loadbot.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Loadbot *LoadbotCallerSession) Registry() (common.Address, error) {
	return _Loadbot.Contract.Registry(&_Loadbot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Loadbot *LoadbotCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Loadbot *LoadbotSession) Symbol() (string, error) {
	return _Loadbot.Contract.Symbol(&_Loadbot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Loadbot *LoadbotCallerSession) Symbol() (string, error) {
	return _Loadbot.Contract.Symbol(&_Loadbot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Loadbot *LoadbotCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Loadbot *LoadbotSession) TotalSupply() (*big.Int, error) {
	return _Loadbot.Contract.TotalSupply(&_Loadbot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Loadbot *LoadbotCallerSession) TotalSupply() (*big.Int, error) {
	return _Loadbot.Contract.TotalSupply(&_Loadbot.CallOpts)
}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) view returns(uint256)
func (_Loadbot *LoadbotCaller) UnitsToValue(opts *bind.CallOpts, units *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "unitsToValue", units)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) view returns(uint256)
func (_Loadbot *LoadbotSession) UnitsToValue(units *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.UnitsToValue(&_Loadbot.CallOpts, units)
}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) UnitsToValue(units *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.UnitsToValue(&_Loadbot.CallOpts, units)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Loadbot *LoadbotCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Loadbot *LoadbotSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Loadbot.Contract.ValidatorSignerAddressFromCurrentSet(&_Loadbot.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Loadbot *LoadbotCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Loadbot.Contract.ValidatorSignerAddressFromCurrentSet(&_Loadbot.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Loadbot *LoadbotCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Loadbot *LoadbotSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Loadbot.Contract.ValidatorSignerAddressFromSet(&_Loadbot.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Loadbot *LoadbotCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Loadbot.Contract.ValidatorSignerAddressFromSet(&_Loadbot.CallOpts, index, blockNumber)
}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) view returns(uint256)
func (_Loadbot *LoadbotCaller) ValueToUnits(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loadbot.contract.Call(opts, &out, "valueToUnits", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) view returns(uint256)
func (_Loadbot *LoadbotSession) ValueToUnits(value *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.ValueToUnits(&_Loadbot.CallOpts, value)
}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) view returns(uint256)
func (_Loadbot *LoadbotCallerSession) ValueToUnits(value *big.Int) (*big.Int, error) {
	return _Loadbot.Contract.ValueToUnits(&_Loadbot.CallOpts, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Approve(&_Loadbot.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Approve(&_Loadbot.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Burn(&_Loadbot.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Burn(&_Loadbot.TransactOpts, value)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_Loadbot *LoadbotTransactor) CreditGasFees(opts *bind.TransactOpts, from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "creditGasFees", from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_Loadbot *LoadbotSession) CreditGasFees(from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.CreditGasFees(&_Loadbot.TransactOpts, from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_Loadbot *LoadbotTransactorSession) CreditGasFees(from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.CreditGasFees(&_Loadbot.TransactOpts, from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_Loadbot *LoadbotTransactor) DebitGasFees(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "debitGasFees", from, value)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_Loadbot *LoadbotSession) DebitGasFees(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.DebitGasFees(&_Loadbot.TransactOpts, from, value)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_Loadbot *LoadbotTransactorSession) DebitGasFees(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.DebitGasFees(&_Loadbot.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "decreaseAllowance", spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.DecreaseAllowance(&_Loadbot.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.DecreaseAllowance(&_Loadbot.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "increaseAllowance", spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.IncreaseAllowance(&_Loadbot.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.IncreaseAllowance(&_Loadbot.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e4f0e03.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues, string exchangeIdentifier) returns()
func (_Loadbot *LoadbotTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int, exchangeIdentifier string) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "initialize", _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues, exchangeIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e4f0e03.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues, string exchangeIdentifier) returns()
func (_Loadbot *LoadbotSession) Initialize(_name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int, exchangeIdentifier string) (*types.Transaction, error) {
	return _Loadbot.Contract.Initialize(&_Loadbot.TransactOpts, _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues, exchangeIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e4f0e03.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues, string exchangeIdentifier) returns()
func (_Loadbot *LoadbotTransactorSession) Initialize(_name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int, exchangeIdentifier string) (*types.Transaction, error) {
	return _Loadbot.Contract.Initialize(&_Loadbot.TransactOpts, _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues, exchangeIdentifier)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Mint(&_Loadbot.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Mint(&_Loadbot.TransactOpts, to, value)
}

// MintForLoadBot is a paid mutator transaction binding the contract method 0x7de02189.
//
// Solidity: function mintForLoadBot(address to, uint256 value) returns()
func (_Loadbot *LoadbotTransactor) MintForLoadBot(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "mintForLoadBot", to, value)
}

// MintForLoadBot is a paid mutator transaction binding the contract method 0x7de02189.
//
// Solidity: function mintForLoadBot(address to, uint256 value) returns()
func (_Loadbot *LoadbotSession) MintForLoadBot(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.MintForLoadBot(&_Loadbot.TransactOpts, to, value)
}

// MintForLoadBot is a paid mutator transaction binding the contract method 0x7de02189.
//
// Solidity: function mintForLoadBot(address to, uint256 value) returns()
func (_Loadbot *LoadbotTransactorSession) MintForLoadBot(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.MintForLoadBot(&_Loadbot.TransactOpts, to, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Loadbot *LoadbotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Loadbot *LoadbotSession) RenounceOwnership() (*types.Transaction, error) {
	return _Loadbot.Contract.RenounceOwnership(&_Loadbot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Loadbot *LoadbotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Loadbot.Contract.RenounceOwnership(&_Loadbot.TransactOpts)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_Loadbot *LoadbotTransactor) SetInflationParameters(opts *bind.TransactOpts, rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "setInflationParameters", rate, updatePeriod)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_Loadbot *LoadbotSession) SetInflationParameters(rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.SetInflationParameters(&_Loadbot.TransactOpts, rate, updatePeriod)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_Loadbot *LoadbotTransactorSession) SetInflationParameters(rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.SetInflationParameters(&_Loadbot.TransactOpts, rate, updatePeriod)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Loadbot *LoadbotTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Loadbot *LoadbotSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Loadbot.Contract.SetRegistry(&_Loadbot.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Loadbot *LoadbotTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Loadbot.Contract.SetRegistry(&_Loadbot.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Transfer(&_Loadbot.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.Transfer(&_Loadbot.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferFrom(&_Loadbot.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Loadbot *LoadbotTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferFrom(&_Loadbot.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Loadbot *LoadbotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Loadbot *LoadbotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferOwnership(&_Loadbot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Loadbot *LoadbotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferOwnership(&_Loadbot.TransactOpts, newOwner)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_Loadbot *LoadbotTransactor) TransferWithComment(opts *bind.TransactOpts, to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _Loadbot.contract.Transact(opts, "transferWithComment", to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_Loadbot *LoadbotSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferWithComment(&_Loadbot.TransactOpts, to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_Loadbot *LoadbotTransactorSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _Loadbot.Contract.TransferWithComment(&_Loadbot.TransactOpts, to, value, comment)
}

// LoadbotApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Loadbot contract.
type LoadbotApprovalIterator struct {
	Event *LoadbotApproval // Event containing the contract specifics and raw log

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
func (it *LoadbotApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotApproval)
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
		it.Event = new(LoadbotApproval)
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
func (it *LoadbotApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotApproval represents a Approval event raised by the Loadbot contract.
type LoadbotApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Loadbot *LoadbotFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LoadbotApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LoadbotApprovalIterator{contract: _Loadbot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Loadbot *LoadbotFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LoadbotApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotApproval)
				if err := _Loadbot.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Loadbot *LoadbotFilterer) ParseApproval(log types.Log) (*LoadbotApproval, error) {
	event := new(LoadbotApproval)
	if err := _Loadbot.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotInflationFactorUpdatedIterator is returned from FilterInflationFactorUpdated and is used to iterate over the raw logs and unpacked data for InflationFactorUpdated events raised by the Loadbot contract.
type LoadbotInflationFactorUpdatedIterator struct {
	Event *LoadbotInflationFactorUpdated // Event containing the contract specifics and raw log

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
func (it *LoadbotInflationFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotInflationFactorUpdated)
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
		it.Event = new(LoadbotInflationFactorUpdated)
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
func (it *LoadbotInflationFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotInflationFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotInflationFactorUpdated represents a InflationFactorUpdated event raised by the Loadbot contract.
type LoadbotInflationFactorUpdated struct {
	Factor      *big.Int
	LastUpdated *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInflationFactorUpdated is a free log retrieval operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) FilterInflationFactorUpdated(opts *bind.FilterOpts) (*LoadbotInflationFactorUpdatedIterator, error) {

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "InflationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &LoadbotInflationFactorUpdatedIterator{contract: _Loadbot.contract, event: "InflationFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchInflationFactorUpdated is a free log subscription operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) WatchInflationFactorUpdated(opts *bind.WatchOpts, sink chan<- *LoadbotInflationFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "InflationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotInflationFactorUpdated)
				if err := _Loadbot.contract.UnpackLog(event, "InflationFactorUpdated", log); err != nil {
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

// ParseInflationFactorUpdated is a log parse operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) ParseInflationFactorUpdated(log types.Log) (*LoadbotInflationFactorUpdated, error) {
	event := new(LoadbotInflationFactorUpdated)
	if err := _Loadbot.contract.UnpackLog(event, "InflationFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotInflationParametersUpdatedIterator is returned from FilterInflationParametersUpdated and is used to iterate over the raw logs and unpacked data for InflationParametersUpdated events raised by the Loadbot contract.
type LoadbotInflationParametersUpdatedIterator struct {
	Event *LoadbotInflationParametersUpdated // Event containing the contract specifics and raw log

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
func (it *LoadbotInflationParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotInflationParametersUpdated)
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
		it.Event = new(LoadbotInflationParametersUpdated)
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
func (it *LoadbotInflationParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotInflationParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotInflationParametersUpdated represents a InflationParametersUpdated event raised by the Loadbot contract.
type LoadbotInflationParametersUpdated struct {
	Rate         *big.Int
	UpdatePeriod *big.Int
	LastUpdated  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInflationParametersUpdated is a free log retrieval operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) FilterInflationParametersUpdated(opts *bind.FilterOpts) (*LoadbotInflationParametersUpdatedIterator, error) {

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "InflationParametersUpdated")
	if err != nil {
		return nil, err
	}
	return &LoadbotInflationParametersUpdatedIterator{contract: _Loadbot.contract, event: "InflationParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchInflationParametersUpdated is a free log subscription operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) WatchInflationParametersUpdated(opts *bind.WatchOpts, sink chan<- *LoadbotInflationParametersUpdated) (event.Subscription, error) {

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "InflationParametersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotInflationParametersUpdated)
				if err := _Loadbot.contract.UnpackLog(event, "InflationParametersUpdated", log); err != nil {
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

// ParseInflationParametersUpdated is a log parse operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_Loadbot *LoadbotFilterer) ParseInflationParametersUpdated(log types.Log) (*LoadbotInflationParametersUpdated, error) {
	event := new(LoadbotInflationParametersUpdated)
	if err := _Loadbot.contract.UnpackLog(event, "InflationParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Loadbot contract.
type LoadbotOwnershipTransferredIterator struct {
	Event *LoadbotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoadbotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotOwnershipTransferred)
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
		it.Event = new(LoadbotOwnershipTransferred)
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
func (it *LoadbotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotOwnershipTransferred represents a OwnershipTransferred event raised by the Loadbot contract.
type LoadbotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Loadbot *LoadbotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoadbotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoadbotOwnershipTransferredIterator{contract: _Loadbot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Loadbot *LoadbotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoadbotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotOwnershipTransferred)
				if err := _Loadbot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Loadbot *LoadbotFilterer) ParseOwnershipTransferred(log types.Log) (*LoadbotOwnershipTransferred, error) {
	event := new(LoadbotOwnershipTransferred)
	if err := _Loadbot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Loadbot contract.
type LoadbotRegistrySetIterator struct {
	Event *LoadbotRegistrySet // Event containing the contract specifics and raw log

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
func (it *LoadbotRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotRegistrySet)
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
		it.Event = new(LoadbotRegistrySet)
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
func (it *LoadbotRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotRegistrySet represents a RegistrySet event raised by the Loadbot contract.
type LoadbotRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Loadbot *LoadbotFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*LoadbotRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &LoadbotRegistrySetIterator{contract: _Loadbot.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Loadbot *LoadbotFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *LoadbotRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotRegistrySet)
				if err := _Loadbot.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Loadbot *LoadbotFilterer) ParseRegistrySet(log types.Log) (*LoadbotRegistrySet, error) {
	event := new(LoadbotRegistrySet)
	if err := _Loadbot.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Loadbot contract.
type LoadbotTransferIterator struct {
	Event *LoadbotTransfer // Event containing the contract specifics and raw log

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
func (it *LoadbotTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotTransfer)
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
		it.Event = new(LoadbotTransfer)
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
func (it *LoadbotTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotTransfer represents a Transfer event raised by the Loadbot contract.
type LoadbotTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Loadbot *LoadbotFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LoadbotTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LoadbotTransferIterator{contract: _Loadbot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Loadbot *LoadbotFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LoadbotTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotTransfer)
				if err := _Loadbot.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Loadbot *LoadbotFilterer) ParseTransfer(log types.Log) (*LoadbotTransfer, error) {
	event := new(LoadbotTransfer)
	if err := _Loadbot.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadbotTransferCommentIterator is returned from FilterTransferComment and is used to iterate over the raw logs and unpacked data for TransferComment events raised by the Loadbot contract.
type LoadbotTransferCommentIterator struct {
	Event *LoadbotTransferComment // Event containing the contract specifics and raw log

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
func (it *LoadbotTransferCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadbotTransferComment)
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
		it.Event = new(LoadbotTransferComment)
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
func (it *LoadbotTransferCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadbotTransferCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadbotTransferComment represents a TransferComment event raised by the Loadbot contract.
type LoadbotTransferComment struct {
	Comment string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferComment is a free log retrieval operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_Loadbot *LoadbotFilterer) FilterTransferComment(opts *bind.FilterOpts) (*LoadbotTransferCommentIterator, error) {

	logs, sub, err := _Loadbot.contract.FilterLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return &LoadbotTransferCommentIterator{contract: _Loadbot.contract, event: "TransferComment", logs: logs, sub: sub}, nil
}

// WatchTransferComment is a free log subscription operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_Loadbot *LoadbotFilterer) WatchTransferComment(opts *bind.WatchOpts, sink chan<- *LoadbotTransferComment) (event.Subscription, error) {

	logs, sub, err := _Loadbot.contract.WatchLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadbotTransferComment)
				if err := _Loadbot.contract.UnpackLog(event, "TransferComment", log); err != nil {
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

// ParseTransferComment is a log parse operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_Loadbot *LoadbotFilterer) ParseTransferComment(log types.Log) (*LoadbotTransferComment, error) {
	event := new(LoadbotTransferComment)
	if err := _Loadbot.contract.UnpackLog(event, "TransferComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
