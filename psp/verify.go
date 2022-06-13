package psp

import (
	"encoding/json"
	"io/ioutil"
)

type VerifyRequest struct {
	Type              PaymentInstrumentVerifyType `json:"type"`  // Type is the type of verification we want to perform e.g. identify, challenge
	Value             string                      `json:"value"` // Value to verify
	Amount            Amount                      `json:"amount"`
	BillPayer         Person                      `json:"billPayer"`
	BillingProfileID  string                      `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument           `json:"paymentInstrument"`
}

type VerifyResponse struct {
	ThreeDSResult
	TransactionResponse
}

func (r VerifyRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/verify"
}

func (r VerifyRequest) Do(conn Connection) (resp VoidResponse, err error) {
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
