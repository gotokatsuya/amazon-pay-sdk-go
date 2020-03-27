package amazonpay

import (
	"context"
	"net/http"
)

// ConfirmBillingAgreement method
// 制約が無く、全ての必須情報がセットされたBilling Agreementを承認します。
func (c *Client) ConfirmBillingAgreement(ctx context.Context, req *ConfirmBillingAgreementRequest) (*ConfirmBillingAgreementResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("ConfirmBillingAgreement", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(ConfirmBillingAgreementResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// ConfirmBillingAgreementRequest type
type ConfirmBillingAgreementRequest struct {
	AmazonBillingAgreementID string `form:"AmazonBillingAgreementId"`
}

// ConfirmBillingAgreementResponse type
type ConfirmBillingAgreementResponse struct {
	ResponseMetadata ResponseMetadata
}
