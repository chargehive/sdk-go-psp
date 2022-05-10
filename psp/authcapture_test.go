package psp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestAuthCapture(t *testing.T) {
	h := testAuthCapHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := AuthCaptureRequest{AuthorizeRequest{Amount: NewAmount(123, "USD")}}
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

type testAuthCapHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testAuthCapHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testAuthCapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := AuthCaptureRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := AuthCaptureResponse{
		Authorize: AuthorizeResponse{AmountAuthorized: req.Amount},
		Capture:   CaptureResponse{AmountCaptured: req.Amount},
	}
	j, err := json.Marshal(resp)
	_, _ = w.Write(j)
}
