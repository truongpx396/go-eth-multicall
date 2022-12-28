// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiCall

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CustomMulticall2Call is an auto generated low-level Go binding around an user-defined struct.
type CustomMulticall2Call struct {
	Target   common.Address
	CallData []byte
}

// CustomMulticall2Result is an auto generated low-level Go binding around an user-defined struct.
type CustomMulticall2Result struct {
	Success    bool
	ReturnData []byte
}

// MultiCallABI is the input ABI used to generate the binding from.
const MultiCallABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"tryAggregateBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"userNativeBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structCustomMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MultiCall is an auto generated Go binding around an Ethereum contract.
type MultiCall struct {
	MultiCallCaller     // Read-only binding to the contract
	MultiCallTransactor // Write-only binding to the contract
	MultiCallFilterer   // Log filterer for contract events
}

// MultiCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCallSession struct {
	Contract     *MultiCall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallCallerSession struct {
	Contract *MultiCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultiCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCallTransactorSession struct {
	Contract     *MultiCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultiCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCallRaw struct {
	Contract *MultiCall // Generic contract binding to access the raw methods on
}

// MultiCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallCallerRaw struct {
	Contract *MultiCallCaller // Generic read-only contract binding to access the raw methods on
}

// MultiCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCallTransactorRaw struct {
	Contract *MultiCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCall creates a new instance of MultiCall, bound to a specific deployed contract.
func NewMultiCall(address common.Address, backend bind.ContractBackend) (*MultiCall, error) {
	contract, err := bindMultiCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCall{MultiCallCaller: MultiCallCaller{contract: contract}, MultiCallTransactor: MultiCallTransactor{contract: contract}, MultiCallFilterer: MultiCallFilterer{contract: contract}}, nil
}

// NewMultiCallCaller creates a new read-only instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallCaller(address common.Address, caller bind.ContractCaller) (*MultiCallCaller, error) {
	contract, err := bindMultiCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallCaller{contract: contract}, nil
}

// NewMultiCallTransactor creates a new write-only instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiCallTransactor, error) {
	contract, err := bindMultiCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallTransactor{contract: contract}, nil
}

// NewMultiCallFilterer creates a new log filterer instance of MultiCall, bound to a specific deployed contract.
func NewMultiCallFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiCallFilterer, error) {
	contract, err := bindMultiCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCallFilterer{contract: contract}, nil
}

// bindMultiCall binds a generic wrapper to an already deployed contract.
func bindMultiCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall *MultiCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall.Contract.MultiCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall *MultiCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall.Contract.MultiCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall *MultiCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall.Contract.MultiCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall *MultiCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall *MultiCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall *MultiCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall *MultiCallCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall *MultiCallSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultiCall.Contract.GetBlockHash(&_MultiCall.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall *MultiCallCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultiCall.Contract.GetBlockHash(&_MultiCall.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall *MultiCallCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall *MultiCallSession) GetBlockNumber() (*big.Int, error) {
	return _MultiCall.Contract.GetBlockNumber(&_MultiCall.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall *MultiCallCallerSession) GetBlockNumber() (*big.Int, error) {
	return _MultiCall.Contract.GetBlockNumber(&_MultiCall.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall *MultiCallCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall *MultiCallSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultiCall.Contract.GetCurrentBlockCoinbase(&_MultiCall.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall *MultiCallCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultiCall.Contract.GetCurrentBlockCoinbase(&_MultiCall.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall *MultiCallCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall *MultiCallSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockDifficulty(&_MultiCall.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall *MultiCallCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockDifficulty(&_MultiCall.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall *MultiCallCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall *MultiCallSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockGasLimit(&_MultiCall.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall *MultiCallCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockGasLimit(&_MultiCall.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall *MultiCallCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall *MultiCallSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockTimestamp(&_MultiCall.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall *MultiCallCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultiCall.Contract.GetCurrentBlockTimestamp(&_MultiCall.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall *MultiCallCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall *MultiCallSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultiCall.Contract.GetEthBalance(&_MultiCall.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall *MultiCallCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultiCall.Contract.GetEthBalance(&_MultiCall.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall *MultiCallCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultiCall.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall *MultiCallSession) GetLastBlockHash() ([32]byte, error) {
	return _MultiCall.Contract.GetLastBlockHash(&_MultiCall.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall *MultiCallCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MultiCall.Contract.GetLastBlockHash(&_MultiCall.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall *MultiCallTransactor) Aggregate(opts *bind.TransactOpts, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall *MultiCallSession) Aggregate(calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.Aggregate(&_MultiCall.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall *MultiCallTransactorSession) Aggregate(calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.Aggregate(&_MultiCall.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactor) BlockAndAggregate(opts *bind.TransactOpts, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallSession) BlockAndAggregate(calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.BlockAndAggregate(&_MultiCall.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactorSession) BlockAndAggregate(calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.BlockAndAggregate(&_MultiCall.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall *MultiCallSession) TryAggregate(requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.TryAggregate(&_MultiCall.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactorSession) TryAggregate(requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.TryAggregate(&_MultiCall.TransactOpts, requireSuccess, calls)
}

// TryAggregateBalances is a paid mutator transaction binding the contract method 0xc92d8490.
//
// Solidity: function tryAggregateBalances(bool requireSuccess, (address,bytes)[] calls, address userAddress) returns((bool,bytes)[] returnData, uint256 userNativeBalance)
func (_MultiCall *MultiCallTransactor) TryAggregateBalances(opts *bind.TransactOpts, requireSuccess bool, calls []CustomMulticall2Call, userAddress common.Address) (*types.Transaction, error) {
	return _MultiCall.contract.Transact(opts, "tryAggregateBalances", requireSuccess, calls, userAddress)
}

// TryAggregateBalances is a paid mutator transaction binding the contract method 0xc92d8490.
//
// Solidity: function tryAggregateBalances(bool requireSuccess, (address,bytes)[] calls, address userAddress) returns((bool,bytes)[] returnData, uint256 userNativeBalance)
func (_MultiCall *MultiCallSession) TryAggregateBalances(requireSuccess bool, calls []CustomMulticall2Call, userAddress common.Address) (*types.Transaction, error) {
	return _MultiCall.Contract.TryAggregateBalances(&_MultiCall.TransactOpts, requireSuccess, calls, userAddress)
}

// TryAggregateBalances is a paid mutator transaction binding the contract method 0xc92d8490.
//
// Solidity: function tryAggregateBalances(bool requireSuccess, (address,bytes)[] calls, address userAddress) returns((bool,bytes)[] returnData, uint256 userNativeBalance)
func (_MultiCall *MultiCallTransactorSession) TryAggregateBalances(requireSuccess bool, calls []CustomMulticall2Call, userAddress common.Address) (*types.Transaction, error) {
	return _MultiCall.Contract.TryAggregateBalances(&_MultiCall.TransactOpts, requireSuccess, calls, userAddress)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallSession) TryBlockAndAggregate(requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.TryBlockAndAggregate(&_MultiCall.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall *MultiCallTransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []CustomMulticall2Call) (*types.Transaction, error) {
	return _MultiCall.Contract.TryBlockAndAggregate(&_MultiCall.TransactOpts, requireSuccess, calls)
}
