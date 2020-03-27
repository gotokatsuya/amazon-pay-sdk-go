package amazonpay

import (
	"context"
	"net/http"
)

// GetBillingAgreementDetails method
// Billing Agreementオブジェクトの詳細と現在の状態を返します。
// GetBillingAgreementDetails処理は、Billing Agreementオブジェクトの詳細と現在の状態を返します。Billing Agreementオブジェクトは次の情報を提供します。
// 購入者
// 説明
// 説明（オプション）
// 販売事業者のBilling Agreement詳細（オプション）
// 制約のリスト（オプション）
// 本番環境では、この処理の最大リクエストクォーターは20であり、回復レートは1秒間に2回です。SANDBOX環境では、最大リクエストクォーターは5であり、回復レートは1秒間に1回です。
func (c *Client) GetBillingAgreementDetails(ctx context.Context, req *GetBillingAgreementDetailsRequest) (*GetBillingAgreementDetailsResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("GetBillingAgreementDetails", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetBillingAgreementDetailsResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// GetBillingAgreementDetailsRequest type
type GetBillingAgreementDetailsRequest struct {
	AmazonBillingAgreementID string `form:"AmazonBillingAgreementId"`
	AccessToken              string `form:",omitempty"`
}

// GetBillingAgreementDetailsResponse type
type GetBillingAgreementDetailsResponse struct {
	GetBillingAgreementDetailsResult struct {
		BillingAgreementDetails BillingAgreementDetails
	}
	ResponseMetadata ResponseMetadata
}
