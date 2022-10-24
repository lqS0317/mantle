package go_tests

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	l1bit "github.com/mantlenetworkio/mantle/go-test/contracts/L1/local/TestBitToken.sol"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestMintBit(t *testing.T) {
	l1Client, err := ethclient.Dial(l1url)
	require.NoError(t, err)
	require.NotNil(t, l1Client)

	//// deploy bit
	//deployAuth := buildAuth(t, l1Client, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", big.NewInt(0))
	//addr, tx, bit, err := l1bit.DeployBIT(deployAuth, l1Client, "Bit", "Bit Token")
	//require.NoError(t, err)
	//require.NotNil(t, bit)
	//t.Log("tx hash is: ", tx.Hash())
	//t.Log("bit address is: ", addr.Hex())

	// query balance
	code, err := l1Client.CodeAt(context.Background(), common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), nil)
	require.NoError(t, err)
	require.True(t, len(code) > 0)
	bit, err := l1bit.NewBitTokenERC20(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), l1Client)
	require.NoError(t, err)

	balance, err := bit.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	require.NoError(t, err)
	t.Log("balance is: ", balance.String())
	// mint
	mintAuth := buildAuth(t, l1Client, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", big.NewInt(1000000000000000000))
	tx, err := bit.Mint(mintAuth)
	require.NoError(t, err)
	t.Log("tx hash is: ", tx.Hash())
	// query balance
	balanceAfterMint, err := bit.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	require.NoError(t, err)
	t.Log("balance is: ", balanceAfterMint.String())
}
