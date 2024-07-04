package psp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestCapture(t *testing.T) {
	h := testCaptureHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := CaptureRequest{BaseTransactionRequest: BaseTransactionRequest{Amount: NewAmount(123, "USD")}}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.AmountAuthorized == nil || (resp.AmountAuthorized.Units != req.Amount.Units && resp.AmountAuthorized.Currency != req.Amount.Currency) {
		t.Error("authorize amount doesn't match")
	}

	if resp.AmountCaptured != req.Amount {
		t.Error("capture amount doesn't match")
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := NewCaptureResponse(req.Amount.Currency)
	resp.AmountCaptured = req.Amount
	resp.AmountAuthorized = &req.Amount

	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
}
