package psp

import (
	"encoding/json"
)

type CaptureAuthRequest struct {
	correlationID string

	AuthorizeID       string            `json:"authorizeId"`
	MerchantReference string            `json:"merchantReference"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
}

type CaptureAuthResponse struct {
	BaseResponse
	TransactionResponse
	AmountCaptured Amount `json:"amountCaptured"`
}

func NewCaptureAuthResponse(currency string) CaptureAuthResponse {
	return CaptureAuthResponse{
		AmountCaptured: NewAmount(0, currency),
	}
}

func (r *CaptureAuthRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture-auth"
}

func (r *CaptureAuthRequest) Do(conn Connection) (resp CaptureAuthResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

func (r *CaptureAuthRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *CaptureAuthRequest) GetCorrelationID() string {
	return r.correlationID
}
