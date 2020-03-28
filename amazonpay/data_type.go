package amazonpay

import (
	"encoding/xml"
	"time"
)

type ResponseError struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Type    string   `xml:"Error>Type"`
	Code    string   `xml:"Error>Code"`
	Message string   `xml:"Error>Message"`
}

func (re ResponseError) Error() string {
	return re.Message
}

// ResponseMetadata type
type ResponseMetadata struct {
	RequestID string `xml:"RequestId"`
}

// BillingAgreementDetails type
type BillingAgreementDetails struct {
	AmazonBillingAgreementID         string `form:"AmazonBillingAgreementId" xml:"AmazonBillingAgreementId"`
	BillingAgreementLimits           BillingAgreementLimits
	Buyer                            Buyer
	SellerNote                       string
	PlatformID                       string `form:"PlatformId" xml:"PlatformId"`
	Destination                      Destination
	ReleaseEnvironment               string
	SellerBillingAgreementAttributes SellerBillingAgreementAttributes
	BillingAgreementStatus           BillingAgreementStatus
	Constraints                      []Constraint
	CreationTimestamp                *time.Time
	BillingAgreementConsent          bool
}

type Price struct {
	Amount       string
	CurrencyCode string
}

// BillingAgreementLimits type
type BillingAgreementLimits struct {
	AmountLimitPerTimePeriod Price
	TimePeriodStartDate      *time.Time
	TimePeriodEndDate        *time.Time
	CurrentRemainingBalance  Price
}

// SellerBillingAgreementAttributes type
type SellerBillingAgreementAttributes struct {
	SellerBillingAgreementID string `form:"SellerBillingAgreementId,omitempty" xml:"SellerBillingAgreementId"`
	StoreName                string `form:",omitempty"`
	CustomInformation        string `form:",omitempty"`
}

// BillingAgreementStatus type
type BillingAgreementStatus struct {
	State                string
	LastUpdatedTimestamp *time.Time
	ReasonCode           string
	ReasonDescription    string
}

// BillingAgreementAttributes type
type BillingAgreementAttributes struct {
	PlatformID                       string `form:"PlatformId,omitempty" xml:"PlatformId"`
	SellerNote                       string `form:",omitempty"`
	SellerBillingAgreementAttributes SellerBillingAgreementAttributes
}

// Buyer type
type Buyer struct {
	Name  string
	Email string
	Phone string
}

// Destination type
type Destination struct {
	DestinationType     string
	PhysicalDestination Address
}

// Address type
type Address struct {
	Name                                     string
	AddressLine1, AddressLine2, AddressLine3 string
	City                                     string
	Country                                  string
	District                                 string
	StateOrRegion                            string
	PostalCode                               string
	CountryCode                              string
	Phone                                    string
}

// Constraint type
type Constraint struct {
	ConstraintID string `form:"ConstraintId" xml:"ConstraintId"`
	Description  string
}

// Status type
type Status struct {
	State               string
	LastUpdateTimestamp *time.Time
	ReasonCode          string
	ReasonDescription   string
}

// AuthorizationDetails type
type AuthorizationDetails struct {
	AmazonAuthorizationID    string `form:"AmazonAuthorizationId" xml:"AmazonAuthorizationId"`
	AuthorizationReferenceID string `form:"AuthorizationReferenceId" xml:"AuthorizationReferenceId"`
	SellerAuthorizationNote  string
	AuthorizationAmount      Price
	CaptureAmount            Price
	AuthorizationFee         Price
	IDList                   struct {
		Member []string `form:"member" xml:"member"`
	} `form:"IdList" xml:"IdList"`
	CreationTimestamp   *time.Time
	ExpirationTimestamp *time.Time
	AuthorizationStatus Status
	SoftDecline         bool
	CaptureNow          bool
}

// RefundDetails type
type RefundDetails struct {
	AmazonRefundID    string `form:"AmazonRefundId" xml:"AmazonRefundId"`
	RefundReferenceID string `form:"RefundReferenceId" xml:"RefundReferenceId"`
	SellerRefundNote  string
	RefundType        string
	RefundAmount      Price
	FeeRefunded       Price
	CreationTimestamp *time.Time
	RefundStatus      Status
	SoftDescriptor    string
}

// CaptureDetails type
type CaptureDetails struct {
	AmazonCaptureID    string `form:"AmazonCaptureId" xml:"AmazonCaptureId"`
	CaptureReferenceID string `form:"CaptureReferenceId" xml:"CaptureReferenceId"`
	SellerCaptureNote  string
	CaptureAmount      Price
	RefundAmount       Price
	CaptureFee         Price
	IDList             struct {
		Member []string `form:"member" xml:"member"`
	} `form:"IdList" xml:"IdList"`
	CreationTimestamp *time.Time
	CaptureStatus     Status
	SoftDescriptor    string
}

// SellerOrderAttributes type
type SellerOrderAttributes struct {
	SellerOrderID     string `form:"SellerOrderId" xml:"SellerOrderId"`
	StoreName         string
	CustomInformation string
}
