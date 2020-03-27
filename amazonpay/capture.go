package amazonpay

import (
	"context"
	"net/http"
)

// Capture method
func (c *Client) Capture(ctx context.Context, req *CaptureRequest) (*CaptureResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("Capture", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CaptureResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// CaptureRequest type
type CaptureRequest struct {
	AmazonAuthorizationID string `form:"AmazonAuthorizationId"`
	CaptureReferenceID    string `form:"CaptureReferenceId"`
	CaptureAmount         Price
}

// CaptureResponse type
type CaptureResponse struct {
	CaptureResult struct {
		CaptureDetails CaptureDetails
	}
	ResponseMetadata ResponseMetadata
}
