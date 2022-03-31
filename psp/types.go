package psp

import (
	"time"
)

type Amount struct {
	Units    int64  `json:"units"`
	Currency string `json:"currency"`
}

//goland:noinspection GoUnusedExportedFunction
func NewAmount(units int64, currency string) Amount {
	return Amount{
		Units:    units,
		Currency: currency,
	}
}

type Person struct {
	Title       string      `json:"title"`
	First       string      `json:"first"`
	Last        string      `json:"last"`
	FullName    string      `json:"fullName"`
	Email       Email       `json:"email"`
	PhoneNumber PhoneNumber `json:"phoneNumber"`
	Language    Language    `json:"language"`
}

func (p Person) Name() string {
	if p.FullName == "" {
		return p.First + " " + p.Last
	}
	return p.FullName
}

type Language string
type Email string
type PhoneNumber string
type Country string

type Company struct {
	Name        string      `json:"name"`
	Email       Email       `json:"email"`
	PhoneNumber PhoneNumber `json:"phoneNumber"`
}

type Address struct {
	Fao     Person  `json:"fao"`
	Company Company `json:"company"`
	Line1   string  `json:"line1"`
	Line2   string  `json:"line2"`
	Line3   string  `json:"line3"`
	Town    string  `json:"town"`
	County  string  `json:"county"`
	Postal  string  `json:"postal"`
	Country Country `json:"country"`
}

type TransactionSource struct {
	IPAddress string `json:"ipAddress"`
	UserAgent string `json:"userAgent"`
}

type Meta struct {
	Source          TransactionSource `json:"source"`
	BillingAddress  Address           `json:"billingAddress"`
	ShippingAddress Address           `json:"shippingAddress"`
	CustomData      map[string]string `json:"customData"`
}

type PaymentInstrument struct {
	HVT          string                          `json:"hvt"`
	Verification []PaymentInstrumentVerification `json:"verification"`
}

type PaymentInstrumentVerifyType int32

//goland:noinspection GoUnusedConst
const (
	PaymentInstrumentVerifyTypeCvv               PaymentInstrumentVerifyType = 1
	PaymentInstrumentVerifyTypeDeviceDetails     PaymentInstrumentVerifyType = 2
	PaymentInstrumentVerifyTypeIdentifyResponse  PaymentInstrumentVerifyType = 3
	PaymentInstrumentVerifyTypeChallengeResponse PaymentInstrumentVerifyType = 4
)

type PaymentInstrumentVerification struct {
	Type  PaymentInstrumentVerifyType `json:"type"`
	Value []byte                      `json:"value"`
}

type TransactionIdentifier struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

type TransactionResponse struct {
	TransactionID          string                  `json:"transactionId"`
	TransactionTime        time.Time               `json:"transactionTime"`
	TransactionStatus      TransactionStatus       `json:"transactionStatus"`
	TransactionIdentifiers []TransactionIdentifier `json:"transactions"`
	FailureCategory        FailureCategory         `json:"responseCategory"`
	FailureType            FailureType             `json:"failureType"`
	MerchantMessage        string                  `json:"merchantMessage"`
	Tags                   []string                `json:"tags"`
}

type TransactionStatus string

//goland:noinspection GoUnusedConst
const (
	// TransactionStatusUnknown The status is not known or has not been set
	TransactionStatusUnknown TransactionStatus = ""

	// TransactionStatusReceived The transaction request has been received and is awaiting processing, primarily for asynchronous requests.
	TransactionStatusReceived TransactionStatus = "received"

	// TransactionStatusInProgress The transaction is currently processing with the gateway, acquirer or network. No definitive result is available yet.
	TransactionStatusInProgress TransactionStatus = "in-progress"

	// TransactionStatusSuccess The transaction was successful. This result is final.
	TransactionStatusSuccess TransactionStatus = "success"

	// TransactionStatusFailed The transaction has failed. This result is final.
	TransactionStatusFailed TransactionStatus = "failed"

	// TransactionStatusCancelled The transaction request has been cancelled. This result is final.
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

type Request interface {
	GetPath(credentialID string) string
}

type Liability int

//goland:noinspection GoUnusedConst
const (
	LiabilityInvalid Liability = iota
	LiabilityUnknown
	LiabilityMerchant
	LiabilityIssuer
)

type FailureCategory int

//goland:noinspection GoUnusedConst
const (
	FailureCategoryInvalid FailureCategory = iota
	FailureCategoryNone
	FailureCategoryPayload
	FailureCategoryMethod
	FailureCategoryPerson
	FailureCategoryConfiguration
	FailureCategoryConnectivity
	FailureCategoryFraud
	FailureCategoryVerification
	FailureCategoryProcessing
	FailureCategoryUnknown
)

type FailureType int

//goland:noinspection GoUnusedConst
const (
	FailureTypeInvalid FailureType = iota
	FailureTypeNone
	FailureTypeRetry
	FailureTypeSoft
	FailureTypeHard
)
