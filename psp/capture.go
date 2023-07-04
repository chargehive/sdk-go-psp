package psp

import (
	"encoding/json"
)

type CaptureRequest struct {
	BaseTransactionRequest
}

type CaptureResponse struct {
	BaseResponse
	TransactionResponse

	// Deprecated: AmountCaptured is now inline
	Capture        CaptureAuthResponse
	AmountCaptured Amount `json:"amountCaptured"`

	// Deprecated: Authorize fields are now inline
	Authorize            AuthorizeResponse
	AuthorizeTransaction *TransactionResponse `json:"authorizeTransaction,omitempty"`
	ThreeDSResult        *ThreeDSResult       `json:"3dsResult"`
	AmountAuthorized     Amount               `json:"amountAuthorized"`
	AuthCode             string               `json:"authCode"`
	CVVResponse          string               `json:"cvvResponse"`
	AVS                  string               `json:"avs"`
	ECI                  string               `json:"eci"`
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

func (r *CaptureRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *CaptureRequest) GetCorrelationID() string {
	return r.correlationID
}
