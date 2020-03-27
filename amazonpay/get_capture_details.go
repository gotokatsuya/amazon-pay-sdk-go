package amazonpay

import (
	"context"
	"net/http"
)

// GetCaptureDetails method
func (c *Client) GetCaptureDetails(ctx context.Context, req *GetCaptureDetailsRequest) (*GetCaptureDetailsResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("GetCaptureDetails", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetCaptureDetailsResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// GetCaptureDetailsRequest type
type GetCaptureDetailsRequest struct {
	AmazonRefundID string `form:"AmazonRefundId"`
}

// GetCaptureDetailsResponse type
type GetCaptureDetailsResponse struct {
	GetCaptureDetailsResult struct {
		CaptureDetails CaptureDetails
	}
	ResponseMetadata ResponseMetadata
}
