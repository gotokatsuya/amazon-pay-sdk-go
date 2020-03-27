package amazonpay

import (
	"context"
	"net/http"
)

// Refund method
// 以前に売上請求された資金を返金します。
func (c *Client) Refund(ctx context.Context, req *RefundRequest) (*RefundResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("Refund", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(RefundResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// RefundRequest type
type RefundRequest struct {
	AmazonCaptureID   string `form:"AmazonCaptureId"`
	RefundReferenceID string `form:"RefundReferenceId"`
	RefundAmount      Price
}

// RefundResponse type
type RefundResponse struct {
	RefundResult struct {
		RefundDetails RefundDetails
	}
	ResponseMetadata ResponseMetadata
}
