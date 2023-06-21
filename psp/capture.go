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
	Capture   CaptureAuthResponse
	Authorize AuthorizeResponse
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
