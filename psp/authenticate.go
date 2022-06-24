package psp

import (
	"encoding/json"
)

type AuthenticateRequestType string

//goland:noinspection GoUnusedConst
const (
	AuthenticateRequestTypeIdentify  AuthenticateRequestType = "identify"
	AuthenticateRequestTypeChallenge AuthenticateRequestType = "challenge"
)

type AuthenticateRequest struct {
	Type              AuthenticateRequestType `json:"type"` // Type is the type of authentication we want to perform e.g. identify, challenge
	Amount            Amount                  `json:"amount"`
	BillPayer         Person                  `json:"billPayer"`
	BillingProfileID  string                  `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument       `json:"paymentInstrument"`
}

type AuthenticateResponse struct {
	ThreeDSResult
	TransactionResponse
}

func (r AuthenticateRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authenticate"
}

func (r AuthenticateRequest) Do(conn Connection) (resp AuthenticateResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}
