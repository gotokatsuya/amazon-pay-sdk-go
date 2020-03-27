package amazonpay

import (
	"context"
	"net/http"
)

// SetBillingAgreementDetails method
// Billing Agreementの説明とその他の販売事業者の情報などの詳細をBilling Agreementにセットします。
// Billing Agreementの説明とその他の販売事業者の情報などの詳細をBilling Agreementに指定するためにSetBillingAgreementDetails処理を呼び出します。
// 本番環境では、この処理の最大リクエストクォーターは10であり、回復レートは1秒間に1回です。SANDBOX環境では、最大リクエストクォーターは2であり、回復レートは2秒間に1回です。
func (c *Client) SetBillingAgreementDetails(ctx context.Context, req *SetBillingAgreementDetailsRequest) (*SetBillingAgreementDetailsResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("SetBillingAgreementDetails", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(SetBillingAgreementDetailsResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// SetBillingAgreementDetailsRequest type
type SetBillingAgreementDetailsRequest struct {
	AmazonBillingAgreementID   string `form:"AmazonBillingAgreementId"`
	BillingAgreementAttributes BillingAgreementAttributes
}

// SetBillingAgreementDetailsResponse type
type SetBillingAgreementDetailsResponse struct {
	SetBillingAgreementDetailsResult struct {
		BillingAgreementDetails BillingAgreementDetails
	}
	ResponseMetadata ResponseMetadata
}
