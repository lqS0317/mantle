// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BitTokenERC20

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

// BitTokenERC20MetaData contains all meta data concerning the BitTokenERC20 contract.
var BitTokenERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001fcb38038062001fcb83398181016040528101906200003791906200048e565b818181600390805190602001906200005192919062000241565b5080600490805190602001906200006a92919062000241565b505050620000ad3362000082620000b560201b60201c565b60ff16600a620000939190620006a0565b6064620000a19190620006f1565b620000be60201b60201c565b5050620008c5565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000131576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200012890620007b3565b60405180910390fd5b62000145600083836200023760201b60201c565b8060026000828254620001599190620007d5565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001b09190620007d5565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405162000217919062000843565b60405180910390a362000233600083836200023c60201b60201c565b5050565b505050565b505050565b8280546200024f906200088f565b90600052602060002090601f016020900481019282620002735760008555620002bf565b82601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b5b80821115620002ed576000816000905550600101620002d3565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200035a826200030f565b810181811067ffffffffffffffff821117156200037c576200037b62000320565b5b80604052505050565b600062000391620002f1565b90506200039f82826200034f565b919050565b600067ffffffffffffffff821115620003c257620003c162000320565b5b620003cd826200030f565b9050602081019050919050565b60005b83811015620003fa578082015181840152602081019050620003dd565b838111156200040a576000848401525b50505050565b6000620004276200042184620003a4565b62000385565b9050828152602081018484840111156200044657620004456200030a565b5b62000453848285620003da565b509392505050565b600082601f83011262000473576200047262000305565b5b81516200048584826020860162000410565b91505092915050565b60008060408385031215620004a857620004a7620002fb565b5b600083015167ffffffffffffffff811115620004c957620004c862000300565b5b620004d7858286016200045b565b925050602083015167ffffffffffffffff811115620004fb57620004fa62000300565b5b62000509858286016200045b565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115620005a15780860481111562000579576200057862000513565b5b6001851615620005895780820291505b8081029050620005998562000542565b945062000559565b94509492505050565b600082620005bc57600190506200068f565b81620005cc57600090506200068f565b8160018114620005e55760028114620005f05762000626565b60019150506200068f565b60ff84111562000605576200060462000513565b5b8360020a9150848211156200061f576200061e62000513565b5b506200068f565b5060208310610133831016604e8410600b8410161715620006605782820a9050838111156200065a576200065962000513565b5b6200068f565b6200066f84848460016200054f565b9250905081840481111562000689576200068862000513565b5b81810290505b9392505050565b6000819050919050565b6000620006ad8262000696565b9150620006ba8362000696565b9250620006e97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484620005aa565b905092915050565b6000620006fe8262000696565b91506200070b8362000696565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000747576200074662000513565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200079b601f8362000752565b9150620007a88262000763565b602082019050919050565b60006020820190508181036000830152620007ce816200078c565b9050919050565b6000620007e28262000696565b9150620007ef8362000696565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000827576200082662000513565b5b828201905092915050565b6200083d8162000696565b82525050565b60006020820190506200085a600083018462000832565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620008a857607f821691505b60208210811415620008bf57620008be62000860565b5b50919050565b6116f680620008d56000396000f3fe6080604052600436106100a75760003560e01c8063395093511161006457806339509351146101b157806370a08231146101ee57806395d89b411461022b578063a457c2d714610256578063a9059cbb14610293578063dd62ed3e146102d0576100a7565b806306fdde03146100ac578063095ea7b3146100d75780631249c58b1461011457806318160ddd1461011e57806323b872dd14610149578063313ce56714610186575b600080fd5b3480156100b857600080fd5b506100c161030d565b6040516100ce9190610ec4565b60405180910390f35b3480156100e357600080fd5b506100fe60048036038101906100f99190610f7f565b61039f565b60405161010b9190610fda565b60405180910390f35b61011c6103bd565b005b34801561012a57600080fd5b5061013361044c565b6040516101409190611004565b60405180910390f35b34801561015557600080fd5b50610170600480360381019061016b919061101f565b610456565b60405161017d9190610fda565b60405180910390f35b34801561019257600080fd5b5061019b61054e565b6040516101a8919061108e565b60405180910390f35b3480156101bd57600080fd5b506101d860048036038101906101d39190610f7f565b610557565b6040516101e59190610fda565b60405180910390f35b3480156101fa57600080fd5b50610215600480360381019061021091906110a9565b610603565b6040516102229190611004565b60405180910390f35b34801561023757600080fd5b5061024061064b565b60405161024d9190610ec4565b60405180910390f35b34801561026257600080fd5b5061027d60048036038101906102789190610f7f565b6106dd565b60405161028a9190610fda565b60405180910390f35b34801561029f57600080fd5b506102ba60048036038101906102b59190610f7f565b6107c8565b6040516102c79190610fda565b60405180910390f35b3480156102dc57600080fd5b506102f760048036038101906102f291906110d6565b6107e6565b6040516103049190611004565b60405180910390f35b60606003805461031c90611145565b80601f016020809104026020016040519081016040528092919081815260200182805461034890611145565b80156103955780601f1061036a57610100808354040283529160200191610395565b820191906000526020600020905b81548152906001019060200180831161037857829003601f168201915b5050505050905090565b60006103b36103ac61086d565b8484610875565b6001905092915050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561042d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610424906111c3565b60405180910390fd5b60006103e83461043d9190611212565b90506104493382610a40565b50565b6000600254905090565b6000610463848484610ba0565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006104ae61086d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561052e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610525906112de565b60405180910390fd5b6105428561053a61086d565b858403610875565b60019150509392505050565b60006012905090565b60006105f961056461086d565b84846001600061057261086d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105f491906112fe565b610875565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606004805461065a90611145565b80601f016020809104026020016040519081016040528092919081815260200182805461068690611145565b80156106d35780601f106106a8576101008083540402835291602001916106d3565b820191906000526020600020905b8154815290600101906020018083116106b657829003601f168201915b5050505050905090565b600080600160006106ec61086d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156107a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a0906113c6565b60405180910390fd5b6107bd6107b461086d565b85858403610875565b600191505092915050565b60006107dc6107d561086d565b8484610ba0565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156108e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dc90611458565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094c906114ea565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610a339190611004565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ab0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa7906111c3565b60405180910390fd5b610abc60008383610e21565b8060026000828254610ace91906112fe565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b2391906112fe565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610b889190611004565b60405180910390a3610b9c60008383610e26565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c079061157c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c779061160e565b60405180910390fd5b610c8b838383610e21565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610d11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d08906116a0565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610da491906112fe565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e089190611004565b60405180910390a3610e1b848484610e26565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e65578082015181840152602081019050610e4a565b83811115610e74576000848401525b50505050565b6000601f19601f8301169050919050565b6000610e9682610e2b565b610ea08185610e36565b9350610eb0818560208601610e47565b610eb981610e7a565b840191505092915050565b60006020820190508181036000830152610ede8184610e8b565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f1682610eeb565b9050919050565b610f2681610f0b565b8114610f3157600080fd5b50565b600081359050610f4381610f1d565b92915050565b6000819050919050565b610f5c81610f49565b8114610f6757600080fd5b50565b600081359050610f7981610f53565b92915050565b60008060408385031215610f9657610f95610ee6565b5b6000610fa485828601610f34565b9250506020610fb585828601610f6a565b9150509250929050565b60008115159050919050565b610fd481610fbf565b82525050565b6000602082019050610fef6000830184610fcb565b92915050565b610ffe81610f49565b82525050565b60006020820190506110196000830184610ff5565b92915050565b60008060006060848603121561103857611037610ee6565b5b600061104686828701610f34565b935050602061105786828701610f34565b925050604061106886828701610f6a565b9150509250925092565b600060ff82169050919050565b61108881611072565b82525050565b60006020820190506110a3600083018461107f565b92915050565b6000602082840312156110bf576110be610ee6565b5b60006110cd84828501610f34565b91505092915050565b600080604083850312156110ed576110ec610ee6565b5b60006110fb85828601610f34565b925050602061110c85828601610f34565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061115d57607f821691505b6020821081141561117157611170611116565b5b50919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006111ad601f83610e36565b91506111b882611177565b602082019050919050565b600060208201905081810360008301526111dc816111a0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061121d82610f49565b915061122883610f49565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611261576112606111e3565b5b828202905092915050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b60006112c8602883610e36565b91506112d38261126c565b604082019050919050565b600060208201905081810360008301526112f7816112bb565b9050919050565b600061130982610f49565b915061131483610f49565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611349576113486111e3565b5b828201905092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006113b0602583610e36565b91506113bb82611354565b604082019050919050565b600060208201905081810360008301526113df816113a3565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611442602483610e36565b915061144d826113e6565b604082019050919050565b6000602082019050818103600083015261147181611435565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006114d4602283610e36565b91506114df82611478565b604082019050919050565b60006020820190508181036000830152611503816114c7565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611566602583610e36565b91506115718261150a565b604082019050919050565b6000602082019050818103600083015261159581611559565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006115f8602383610e36565b91506116038261159c565b604082019050919050565b60006020820190508181036000830152611627816115eb565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b600061168a602683610e36565b91506116958261162e565b604082019050919050565b600060208201905081810360008301526116b98161167d565b905091905056fea26469706673582212201b7949a07355bf596f0435c1dd0bd9e053b51e0bfa20306e37c7323f7435ffe364736f6c63430008090033",
}

// DeployBIT deploys a new Ethereum contract, binding an instance of BIT to it.
func DeployBIT(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BitTokenERC20, error) {
	parsed, err := BitTokenERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BitTokenERC20MetaData.Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BitTokenERC20{BitTokenERC20Caller: BitTokenERC20Caller{contract: contract}, BitTokenERC20Transactor: BitTokenERC20Transactor{contract: contract}, BitTokenERC20Filterer: BitTokenERC20Filterer{contract: contract}}, nil
}


// BitTokenERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BitTokenERC20MetaData.ABI instead.
var BitTokenERC20ABI = BitTokenERC20MetaData.ABI

// BitTokenERC20 is an auto generated Go binding around an Ethereum contract.
type BitTokenERC20 struct {
	BitTokenERC20Caller     // Read-only binding to the contract
	BitTokenERC20Transactor // Write-only binding to the contract
	BitTokenERC20Filterer   // Log filterer for contract events
}

// BitTokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BitTokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BitTokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitTokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitTokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitTokenERC20Session struct {
	Contract     *BitTokenERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitTokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitTokenERC20CallerSession struct {
	Contract *BitTokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BitTokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitTokenERC20TransactorSession struct {
	Contract     *BitTokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BitTokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BitTokenERC20Raw struct {
	Contract *BitTokenERC20 // Generic contract binding to access the raw methods on
}

// BitTokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitTokenERC20CallerRaw struct {
	Contract *BitTokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// BitTokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitTokenERC20TransactorRaw struct {
	Contract *BitTokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBitTokenERC20 creates a new instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20(address common.Address, backend bind.ContractBackend) (*BitTokenERC20, error) {
	contract, err := bindBitTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20{BitTokenERC20Caller: BitTokenERC20Caller{contract: contract}, BitTokenERC20Transactor: BitTokenERC20Transactor{contract: contract}, BitTokenERC20Filterer: BitTokenERC20Filterer{contract: contract}}, nil
}

// NewBitTokenERC20Caller creates a new read-only instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*BitTokenERC20Caller, error) {
	contract, err := bindBitTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Caller{contract: contract}, nil
}

// NewBitTokenERC20Transactor creates a new write-only instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BitTokenERC20Transactor, error) {
	contract, err := bindBitTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Transactor{contract: contract}, nil
}

// NewBitTokenERC20Filterer creates a new log filterer instance of BitTokenERC20, bound to a specific deployed contract.
func NewBitTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BitTokenERC20Filterer, error) {
	contract, err := bindBitTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20Filterer{contract: contract}, nil
}

// bindBitTokenERC20 binds a generic wrapper to an already deployed contract.
func bindBitTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitTokenERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitTokenERC20 *BitTokenERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitTokenERC20.Contract.BitTokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitTokenERC20 *BitTokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.BitTokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitTokenERC20 *BitTokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.BitTokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitTokenERC20 *BitTokenERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitTokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitTokenERC20 *BitTokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitTokenERC20 *BitTokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.Allowance(&_BitTokenERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.Allowance(&_BitTokenERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.BalanceOf(&_BitTokenERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BitTokenERC20.Contract.BalanceOf(&_BitTokenERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20Session) Decimals() (uint8, error) {
	return _BitTokenERC20.Contract.Decimals(&_BitTokenERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Decimals() (uint8, error) {
	return _BitTokenERC20.Contract.Decimals(&_BitTokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Session) Name() (string, error) {
	return _BitTokenERC20.Contract.Name(&_BitTokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Name() (string, error) {
	return _BitTokenERC20.Contract.Name(&_BitTokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20Session) Symbol() (string, error) {
	return _BitTokenERC20.Contract.Symbol(&_BitTokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BitTokenERC20 *BitTokenERC20CallerSession) Symbol() (string, error) {
	return _BitTokenERC20.Contract.Symbol(&_BitTokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BitTokenERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20Session) TotalSupply() (*big.Int, error) {
	return _BitTokenERC20.Contract.TotalSupply(&_BitTokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BitTokenERC20 *BitTokenERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BitTokenERC20.Contract.TotalSupply(&_BitTokenERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Approve(&_BitTokenERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Approve(&_BitTokenERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.DecreaseAllowance(&_BitTokenERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.DecreaseAllowance(&_BitTokenERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.IncreaseAllowance(&_BitTokenERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.IncreaseAllowance(&_BitTokenERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_BitTokenERC20 *BitTokenERC20Transactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_BitTokenERC20 *BitTokenERC20Session) Mint() (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Mint(&_BitTokenERC20.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Mint() (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Mint(&_BitTokenERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Transfer(&_BitTokenERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.Transfer(&_BitTokenERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.TransferFrom(&_BitTokenERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BitTokenERC20 *BitTokenERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BitTokenERC20.Contract.TransferFrom(&_BitTokenERC20.TransactOpts, sender, recipient, amount)
}

// BitTokenERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BitTokenERC20 contract.
type BitTokenERC20ApprovalIterator struct {
	Event *BitTokenERC20Approval // Event containing the contract specifics and raw log

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
func (it *BitTokenERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitTokenERC20Approval)
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
		it.Event = new(BitTokenERC20Approval)
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
func (it *BitTokenERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitTokenERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitTokenERC20Approval represents a Approval event raised by the BitTokenERC20 contract.
type BitTokenERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BitTokenERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BitTokenERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20ApprovalIterator{contract: _BitTokenERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BitTokenERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BitTokenERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitTokenERC20Approval)
				if err := _BitTokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BitTokenERC20 *BitTokenERC20Filterer) ParseApproval(log types.Log) (*BitTokenERC20Approval, error) {
	event := new(BitTokenERC20Approval)
	if err := _BitTokenERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BitTokenERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BitTokenERC20 contract.
type BitTokenERC20TransferIterator struct {
	Event *BitTokenERC20Transfer // Event containing the contract specifics and raw log

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
func (it *BitTokenERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitTokenERC20Transfer)
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
		it.Event = new(BitTokenERC20Transfer)
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
func (it *BitTokenERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitTokenERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitTokenERC20Transfer represents a Transfer event raised by the BitTokenERC20 contract.
type BitTokenERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BitTokenERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitTokenERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BitTokenERC20TransferIterator{contract: _BitTokenERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BitTokenERC20 *BitTokenERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BitTokenERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BitTokenERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitTokenERC20Transfer)
				if err := _BitTokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BitTokenERC20 *BitTokenERC20Filterer) ParseTransfer(log types.Log) (*BitTokenERC20Transfer, error) {
	event := new(BitTokenERC20Transfer)
	if err := _BitTokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
