package amazonpay

import (
	"context"
	"net/http"
)

// CloseBillingAgreement method
// 購入者のBilling Agreementを終了させるために承認し、このBilling Agreementから新しいOrder Referenceまたはオーソリを生成できないようにします。
func (c *Client) CloseBillingAgreement(ctx context.Context, req *CloseBillingAgreementRequest) (*CloseBillingAgreementResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("CloseBillingAgreement", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CloseBillingAgreementResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// CloseBillingAgreementRequest type
type CloseBillingAgreementRequest struct {
	AmazonBillingAgreementID string `form:"AmazonBillingAgreementId"`
	ClosureReason            string
}

// CloseBillingAgreementResponse type
type CloseBillingAgreementResponse struct {
	ResponseMetadata ResponseMetadata
}
