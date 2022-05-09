package psp

type AuthCaptureRequest struct {
	AuthorizeRequest
}

type AuthCaptureResponse struct {
	TransactionResponse
	Authorize AuthorizeResponse
	Capture   CaptureResponse
}

func (r AuthCaptureRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/payment/auth-capture"
}
