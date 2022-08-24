// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package backends

import (
	"bytes"
	"context"
	"github.com/bitdao-io/bitnetwork/l2geth/rollup/rcfg"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
	"time"

	ethereum "github.com/bitdao-io/bitnetwork/l2geth"
	"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi"
	"github.com/bitdao-io/bitnetwork/l2geth/accounts/abi/bind"
	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/core"
	"github.com/bitdao-io/bitnetwork/l2geth/core/types"
	"github.com/bitdao-io/bitnetwork/l2geth/crypto"
	"github.com/bitdao-io/bitnetwork/l2geth/params"
)

func TestSimulatedBackend(t *testing.T) {
	var gasLimit uint64 = 8000029
	key, _ := crypto.GenerateKey() // nolint: gosec
	auth := bind.NewKeyedTransactor(key)
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}

	sim := NewSimulatedBackend(genAlloc, gasLimit)
	defer sim.Close()

	// should return an error if the tx is not found
	txHash := common.HexToHash("2")
	_, isPending, err := sim.TransactionByHash(context.Background(), txHash)

	if isPending {
		t.Fatal("transaction should not be pending")
	}
	if err != ethereum.NotFound {
		t.Fatalf("err should be `ethereum.NotFound` but received %v", err)
	}

	// generate a transaction and confirm you can retrieve it
	code := `6060604052600a8060106000396000f360606040526008565b00`
	var gas uint64 = 3000000
	tx := types.NewContractCreation(0, big.NewInt(0), gas, big.NewInt(1), common.FromHex(code))
	tx, _ = types.SignTx(tx, types.HomesteadSigner{}, key)

	err = sim.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal("error sending transaction")
	}

	txHash = tx.Hash()
	_, isPending, err = sim.TransactionByHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("error getting transaction with hash: %v", txHash.String())
	}
	if !isPending {
		t.Fatal("transaction should have pending status")
	}

	sim.Commit()
	_, isPending, err = sim.TransactionByHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("error getting transaction with hash: %v", txHash.String())
	}
	if isPending {
		t.Fatal("transaction should not have pending status")
	}
}

var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

//  the following is based on this contract:
//  contract T {
//  	event received(address sender, uint amount, bytes memo);
//  	event receivedAddr(address sender);
//
//  	function receive(bytes calldata memo) external payable returns (string memory res) {
//  		emit received(msg.sender, msg.value, memo);
//  		emit receivedAddr(msg.sender);
//		    return "hello world";
//  	}
//  }
const abiJSON = `[ { "constant": false, "inputs": [ { "name": "memo", "type": "bytes" } ], "name": "receive", "outputs": [ { "name": "res", "type": "string" } ], "payable": true, "stateMutability": "payable", "type": "function" }, { "anonymous": false, "inputs": [ { "indexed": false, "name": "sender", "type": "address" }, { "indexed": false, "name": "amount", "type": "uint256" }, { "indexed": false, "name": "memo", "type": "bytes" } ], "name": "received", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": false, "name": "sender", "type": "address" } ], "name": "receivedAddr", "type": "event" } ]`
const abiBin = `0x608060405234801561001057600080fd5b506102a0806100206000396000f3fe60806040526004361061003b576000357c010000000000000000000000000000000000000000000000000000000090048063a69b6ed014610040575b600080fd5b6100b76004803603602081101561005657600080fd5b810190808035906020019064010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b9091929391929390505050610132565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f75780820151818401526020810190506100dc565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60607f75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed33348585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509550505050505060405180910390a17f46923992397eac56cf13058aced2a1871933622717e27b24eabc13bf9dd329c833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a16040805190810160405280600b81526020017f68656c6c6f20776f726c6400000000000000000000000000000000000000000081525090509291505056fea165627a7a72305820ff0c57dad254cfeda48c9cfb47f1353a558bccb4d1bc31da1dae69315772d29e0029`
const deployedCode = `60806040526004361061003b576000357c010000000000000000000000000000000000000000000000000000000090048063a69b6ed014610040575b600080fd5b6100b76004803603602081101561005657600080fd5b810190808035906020019064010000000081111561007357600080fd5b82018360208201111561008557600080fd5b803590602001918460018302840111640100000000831117156100a757600080fd5b9091929391929390505050610132565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f75780820151818401526020810190506100dc565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60607f75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed33348585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509550505050505060405180910390a17f46923992397eac56cf13058aced2a1871933622717e27b24eabc13bf9dd329c833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a16040805190810160405280600b81526020017f68656c6c6f20776f726c6400000000000000000000000000000000000000000081525090509291505056fea165627a7a72305820ff0c57dad254cfeda48c9cfb47f1353a558bccb4d1bc31da1dae69315772d29e0029`

// expected return value contains "hello world"
var expectedReturn = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func TestNewSimulatedBackend(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	expectedBal := big.NewInt(10000000000)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: expectedBal},
		}, 10000000,
	)
	defer sim.Close()

	if sim.config != params.AllEthashProtocolChanges {
		t.Errorf("expected sim config to equal params.AllEthashProtocolChanges, got %v", sim.config)
	}

	if sim.blockchain.Config() != params.AllEthashProtocolChanges {
		t.Errorf("expected sim blockchain config to equal params.AllEthashProtocolChanges, got %v", sim.config)
	}

	statedb, _ := sim.blockchain.State()
	bal := statedb.GetBalance(testAddr)
	if bal.Cmp(expectedBal) != 0 {
		t.Errorf("expected balance for test address not received. expected: %v actual: %v", expectedBal, bal)
	}
}

func TestSimulatedBackend_AdjustTime(t *testing.T) {
	sim := NewSimulatedBackend(
		core.GenesisAlloc{}, 10000000,
	)
	defer sim.Close()

	prevTime := sim.pendingBlock.Time()
	err := sim.AdjustTime(time.Second)
	if err != nil {
		t.Error(err)
	}
	newTime := sim.pendingBlock.Time()

	if newTime-prevTime != uint64(time.Second.Seconds()) {
		t.Errorf("adjusted time not equal to a second. prev: %v, new: %v", prevTime, newTime)
	}
}

func TestSimulatedBackend_BalanceAt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	expectedBal := big.NewInt(10000000000)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: expectedBal},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	bal, err := sim.BalanceAt(bgCtx, testAddr, nil)
	if err != nil {
		t.Error(err)
	}

	if bal.Cmp(expectedBal) != 0 {
		t.Errorf("expected balance for test address not received. expected: %v actual: %v", expectedBal, bal)
	}
}

func TestSimulatedBackend_BlockByHash(t *testing.T) {
	sim := NewSimulatedBackend(
		core.GenesisAlloc{}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	block, err := sim.BlockByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}
	blockByHash, err := sim.BlockByHash(bgCtx, block.Hash())
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}

	if block.Hash() != blockByHash.Hash() {
		t.Errorf("did not get expected block")
	}
}

func TestSimulatedBackend_BlockByNumber(t *testing.T) {
	sim := NewSimulatedBackend(
		core.GenesisAlloc{}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	block, err := sim.BlockByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}
	if block.NumberU64() != 0 {
		t.Errorf("did not get most recent block, instead got block number %v", block.NumberU64())
	}

	// create one block
	sim.Commit()

	block, err = sim.BlockByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}
	if block.NumberU64() != 1 {
		t.Errorf("did not get most recent block, instead got block number %v", block.NumberU64())
	}

	blockByNumber, err := sim.BlockByNumber(bgCtx, big.NewInt(1))
	if err != nil {
		t.Errorf("could not get block by number: %v", err)
	}
	if blockByNumber.Hash() != block.Hash() {
		t.Errorf("did not get the same block with height of 1 as before")
	}
}

func TestSimulatedBackend_NonceAt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	nonce, err := sim.NonceAt(bgCtx, testAddr, big.NewInt(0))
	if err != nil {
		t.Errorf("could not get nonce for test addr: %v", err)
	}

	if nonce != uint64(0) {
		t.Errorf("received incorrect nonce. expected 0, got %v", nonce)
	}

	// create a signed transaction to send
	tx := types.NewTransaction(nonce, testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}
	sim.Commit()

	newNonce, err := sim.NonceAt(bgCtx, testAddr, big.NewInt(1))
	if err != nil {
		t.Errorf("could not get nonce for test addr: %v", err)
	}

	if newNonce != nonce+uint64(1) {
		t.Errorf("received incorrect nonce. expected 1, got %v", nonce)
	}
}

func TestSimulatedBackend_SendTransaction(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}
	sim.Commit()

	block, err := sim.BlockByNumber(bgCtx, big.NewInt(1))
	if err != nil {
		t.Errorf("could not get block at height 1: %v", err)
	}

	if signedTx.Hash() != block.Transactions()[0].Hash() {
		t.Errorf("did not commit sent transaction. expected hash %v got hash %v", block.Transactions()[0].Hash(), signedTx.Hash())
	}
}

func TestSimulatedBackend_TransactionByHash(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}

	// ensure tx is committed pending
	receivedTx, pending, err := sim.TransactionByHash(bgCtx, signedTx.Hash())
	if err != nil {
		t.Errorf("could not get transaction by hash %v: %v", signedTx.Hash(), err)
	}
	if !pending {
		t.Errorf("expected transaction to be in pending state")
	}
	if receivedTx.Hash() != signedTx.Hash() {
		t.Errorf("did not received committed transaction. expected hash %v got hash %v", signedTx.Hash(), receivedTx.Hash())
	}

	sim.Commit()

	// ensure tx is not and committed pending
	receivedTx, pending, err = sim.TransactionByHash(bgCtx, signedTx.Hash())
	if err != nil {
		t.Errorf("could not get transaction by hash %v: %v", signedTx.Hash(), err)
	}
	if pending {
		t.Errorf("expected transaction to not be in pending state")
	}
	if receivedTx.Hash() != signedTx.Hash() {
		t.Errorf("did not received committed transaction. expected hash %v got hash %v", signedTx.Hash(), receivedTx.Hash())
	}
}

func TestSimulatedBackend_EstimateGas(t *testing.T) {
	t.Skip("BVM breaks this because gas consumption is not yet standardized")

	sim := NewSimulatedBackend(
		core.GenesisAlloc{}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	gas, err := sim.EstimateGas(bgCtx, ethereum.CallMsg{
		From:  testAddr,
		To:    &testAddr,
		Value: big.NewInt(1000),
		Data:  []byte{},
	})
	if err != nil {
		t.Errorf("could not estimate gas: %v", err)
	}

	if gas != params.TxGas {
		t.Errorf("expected 21000 gas cost for a transaction got %v", gas)
	}
}

func TestSimulatedBackend_HeaderByHash(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	header, err := sim.HeaderByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}
	headerByHash, err := sim.HeaderByHash(bgCtx, header.Hash())
	if err != nil {
		t.Errorf("could not get recent block: %v", err)
	}

	if header.Hash() != headerByHash.Hash() {
		t.Errorf("did not get expected block")
	}
}

func TestSimulatedBackend_HeaderByNumber(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	latestBlockHeader, err := sim.HeaderByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get header for tip of chain: %v", err)
	}
	if latestBlockHeader == nil {
		t.Errorf("received a nil block header")
	}
	if latestBlockHeader.Number.Uint64() != uint64(0) {
		t.Errorf("expected block header number 0, instead got %v", latestBlockHeader.Number.Uint64())
	}

	sim.Commit()

	latestBlockHeader, err = sim.HeaderByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get header for blockheight of 1: %v", err)
	}

	blockHeader, err := sim.HeaderByNumber(bgCtx, big.NewInt(1))
	if err != nil {
		t.Errorf("could not get header for blockheight of 1: %v", err)
	}

	if blockHeader.Hash() != latestBlockHeader.Hash() {
		t.Errorf("block header and latest block header are not the same")
	}
	if blockHeader.Number.Int64() != int64(1) {
		t.Errorf("did not get blockheader for block 1. instead got block %v", blockHeader.Number.Int64())
	}

	block, err := sim.BlockByNumber(bgCtx, big.NewInt(1))
	if err != nil {
		t.Errorf("could not get block for blockheight of 1: %v", err)
	}

	if block.Hash() != blockHeader.Hash() {
		t.Errorf("block hash and block header hash do not match. expected %v, got %v", block.Hash(), blockHeader.Hash())
	}
}

func TestSimulatedBackend_TransactionCount(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()
	currentBlock, err := sim.BlockByNumber(bgCtx, nil)
	if err != nil || currentBlock == nil {
		t.Error("could not get current block")
	}

	count, err := sim.TransactionCount(bgCtx, currentBlock.Hash())
	if err != nil {
		t.Error("could not get current block's transaction count")
	}

	if count != 0 {
		t.Errorf("expected transaction count of %v does not match actual count of %v", 0, count)
	}

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}

	sim.Commit()

	lastBlock, err := sim.BlockByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get header for tip of chain: %v", err)
	}

	count, err = sim.TransactionCount(bgCtx, lastBlock.Hash())
	if err != nil {
		t.Error("could not get current block's transaction count")
	}

	if count != 1 {
		t.Errorf("expected transaction count of %v does not match actual count of %v", 1, count)
	}
}

func TestSimulatedBackend_TransactionInBlock(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	transaction, err := sim.TransactionInBlock(bgCtx, sim.pendingBlock.Hash(), uint(0))
	if err == nil && err != errTransactionDoesNotExist {
		t.Errorf("expected a transaction does not exist error to be received but received %v", err)
	}
	if transaction != nil {
		t.Errorf("expected transaction to be nil but received %v", transaction)
	}

	// expect pending nonce to be 0 since account has not been used
	pendingNonce, err := sim.PendingNonceAt(bgCtx, testAddr)
	if err != nil {
		t.Errorf("did not get the pending nonce: %v", err)
	}

	if pendingNonce != uint64(0) {
		t.Errorf("expected pending nonce of 0 got %v", pendingNonce)
	}

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}

	sim.Commit()

	lastBlock, err := sim.BlockByNumber(bgCtx, nil)
	if err != nil {
		t.Errorf("could not get header for tip of chain: %v", err)
	}

	transaction, err = sim.TransactionInBlock(bgCtx, lastBlock.Hash(), uint(1))
	if err == nil && err != errTransactionDoesNotExist {
		t.Errorf("expected a transaction does not exist error to be received but received %v", err)
	}
	if transaction != nil {
		t.Errorf("expected transaction to be nil but received %v", transaction)
	}

	transaction, err = sim.TransactionInBlock(bgCtx, lastBlock.Hash(), uint(0))
	if err != nil {
		t.Errorf("could not get transaction in the lastest block with hash %v: %v", lastBlock.Hash().String(), err)
	}

	if signedTx.Hash().String() != transaction.Hash().String() {
		t.Errorf("received transaction that did not match the sent transaction. expected hash %v, got hash %v", signedTx.Hash().String(), transaction.Hash().String())
	}
}

func TestSimulatedBackend_PendingNonceAt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	// expect pending nonce to be 0 since account has not been used
	pendingNonce, err := sim.PendingNonceAt(bgCtx, testAddr)
	if err != nil {
		t.Errorf("did not get the pending nonce: %v", err)
	}

	if pendingNonce != uint64(0) {
		t.Errorf("expected pending nonce of 0 got %v", pendingNonce)
	}

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}

	// expect pending nonce to be 1 since account has submitted one transaction
	pendingNonce, err = sim.PendingNonceAt(bgCtx, testAddr)
	if err != nil {
		t.Errorf("did not get the pending nonce: %v", err)
	}

	if pendingNonce != uint64(1) {
		t.Errorf("expected pending nonce of 1 got %v", pendingNonce)
	}

	// make a new transaction with a nonce of 1
	tx = types.NewTransaction(uint64(1), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err = types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not send tx: %v", err)
	}

	// expect pending nonce to be 2 since account now has two transactions
	pendingNonce, err = sim.PendingNonceAt(bgCtx, testAddr)
	if err != nil {
		t.Errorf("did not get the pending nonce: %v", err)
	}

	if pendingNonce != uint64(2) {
		t.Errorf("expected pending nonce of 2 got %v", pendingNonce)
	}
}

func TestSimulatedBackend_TransactionReceipt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)

	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		}, 10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	// create a signed transaction to send
	tx := types.NewTransaction(uint64(0), testAddr, big.NewInt(1000), params.TxGas, big.NewInt(1), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testKey)
	if err != nil {
		t.Errorf("could not sign tx: %v", err)
	}

	// send tx to simulated backend
	err = sim.SendTransaction(bgCtx, signedTx)
	if err != nil {
		t.Errorf("could not add tx to pending block: %v", err)
	}
	sim.Commit()

	receipt, err := sim.TransactionReceipt(bgCtx, signedTx.Hash())
	if err != nil {
		t.Errorf("could not get transaction receipt: %v", err)
	}

	if receipt.ContractAddress != testAddr && receipt.TxHash != signedTx.Hash() {
		t.Errorf("received receipt is not correct: %v", receipt)
	}
}

func TestSimulatedBackend_SuggestGasPrice(t *testing.T) {
	sim := NewSimulatedBackend(
		core.GenesisAlloc{},
		10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()
	gasPrice, err := sim.SuggestGasPrice(bgCtx)
	if err != nil {
		t.Errorf("could not get gas price: %v", err)
	}
	if gasPrice.Uint64() != uint64(1) {
		t.Errorf("gas price was not expected value of 1. actual: %v", gasPrice.Uint64())
	}
}

func TestSimulatedBackend_PendingCodeAt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		},
		10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()
	code, err := sim.CodeAt(bgCtx, testAddr, nil)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	if len(code) != 0 {
		t.Errorf("got code for account that does not have contract code")
	}

	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	auth := bind.NewKeyedTransactor(testKey)
	contractAddr, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(abiBin), sim)
	if err != nil {
		t.Errorf("could not deploy contract: %v tx: %v contract: %v", err, tx, contract)
	}

	code, err = sim.PendingCodeAt(bgCtx, contractAddr)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	if len(code) == 0 {
		t.Errorf("did not get code for account that has contract code")
	}
	// ensure code received equals code deployed
	if !bytes.Equal(code, common.FromHex(deployedCode)) {
		t.Errorf("code received did not match expected deployed code:\n expected %v\n actual %v", common.FromHex(deployedCode), code)
	}
}

func TestSimulatedBackend_CodeAt(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		},
		10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()
	code, err := sim.CodeAt(bgCtx, testAddr, nil)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	if len(code) != 0 {
		t.Errorf("got code for account that does not have contract code")
	}

	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	auth := bind.NewKeyedTransactor(testKey)
	contractAddr, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(abiBin), sim)
	if err != nil {
		t.Errorf("could not deploy contract: %v tx: %v contract: %v", err, tx, contract)
	}

	sim.Commit()
	code, err = sim.CodeAt(bgCtx, contractAddr, nil)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	if len(code) == 0 {
		t.Errorf("did not get code for account that has contract code")
	}
	// ensure code received equals code deployed
	if !bytes.Equal(code, common.FromHex(deployedCode)) {
		t.Errorf("code received did not match expected deployed code:\n expected %v\n actual %v", common.FromHex(deployedCode), code)
	}
}

// When receive("X") is called with sender 0x00... and value 1, it produces this tx receipt:
//   receipt{status=1 cgas=23949 bloom=00000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000040200000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 logs=[log: b6818c8064f645cd82d99b59a1a267d6d61117ef [75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed] 000000000000000000000000376c47978271565f56deb45495afa69e59c16ab200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000158 9ae378b6d4409eada347a5dc0c180f186cb62dc68fcc0f043425eb917335aa28 0 95d429d309bb9d753954195fe2d69bd140b4ae731b9b5b605c34323de162cf00 0]}
func TestSimulatedBackend_PendingAndCallContract(t *testing.T) {
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: big.NewInt(10000000000)},
		},
		10000000,
	)
	defer sim.Close()
	bgCtx := context.Background()

	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}
	contractAuth := bind.NewKeyedTransactor(testKey)
	addr, _, _, err := bind.DeployContract(contractAuth, parsed, common.FromHex(abiBin), sim)
	if err != nil {
		t.Errorf("could not deploy contract: %v", err)
	}

	input, err := parsed.Pack("receive", []byte("X"))
	if err != nil {
		t.Errorf("could pack receive function on contract: %v", err)
	}

	// make sure you can call the contract in pending state
	res, err := sim.PendingCallContract(bgCtx, ethereum.CallMsg{
		From: testAddr,
		To:   &addr,
		Data: input,
	})
	if err != nil {
		t.Errorf("could not call receive method on contract: %v", err)
	}
	if len(res) == 0 {
		t.Errorf("result of contract call was empty: %v", res)
	}

	// while comparing against the byte array is more exact, also compare against the human readable string for readability
	if !bytes.Equal(res, expectedReturn) || !strings.Contains(string(res), "hello world") {
		t.Errorf("response from calling contract was expected to be 'hello world' instead received %v", string(res))
	}

	sim.Commit()

	// make sure you can call the contract
	res, err = sim.CallContract(bgCtx, ethereum.CallMsg{
		From: testAddr,
		To:   &addr,
		Data: input,
	}, nil)
	if err != nil {
		t.Errorf("could not call receive method on contract: %v", err)
	}
	if len(res) == 0 {
		t.Errorf("result of contract call was empty: %v", res)
	}

	if !bytes.Equal(res, expectedReturn) || !strings.Contains(string(res), "hello world") {
		t.Errorf("response from calling contract was expected to be 'hello world' instead received %v", string(res))
	}
}

const tssAddr = "0x4200000000000000000000000000000000000020"
const tssAbi = "[\n    {\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"_deadAddress\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"address payable\",\n          \"name\": \"_owner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"constructor\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"blockStartHeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"length\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"address[]\",\n          \"name\": \"tssMembers\",\n          \"type\": \"address[]\"\n        }\n      ],\n      \"name\": \"DistributeTssReward\",\n      \"type\": \"event\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"bestBlockID\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"_blockStartHeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint32\",\n          \"name\": \"_length\",\n          \"type\": \"uint32\"\n        },\n        {\n          \"internalType\": \"address[]\",\n          \"name\": \"_tssMembers\",\n          \"type\": \"address[]\"\n        }\n      ],\n      \"name\": \"claimReward\",\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"deadAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"dust\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"ledger\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"owner\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address payable\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"queryReward\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"totalAmount\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"_blockID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"_amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"updateReward\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"withdraw\",\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"inputs\": [],\n      \"name\": \"withdrawDust\",\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    }\n  ]"
const tssBin = "0x608060405234801561001057600080fd5b50604051610dc9380380610dc983398101604081905261002f91610078565b600180546001600160a01b039384166001600160a01b031991821617909155600280549290931691161790556100b2565b6001600160a01b038116811461007557600080fd5b50565b6000806040838503121561008b57600080fd5b825161009681610060565b60208401519092506100a781610060565b809150509250929050565b610d08806100c16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80632c79db1111610081578063cfb550f11161005b578063cfb550f1146101ad578063e04f6e35146101b5578063fad9aba3146101c857600080fd5b80632c79db111461017b5780633ccfd60b146101835780638da5cb5b1461018d57600080fd5b806319d509a1116100b257806319d509a1146101245780631a39d8ef1461012d57806327c8f8351461013657600080fd5b80630b50cd3e146100ce57806310a7fd7b146100f6575b600080fd5b6100e16100dc366004610a4a565b6101d1565b60405190151581526020015b60405180910390f35b610116610104366004610a6c565b60006020819052908152604090205481565b6040519081526020016100ed565b61011660035481565b61011660055481565b6001546101569073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b6101166103b3565b61018b61044c565b005b6002546101569073ffffffffffffffffffffffffffffffffffffffff1681565b61018b6105da565b61018b6101c3366004610a85565b61077e565b61011660045481565b60015460009073ffffffffffffffffffffffffffffffffffffffff163314610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f747373207265776172642063616c6c206d65737361676520756e61757468656e60448201527f746963617465640000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600554471015610312576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b600354610320906001610b49565b8314610388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f626c6f636b2069642075706461746520696c6c6567616c0000000000000000006044820152606401610277565b600383905560055461039a9083610a1f565b6005555060009182526020829052604090912055600190565b6000600554471015610447576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b504790565b60025473ffffffffffffffffffffffffffffffffffffffff1633146104f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460448201527f68697320636f6e747261637400000000000000000000000000000000000000006064820152608401610277565b600554471015610585576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b600060055547156105d85760025460405173ffffffffffffffffffffffffffffffffffffffff909116904780156108fc02916000818181858888f193505050501580156105d6573d6000803e3d6000fd5b505b565b60025473ffffffffffffffffffffffffffffffffffffffff163314610681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460448201527f68697320636f6e747261637400000000000000000000000000000000000000006064820152608401610277565b600554471015610713576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b6004546005546107239082610a32565b600555600060045580156105d65760025460045460405173ffffffffffffffffffffffffffffffffffffffff9092169181156108fc0291906000818181858888f1935050505015801561077a573d6000803e3d6000fd5b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610825576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f747373207265776172642063616c6c206d65737361676520756e61757468656e60448201527f74696361746564000000000000000000000000000000000000000000000000006064820152608401610277565b6005544710156108b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f62616c616e6365207265636f726420616e6420636f6e74726163742062616c6160448201527f6e636520617265206e6f7420657175616c0000000000000000000000000000006064820152608401610277565b6000848152602081905260408120549080865b8663ffffffff168110156109d8576108e28486610a3e565b60008981526020819052604081208190559093505b8581101561099f57600087878381811061091357610913610b61565b90506020020160208101906109289190610bb9565b90506109348486610a1f565b6005549094506109449086610a32565b60055560405173ffffffffffffffffffffffffffffffffffffffff82169086156108fc029087906000818181858888f1935050505015801561098a573d6000803e3d6000fd5b5050808061099790610bd4565b9150506108f7565b5060006109ac8584610a32565b905080156109c5576004546109c19082610a1f565b6004555b50806109d081610bd4565b9150506108ca565b507ff630cba6d450d736e85735388d4fe67a177b8a3685cdd7dee2bea7727b47860a87878787604051610a0e9493929190610c0d565b60405180910390a150505050505050565b6000610a2b8284610b49565b9392505050565b6000610a2b8284610c80565b6000610a2b8284610c97565b60008060408385031215610a5d57600080fd5b50508035926020909101359150565b600060208284031215610a7e57600080fd5b5035919050565b60008060008060608587031215610a9b57600080fd5b84359350602085013563ffffffff81168114610ab657600080fd5b9250604085013567ffffffffffffffff80821115610ad357600080fd5b818701915087601f830112610ae757600080fd5b813581811115610af657600080fd5b8860208260051b8501011115610b0b57600080fd5b95989497505060200194505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610b5c57610b5c610b1a565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610bb457600080fd5b919050565b600060208284031215610bcb57600080fd5b610a2b82610b90565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610c0657610c06610b1a565b5060010190565b600060608201868352602063ffffffff871681850152606060408501528185835260808501905086925060005b86811015610c735773ffffffffffffffffffffffffffffffffffffffff610c6085610b90565b1682529282019290820190600101610c3a565b5098975050505050505050565b600082821015610c9257610c92610b1a565b500390565b600082610ccd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220a51ec743c0bb2b95aac6f7675b0aa3cbda567cf1153f306a8933bd643a1dc0cd64736f6c63430008090033"

func TestNewSimulatedBackend_SystemCall(t *testing.T) {
	rcfg.UsingBVM = true
	testAddr := crypto.PubkeyToAddress(testKey.PublicKey)
	toAddr := common.HexToAddress(tssAddr)
	expectedBal := big.NewInt(10000000000)
	sim := NewSimulatedBackend(
		core.GenesisAlloc{
			testAddr: {Balance: expectedBal},
		}, 10000000,
	)
	defer sim.Close()
	// init statedb
	statedb, err := sim.blockchain.State()
	require.NoError(t, err)
	// set contract
	statedb.SetCode(toAddr, []byte(tssBin))
	// set balance
	statedb.SetBalance(toAddr, big.NewInt(100000000))
	sim.Commit()

	// construct
	parsed, err := abi.JSON(strings.NewReader(tssAbi))
	require.NoError(t, err)
	input, err := parsed.Pack("queryReward")
	require.NoError(t, err)

	bgCtx := context.Background()
	res, err := sim.CallContract(bgCtx, ethereum.CallMsg{
		From: testAddr,
		To:   &toAddr,
		Data: input,
	}, nil)
	t.Log("res: ", res)

	bal := statedb.GetBalance(toAddr)
	t.Log("balance is: ", bal)

	bgCtx = context.Background()
	res, err = sim.CallContract(bgCtx, ethereum.CallMsg{
		From: testAddr,
		To:   &toAddr,
		Data: input,
	}, nil)
	t.Log("res: ", res)

	// test update reward
	input, err = parsed.Pack("updateReward", big.NewInt(1), big.NewInt(10000))
	require.NoError(t, err)
	res, err = sim.CallContract(bgCtx, ethereum.CallMsg{
		From: testAddr,
		To:   &toAddr,
		Data: input,
	}, nil)
	t.Log("res: ", res)

}
