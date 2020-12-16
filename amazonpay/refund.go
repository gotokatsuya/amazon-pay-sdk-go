package amazonpay

import (
	"context"
	"fmt"
	"net/http"
)

type CreateRefundRequest struct {
	ChargeID       string `json:"chargeId,omitempty"`
	RefundAmount   *Price `json:"refundAmount,omitempty"`
	SoftDescriptor string `json:"softDescriptor,omitempty"`
}

type RefundResponse struct {
	ErrorResponse
	RefundID           string         `json:"refundId,omitempty"`
	ChargeID           string         `json:"chargeId,omitempty"`
	RefundAmount       *Price         `json:"refundAmount,omitempty"`
	SoftDescriptor     string         `json:"softDescriptor,omitempty"`
	CreationTimestamp  string         `json:"creationTimestamp,omitempty"`
	StatusDetails      *StatusDetails `json:"statusDetails,omitempty"`
	ReleaseEnvironment string         `json:"releaseEnvironment,omitempty"`
}

type CreateRefundResponse RefundResponse

func (c *Client) CreateRefund(ctx context.Context, req *CreateRefundRequest) (*CreateRefundResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/refunds", APIVersion)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CreateRefundResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

type GetRefundResponse RefundResponse

func (c *Client) GetRefund(ctx context.Context, refundID string) (*GetRefundResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/refunds/%s", APIVersion, refundID)
	httpReq, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetRefundResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}
