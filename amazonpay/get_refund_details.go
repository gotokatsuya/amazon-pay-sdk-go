package amazonpay

import (
	"context"
	"net/http"
)

// GetRefundDetails method
// 返金の詳細なステータスを返します。
func (c *Client) GetRefundDetails(ctx context.Context, req *GetRefundDetailsRequest) (*GetRefundDetailsResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("GetRefundDetails", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetRefundDetailsResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// GetRefundDetailsRequest type
type GetRefundDetailsRequest struct {
	AmazonRefundID string `form:"AmazonRefundId"`
}

// GetRefundDetailsResponse type
type GetRefundDetailsResponse struct {
	GetRefundDetailsResult struct {
		RefundDetails RefundDetails
	}
	ResponseMetadata ResponseMetadata
}
