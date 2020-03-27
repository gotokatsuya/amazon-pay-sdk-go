package amazonpay

import (
	"encoding/xml"
	"time"
)

type APIError struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Type    string   `xml:"Error>Type"`
	Code    string   `xml:"Error>Code"`
	Message string   `xml:"Error>Message"`
}

func (apiError APIError) Error() string {
	return apiError.Message
}

// ResponseMetadata type
type ResponseMetadata struct {
	RequestID string `xml:"RequestId"`
}

// BillingAgreementDetails type
type BillingAgreementDetails struct {
	AmazonBillingAgreementID         string `xml:"AmazonBillingAgreementId"`
	BillingAgreementLimits           BillingAgreementLimits
	Buyer                            Buyer
	SellerNote                       string
	PlatformID                       string `xml:"PlatformId"`
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
	SellerBillingAgreementID string `form:",omitempty" xml:"SellerBillingAgreementId"`
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
	PlatformID                       string `form:"PlatformId" xml:"PlatformId"`
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
	ConstraintID string `xml:"ConstraintId"`
	Description  string
}

// AuthorizationDetails type
type AuthorizationDetails struct {
	AmazonAuthorizationID    string `xml:"AmazonAuthorizationId"`
	AuthorizationReferenceID string `xml:"AuthorizationReferenceId"`
	SellerAuthorizationNote  string
	AuthorizationAmount      Price
	CaptureAmount            Price
	AuthorizationFee         Price
	IDList                   []string `xml:"IdList"`
	CreationTimestamp        *time.Time
	ExpirationTimestamp      *time.Time
	AuthorizationStatus      struct {
		State               string
		LastUpdateTimestamp *time.Time
		ReasonCode          string
		ReasonDescription   string
	}
	SoftDecline bool
	CaptureNow  bool
}
