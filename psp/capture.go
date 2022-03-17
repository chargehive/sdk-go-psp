package psp

type CaptureRequest struct {
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type CaptureResponse struct {
	TransactionResponse
	AmountCaptured Amount `json:"amountCaptured"`
}
