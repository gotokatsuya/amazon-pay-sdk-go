package amazonpay

import (
	"context"
	"net/http"
)

// ValidateBillingAgreement method
// Billing Agreementオブジェクトのステータスと関連された支払方法を確認します。
func (c *Client) ValidateBillingAgreement(ctx context.Context, req *ValidateBillingAgreementRequest) (*ValidateBillingAgreementResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("ValidateBillingAgreement", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(ValidateBillingAgreementResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// ValidateBillingAgreementRequest type
type ValidateBillingAgreementRequest struct {
	AmazonBillingAgreementID string `form:"AmazonBillingAgreementId"`
}

// ValidateBillingAgreementResponse type
type ValidateBillingAgreementResponse struct {
	ValidateBillingAgreementResult struct {
		ValidationResult       string
		FailureReasonCode      string
		BillingAgreementStatus BillingAgreementStatus
	}
	ResponseMetadata ResponseMetadata
}
