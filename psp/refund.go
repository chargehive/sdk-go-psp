package psp

import (
	"encoding/json"
)

type RefundRequest struct {
	BaseRequest

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

func NewRefundResponse(currency string) RefundResponse {
	return RefundResponse{
		AmountRefunded: NewAmount(0, currency),
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
