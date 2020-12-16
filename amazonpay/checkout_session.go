package amazonpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gotokatsuya/amazon-pay-sdk-go/amazonpay/signing"
)

type CreateCheckoutSessionRequest struct {
	WebCheckoutDetails     *WebCheckoutDetails     `json:"webCheckoutDetails,omitempty"`
	StoreID                string                  `json:"storeId,omitempty"`
	ChargePermissionType   string                  `json:"chargePermissionType,omitempty"`
	RecurringMetadata      *RecurringMetadata      `json:"recurringMetadata,omitempty"`
	DeliverySpecifications *DeliverySpecifications `json:"deliverySpecifications,omitempty"`
	PaymentDetails         *PaymentDetails         `json:"paymentDetails,omitempty"`
	MerchantMetadata       *MerchantMetadata       `json:"merchantMetadata,omitempty"`
	PlatformID             string                  `json:"platformId,omitempty"`
	ProviderMetadata       *ProviderMetadata       `json:"providerMetadata,omitempty"`
	AddressDetails         *AddressDetails         `json:"addressDetails,omitempty"`
}

func (c *CreateCheckoutSessionRequest) ToPayload() (string, error) {
	payloadByte, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(payloadByte), nil
}

// GenerateButtonSignature method
func (c *Client) GenerateButtonSignature(payload string) (string, error) {
	stringToSign, err := signing.StringToSign(payload)
	if err != nil {
		return "", err
	}
	signature, err := signing.Sign(c.PrivateKey, stringToSign)
	if err != nil {
		return "", err
	}
	return signature, nil
}

type CheckoutSessionResponse struct {
	ErrorResponse
	CheckoutSessionID      string                  `json:"checkoutSessionId,omitempty"`
	WebCheckoutDetails     *WebCheckoutDetails     `json:"webCheckoutDetails,omitempty"`
	ChargePermissionType   string                  `json:"chargePermissionType,omitempty"`
	RecurringMetadata      *RecurringMetadata      `json:"recurringMetadata,omitempty"`
	ProductType            string                  `json:"productType,omitempty"`
	PaymentDetails         *PaymentDetails         `json:"paymentDetails,omitempty"`
	MerchantMetadata       *MerchantMetadata       `json:"merchantMetadata,omitempty"`
	Buyer                  *Buyer                  `json:"buyer,omitempty"`
	BillingAddress         *AddressDetails         `json:"billingAddress,omitempty"`
	PaymentPreferences     []PaymentPreference     `json:"paymentPreferences,omitempty"`
	StatusDetails          *StatusDetails          `json:"statusDetails,omitempty"`
	ShippingAddress        *AddressDetails         `json:"shippingAddress,omitempty"`
	PlatformID             string                  `json:"platformId,omitempty"`
	ChargePermissionID     string                  `json:"chargePermissionId,omitempty"`
	ChargeID               string                  `json:"chargeId,omitempty"`
	Constraints            []Constraint            `json:"constraints,omitempty"`
	CreationTimestamp      string                  `json:"creationTimestamp,omitempty"`
	ExpirationTimestamp    string                  `json:"expirationTimestamp,omitempty"`
	StoreID                string                  `json:"storeId,omitempty"`
	DeliverySpecifications *DeliverySpecifications `json:"deliverySpecifications,omitempty"`
	ProviderMetadata       *ProviderMetadata       `json:"providerMetadata,omitempty"`
	ReleaseEnvironment     string                  `json:"releaseEnvironment,omitempty"`
}

type GetCheckoutSessionResponse CheckoutSessionResponse

func (c *Client) GetCheckoutSession(ctx context.Context, checkoutSessionID string) (*GetCheckoutSessionResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/checkoutSessions/%s", APIVersion, checkoutSessionID)
	httpReq, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	resp := new(GetCheckoutSessionResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

type UpdateCheckoutSessionRequest struct {
	WebCheckoutDetails *WebCheckoutDetails `json:"webCheckoutDetails,omitempty"`
	PaymentDetails     *PaymentDetails     `json:"paymentDetails,omitempty"`
	MerchantMetadata   *MerchantMetadata   `json:"merchantMetadata,omitempty"`
}

type UpdateCheckoutSessionResponse CheckoutSessionResponse

func (c *Client) UpdateCheckoutSession(ctx context.Context, checkoutSessionID string, req *UpdateCheckoutSessionRequest) (*UpdateCheckoutSessionResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/checkoutSessions/%s", APIVersion, checkoutSessionID)
	httpReq, err := c.NewRequest(http.MethodPatch, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(UpdateCheckoutSessionResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

type CompleteCheckoutSessionRequest struct {
	ChargeAmount *Price `json:"chargeAmount,omitempty"`
}

type CompleteCheckoutSessionResponse CheckoutSessionResponse

func (c *Client) CompleteCheckoutSession(ctx context.Context, checkoutSessionID string, req *CompleteCheckoutSessionRequest) (*CompleteCheckoutSessionResponse, *http.Response, error) {
	path := fmt.Sprintf("%s/checkoutSessions/%s/complete", APIVersion, checkoutSessionID)
	httpReq, err := c.NewRequest(http.MethodPost, path, req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(CompleteCheckoutSessionResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}
