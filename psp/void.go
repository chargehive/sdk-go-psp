package psp

import (
	"encoding/json"
)

type VoidRequest struct {
	correlationID string

	AuthorizeID       string            `json:"authorizeId"`
	MerchantReference string            `json:"merchantReference"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type VoidResponse struct {
	BaseResponse
	TransactionResponse
}

func NewVoidResponse() VoidResponse {
	return VoidResponse{}
}

func (r *VoidRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/void"
}

func (r *VoidRequest) Do(conn Connection) (resp VoidResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

func (r *VoidRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *VoidRequest) GetCorrelationID() string {
	return r.correlationID
}
