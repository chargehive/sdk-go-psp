package psp

import (
	"encoding/json"
)

type RefundRequest struct {
	CaptureID         string            `json:"captureId"`
	MerchantReference string            `json:"merchantReference"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type RefundResponse struct {
	TransactionResponse
	AmountRefunded Amount `json:"amount"`
}

func (r RefundRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/refund"
}

func (r RefundRequest) Do(conn Connection) (resp RefundResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}
