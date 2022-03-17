package psp

type AuthCaptureRequest struct {
	AuthorizeRequest
}

type AuthCaptureResponse struct {
	TransactionResponse
	Authorize AuthorizeResponse
	Capture   CaptureResponse
}
