package amazonpay

type Price struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type MerchantMetadata struct {
	MerchantReferenceID string `json:"merchantReferenceId,omitempty"`
	MerchantStoreName   string `json:"merchantStoreName,omitempty"`
	NoteToBuyer         string `json:"noteToBuyer,omitempty"`
	CustomInformation   string `json:"customInformation,omitempty"`
}

type ProviderMetadata struct {
	ProviderReferenceID string `json:"providerReferenceId,omitempty"`
}

type Buyer struct {
	BuyerID         string `json:"buyerId,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	ShippingAddress AddressDetails
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}

type AddressDetails struct {
	Name          string `json:"name,omitempty"`
	AddressLine1  string `json:"addressLine1,omitempty"`
	AddressLine2  string `json:"addressLine2,omitempty"`
	AddressLine3  string `json:"addressLine3,omitempty"`
	City          string `json:"city,omitempty"`
	County        string `json:"county,omitempty"`
	District      string `json:"district,omitempty"`
	StateOrRegion string `json:"stateOrRegion,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	PhoneNumber   string `json:"phoneNumber,omitempty"`
}

type PaymentPreference struct {
	PaymentDescriptor string `json:"paymentDescriptor,omitempty"`
}

type StatusDetails struct {
	State string `json:"state,omitempty"`
	// https://amazonpaycheckoutintegrationguide.s3.amazonaws.com/amazon-pay-api-v2/checkout-session.html#type-statusdetails
	ReasonCode        string `json:"reasonCode,omitempty"`
	ReasonDescription string `json:"reasonDescription,omitempty"`
	// https://amazonpaycheckoutintegrationguide.s3.amazonaws.com/amazon-pay-api-v2/charge-permission.html#type-statusdetails
	Reasons              []Reason `json:"reasons,omitempty"`
	LastUpdatedTimestamp string   `json:"lastUpdatedTimestamp,omitempty"`
}

type Reason struct {
	ReasonCode        string `json:"reasonCode,omitempty"`
	ReasonDescription string `json:"reasonDescription,omitempty"`
}
