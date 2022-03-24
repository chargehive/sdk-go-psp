package psp

type FraudScanRequest struct {
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type FraudScanResponse struct {
	TransactionResponse
}

func (a FraudScanRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/fraud/scan"
}
