package psp

type RefundRequest struct {
	CaptureId string `json:"captureId"`
	Amount    Amount `json:"amount"`
}

type RefundResponse struct {
}
