package psp

import (
	"encoding/json"
)

type CaptureAuthRequest struct {
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type CaptureAuthResponse struct {
	TransactionResponse
	AmountCaptured Amount `json:"amountCaptured"`
}

func (r CaptureAuthRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture-auth"
}

func (r CaptureAuthRequest) Do(conn Connection) (resp CaptureAuthResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}
