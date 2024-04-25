package psp

import (
	"encoding/json"
)

type AuthenticateRequest struct {
	BaseTransactionRequest
	SCAChallengeRequest SCAChallengePreference `json:"scaChallengeRequest"`
}

type AuthenticateResponse struct {
	BaseResponse
	TransactionResponse
	ThreeDSResult *ThreeDSResult `json:"3dsResult"`
	ThreeRI       *ThreeRI       `json:"threeRI"`
}

func (r *AuthenticateRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authenticate"
}

func (r *AuthenticateRequest) Do(conn Connection) (resp AuthenticateResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

func (r *AuthenticateRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *AuthenticateRequest) GetCorrelationID() string {
	return r.correlationID
}
