package psp

import "encoding/json"

type AuthorizeRequest struct {
	Amount            Amount            `json:"amount"`
	MerchantReference string            `json:"merchantReference"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
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
	body, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
	}
	return
}
