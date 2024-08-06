package psp

import (
	"encoding/json"

	"github.com/chargehive/sdk-go-core/payment"
)

type NetworkTokenRequest struct {
	BaseTransactionRequest
}

type NetworkTokenResponse struct {
	TransactionResponse
	NetworkTokenStatusResponse
	Token string `json:"token"`
}

func (r *NetworkTokenRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/network-token/tokenize"
}

func (r *NetworkTokenRequest) Do(conn Connection) (resp NetworkTokenResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

type NetworkTokenManageAction string

const (
	NetworkTokenManageActionSuspend NetworkTokenManageAction = "suspend"
	NetworkTokenManageActionResume  NetworkTokenManageAction = "resume"
	NetworkTokenManageActionDelete  NetworkTokenManageAction = "delete"
)

type NetworkTokenManageRequest struct {
	BaseTransactionRequest
	Action NetworkTokenManageAction `json:"action"`
	Reason string                   `json:"reason"`
}

type NetworkTokenManageResponse struct {
	TransactionResponse
}

func (r *NetworkTokenManageRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/network-token/manage"
}

func (r *NetworkTokenManageRequest) Do(conn Connection) (resp NetworkTokenManageResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

type NetworkTokenStatusRequest struct {
	BaseTransactionRequest
}

type NetworkTokenStatusResponse struct {
	TransactionResponse
	CardNetwork      payment.CardNetwork `json:"cardNetwork"`
	TokenReferenceID string              `json:"tokenReferenceId"`
	TokenState       NetworkTokenState   `json:"tokenState"`
	TokenExpiryMonth int32               `json:"tokenExpiryMonth"`
	TokenExpiryYear  int32               `json:"tokenExpiryYear"`
	TokenRequestorID string              `json:"tokenRequestorId"`

	PaymentAccountReference string `json:"paymentAccountReference"`

	PanExpiryMonth int32  `json:"panExpiryMonth"`
	PanExpiryYear  int32  `json:"panExpiryYear"`
	PanLast4       string `json:"panLast4"`

	MetaData map[string]string `json:"metaData"`
}

func (r *NetworkTokenStatusRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/network-token/token-status"
}

func (r *NetworkTokenStatusRequest) Do(conn Connection) (resp NetworkTokenStatusResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

type CryptogramRequest struct {
	BaseTransactionRequest
	Cryptogram string `json:"cryptogram"`
	ECI        string `json:"eci"`
}

type CryptogramResponse struct {
	TransactionResponse
	Cryptogram string `json:"cryptogram"`
	ECI        string `json:"eci"`
}

func (r *CryptogramRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/network-token/cryptogram"
}

func (r *CryptogramRequest) Do(conn Connection) (resp CryptogramResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

type NetworkTokenAssetRequest struct {
	BaseTransactionRequest
}

type NetworkTokenAssetResponse struct {
	TransactionResponse
	CardArtURL string `json:"cardImageUrl"`
}

func (r *NetworkTokenAssetRequest) GetPath(credentialID string) string {
	return "/v1/" + credentialID + "/network-token/asset"
}

func (r *NetworkTokenAssetRequest) Do(conn Connection) (resp NetworkTokenAssetResponse, err error) {
	body, headers, err := conn.Do(r)

	if err == nil {
		err = json.Unmarshal(body, &resp)
		resp.RequestID = headers.Get(RequestHeaderRequestID)
	}

	return
}

type NetworkTokenState string

const (
	TokenStateActive    NetworkTokenState = "ACTIVE"
	TokenStateInactive  NetworkTokenState = "INACTIVE"
	TokenStateSuspended NetworkTokenState = "SUSPENDED"
	TokenStateDeleted   NetworkTokenState = "DELETED"
	TokenStateConsumed  NetworkTokenState = "CONSUMED"
)
