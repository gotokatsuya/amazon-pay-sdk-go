package amazonpay

import (
	"context"
	"fmt"
	"net/http"
)

type ChargePermissionResponse struct {
	ErrorResponse
	ChargePermissionID          string              `json:"chargePermissionId,omitempty"`
	ChargePermissionReferenceID string              `json:"chargePermissionReferenceId,omitempty"`
	Buyer                       *Buyer              `json:"buyer,omitempty"`
	ReleaseEnvironment          string              `json:"releaseEnvironment,omitempty"`
	ShippingAddress             *AddressDetails     `json:"shippingAddress,omitempty"`
	BillingAddress              *AddressDetails     `json:"billingAddress,omitempty"`
	PaymentPreferences          []PaymentPreference `json:"paymentPreferences,omitempty"`
	StatusDetails               *StatusDetails      `json:"statusDetails,omitempty"`
	CreationTimestamp           string              `json:"creationTimestamp,omitempty"`
	ExpirationTimestamp         string              `json:"expirationTimestamp,omitempty"`
	MerchantMetadata            *MerchantMetadata   `json:"merchantMetadata,omitempty"`
	PlatformID                  string              `json:"platformId,omitempty"`
	Limits                      *Limits             `json:"limits,omitempty"`
	PresentmentCurrency         string              `json:"presentmentCurrency,omitempty"`
}

type GetChargePermissionResponse ChargePermissionResponse

func (c *Client) GetChargePermission(ctx context.Context, chargePermissionID string) (*GetChargePermissionResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/chargePermissions/%s", APIVersion, chargePermissionID)
	httpReq, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetChargePermissionResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

type CloseChargePermissionRequest struct {
	ClosureReason        string `json:"closureReason,omitempty"`
	CancelPendingCharges *bool  `json:"cancelPendingCharges,omitempty"`
}

type CloseChargePermissionResponse ChargePermissionResponse

func (c *Client) CloseChargePermission(ctx context.Context, chargePermissionID string, req *CloseChargePermissionRequest) (*CloseChargePermissionResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/chargePermissions/%s/close", APIVersion, chargePermissionID)
	httpReq, err := c.NewRequest(http.MethodDelete, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CloseChargePermissionResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}
