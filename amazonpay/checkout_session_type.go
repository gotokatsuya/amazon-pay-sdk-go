package amazonpay

type WebCheckoutDetails struct {
	CheckoutReviewReturnURL string `json:"checkoutReviewReturnUrl,omitempty"`
	CheckoutResultReturnURL string `json:"checkoutResultReturnUrl,omitempty"`
	AmazonPayRedirectURL    string `json:"amazonPayRedirectUrl,omitempty"`
}

type Frequency struct {
	Unit  string `json:"unit,omitempty"`
	Value string `json:"value,omitempty"`
}

type RecurringMetadata struct {
	Frequency *Frequency `json:"frequency,omitempty"`
	Amount    *Price     `json:"amount,omitempty"`
}

type PaymentDetails struct {
	PaymentIntent                 string `json:"paymentIntent,omitempty"`
	CanHandlePendingAuthorization *bool  `json:"canHandlePendingAuthorization,omitempty"`
	ChargeAmount                  *Price `json:"chargeAmount,omitempty"`
	TotalOrderAmount              *Price `json:"totalOrderAmount,omitempty"`
	SoftDescriptor                string `json:"softDescriptor,omitempty"`
	PresentmentCurrency           string `json:"presentmentCurrency,omitempty"`
	AllowOvercharge               *bool  `json:"allowOvercharge,omitempty"`
	ExtendExpiration              *bool  `json:"extendExpiration,omitempty"`
}

type Constraint struct {
	ConstraintID string `json:"constraintId,omitempty"`
	Description  string `json:"description,omitempty"`
}

type DeliverySpecifications struct {
	SpecialRestrictions []string `json:"specialRestrictions,omitempty"`
	AddressRestrictions struct {
		Type         string `json:"type,omitempty"`
		Restrictions struct {
			US struct {
				StatesOrRegions []string `json:"statesOrRegions,omitempty"`
				ZipCodes        []string `json:"zipCodes,omitempty"`
			} `json:"US,omitempty"`
			GB struct {
				ZipCodes []string `json:"zipCodes,omitempty"`
			} `json:"GB,omitempty"`
			IN struct {
				StatesOrRegions []string `json:"statesOrRegions,omitempty"`
			} `json:"IN,omitempty"`
			JP struct {
			} `json:"JP,omitempty"`
		} `json:"restrictions,omitempty"`
	} `json:"addressRestrictions,omitempty"`
}
