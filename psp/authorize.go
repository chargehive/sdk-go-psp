package psp

import (
	"encoding/json"
)

type AuthorizeRequest struct {
	BaseTransactionRequest
	PerformSCA    bool           `json:"performSca"`
	ThreeDSResult *ThreeDSResult `json:"3dsResult"`
}

type AuthorizeResponse struct {
	BaseResponse
	TransactionResponse
	ThreeDSResult    *ThreeDSResult `json:"3dsResult"`
	AmountAuthorized Amount         `json:"amountAuthorized"`
	AuthCode         string         `json:"authCode"`
	CVVResponse      string         `json:"cvvResponse"`
	AVS              string         `json:"avs"`
	ECI              string         `json:"eci"`
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

func (r *AuthorizeRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *AuthorizeRequest) GetCorrelationID() string {
	return r.correlationID
}
