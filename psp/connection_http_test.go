package psp

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnection(t *testing.T) {
	h := echoHandler{t: t, c: testHandlerCredentials{"abc", "123"}}
	con := NewTestConnection(h)
	resp, _, err := con.Do(testReq{Data: "this is a test"})

	if err != nil {
		t.Error(err)
	}

	if string(resp) != `/v1/abc/my/test/path ~ {"Data":"this is a test"}` {
		t.Errorf("unexpected response: %s", resp)
	}
}

type echoHandler struct {
	testHandler
	t *testing.T
	c testHandlerCredentials
}

func (h echoHandler) GetHandlerCredentials() testHandlerCredentials {
	return h.c
}

func (h echoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)

	if r.Header.Get("authorization") != "Basic "+base64.StdEncoding.EncodeToString([]byte(h.GetHandlerCredentials().id+":"+h.GetHandlerCredentials().token)) {
		h.t.Errorf("incorrect auth header: %s", r.Header.Get("authorization"))
	}

	_, _ = w.Write([]byte(r.URL.String() + " ~ " + string(b)))
}

type testReq struct {
	Data string
}

func (t testReq) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/my/test/path"
}

func (t testReq) SetCorrelationID(_ string) {}

func (t testReq) GetCorrelationID() string { return "" }

func NewTestConnection(h testHandler) *HttpConnection {
	srv := httptest.NewServer(h)
	con := NewHTTPConnection(h.GetHandlerCredentials().id, h.GetHandlerCredentials().token)
	con.SetHost(srv.URL)
	con.newRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return httptest.NewRequest(method, url, body), nil
	}
	return &con
}

type testHandler interface {
	http.Handler
	GetHandlerCredentials() testHandlerCredentials
}

type testHandlerCredentials struct {
	id    string
	token string
}
