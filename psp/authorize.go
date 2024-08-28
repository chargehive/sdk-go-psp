package psp

import (
	"encoding/json"
)

type AuthorizeRequest struct {
	BaseTransactionRequest
	PerformSCA          bool                   `json:"performSca"`
	SCAChallengeRequest SCAChallengePreference `json:"scaChallengeRequest"`
	ThreeDSResult       *ThreeDSResult         `json:"3dsResult"`
}

type AuthorizeResponse struct {
	TransactionResponse
	ThreeDSResult    *ThreeDSResult `json:"3dsResult"`
	AmountAuthorized Amount         `json:"amountAuthorized"`
}

func NewAuthorizeResponse(currency string) AuthorizeResponse {
	return AuthorizeResponse{
		AmountAuthorized: NewAmount(0, currency),
	}
}

func (r *AuthorizeRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authorize"
}

func (r *AuthorizeRequest) Do(conn Connection) (resp AuthorizeResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}
