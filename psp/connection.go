package psp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Connection struct {
	credentialID string
	authHeader   string
	host         string
	httpClient   *http.Client
	newRequest   func(method, url string, body io.Reader) (*http.Request, error)
}

func NewConnection(credentialID string, credentialToken string) Connection {
	c := Connection{
		host:         "https://psp.api.chargehive.com",
		credentialID: credentialID,
		authHeader:   "Basic " + base64.StdEncoding.EncodeToString([]byte(credentialID+":"+credentialToken)),
		newRequest:   http.NewRequest,
	}

	return c
}

func (c Connection) SetHost(host string) {
	c.host = host
}

func (c Connection) SetClient(client *http.Client) {
	c.httpClient = client
}

func (c Connection) Do(r Request) ([]byte, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.host+r.GetPath(c.credentialID), bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.authHeader)

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	rawResp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
