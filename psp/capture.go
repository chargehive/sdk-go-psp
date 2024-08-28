package psp

import (
	"encoding/json"
)

type CaptureRequest struct {
	BaseTransactionRequest
	ThreeDSResult *ThreeDSResult `json:"3dsResult"`
}

type CaptureResponse struct {
	TransactionResponse

	AmountCaptured Amount `json:"amountCaptured"`

	AuthorizeTransaction *TransactionResponse `json:"authorizeTransaction,omitempty"`
	ThreeDSResult        *ThreeDSResult       `json:"3dsResult"`
	AmountAuthorized     *Amount              `json:"amountAuthorized"`
}

func NewCaptureResponse(currency string) CaptureResponse {
	return CaptureResponse{
		AmountCaptured: NewAmount(0, currency),
	}
}

func (r *CaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture"
}

func (r *CaptureRequest) Do(conn Connection) (resp CaptureResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}
	return
}
