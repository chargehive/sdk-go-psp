package psp

import (
	"encoding/json"
)

type CaptureRequest struct {
	BaseTransactionRequest
}

type CaptureResponse struct {
	TransactionResponse
	Authorize    AuthorizeResponse
	Capture      CaptureAuthResponse
	Asynchronous bool
}

func (r *CaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture"
}

func (r *CaptureRequest) Do(conn Connection) (resp CaptureResponse, err error) {
	body, _, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}

	return
}

func (r *CaptureRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *CaptureRequest) GetCorrelationID() string {
	return r.correlationID
}
