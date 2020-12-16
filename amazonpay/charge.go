package amazonpay

import (
	"context"
	"fmt"
	"net/http"
)

type CreateChargeRequest struct {
	ChargePermissionID            string `json:"chargePermissionId,omitempty"`
	ChargeAmount                  *Price `json:"chargeAmount,omitempty"`
	CaptureNow                    *bool  `json:"captureNow,omitempty"`
	SoftDescriptor                string `json:"softDescriptor,omitempty"`
	CanHandlePendingAuthorization *bool  `json:"canHandlePendingAuthorization,omitempty"`
}

type CreateChargeResponse struct {
	ErrorResponse
	ChargeID            string            `json:"chargeId,omitempty"`
	ChargePermissionID  string            `json:"chargePermissionId,omitempty"`
	ChargeAmount        *Price            `json:"chargeAmount,omitempty"`
	CaptureAmount       *Price            `json:"captureAmount,omitempty"`
	RefundedAmount      *Price            `json:"refundedAmount,omitempty"`
	ConvertedAmount     string            `json:"convertedAmount,omitempty"`
	ConversionRate      string            `json:"conversionRate,omitempty"`
	SoftDescriptor      string            `json:"softDescriptor,omitempty"`
	MerchantMetadata    *MerchantMetadata `json:"merchantMetadata,omitempty"`
	ProviderMetadata    *ProviderMetadata `json:"providerMetadata,omitempty"`
	StatusDetails       *StatusDetails    `json:"statusDetails,omitempty"`
	CreationTimestamp   string            `json:"creationTimestamp,omitempty"`
	ExpirationTimestamp string            `json:"expirationTimestamp,omitempty"`
	ReleaseEnvironment  string            `json:"releaseEnvironment,omitempty"`
}

func (c *Client) CreateCharge(ctx context.Context, req *CreateChargeRequest) (*CreateChargeResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/charges", APIVersion)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CreateChargeResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}
