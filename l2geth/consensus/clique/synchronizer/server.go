package synchronizer

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	binding "github.com/mantlenetworkio/mantle/l2geth/consensus/clique/synchronizer/bindings"
)

const (
	ctxTimeout = 100 * time.Second

	ENV_SEQUENCER_CONTRACT_ADDRESS = "SEQUENCER_CONTRACT_ADDRESS"
	ENV_SEQUENCER_L1_RPC           = "SEQUENCER_L1_RPC"
)

// set infos for sort
type SequencerSequencerInfos []binding.SequencerSequencerInfo

func (seqs SequencerSequencerInfos) Len() int {
	return len(seqs)
}

func (s SequencerSequencerInfos) Less(i, j int) bool {
	return s[i].Amount.Int64() < s[j].Amount.Int64()
}

func (s SequencerSequencerInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Synchronizer struct {
	sequencerContractAddr common.Address
	rpcUrl                string
}

func NewSynchronizer() *Synchronizer {
	seqAddr := os.Getenv(ENV_SEQUENCER_CONTRACT_ADDRESS)
	rpcUrl := os.Getenv(ENV_SEQUENCER_L1_RPC)
	return &Synchronizer{
		sequencerContractAddr: common.HexToAddress(seqAddr),
		rpcUrl:                rpcUrl,
	}
}

func (sync *Synchronizer) GetSchedulerAddr() (common.Address, error) {
	// Create eth client
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, sync.rpcUrl)
	if err != nil {
		return common.Address{}, fmt.Errorf("create rpcClient failed: %s", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	seqContractInst, err := binding.NewSequencer(sync.sequencerContractAddr, ethClient)
	if err != nil {
		return common.Address{}, fmt.Errorf("create rpcClient failed: %s", err)
	}
	return seqContractInst.Scheduler(nil)
}

// GetSequencerSet will return the validator set
func (sync *Synchronizer) GetSequencerSet() (SequencerSequencerInfos, error) {
	// Create eth client
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rpcClient, err := rpc.DialContext(ctx, sync.rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("create rpcClient failed: %s", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	seqContractInst, err := binding.NewSequencer(sync.sequencerContractAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("create sequencer contract instance failed: %s", err)
	}
	num, err := seqContractInst.SequencerLimit(nil)
	if err != nil {
		return nil, err
	}
	// get sequencers from sequencer contract
	var seqInfos SequencerSequencerInfos
	seqInfos, err = seqContractInst.GetSequencers(nil)
	if err != nil {
		return nil, err
	}
	if len(seqInfos) == 0 {
		return nil, fmt.Errorf("empty sequencer set")
	}
	// sort by deposit amount ascending
	sort.Stable(seqInfos)
	// find users deposit and infos from the limit
	if seqInfos.Len() > int(num) {
		seqInfos = seqInfos[seqInfos.Len()-int(num):]
	}
	return seqInfos, nil
}