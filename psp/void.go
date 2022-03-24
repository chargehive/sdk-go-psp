package psp

type VoidRequest struct {
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type VoidResponse struct {
	TransactionResponse
}

func (a VoidRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/void"
}
