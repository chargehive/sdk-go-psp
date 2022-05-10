package psp

import "encoding/json"

type CaptureRequest struct {
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type CaptureResponse struct {
	TransactionResponse
	AmountCaptured Amount `json:"amountCaptured"`
}

func (r CaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture"
}

func (r CaptureRequest) Do(conn Connection) (resp CaptureResponse, err error) {
	body, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}
	return
}
