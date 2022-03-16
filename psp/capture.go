package psp

type CaptureRequest struct {
	MerchantReference string            `json:"merchantReference"`
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type CaptureResponse struct {
	AmountCaptured Amount `json:"amountCaptured"`
}
