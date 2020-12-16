package amazonpay

type Limits struct {
	AmountLimit   *Price `json:"amountLimit,omitempty"`
	AmountBalance *Price `json:"amountBalance,omitempty"`
}
