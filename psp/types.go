package psp

import "time"

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
	ID     string `json:"id"`
}

type TransactionResponse struct {
	TransactionID          string                  `json:"transactionId"`
	TransactionTime        time.Time               `json:"transactionTime"`
	TransactionIdentifiers []TransactionIdentifier `json:"transactions"`
}
