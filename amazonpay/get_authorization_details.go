package amazonpay

import (
	"context"
	"net/http"
)

// GetAuthorizationDetails method
// 詳細なオーソリステータスとオーソリで売上請求された合計金額を返します。
func (c *Client) GetAuthorizationDetails(ctx context.Context, req *GetAuthorizationDetailsRequest) (*GetAuthorizationDetailsResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("GetAuthorizationDetails", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetAuthorizationDetailsResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// GetAuthorizationDetailsRequest type
type GetAuthorizationDetailsRequest struct {
	AmazonAuthorizationID string `form:"AmazonAuthorizationId"`
}

// GetAuthorizationDetailsResponse type
type GetAuthorizationDetailsResponse struct {
	GetAuthorizationDetailsResult struct {
		AuthorizationDetails AuthorizationDetails
	}
	ResponseMetadata ResponseMetadata
}
