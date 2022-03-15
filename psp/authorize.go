package psp

type AuthorizeRequest struct {
	MerchantReference string            `json:"merchantReference"`
	Amount            Amount            `json:"amount"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type AuthorizeResponse struct {
	AmountAuthorized Amount `json:"amountAuthorized"`
}
