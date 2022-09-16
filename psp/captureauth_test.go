package psp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestCaptureAuth(t *testing.T) {
	h := testCaptureAuthHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := CaptureAuthRequest{AuthorizeID: "abc123", Amount: NewAmount(123, "USD")}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.AmountCaptured != req.Amount {
		t.Error("amount doesnt match")
	}

	if resp.GatewayTransactionID != req.AuthorizeID+":captured" {
		t.Error("unexpected gateway transaction id")
	}
}

type testCaptureAuthHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testCaptureAuthHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testCaptureAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := CaptureAuthRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := CaptureAuthResponse{AmountCaptured: req.Amount}
	resp.GatewayTransactionID = req.AuthorizeID + ":captured"
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
}
