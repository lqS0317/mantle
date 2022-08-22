package manager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/bitdao-io/bitnetwork/l2geth/log"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	tss "github.com/bitdao-io/bitnetwork/tss/types"
	"github.com/bitdao-io/bitnetwork/tss/ws/server"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

const askTimeOutSeconds = 10

func (m Manager) agreement(ctx types.Context, request tss.SignStateRequest) (types.Context, error) {
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})
	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		log.Error("failed to register response channel", "step", "agreement", "err", err)
		return types.Context{}, err
	}
	requestBz, err := json.Marshal(request)
	if err != nil {
		return types.Context{}, err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	errSendChan := make(chan struct{})
	expectedResponseCount := len(ctx.AvailableNodes())
	maxAllowedLostCount := len(ctx.AvailableNodes()) - ctx.TssInfos().Threshold - 1
	results := make(map[string]bool) // node -> true/false
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), askTimeOutSeconds*time.Second)
		defer func() {
			log.Info("exit accept agreement response goroutine")
			cancel()
			close(stopChan)
			wg.Done()
		}()
		var errSend int
		errResp := make(map[string]struct{})
		log.Info("start to accept agreement response")
		for {
			select {
			case resp := <-respChan:
				log.Info("received ask response", "response", resp.RpcResponse.String(), "ndoe", resp.SourceNode)
				if resp.RpcResponse.Error != nil {
					errResp[resp.SourceNode] = struct{}{}
					if len(errResp)+errSend > maxAllowedLostCount {
						log.Error("maxAllowedLostCount exceed.")
						return
					}
					if len(errResp)+len(results) == expectedResponseCount {
						return
					}
					continue
				}
				var askResponse tss.AskResponse
				if err := tmjson.Unmarshal(resp.RpcResponse.Result, &askResponse); err != nil {
					log.Error("failed to unmarshal ask response", err)
					errResp[resp.SourceNode] = struct{}{}
					if len(errResp)+errSend > maxAllowedLostCount {
						log.Error("maxAllowedLostCount exceed.")
						return
					}
					if len(errResp)+len(results) == expectedResponseCount {
						return
					}
					continue
				}
				results[resp.SourceNode] = askResponse.Result
				if len(errResp)+len(results) == expectedResponseCount {
					return
				}
			case <-errSendChan:
				if errSend == maxAllowedLostCount {
					log.Error("maxAllowedLostCount exceed")
					return
				}
				expectedResponseCount--
				errSend++
			case <-cctx.Done():
				log.Error("wait nodes for ask result response timeout")
				return
			default:
				if len(errResp)+len(results) >= expectedResponseCount {
					return
				}
			}
		}
	}()

	m.askNodes(ctx, requestBz, stopChan, errSendChan)
	wg.Wait()

	if len(results) < ctx.TssInfos().Threshold+1 {
		return types.Context{}, errors.New(fmt.Sprintf("not enough response, %d nodes response for the ask result, we need at least %d nodes to complete the signing process", len(results), ctx.TssInfos().Threshold+1))
	}

	approvers := make([]string, 0)
	for node, pass := range results {
		if pass {
			approvers = append(approvers, node)
		}
	}
	ctx = ctx.WithApprovers(approvers)

	return ctx, nil
}

func (m Manager) askNodes(ctx types.Context, request []byte, stopChan chan struct{}, errSendChan chan struct{}) {
	log.Info("start to sendTonNodes", "number", len(ctx.AvailableNodes()))
	nodes := ctx.AvailableNodes()
	for i := range nodes {
		node := nodes[i]
		go func() {
			select {
			case <-stopChan:
				return
			default:
				requestMsg := server.RequestMsg{
					TargetNode: node,
					RpcRequest: tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(ctx.RequestId()), "ask", request),
				}
				if err := m.wsServer.SendMsg(requestMsg); err != nil {
					log.Error("fail to send ask request", "requestId", ctx.RequestId(), "node", node, "err", err)
					errSendChan <- struct{}{}
				}
			}
		}()
	}
}