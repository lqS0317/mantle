package tss_client

import (
	"fmt"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	"github.com/go-resty/resty/v2"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"math/big"
)

var errTssHTTPError = errors.New("tss http error")

type SignStateRequest struct {
	StartBlock          big.Int `json:"start_block"`
	OffsetStartsAtIndex big.Int `json:"offset_starts_at_index"`
	StateRoots          string  `json:"state_roots"`
}

type TssClient interface {
	GetSignStateBatch(BatchData types.SignStateRequest) ([]byte, error)
}

type Client struct {
	client *resty.Client
}

func NewClient(url string) *Client {
	client := resty.New()
	client.SetHostURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errTssHTTPError)
		}
		return nil
	})
	return &Client{
		client: client,
	}
}

func (c *Client) GetSignStateBatch(BatchData types.SignStateRequest) ([]byte, error) {
	var signature []byte
	response, err := c.client.R().
		SetBody(map[string]interface{}{"start_block": BatchData.StartBlock, "offset_starts_at_index": BatchData.OffsetStartsAtIndex, "state_roots": BatchData.StateRoots}).
		SetResult(signature).
		Post("/api/v1/sign/state")
	if err != nil {
		return nil, fmt.Errorf("cannot get signature: %w", err)
	}
	if response.StatusCode() == 200 {
		return response.Body(), nil
	} else {
		return nil, errors.New("fetch tss manager signature faill")
	}
}