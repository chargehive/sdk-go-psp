package psp

type FraudScanRequest struct {
	Amount            Amount            `json:"amount"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type FraudScanResponse struct {
}
