package psp

import (
	"encoding/json"
)

type AuthenticateRequest struct {
	BaseTransactionRequest
}

type AuthenticateResponse struct {
	TransactionResponse
	ThreeDSResult *ThreeDSResult
}

func (r *AuthenticateRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authenticate"
}

func (r *AuthenticateRequest) Do(conn Connection) (resp AuthenticateResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}

func (r *AuthenticateRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *AuthenticateRequest) GetCorrelationID() string {
	return r.correlationID
}
