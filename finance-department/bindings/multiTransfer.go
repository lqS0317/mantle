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

// MultiTransferBITContractMetaData contains all meta data concerning the MultiTransferBITContract contract.
var MultiTransferBITContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"MultiTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"multiTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MultiTransferBITContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiTransferBITContractMetaData.ABI instead.
var MultiTransferBITContractABI = MultiTransferBITContractMetaData.ABI

// MultiTransferBITContract is an auto generated Go binding around an Ethereum contract.
type MultiTransferBITContract struct {
	MultiTransferBITContractCaller     // Read-only binding to the contract
	MultiTransferBITContractTransactor // Write-only binding to the contract
	MultiTransferBITContractFilterer   // Log filterer for contract events
}

// MultiTransferBITContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiTransferBITContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTransferBITContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiTransferBITContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTransferBITContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiTransferBITContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTransferBITContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiTransferBITContractSession struct {
	Contract     *MultiTransferBITContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultiTransferBITContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiTransferBITContractCallerSession struct {
	Contract *MultiTransferBITContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MultiTransferBITContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiTransferBITContractTransactorSession struct {
	Contract     *MultiTransferBITContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MultiTransferBITContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiTransferBITContractRaw struct {
	Contract *MultiTransferBITContract // Generic contract binding to access the raw methods on
}

// MultiTransferBITContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiTransferBITContractCallerRaw struct {
	Contract *MultiTransferBITContractCaller // Generic read-only contract binding to access the raw methods on
}

// MultiTransferBITContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiTransferBITContractTransactorRaw struct {
	Contract *MultiTransferBITContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiTransferBITContract creates a new instance of MultiTransferBITContract, bound to a specific deployed contract.
func NewMultiTransferBITContract(address common.Address, backend bind.ContractBackend) (*MultiTransferBITContract, error) {
	contract, err := bindMultiTransferBITContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiTransferBITContract{MultiTransferBITContractCaller: MultiTransferBITContractCaller{contract: contract}, MultiTransferBITContractTransactor: MultiTransferBITContractTransactor{contract: contract}, MultiTransferBITContractFilterer: MultiTransferBITContractFilterer{contract: contract}}, nil
}

// NewMultiTransferBITContractCaller creates a new read-only instance of MultiTransferBITContract, bound to a specific deployed contract.
func NewMultiTransferBITContractCaller(address common.Address, caller bind.ContractCaller) (*MultiTransferBITContractCaller, error) {
	contract, err := bindMultiTransferBITContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTransferBITContractCaller{contract: contract}, nil
}

// NewMultiTransferBITContractTransactor creates a new write-only instance of MultiTransferBITContract, bound to a specific deployed contract.
func NewMultiTransferBITContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTransferBITContractTransactor, error) {
	contract, err := bindMultiTransferBITContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTransferBITContractTransactor{contract: contract}, nil
}

// NewMultiTransferBITContractFilterer creates a new log filterer instance of MultiTransferBITContract, bound to a specific deployed contract.
func NewMultiTransferBITContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiTransferBITContractFilterer, error) {
	contract, err := bindMultiTransferBITContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiTransferBITContractFilterer{contract: contract}, nil
}

// bindMultiTransferBITContract binds a generic wrapper to an already deployed contract.
func bindMultiTransferBITContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiTransferBITContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiTransferBITContract *MultiTransferBITContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTransferBITContract.Contract.MultiTransferBITContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiTransferBITContract *MultiTransferBITContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.MultiTransferBITContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiTransferBITContract *MultiTransferBITContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.MultiTransferBITContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiTransferBITContract *MultiTransferBITContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiTransferBITContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiTransferBITContract *MultiTransferBITContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiTransferBITContract *MultiTransferBITContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.contract.Transact(opts, method, params...)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xa16a3179.
//
// Solidity: function multiTransfer(address[] _addresses, uint256 _amount) payable returns()
func (_MultiTransferBITContract *MultiTransferBITContractTransactor) MultiTransfer(opts *bind.TransactOpts, _addresses []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MultiTransferBITContract.contract.Transact(opts, "multiTransfer", _addresses, _amount)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xa16a3179.
//
// Solidity: function multiTransfer(address[] _addresses, uint256 _amount) payable returns()
func (_MultiTransferBITContract *MultiTransferBITContractSession) MultiTransfer(_addresses []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.MultiTransfer(&_MultiTransferBITContract.TransactOpts, _addresses, _amount)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xa16a3179.
//
// Solidity: function multiTransfer(address[] _addresses, uint256 _amount) payable returns()
func (_MultiTransferBITContract *MultiTransferBITContractTransactorSession) MultiTransfer(_addresses []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MultiTransferBITContract.Contract.MultiTransfer(&_MultiTransferBITContract.TransactOpts, _addresses, _amount)
}

// MultiTransferBITContractMultiTransferIterator is returned from FilterMultiTransfer and is used to iterate over the raw logs and unpacked data for MultiTransfer events raised by the MultiTransferBITContract contract.
type MultiTransferBITContractMultiTransferIterator struct {
	Event *MultiTransferBITContractMultiTransfer // Event containing the contract specifics and raw log

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
func (it *MultiTransferBITContractMultiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiTransferBITContractMultiTransfer)
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
		it.Event = new(MultiTransferBITContractMultiTransfer)
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
func (it *MultiTransferBITContractMultiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiTransferBITContractMultiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiTransferBITContractMultiTransfer represents a MultiTransfer event raised by the MultiTransferBITContract contract.
type MultiTransferBITContractMultiTransfer struct {
	From   common.Address
	Value  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMultiTransfer is a free log retrieval operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_MultiTransferBITContract *MultiTransferBITContractFilterer) FilterMultiTransfer(opts *bind.FilterOpts, _from []common.Address, _value []*big.Int) (*MultiTransferBITContractMultiTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _MultiTransferBITContract.contract.FilterLogs(opts, "MultiTransfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &MultiTransferBITContractMultiTransferIterator{contract: _MultiTransferBITContract.contract, event: "MultiTransfer", logs: logs, sub: sub}, nil
}

// WatchMultiTransfer is a free log subscription operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_MultiTransferBITContract *MultiTransferBITContractFilterer) WatchMultiTransfer(opts *bind.WatchOpts, sink chan<- *MultiTransferBITContractMultiTransfer, _from []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _MultiTransferBITContract.contract.WatchLogs(opts, "MultiTransfer", _fromRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiTransferBITContractMultiTransfer)
				if err := _MultiTransferBITContract.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
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

// ParseMultiTransfer is a log parse operation binding the contract event 0x319e0008dcdeba1f31169497fb0f35d31b2b2f481c714d1f50640e86ac6c3bb0.
//
// Solidity: event MultiTransfer(address indexed _from, uint256 indexed _value, address _to, uint256 _amount)
func (_MultiTransferBITContract *MultiTransferBITContractFilterer) ParseMultiTransfer(log types.Log) (*MultiTransferBITContractMultiTransfer, error) {
	event := new(MultiTransferBITContractMultiTransfer)
	if err := _MultiTransferBITContract.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
