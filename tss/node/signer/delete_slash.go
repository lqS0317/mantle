package signer

import (
	"context"

	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss/slash"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"
)

func (p *Processor) deleteSlashing() {
	queryTicker := time.NewTicker(p.taskInterval)
	for {
		signingInfos := p.nodeStore.ListSlashingInfo()
		for _, si := range signingInfos {
			p.handleSlashing(si)
		}
		select {
		case <-p.stopChan:
			return
		case <-queryTicker.C:
		}
	}
}

func (p *Processor) handleSlashing(si slash.SlashingInfo) {
	currentBlockNumber, err := p.l1Client.BlockNumber(context.Background())
	if err != nil {
		log.Error("failed to query block number", "err", err)
		return
	}
	found, err := p.tssStakingSlashingCaller.GetSlashRecord(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber)}, new(big.Int).SetUint64(si.BatchIndex), si.Address)
	if err != nil {
		log.Error("failed to GetSlashRecord", "err", err)
		return
	}
	if found { // is submitted to ethereum
		found, err = p.tssStakingSlashingCaller.GetSlashRecord(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber - uint64(p.l1ConfirmBlocks))}, new(big.Int).SetUint64(si.BatchIndex), si.Address)
		if err != nil {
			log.Error("failed to GetSlashRecord", "err", err)
			return
		}
		if found { // this slashing is confirmed on ethereum
			p.nodeStore.RemoveSlashingInfo(si.Address, si.BatchIndex)
		}
		return
	}

	currentTssInfo, err := p.tssQueryService.QueryActiveInfo()
	if err != nil {
		log.Error("failed to query active tss info", "err", err)
		return
	}

	if si.ElectionId != currentTssInfo.ElectionId {
		log.Error("the election which this node supposed to be slashed is expired, ignore the slash",
			"node", si.Address.String(), "electionId", si.ElectionId, "batch index", si.BatchIndex)
		p.nodeStore.RemoveSlashingInfo(si.Address, si.BatchIndex)
		return
	}
}