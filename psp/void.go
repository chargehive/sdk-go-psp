package psp

import (
	"encoding/json"
	"io/ioutil"
)

type VoidRequest struct {
	AuthorizeID       string            `json:"authorizeId"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type VoidResponse struct {
	TransactionResponse
}

func (r VoidRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/void"
}

func (r VoidRequest) Do(conn Connection) (resp VoidResponse, err error) {
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
