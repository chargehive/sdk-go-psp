package psp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestCapture(t *testing.T) {
	h := testCaptureHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := CaptureRequest{BaseTransactionRequest{Amount: NewAmount(123, "USD")}}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.Authorize.AmountAuthorized != req.Amount {
		t.Error("authorize amount doesnt match")
	}

	if resp.Capture.AmountCaptured != req.Amount {
		t.Error("capture amount doesnt match")
	}
}

type testCaptureHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testCaptureHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testCaptureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := CaptureRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := CaptureResponse{
		Authorize: AuthorizeResponse{AmountAuthorized: req.Amount},
		Capture:   CaptureAuthResponse{AmountCaptured: req.Amount},
	}
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
}
