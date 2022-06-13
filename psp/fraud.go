package psp

import (
	"encoding/json"
)

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

func (r FraudScanRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/fraud/scan"
}

func (r FraudScanRequest) Do(conn Connection) (resp FraudScanResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}
