package psp

import (
	"encoding/json"
	"io/ioutil"
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
	httpResp, err := conn.Do(r)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}
