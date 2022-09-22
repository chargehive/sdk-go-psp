package psp

import (
	"encoding/json"
)

type RefreshRequest struct {
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	CorrelationID     string            `json:"correlationID"`
}

type RefreshResponse struct {
	TransactionResponse
	PaymentInstrument PaymentInstrument   `json:"paymentInstrument"`
	MethodStatus      MethodRefreshStatus `json:"methodStatus"`
}

func (r *RefreshRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/refresh"
}

func (r *RefreshRequest) Do(conn Connection) (resp RefreshResponse, err error) {
	body, _, err := conn.Do(r)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	return resp, err
}

func (r *RefreshRequest) SetCorrelationID(correlationID string) {
	r.CorrelationID = correlationID
}

func (r *RefreshRequest) GetCorrelationID() string {
	return r.CorrelationID
}
