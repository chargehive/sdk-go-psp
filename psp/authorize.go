package psp

import (
	"encoding/json"
	"io/ioutil"
)

type AuthorizeRequest struct {
	BaseTransactionRequest
}

type AuthorizeResponse struct {
	TransactionResponse
	AmountAuthorized Amount `json:"amountAuthorized"`
	AuthCode         string `json:"authCode"`
	CVVResponse      string `json:"cvvResponse"`
	AVS              string `json:"avs"`
	ECI              string `json:"eci"`
}

func (r AuthorizeRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/authorize"
}

func (r AuthorizeRequest) Do(conn Connection) (resp AuthorizeResponse, err error) {
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
