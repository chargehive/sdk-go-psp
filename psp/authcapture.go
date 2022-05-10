package psp

import "encoding/json"

type AuthCaptureRequest struct {
	AuthorizeRequest
}

type AuthCaptureResponse struct {
	TransactionResponse
	Authorize AuthorizeResponse
	Capture   CaptureResponse
}

func (r AuthCaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/auth-capture"
}

func (r AuthCaptureRequest) Do(conn Connection) (resp AuthCaptureResponse, err error) {
	body, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}
	return
}
