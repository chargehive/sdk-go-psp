package psp

import (
	"encoding/json"
)

type AccountUpdateRequest struct {
	BaseRequest
	BillingProfileID  string            `json:"billingProfileId"`
	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
}

type AccountUpdateResponse struct {
	BaseResponse
	TransactionID     string              `json:"transactionId"`
	TransactionStatus TransactionStatus   `json:"transactionStatus"`
	Status            AccountUpdateStatus `json:"state"`
	Result            AccountUpdateResult `json:"result"`
	ResponseCode      string              `json:"responseCode"`
	ResponseMessage   string              `json:"responseMessage"`
	ErrorCode         string              `json:"errorCode"`
	ErrorMessage      string              `json:"errorMessage"`
}

func NewAccountUpdateResponse() AccountUpdateResponse {
	return AccountUpdateResponse{}
}

func (r *AccountUpdateRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/method/update"
}

func (r *AccountUpdateRequest) Do(conn Connection) (resp AccountUpdateResponse, err error) {
	body, headers, err := conn.Do(r)
	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}
