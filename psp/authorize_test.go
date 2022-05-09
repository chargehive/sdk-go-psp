package psp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestAuthorize(t *testing.T) {
	h := testAuthHandler{c: testHandlerCredentials{"abc", "123"}}

	con := NewTestConnection(h)
	req := AuthorizeRequest{Amount: NewAmount(123, "USD")}
	resp, err := req.Do(con)

	if err != nil {
		t.Error(err)
	}

	if resp.AmountAuthorized != req.Amount {
		t.Error("amount doesnt match")
	}
}

type testAuthHandler struct {
	testHandler
	c testHandlerCredentials
}

func (h testAuthHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h testAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := AuthorizeRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}

	resp := AuthorizeResponse{AmountAuthorized: req.Amount}
	j, err := json.Marshal(resp)
	_, _ = w.Write(j)
}
