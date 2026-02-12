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
}

func NewHTTPConnection(credentialID string, credentialToken string) HttpConnection {
	c := HttpConnection{
		host:         "https://psp.api.chargehive.com",
		credentialID: credentialID,
		authHeader:   "Basic " + base64.StdEncoding.EncodeToString([]byte(credentialID+":"+credentialToken)),
	}

	return c
}

func (c *HttpConnection) SetHost(host string) {
	c.host = host
}

func (c *HttpConnection) SetClient(client *http.Client) {
	c.httpClient = client
}

func (c *HttpConnection) Do(r Request) ([]byte, http.Header, error) {
	j, err := json.Marshal(r)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.host+r.GetPath(c.credentialID), bytes.NewReader(j))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set(RequestHeaderAuthorization, c.authHeader)
	req.Header.Set(RequestHeaderMerchantUUID, r.GetMerchantUUID())
	req.Header.Set(RequestHeaderCorrelationID, r.GetCorrelationID())
	req.Header.Set(RequestHeaderWorkspaceID, r.GetWorkspaceID())

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	rawResp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer rawResp.Body.Close()

	respBody, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, nil, err
	}
	return respBody, rawResp.Header, nil
}
