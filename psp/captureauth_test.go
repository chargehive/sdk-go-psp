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
	req := CaptureAuthRequest{Amount: NewAmount(123, "USD")}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.AmountCaptured != req.Amount {
		t.Error("amount doesnt match")
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
	req := CaptureAuthRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := CaptureAuthResponse{AmountCaptured: req.Amount}
	j, err := json.Marshal(resp)
	_, _ = w.Write(j)
}
