package psp

import "encoding/json"

type CaptureRequest struct {
	BaseTransactionRequest
}

type CaptureResponse struct {
	TransactionResponse
	Capture   CaptureAuthResponse
	Authorize AuthorizeResponse
}

func (r CaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/auth-capture"
}

func (r CaptureRequest) Do(conn Connection) (resp CaptureResponse, err error) {
	body, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}
	return
}
