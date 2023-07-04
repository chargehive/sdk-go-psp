package psp

import (
	"encoding/json"
)

type RefundRequest struct {
	correlationID string

	AuthorizeID       string            `json:"authorizeId"`
	CaptureID         string            `json:"captureId"`
	MerchantReference string            `json:"merchantReference"`
	Amount            Amount            `json:"amount"`
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type RefundResponse struct {
	BaseResponse
	TransactionResponse
	AmountRefunded Amount `json:"amount"`
}

func NewRefundResponse(amountRefunded Amount) RefundResponse {
	return RefundResponse{
		AmountRefunded: amountRefunded,
	}
}

func (r *RefundRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/refund"
}

func (r *RefundRequest) Do(conn Connection) (resp RefundResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

func (r *RefundRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *RefundRequest) GetCorrelationID() string {
	return r.correlationID
}
