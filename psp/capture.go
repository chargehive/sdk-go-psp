package psp

import (
	"encoding/json"
	"io/ioutil"
)

type CaptureRequest struct {
	BaseTransactionRequest
}

type CaptureResponse struct {
	TransactionResponse
	Capture   CaptureAuthResponse
	Authorize AuthorizeResponse
}

func (r CaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/capture"
}

func (r CaptureRequest) Do(conn Connection) (resp CaptureResponse, err error) {
	httpResp, err := conn.Do(r)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}
