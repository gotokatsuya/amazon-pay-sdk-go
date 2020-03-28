package amazonpay

import (
	"context"
	"net/http"
)

// AuthorizeOnBillingAgreement method
// Billing Agreementに保存された支払方法に対して指定の金額を確保します。
// AuthorizeOnBillingAgreement処理は、Billing Agreementに保存された支払方法に対して指定の金額を確保します。
// 支払方法で請求するためには、 CaptureNowリクエストパラメータにtrueをセットするか、この処理の後でCapture処理を呼び出さなければなりません。オーソリはこの処理で返された一定期間のみ有効です。
// インスタント支払通知（IPN）をセットアップしている場合は、この期限の終わりにオーソリが期限切れと通知を送信します。 インスタント支払通知（IPN）の詳しい情報については、Amazon Payインテグレーションガイドを参照してください。
// オーソリの詳細はGetAuthorizationDetails処理で要求することができます。
func (c *Client) AuthorizeOnBillingAgreement(ctx context.Context, req *AuthorizeOnBillingAgreementRequest) (*AuthorizeOnBillingAgreementResponse, *http.Response, error) {
	httpReq, err := c.NewRequest("AuthorizeOnBillingAgreement", req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(AuthorizeOnBillingAgreementResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// AuthorizeOnBillingAgreementRequest type
type AuthorizeOnBillingAgreementRequest struct {
	AmazonBillingAgreementID string `form:"AmazonBillingAgreementId"`
	AuthorizationReferenceID string `form:"AuthorizationReferenceId"`
	AuthorizationAmount      Price
	SellerAuthorizationNote  string
	TransactionTimeout       uint
	CaptureNow               bool
	SoftDescriptor           string
	SellerNote               string
	PlatformID               string
	SellerOrderAttributes    SellerOrderAttributes
	InheritShippingAddress   bool
}

// AuthorizeOnBillingAgreementResponse type
type AuthorizeOnBillingAgreementResponse struct {
	AuthorizeOnBillingAgreementResult struct {
		AuthorizationDetails   AuthorizationDetails
		AmazonOrderReferenceID string `xml:"AmazonOrderReferenceId"`
	}
	ResponseMetadata ResponseMetadata
}
