package psp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestVoid(t *testing.T) {
	h := testVoidHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := VoidRequest{AuthorizeID: "abc123", Amount: NewAmount(123, "USD")}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.GatewayTransactionID != req.AuthorizeID+":voided" {
		t.Error("unexpected gateway transaction id")
	}
}

type testVoidHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testVoidHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testVoidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := VoidRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := NewVoidResponse()
	resp.GatewayTransactionID = req.AuthorizeID + ":voided"
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
}
