package psp

type AuthorizeRequest struct {
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type AuthorizeResponse struct {
	TransactionResponse
	AmountAuthorized Amount `json:"amountAuthorized"`
}
