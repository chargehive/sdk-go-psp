package psp

type CaptureRequest struct {
	MerchantReference string `json:"merchantReference"`
	AuthorizeId       string `json:"authorizeId"`
	Amount            Amount `json:"amount"`
}

type CaptureResponse struct {
	AmountCaptured Amount `json:"amountCaptured"`
}
