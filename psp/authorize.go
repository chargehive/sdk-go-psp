package psp

import (
	"encoding/json"
)

type AuthorizeRequest struct {
	BaseTransactionRequest
}

type AuthorizeResponse struct {
	TransactionResponse
	ThreeDSResult    ThreeDSResult `json:"3dsResult"`
	AmountAuthorized Amount        `json:"amountAuthorized"`
	AuthCode         string        `json:"authCode"`
	CVVResponse      string        `json:"cvvResponse"`
	AVS              string        `json:"avs"`
	ECI              string        `json:"eci"`
}

func (r AuthorizeRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authorize"
}

func (r AuthorizeRequest) Do(conn Connection) (resp AuthorizeResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}
