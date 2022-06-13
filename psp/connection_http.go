package psp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type HttpConnection struct {
	credentialID string
	authHeader   string
	host         string
	httpClient   *http.Client
	newRequest   func(method, url string, body io.Reader) (*http.Request, error)
}

func NewHTTPConnection(credentialID string, credentialToken string) HttpConnection {
	c := HttpConnection{
		host:         "https://psp.api.chargehive.com",
		credentialID: credentialID,
		authHeader:   "Basic " + base64.StdEncoding.EncodeToString([]byte(credentialID+":"+credentialToken)),
		newRequest:   http.NewRequest,
	}

	return c
}

func (c *HttpConnection) SetHost(host string) {
	c.host = host
}

func (c *HttpConnection) SetClient(client *http.Client) {
	c.httpClient = client
}

func (c *HttpConnection) Do(r Request) (*http.Response, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.host+r.GetPath(c.credentialID), bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set(RequestHeaderAuthorization, c.authHeader)

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return httpClient.Do(req)
}
