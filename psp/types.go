package psp

import (
	"fmt"
	"github.com/pci-bridge/sdk-go/pcib"
	"time"

	"github.com/chargehive/sdk-go-core/payment"
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
	Middle      string      `json:"middle"`
	Last        string      `json:"last"`
	Suffix      string      `json:"suffix"`
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
	TokenType          TokenType         `json:"tokenType"`
	EphemeralToken     string            `json:"ephemeralToken"`
	AuthenticationData map[string]string `json:"authenticationData"`
	AccountHolder      string            `json:"accountHolder"`
}

type PaymentInstrumentResponse struct {
	MethodType MethodType `json:"methodType"`
	TokenType  TokenType  `json:"tokenType"`

	// NewInstrument indicates if this is a new instrument, not a token replacement
	NewInstrument bool `json:"newInstrument"`

	// LVT low value token, useful for retrieval of non-sensitive data
	LVT string `json:"lvt"`
	// HVT high value token, chargeable token
	HVT string `json:"hvt"`
	// MFP merchant's fingerprint for this instrument
	MFP string `json:"mfp"`
	// PFP billing profile's fingerprint for this instrument
	PFP string `json:"pfp"`
	// TokenExpiry expiry date of the token (not to be confused with the card expiry date)
	TokenExpiry string

	AccountHolder string       `json:"accountHolder"`
	Bin           string       `json:"bin"`
	Last4         string       `json:"last4"`
	CardLength    int32        `json:"cardLength"`
	ExpiryMonth   int32        `json:"expiryMonth"`
	ExpiryYear    int32        `json:"expiryYear"`
	BinData       pcib.BinData `json:"binData"`

	// TSID tokenization session id
	TSID string
	// TSIDK tokenization session id key
	TSIDK string
}

type TransactionIdentifier struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

type TransactionResponse struct {
	TransactionID          string                  `json:"transactionId"`
	GatewayTransactionID   string                  `json:"gatewayTransactionId"`
	GatewayStatusCode      string                  `json:"gatewayStatusCode"`
	AcquirerTransactionID  string                  `json:"acquirerTransactionId"`
	NetworkTransactionID   string                  `json:"networkTransactionId"`
	ARN                    string                  `json:"arn"`
	TransactionTime        time.Time               `json:"transactionTime"`
	TransactionStatus      TransactionStatus       `json:"transactionStatus"`
	TransactionIdentifiers []TransactionIdentifier `json:"transactions"`
	FailureCategory        FailureCategory         `json:"failureCategory"`
	FailureType            FailureType             `json:"failureType"`
	ErrorType              ErrorType               `json:"errorType"`
	MerchantMessage        string                  `json:"merchantMessage"`
	Tags                   []string                `json:"tags"`
}

type ThreeDSResult struct {
	ID              string            `json:"id"`
	Eci             string            `json:"eci"`
	Cavv            string            `json:"cavv"`
	Xid             string            `json:"xid"`
	Enrollment      string            `json:"enrollment"`
	Result          ThreeDSResultCode `json:"result"`
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
	Type            ThreeDSResultType `json:"type"`
}

type ThreeDSResultCode string

//goland:noinspection GoUnusedConst
const (
	ThreeDSResultY ThreeDSResultCode = "Y" // ✅ authentication successful
	ThreeDSResultA ThreeDSResultCode = "A" // ✅ attempted, but no response or issuer does not support 3ds

	ThreeDSResultN ThreeDSResultCode = "N" // ❌ failed, DO NOT PROCEED
	ThreeDSResultR ThreeDSResultCode = "R" // ❌ rejected, DO NOT PROCEED

	ThreeDSResultU ThreeDSResultCode = "U" // ❓ 3ds unavailable, service may be down
)

type ThreeDSResultType string

//goland:noinspection GoUnusedConst
const (
	ThreeDSResultTypeChallenge    ThreeDSResultType = "challenge"
	ThreeDSResultTypeFrictionless ThreeDSResultType = "frictionless"
)

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

	SetCorrelationID(string)
	GetCorrelationID() string
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

type FailureLabel string

//goland:noinspection GoUnusedConst
const (
	FailureLabelInvalidPayload       FailureLabel = "invalid-payload"
	FailureLabelInvalidMethod        FailureLabel = "invalid-method"
	FailureLabelInvalidPerson        FailureLabel = "invalid-person"
	FailureLabelUnsupportedMethod    FailureLabel = "unsupported-method"
	FailureLabelInsufficientFunds    FailureLabel = "insufficient-funds"
	FailureLabelFraud                FailureLabel = "fraud"
	FailureLabelExpiredMethod        FailureLabel = "expired-method"
	FailureLabelDuplicateTransaction FailureLabel = "duplicate-transaction"
	FailureLabelTimeout              FailureLabel = "timeout"
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

type ErrorType string

//goland:noinspection GoUnusedConst
const (
	ErrorTypeInvalid        ErrorType = "invalid"
	ErrorTypeNone           ErrorType = "none"
	ErrorTypeAvailableFunds ErrorType = "available-funds"
	ErrorTypePayload        ErrorType = "payload"
	ErrorTypeLimit          ErrorType = "limit"
	ErrorTypeExpired        ErrorType = "expired"
	ErrorTypeUnavailable    ErrorType = "unavailable"
	ErrorTypeUnsupported    ErrorType = "unsupported"
	ErrorTypeLost           ErrorType = "lost"
	ErrorTypeStolen         ErrorType = "stolen"
	ErrorTypeFraud          ErrorType = "fraud"
	ErrorTypePickup         ErrorType = "pickup"
	ErrorTypeVelocity       ErrorType = "velocity"
	ErrorTypeAddress        ErrorType = "address"
	ErrorTypeDuplicate      ErrorType = "duplicate"
	ErrorTypeTimeout        ErrorType = "timeout"
	ErrorTypeNotFound       ErrorType = "not-found"
	ErrorTypeDisputed       ErrorType = "disputed"
	ErrorTypePermission     ErrorType = "permission"
	ErrorTypeDecline        ErrorType = "decline"
	ErrorTypeUserInput      ErrorType = "user-input"
	ErrorTypeUserDevice     ErrorType = "user-device"
	ErrorTypeAlreadyDone    ErrorType = "already-done"
	ErrorTypeRetry          ErrorType = "retry"
	ErrorTypeQueued         ErrorType = "queued"
	ErrorTypeSystem         ErrorType = "system"
	ErrorTypeUnknown        ErrorType = "unknown"
	ErrorTypeNotReady       ErrorType = "not-ready"
	ErrorTypeDisabled       ErrorType = "disabled"
	ErrorTypeCVV            ErrorType = "cvv"
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
	correlationID string

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
	CardNetwork              payment.CardNetwork     `json:"cardNetwork"`
}

type BaseResponse struct {
	RequestID           string                     `json:"requestId"`
	Status              *StatusResponse            `json:"status"`
	MethodUpgradeTokens []string                   `json:"methodUpgradeTokens"`
	Instrument          *PaymentInstrumentResponse `json:"instrument"`
}

func (r *BaseResponse) SetStatus(code int, message string) {
	r.Status = NewStatus(code, message)
}

type StatusResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewStatus(code int, message string) *StatusResponse {
	return &StatusResponse{
		Message: message,
		Code:    code,
	}
}

func (r *StatusResponse) Error() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("[%d] %s", r.GetCode(), r.GetMessage())
}

func (r *StatusResponse) IsError() bool {
	return r.GetCode() >= 300
}

func (r *StatusResponse) GetMessage() string {
	if r == nil {
		return ""
	}
	return r.Message
}

func (r *StatusResponse) GetCode() int {
	if r == nil {
		return 200
	}
	return r.Code
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

//goland:noinspection GoUnusedConst
const (
	TokenTypePCIB      TokenType = "pcib"
	TokenTypeConnector TokenType = "connector"
	TokenTypeGooglePay TokenType = "googlePay"
	TokenTypeApplePay  TokenType = "applePay"
	DefaultTokenType             = TokenTypePCIB
)

type MethodType string

//goland:noinspection GoUnusedConst
const (
	MethodTypeCard      MethodType = "card"
	MethodTypeGooglePay MethodType = "googlePay"
	MethodTypeApplePay  MethodType = "applePay"
)
