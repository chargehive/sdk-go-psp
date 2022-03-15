package psp

type AuthCaptureRequest struct {
	AuthorizeRequest
}

type AuthCaptureResponse struct {
	AuthorizeResponse
	CaptureResponse
}
