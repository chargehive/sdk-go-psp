package psp

type RefundRequest struct {
	CaptureID         string            `json:"captureId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type RefundResponse struct {
	TransactionResponse
	AmountRefunded Amount `json:"amount"`
}

func (a RefundRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payments/refund"
}
