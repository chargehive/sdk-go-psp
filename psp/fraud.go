package psp

import (
	"encoding/json"
	"time"
)

type FraudScanRequest struct {
	correlationID string

	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	MerchantReference string            `json:"merchantReference"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type FraudScanResponse struct {
	BaseResponse
	TransactionResponse
	FraudScore      float32
	RiskLevel       RiskLevel
	ScanTime        time.Time
	SuggestedAction SuggestedAction
	Summary         string
	AdditionalData  map[string]string
}

func (r *FraudScanRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/fraud/scan"
}

func (r *FraudScanRequest) Do(conn Connection) (resp FraudScanResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}

func (r *FraudScanRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *FraudScanRequest) GetCorrelationID() string {
	return r.correlationID
}
