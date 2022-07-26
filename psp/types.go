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

type Item struct {
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	ProductCode    string      `json:"productCode"`
	SkuCode        string      `json:"skuCode"`
	TermUnits      int64       `json:"termUnits"`
	TermType       TermType    `json:"termType"`
	ProductType    ProductType `json:"productType"`
	SkuType        SKUType     `json:"skuType"`
	Quantity       int64       `json:"quantity"`
	UnitPrice      Amount      `json:"unitPrice"`
	TaxAmount      Amount      `json:"taxAmount"`
	DiscountAmount Amount      `json:"discountAmount"`
	StartDate      *time.Time  `json:"startDate"`
	EndDate        *time.Time  `json:"endDate"`
	RenewalNumber  int32       `json:"renewalNumber"`
	Reference      string      `json:"reference"`
}

type Meta struct {
	Source          TransactionSource `json:"source"`
	BillingAddress  Address           `json:"billingAddress"`
	ShippingAddress Address           `json:"shippingAddress"`
	CustomData      map[string]string `json:"customData"`
	Items           []Item            `json:"items"`
}

type PaymentInstrument struct {
	LVT                string            `json:"lvt"`
	HVT                string            `json:"hvt"`
	Type               TokenType         `json:"tokenType"`
	EphemeralToken     string            `json:"ephemeralToken"`
	AuthenticationData map[string]string `json:"authenticationData"`
}

type TransactionIdentifier struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

type TransactionResponse struct {
	TransactionID          string                  `json:"transactionId"`
	GatewayTransactionID   string                  `json:"gatewayTransactionId"`
	NetworkTransactionID   string                  `json:"networkTransactionId"`
	ARN                    string                  `json:"arn"`
	TransactionTime        time.Time               `json:"transactionTime"`
	TransactionStatus      TransactionStatus       `json:"transactionStatus"`
	TransactionIdentifiers []TransactionIdentifier `json:"transactions"`
	FailureCategory        FailureCategory         `json:"failureCategory"`
	FailureType            FailureType             `json:"failureType"`
	MerchantMessage        string                  `json:"merchantMessage"`
	Tags                   []string                `json:"tags"`
}

type ThreeDSResult struct {
	ID              string            `json:"id"`
	Eci             string            `json:"eci"`
	Cavv            string            `json:"cavv"`
	Xid             string            `json:"xid"`
	Enrollment      string            `json:"enrollment"`
	Result          string            `json:"result"`
	SignatureStatus string            `json:"signatureStatus"`
	Status          string            `json:"status"`
	MajorVersion    int32             `json:"majorVersion"`
	Version         string            `json:"version"`
	PaReq           string            `json:"paReq"`
	AcsUrl          string            `json:"acsUrl"`
	PaRes           string            `json:"paRes"`
	Timestamp       time.Time         `json:"timestamp"`
	Liability       Liability         `json:"liability"`
	Data            map[string]string `json:"data"`
}

type Liability string

//goland:noinspection GoUnusedConst
const (
	LiabilityInvalid  Liability = "invalid"
	LiabilityUnknown  Liability = "unknown"
	LiabilityMerchant Liability = "merchant"
	LiabilityIssuer   Liability = "issuer"
)

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

	// TransactionStatusChallenged The transaction request has been challenged and requires authentication information.
	TransactionStatusChallenged TransactionStatus = "challenged"
)

type Request interface {
	GetPath(credentialID string) string
}

type RequestInitiator string

//goland:noinspection GoUnusedConst
const (
	RequestInitiatorMerchant RequestInitiator = "merchant"
	RequestInitiatorCustomer RequestInitiator = "customer"
)

type RequestSubscriptionType string

//goland:noinspection GoUnusedConst
const (
	RequestSubscriptionTypeNone      RequestSubscriptionType = "none"
	RequestSubscriptionTypeSubscribe RequestSubscriptionType = "subscribe"
	RequestSubscriptionTypeRenew     RequestSubscriptionType = "renew"
)

type FailureCategory string

//goland:noinspection GoUnusedConst
const (
	FailureCategoryInvalid            FailureCategory = "invalid"
	FailureCategoryNone               FailureCategory = "none"
	FailureCategoryPayload            FailureCategory = "payload"
	FailureCategoryMethod             FailureCategory = "method"
	FailureCategoryPerson             FailureCategory = "person"
	FailureCategoryConfiguration      FailureCategory = "configuration"
	FailureCategoryConnectivity       FailureCategory = "connectivity"
	FailureCategoryFraud              FailureCategory = "fraud"
	FailureCategoryCardAuthentication FailureCategory = "card-authentication"
	FailureCategoryProcessing         FailureCategory = "processing"
	FailureCategoryUnknown            FailureCategory = "unknown"
)

type FailureType string

//goland:noinspection GoUnusedConst
const (
	FailureTypeInvalid FailureType = "invalid"
	FailureTypeNone    FailureType = "none"
	FailureTypeRetry   FailureType = "retry"
	FailureTypeSoft    FailureType = "soft"
	FailureTypeHard    FailureType = "hard"
)

type ProductType string

//goland:noinspection GoUnusedConst
const (
	ProductTypeInvalid      ProductType = "invalid"
	ProductTypeProduct      ProductType = "product"
	ProductTypeService      ProductType = "service"
	ProductTypeSubscription ProductType = "subscription"
)

type SKUType string

//goland:noinspection GoUnusedConst
const (
	SkuTypeInvalid SKUType = "invalid"
	SkuTypePrimary SKUType = "primary"
	SkuTypeAddon   SKUType = "addon"
	SkuTypeUpsell  SKUType = "upsell"
)

type TermType string

//goland:noinspection GoUnusedConst
const (
	TermTypeInvalid TermType = "invalid"
	TermTypeOneTime TermType = "one-time"
	TermTypeMinute  TermType = "minute"
	TermTypeDay     TermType = "day"
	TermTypeWeek    TermType = "week"
	TermTypeMonth   TermType = "month"
	TermTypeYear    TermType = "year"
	TermTypeNone    TermType = "none"
)

type BaseTransactionRequest struct {
	Amount                   Amount                  `json:"amount"`
	MerchantReference        string                  `json:"merchantReference"`
	BillingProfileID         string                  `json:"billingProfileId"`
	Initiator                RequestInitiator        `json:"initiator"`
	IsMoto                   bool                    `json:"isMoto"`
	SubscriptionType         RequestSubscriptionType `json:"subscriptionType"`
	SubscribeAuthorizationID string                  `json:"subscribeAuthorizationId"`
	PaymentInstrument        PaymentInstrument       `json:"paymentInstrument"`
	BillPayer                Person                  `json:"billPayer"`
	Meta                     Meta                    `json:"meta"`
}

type SuggestedAction string

//goland:noinspection GoUnusedConst
const (
	SuggestedActionInvalid SuggestedAction = "invalid"
	SuggestedActionNone    SuggestedAction = "none"
	SuggestedActionReview  SuggestedAction = "review"
	SuggestedActionAllow   SuggestedAction = "allow"
	SuggestedActionDeny    SuggestedAction = "deny"
)

type RiskLevel string

const (
	RiskLevelInvalid     RiskLevel = "invalid"
	RiskLevelNotRated    RiskLevel = "not-rated"
	RiskLevelNegligible  RiskLevel = "negligible"
	RiskLevelMinor       RiskLevel = "minor"
	RiskLevelModerate    RiskLevel = "moderate"
	RiskLevelSignificant RiskLevel = "significant"
	RiskLevelSevere      RiskLevel = "severe"
)

func (r RiskLevel) ToInt() int {
	switch r {
	case RiskLevelNotRated:
		return 1
	case RiskLevelNegligible:
		return 10
	case RiskLevelMinor:
		return 30
	case RiskLevelModerate:
		return 50
	case RiskLevelSignificant:
		return 70
	case RiskLevelSevere:
		return 90
	case RiskLevelInvalid:
		return 0
	}

	return 0
}

type TokenType string

const (
	TokenTypePCIB      TokenType = "pcib"
	TokenTypeGooglePay TokenType = "googlePay"
	TokenTypeApplePay  TokenType = "applePay"
	DefaultTokenType             = TokenTypePCIB
)
