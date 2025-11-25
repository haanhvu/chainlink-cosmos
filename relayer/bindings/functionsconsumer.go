// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// FunctionsConsumerMetaData contains all meta data concerning the FunctionsConsumer contract.
var FunctionsConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_donId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyArgs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySource\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoInlineSecrets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRouterCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"FunctionsDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"handleOracleFulfillment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_lastError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_lastRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_lastResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"enumFunctionsRequest.Location\",\"name\":\"secretsLocation\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"encryptedSecretsReference\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"bytesArgs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"sendRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newDonId\",\"type\":\"bytes32\"}],\"name\":\"setDonId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FunctionsConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use FunctionsConsumerMetaData.ABI instead.
var FunctionsConsumerABI = FunctionsConsumerMetaData.ABI

// FunctionsConsumer is an auto generated Go binding around an Ethereum contract.
type FunctionsConsumer struct {
	FunctionsConsumerCaller     // Read-only binding to the contract
	FunctionsConsumerTransactor // Write-only binding to the contract
	FunctionsConsumerFilterer   // Log filterer for contract events
}

// FunctionsConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FunctionsConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionsConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FunctionsConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionsConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FunctionsConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionsConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FunctionsConsumerSession struct {
	Contract     *FunctionsConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FunctionsConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FunctionsConsumerCallerSession struct {
	Contract *FunctionsConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FunctionsConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FunctionsConsumerTransactorSession struct {
	Contract     *FunctionsConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FunctionsConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FunctionsConsumerRaw struct {
	Contract *FunctionsConsumer // Generic contract binding to access the raw methods on
}

// FunctionsConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FunctionsConsumerCallerRaw struct {
	Contract *FunctionsConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// FunctionsConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FunctionsConsumerTransactorRaw struct {
	Contract *FunctionsConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFunctionsConsumer creates a new instance of FunctionsConsumer, bound to a specific deployed contract.
func NewFunctionsConsumer(address common.Address, backend bind.ContractBackend) (*FunctionsConsumer, error) {
	contract, err := bindFunctionsConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumer{FunctionsConsumerCaller: FunctionsConsumerCaller{contract: contract}, FunctionsConsumerTransactor: FunctionsConsumerTransactor{contract: contract}, FunctionsConsumerFilterer: FunctionsConsumerFilterer{contract: contract}}, nil
}

// NewFunctionsConsumerCaller creates a new read-only instance of FunctionsConsumer, bound to a specific deployed contract.
func NewFunctionsConsumerCaller(address common.Address, caller bind.ContractCaller) (*FunctionsConsumerCaller, error) {
	contract, err := bindFunctionsConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerCaller{contract: contract}, nil
}

// NewFunctionsConsumerTransactor creates a new write-only instance of FunctionsConsumer, bound to a specific deployed contract.
func NewFunctionsConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*FunctionsConsumerTransactor, error) {
	contract, err := bindFunctionsConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerTransactor{contract: contract}, nil
}

// NewFunctionsConsumerFilterer creates a new log filterer instance of FunctionsConsumer, bound to a specific deployed contract.
func NewFunctionsConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*FunctionsConsumerFilterer, error) {
	contract, err := bindFunctionsConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerFilterer{contract: contract}, nil
}

// bindFunctionsConsumer binds a generic wrapper to an already deployed contract.
func bindFunctionsConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FunctionsConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionsConsumer *FunctionsConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionsConsumer.Contract.FunctionsConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionsConsumer *FunctionsConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.FunctionsConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionsConsumer *FunctionsConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.FunctionsConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionsConsumer *FunctionsConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionsConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionsConsumer *FunctionsConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionsConsumer *FunctionsConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.contract.Transact(opts, method, params...)
}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerCaller) DonId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FunctionsConsumer.contract.Call(opts, &out, "donId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerSession) DonId() ([32]byte, error) {
	return _FunctionsConsumer.Contract.DonId(&_FunctionsConsumer.CallOpts)
}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerCallerSession) DonId() ([32]byte, error) {
	return _FunctionsConsumer.Contract.DonId(&_FunctionsConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionsConsumer *FunctionsConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FunctionsConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionsConsumer *FunctionsConsumerSession) Owner() (common.Address, error) {
	return _FunctionsConsumer.Contract.Owner(&_FunctionsConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionsConsumer *FunctionsConsumerCallerSession) Owner() (common.Address, error) {
	return _FunctionsConsumer.Contract.Owner(&_FunctionsConsumer.CallOpts)
}

// SLastError is a free data retrieval call binding the contract method 0x4b0795a8.
//
// Solidity: function s_lastError() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerCaller) SLastError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionsConsumer.contract.Call(opts, &out, "s_lastError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SLastError is a free data retrieval call binding the contract method 0x4b0795a8.
//
// Solidity: function s_lastError() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerSession) SLastError() ([]byte, error) {
	return _FunctionsConsumer.Contract.SLastError(&_FunctionsConsumer.CallOpts)
}

// SLastError is a free data retrieval call binding the contract method 0x4b0795a8.
//
// Solidity: function s_lastError() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerCallerSession) SLastError() ([]byte, error) {
	return _FunctionsConsumer.Contract.SLastError(&_FunctionsConsumer.CallOpts)
}

// SLastRequestId is a free data retrieval call binding the contract method 0xb1e21749.
//
// Solidity: function s_lastRequestId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerCaller) SLastRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FunctionsConsumer.contract.Call(opts, &out, "s_lastRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SLastRequestId is a free data retrieval call binding the contract method 0xb1e21749.
//
// Solidity: function s_lastRequestId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerSession) SLastRequestId() ([32]byte, error) {
	return _FunctionsConsumer.Contract.SLastRequestId(&_FunctionsConsumer.CallOpts)
}

// SLastRequestId is a free data retrieval call binding the contract method 0xb1e21749.
//
// Solidity: function s_lastRequestId() view returns(bytes32)
func (_FunctionsConsumer *FunctionsConsumerCallerSession) SLastRequestId() ([32]byte, error) {
	return _FunctionsConsumer.Contract.SLastRequestId(&_FunctionsConsumer.CallOpts)
}

// SLastResponse is a free data retrieval call binding the contract method 0x3944ea3a.
//
// Solidity: function s_lastResponse() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerCaller) SLastResponse(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionsConsumer.contract.Call(opts, &out, "s_lastResponse")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SLastResponse is a free data retrieval call binding the contract method 0x3944ea3a.
//
// Solidity: function s_lastResponse() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerSession) SLastResponse() ([]byte, error) {
	return _FunctionsConsumer.Contract.SLastResponse(&_FunctionsConsumer.CallOpts)
}

// SLastResponse is a free data retrieval call binding the contract method 0x3944ea3a.
//
// Solidity: function s_lastResponse() view returns(bytes)
func (_FunctionsConsumer *FunctionsConsumerCallerSession) SLastResponse() ([]byte, error) {
	return _FunctionsConsumer.Contract.SLastResponse(&_FunctionsConsumer.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionsConsumer *FunctionsConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionsConsumer.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionsConsumer *FunctionsConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.AcceptOwnership(&_FunctionsConsumer.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionsConsumer *FunctionsConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.AcceptOwnership(&_FunctionsConsumer.TransactOpts)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactor) HandleOracleFulfillment(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionsConsumer.contract.Transact(opts, "handleOracleFulfillment", requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionsConsumer *FunctionsConsumerSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.HandleOracleFulfillment(&_FunctionsConsumer.TransactOpts, requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactorSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.HandleOracleFulfillment(&_FunctionsConsumer.TransactOpts, requestId, response, err)
}

// SendRequest is a paid mutator transaction binding the contract method 0x231c1619.
//
// Solidity: function sendRequest(string source, uint8 secretsLocation, bytes encryptedSecretsReference, string[] args, bytes[] bytesArgs, uint64 subscriptionId, uint32 callbackGasLimit) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactor) SendRequest(opts *bind.TransactOpts, source string, secretsLocation uint8, encryptedSecretsReference []byte, args []string, bytesArgs [][]byte, subscriptionId uint64, callbackGasLimit uint32) (*types.Transaction, error) {
	return _FunctionsConsumer.contract.Transact(opts, "sendRequest", source, secretsLocation, encryptedSecretsReference, args, bytesArgs, subscriptionId, callbackGasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x231c1619.
//
// Solidity: function sendRequest(string source, uint8 secretsLocation, bytes encryptedSecretsReference, string[] args, bytes[] bytesArgs, uint64 subscriptionId, uint32 callbackGasLimit) returns()
func (_FunctionsConsumer *FunctionsConsumerSession) SendRequest(source string, secretsLocation uint8, encryptedSecretsReference []byte, args []string, bytesArgs [][]byte, subscriptionId uint64, callbackGasLimit uint32) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.SendRequest(&_FunctionsConsumer.TransactOpts, source, secretsLocation, encryptedSecretsReference, args, bytesArgs, subscriptionId, callbackGasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x231c1619.
//
// Solidity: function sendRequest(string source, uint8 secretsLocation, bytes encryptedSecretsReference, string[] args, bytes[] bytesArgs, uint64 subscriptionId, uint32 callbackGasLimit) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactorSession) SendRequest(source string, secretsLocation uint8, encryptedSecretsReference []byte, args []string, bytesArgs [][]byte, subscriptionId uint64, callbackGasLimit uint32) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.SendRequest(&_FunctionsConsumer.TransactOpts, source, secretsLocation, encryptedSecretsReference, args, bytesArgs, subscriptionId, callbackGasLimit)
}

// SetDonId is a paid mutator transaction binding the contract method 0x78ca5de7.
//
// Solidity: function setDonId(bytes32 newDonId) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactor) SetDonId(opts *bind.TransactOpts, newDonId [32]byte) (*types.Transaction, error) {
	return _FunctionsConsumer.contract.Transact(opts, "setDonId", newDonId)
}

// SetDonId is a paid mutator transaction binding the contract method 0x78ca5de7.
//
// Solidity: function setDonId(bytes32 newDonId) returns()
func (_FunctionsConsumer *FunctionsConsumerSession) SetDonId(newDonId [32]byte) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.SetDonId(&_FunctionsConsumer.TransactOpts, newDonId)
}

// SetDonId is a paid mutator transaction binding the contract method 0x78ca5de7.
//
// Solidity: function setDonId(bytes32 newDonId) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactorSession) SetDonId(newDonId [32]byte) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.SetDonId(&_FunctionsConsumer.TransactOpts, newDonId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FunctionsConsumer.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionsConsumer *FunctionsConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.TransferOwnership(&_FunctionsConsumer.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionsConsumer *FunctionsConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionsConsumer.Contract.TransferOwnership(&_FunctionsConsumer.TransactOpts, to)
}

// FunctionsConsumerFunctionsDataUpdatedIterator is returned from FilterFunctionsDataUpdated and is used to iterate over the raw logs and unpacked data for FunctionsDataUpdated events raised by the FunctionsConsumer contract.
type FunctionsConsumerFunctionsDataUpdatedIterator struct {
	Event *FunctionsConsumerFunctionsDataUpdated // Event containing the contract specifics and raw log

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
func (it *FunctionsConsumerFunctionsDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsConsumerFunctionsDataUpdated)
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
		it.Event = new(FunctionsConsumerFunctionsDataUpdated)
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
func (it *FunctionsConsumerFunctionsDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionsConsumerFunctionsDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionsConsumerFunctionsDataUpdated represents a FunctionsDataUpdated event raised by the FunctionsConsumer contract.
type FunctionsConsumerFunctionsDataUpdated struct {
	RequestId [32]byte
	Response  []byte
	Error     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFunctionsDataUpdated is a free log retrieval operation binding the contract event 0x40d2a764992b185672394ab8cfce9f9e86620d0e7489591a69f168c29896a545.
//
// Solidity: event FunctionsDataUpdated(bytes32 indexed requestId, bytes response, bytes error)
func (_FunctionsConsumer *FunctionsConsumerFilterer) FilterFunctionsDataUpdated(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionsConsumerFunctionsDataUpdatedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.FilterLogs(opts, "FunctionsDataUpdated", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerFunctionsDataUpdatedIterator{contract: _FunctionsConsumer.contract, event: "FunctionsDataUpdated", logs: logs, sub: sub}, nil
}

// WatchFunctionsDataUpdated is a free log subscription operation binding the contract event 0x40d2a764992b185672394ab8cfce9f9e86620d0e7489591a69f168c29896a545.
//
// Solidity: event FunctionsDataUpdated(bytes32 indexed requestId, bytes response, bytes error)
func (_FunctionsConsumer *FunctionsConsumerFilterer) WatchFunctionsDataUpdated(opts *bind.WatchOpts, sink chan<- *FunctionsConsumerFunctionsDataUpdated, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.WatchLogs(opts, "FunctionsDataUpdated", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionsConsumerFunctionsDataUpdated)
				if err := _FunctionsConsumer.contract.UnpackLog(event, "FunctionsDataUpdated", log); err != nil {
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

// ParseFunctionsDataUpdated is a log parse operation binding the contract event 0x40d2a764992b185672394ab8cfce9f9e86620d0e7489591a69f168c29896a545.
//
// Solidity: event FunctionsDataUpdated(bytes32 indexed requestId, bytes response, bytes error)
func (_FunctionsConsumer *FunctionsConsumerFilterer) ParseFunctionsDataUpdated(log types.Log) (*FunctionsConsumerFunctionsDataUpdated, error) {
	event := new(FunctionsConsumerFunctionsDataUpdated)
	if err := _FunctionsConsumer.contract.UnpackLog(event, "FunctionsDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionsConsumerOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the FunctionsConsumer contract.
type FunctionsConsumerOwnershipTransferRequestedIterator struct {
	Event *FunctionsConsumerOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *FunctionsConsumerOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsConsumerOwnershipTransferRequested)
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
		it.Event = new(FunctionsConsumerOwnershipTransferRequested)
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
func (it *FunctionsConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionsConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionsConsumerOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the FunctionsConsumer contract.
type FunctionsConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerOwnershipTransferRequestedIterator{contract: _FunctionsConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FunctionsConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionsConsumerOwnershipTransferRequested)
				if err := _FunctionsConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*FunctionsConsumerOwnershipTransferRequested, error) {
	event := new(FunctionsConsumerOwnershipTransferRequested)
	if err := _FunctionsConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionsConsumerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FunctionsConsumer contract.
type FunctionsConsumerOwnershipTransferredIterator struct {
	Event *FunctionsConsumerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FunctionsConsumerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsConsumerOwnershipTransferred)
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
		it.Event = new(FunctionsConsumerOwnershipTransferred)
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
func (it *FunctionsConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionsConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionsConsumerOwnershipTransferred represents a OwnershipTransferred event raised by the FunctionsConsumer contract.
type FunctionsConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionsConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerOwnershipTransferredIterator{contract: _FunctionsConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FunctionsConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionsConsumerOwnershipTransferred)
				if err := _FunctionsConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionsConsumer *FunctionsConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*FunctionsConsumerOwnershipTransferred, error) {
	event := new(FunctionsConsumerOwnershipTransferred)
	if err := _FunctionsConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionsConsumerRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the FunctionsConsumer contract.
type FunctionsConsumerRequestFulfilledIterator struct {
	Event *FunctionsConsumerRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *FunctionsConsumerRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsConsumerRequestFulfilled)
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
		it.Event = new(FunctionsConsumerRequestFulfilled)
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
func (it *FunctionsConsumerRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionsConsumerRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionsConsumerRequestFulfilled represents a RequestFulfilled event raised by the FunctionsConsumer contract.
type FunctionsConsumerRequestFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, id [][32]byte) (*FunctionsConsumerRequestFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.FilterLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerRequestFulfilledIterator{contract: _FunctionsConsumer.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *FunctionsConsumerRequestFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.WatchLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionsConsumerRequestFulfilled)
				if err := _FunctionsConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) ParseRequestFulfilled(log types.Log) (*FunctionsConsumerRequestFulfilled, error) {
	event := new(FunctionsConsumerRequestFulfilled)
	if err := _FunctionsConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionsConsumerRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the FunctionsConsumer contract.
type FunctionsConsumerRequestSentIterator struct {
	Event *FunctionsConsumerRequestSent // Event containing the contract specifics and raw log

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
func (it *FunctionsConsumerRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionsConsumerRequestSent)
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
		it.Event = new(FunctionsConsumerRequestSent)
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
func (it *FunctionsConsumerRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionsConsumerRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionsConsumerRequestSent represents a RequestSent event raised by the FunctionsConsumer contract.
type FunctionsConsumerRequestSent struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) FilterRequestSent(opts *bind.FilterOpts, id [][32]byte) (*FunctionsConsumerRequestSentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.FilterLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return &FunctionsConsumerRequestSentIterator{contract: _FunctionsConsumer.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *FunctionsConsumerRequestSent, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionsConsumer.contract.WatchLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionsConsumerRequestSent)
				if err := _FunctionsConsumer.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionsConsumer *FunctionsConsumerFilterer) ParseRequestSent(log types.Log) (*FunctionsConsumerRequestSent, error) {
	event := new(FunctionsConsumerRequestSent)
	if err := _FunctionsConsumer.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
