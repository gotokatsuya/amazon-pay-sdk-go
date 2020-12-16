package amazonpay

type ErrorResponse struct {
	ReasonCode string `json:"reasonCode,omitempty"`
	Message    string `json:"message,omitempty"`
}
