package psp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestRefund(t *testing.T) {
	h := testRefundHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := RefundRequest{AuthorizeID: "abc123", Amount: NewAmount(123, "USD")}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.AmountRefunded != req.Amount {
		t.Error("amount doesn't match")
	}

	if resp.GatewayTransactionID != req.AuthorizeID+":refunded" {
		t.Error("unexpected gateway transaction id")
	}
}

type testRefundHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testRefundHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testRefundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := RefundRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := NewRefundResponse(req.Amount)
	resp.GatewayTransactionID = req.AuthorizeID + ":refunded"
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
}
