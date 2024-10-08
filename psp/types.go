package psp

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chargehive/sdk-go-core/payment"
	"github.com/pci-bridge/sdk-go/pcib"
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
		return strings.TrimSpace(p.First + " " + p.Last)
	}
	return p.FullName
}

type Language string
type Email string

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (e Email) Valid() bool {
	return emailRegex.MatchString(string(e))
}

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
	URL       string `json:"url"`
}

type ColorDepth int32

func (c ColorDepth) GetNormalized() int32 {
	allowed := []int32{48, 32, 24, 16, 15, 8, 4, 1}
	for _, a := range allowed {
		if int32(c) >= a {
			return a
		}
	}
	return 1
}

type Device struct {
	BrowserAcceptHeader string     `json:"browserAcceptHeader,omitempty"`
	ColorDepth          ColorDepth `json:"colorDepth,omitempty"`
	JavaEnabled         bool       `json:"javaEnabled,omitempty"`
	JavascriptEnabled   bool       `json:"javascriptEnabled,omitempty"`
	Language            string     `json:"language,omitempty"`
	Screen              Dimension  `json:"screen,omitempty"`
	ScreenAvailable     Dimension  `json:"screenAvailable,omitempty"`
	WindowInner         Dimension  `json:"windowInner,omitempty"`
	WindowOuter         Dimension  `json:"windowOuter,omitempty"`
	TimezoneOffsetMins  int32      `json:"timezoneOffsetMins,omitempty"`
	UserAgent           string     `json:"userAgent,omitempty"`
	IpAddress           string     `json:"ipAddress,omitempty"`
	CookiesEnabled      bool       `json:"cookiesEnabled,omitempty"`
	FlashVersion        string     `json:"flashVersion,omitempty"`
	IsTouch             bool       `json:"isTouch,omitempty"`
	Os                  string     `json:"os,omitempty"`
	OsVersion           string     `json:"osVersion,omitempty"`
	Browser             string     `json:"browser,omitempty"`
	BrowserVersion      string     `json:"browserVersion,omitempty"`
	DeviceManufacturer  string     `json:"deviceManufacturer,omitempty"`
	DeviceName          string     `json:"deviceName,omitempty"`
	DeviceVersion       string     `json:"deviceVersion,omitempty"`
	Fingerprint         string     `json:"fingerprint,omitempty"`
	Type                DeviceType `json:"type,omitempty"`
	Timezone            string     `json:"timezone,omitempty"`
}

type Dimension struct {
	Width  int32 `son:"width,omitempty"`
	Height int32 `json:"height,omitempty"`
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
	Device          Device            `json:"device"`
}

type PaymentInstrument struct {
	InstrumentID       string              `json:"instrumentId"`
	LVT                string              `json:"lvt"`
	HVT                string              `json:"hvt"`
	NetworkToken       *NetworkToken       `json:"networkToken,omitempty"`
	TokenType          TokenType           `json:"tokenType"`
	MethodType         MethodType          `json:"methodType"`
	EphemeralToken     string              `json:"ephemeralToken"`
	AuthenticationData map[string]string   `json:"authenticationData"`
	AccountHolder      string              `json:"accountHolder"`
	CardNetwork        payment.CardNetwork `json:"cardNetwork"`
	ExpiryMonth        int32               `json:"expiryMonth"`
	ExpiryYear         int32               `json:"expiryYear"`
	Bin                string              `json:"bin"`
	Last4              string              `json:"last4"`
	BinData            *pcib.BinData       `json:"binData"`
}

type NetworkToken struct {
	Token       string `json:"token"`
	Eci         string `json:"eci"`
	Cryptogram  string `json:"cryptogram"` // aka TAVV
	ExpiryMonth int32  `json:"expiryMonth"`
	ExpiryYear  int32  `json:"expiryYear"`
	RequestorID string `json:"requestorID"`
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

	CardCategory string `json:"cardCategory"` // e.g. Consumer
	Product      string `json:"product"`      // e.g. MDJ - The card issuer or scheme's product identifier
	ProductType  string `json:"productType"`  // e.g. Debit MasterCard (Enhanced) - The card issuer or scheme's product type

	GatewayInstrumentID string `json:"gatewayInstrumentId"` // e.g. Gateway ID for this method

	// PaymentAccountReference Unique reference that links PANs and their affiliated tokens to a single funding account
	PaymentAccountReference string `json:"paymentAccountReference"` // e.g. 1234567890123456
}

type TransactionIdentifier struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

type TransactionResponse struct {
	BaseResponse
	TransactionID           string                    `json:"transactionId"`
	SubGatewayID            string                    `json:"subGatewayId"`
	GatewayTransactionID    string                    `json:"gatewayTransactionId"`
	SubGatewayTransactionID string                    `json:"subGatewayTransactionId"`
	GatewayStatusCode       string                    `json:"gatewayStatusCode"`
	AcquirerTransactionID   string                    `json:"acquirerTransactionId"`
	NetworkTransactionID    string                    `json:"networkTransactionId"`
	ARN                     string                    `json:"arn"`
	TokenType               TokenType                 `json:"tokenType"`
	InitialTransactionType  PreviousTransactionIdType `json:"initialTransactionType"`
	TransactionTime         time.Time                 `json:"transactionTime"`
	TransactionStatus       TransactionStatus         `json:"transactionStatus"`
	TransactionIdentifiers  []TransactionIdentifier   `json:"transactions"`
	FailureCategory         FailureCategory           `json:"failureCategory"`
	FailureType             FailureType               `json:"failureType"`
	ErrorType               ErrorType                 `json:"errorType"`
	MerchantMessage         string                    `json:"merchantMessage"`
	Tags                    []string                  `json:"tags"`
	RetryDelay              *time.Duration            `json:"retryDelay"`             // Delay before retrying the transaction
	RetryMaxAttempts        *int                      `json:"retryMaxAttempts"`       // Only retry this many times
	RetryMaxAttemptsPeriod  *time.Duration            `json:"retryMaxAttemptsPeriod"` // within this time period
	RetryDate               *time.Time                `json:"retryDate"`              // to be used by scheduler

	AuthCode    string `json:"authCode"`
	CVVResponse string `json:"cvvResponse"`
	AVS         string `json:"avs"`
	ECI         string `json:"eci"`
}

type ThreeDSResult struct {
	ID                     string            `json:"id"`
	AcsTransID             string            `json:"acsTransId"`
	DirectoryServerID      string            `json:"directoryServerId"`
	DeviceFingerprintingID string            `json:"deviceFingerprintingId"`
	Eci                    string            `json:"eci"`
	Cavv                   string            `json:"cavv"`
	Xid                    string            `json:"xid"`
	Enrollment             string            `json:"enrollment"`
	Result                 ThreeDSResultCode `json:"result"`
	SignatureStatus        string            `json:"signatureStatus"`
	Status                 string            `json:"status"`
	MajorVersion           int32             `json:"majorVersion"`
	Version                string            `json:"version"`
	PaReq                  string            `json:"paReq"`
	AcsUrl                 string            `json:"acsUrl"`
	PaRes                  string            `json:"paRes"`
	Timestamp              time.Time         `json:"timestamp"`
	Liability              Liability         `json:"liability"`
	Data                   map[string]string `json:"data"`
	Type                   ThreeDSResultType `json:"type"`
}

type ThreeDSResultCode string

// https://www.emvco.com/specifications/emv-3-d-secure-protocol-and-core-functions-specification-6/
//
//goland:noinspection GoUnusedConst
const (
	ThreeDSResultY ThreeDSResultCode = "Y" // ✅ Authentication Verification Successful.
	ThreeDSResultN ThreeDSResultCode = "N" // ❌ Not Authenticated / Account Not Verified; Transaction denied.
	ThreeDSResultU ThreeDSResultCode = "U" // ❓ Authentication / Account Verification Could Not Be Performed; Technical or other problem, as indicated in ARes or RReq.
	ThreeDSResultA ThreeDSResultCode = "A" // ✅ Attempts Processing Performed; Not Authenticated/ Verified, but a proof of attempted authentication/verification is provided (Attempted, but no response or issuer does not support 3ds)
	ThreeDSResultC ThreeDSResultCode = "C" // ❓ Challenge Required; Additional authentication is required using the CReq/CRes.
	ThreeDSResultD ThreeDSResultCode = "D" // ❓ Challenge Required; Decoupled Authentication confirmed.
	ThreeDSResultR ThreeDSResultCode = "R" // ❌ Authentication / Account Verification Rejected; Issuer is rejecting authentication/verification and request that authorisation not be attempted.
	ThreeDSResultI ThreeDSResultCode = "I" // ❓ Informational Only; 3DS Requestor challenge preference acknowledged.
	ThreeDSResultS ThreeDSResultCode = "S" // ❓ Challenge using SPC
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

	//TransactionStatusInfo This transaction is informational only and does not represent a transaction result.
	TransactionStatusInfo TransactionStatus = "info"

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
	SetWorkspaceID(string)
	GetWorkspaceID() string
	SetMerchantUUID(string)
	GetMerchantUUID() string
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
	FailureLabelContactNetwork       FailureLabel = "contact-network"
	FailureLabelContactIssuer        FailureLabel = "contact-issuer"
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
	FailureCategoryCharge             FailureCategory = "charge"
)

type FailureType string

//goland:noinspection GoUnusedConst
const (
	FailureTypeInvalid  FailureType = "invalid"
	FailureTypeNone     FailureType = "none"
	FailureTypeRetry    FailureType = "retry"
	FailureTypeSoft     FailureType = "soft"
	FailureTypeHard     FailureType = "hard"
	FailureTypeInternal FailureType = "internal"
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

type BaseRequest struct {
	correlationID string
	workspaceID   string
	merchantUUID  string
}

func (r *BaseRequest) SetCorrelationID(correlationID string) {
	r.correlationID = correlationID
}
func (r *BaseRequest) GetCorrelationID() string {
	return r.correlationID
}

func (r *BaseRequest) SetWorkspaceID(workspaceID string) {
	r.workspaceID = workspaceID
}
func (r *BaseRequest) GetWorkspaceID() string {
	return r.workspaceID
}

func (r *BaseRequest) SetMerchantUUID(muid string) {
	r.merchantUUID = muid
}
func (r *BaseRequest) GetMerchantUUID() string {
	return r.merchantUUID
}

type BaseTransactionRequest struct {
	BaseRequest

	Amount                Amount                  `json:"amount"`
	MerchantReference     string                  `json:"merchantReference"`
	ChargeEntityID        string                  `json:"chargeEntityId"`
	BillingProfileID      string                  `json:"billingProfileId"`
	Initiator             RequestInitiator        `json:"initiator"`
	IsMoto                bool                    `json:"isMoto"`
	SubscriptionType      RequestSubscriptionType `json:"subscriptionType"`
	SubscriptionStartTime time.Time               `json:"subscriptionTime"`

	// SubscribeAuthorizationID is the gateway transaction id for the original auth in the sequence
	SubscribeAuthorizationID string `json:"subscribeAuthorizationId"`
	// SubscribeAuthorizationNetworkID is the network transaction id for the original auth in the sequence
	SubscribeAuthorizationNetworkID string `json:"subscribeAuthorizationNetworkId"`
	// SubscribeExternalAuthorizationNetworkID is the network transaction id for the original auth in the sequence, provided by different gateway
	SubscribeExternalAuthorizationNetworkID string `json:"subscribeExternalAuthorizationNetworkId"`

	// LastSuccessfulCaptureID is the gateway transaction id for the last successful capture in the sequence
	LastSuccessfulCaptureID string `json:"lastSuccessfulCaptureID"`
	// LastSuccessfulCaptureNetworkID is the network transaction id for the last successful capture in the sequence
	LastSuccessfulCaptureNetworkID string `json:"lastSuccessfulCaptureNetworkID"`
	// LastSuccessfulCaptureExternalNetworkID is the network transaction id for the last successful capture in the sequence, provided by different gateway
	LastSuccessfulCaptureExternalNetworkID string `json:"lastSuccessfulCaptureExternalNetworkID"`

	PaymentInstrument PaymentInstrument `json:"paymentInstrument"`
	BillPayer         Person            `json:"billPayer"`
	Meta              Meta              `json:"meta"`
	LastDecline       *TransactionResponse
}

func (r *BaseTransactionRequest) GetInitialTransactionId(allowLastCapture bool) string {
	tid := r.SubscribeAuthorizationID
	if tid == "" && allowLastCapture {
		tid = r.LastSuccessfulCaptureID
	}
	return tid
}

func (r *BaseTransactionRequest) GetInitialNetworkTransactionID(allowLastCapture bool) string {
	tid := r.SubscribeAuthorizationNetworkID
	if tid == "" && allowLastCapture && r.PaymentInstrument.CardNetwork == payment.CardNetworkVisa {
		tid = r.LastSuccessfulCaptureNetworkID
	}
	return tid
}

func (r *BaseTransactionRequest) GetInitialExternalNetworkTransactionID(allowLastCapture bool) string {
	tid := r.SubscribeExternalAuthorizationNetworkID
	if tid == "" && allowLastCapture && r.PaymentInstrument.CardNetwork == payment.CardNetworkVisa {
		tid = r.LastSuccessfulCaptureExternalNetworkID
	}
	return tid
}

func (r *BaseTransactionRequest) GetInitialAuthTransaction(external bool) (tid string) {
	if external {
		tid = r.SubscribeAuthorizationNetworkID
		if tid == "" {
			tid = r.SubscribeExternalAuthorizationNetworkID
		}
	} else {
		tid = r.SubscribeAuthorizationID
	}
	return tid
}

func (r *BaseTransactionRequest) GetLastCaptureTransaction(external bool) (tid string) {
	if external {
		tid = r.LastSuccessfulCaptureNetworkID
		if tid == "" {
			tid = r.LastSuccessfulCaptureExternalNetworkID
		}
	} else {
		tid = r.LastSuccessfulCaptureID
	}
	return tid
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
	TokenTypeNetwork   TokenType = "network"
	DefaultTokenType             = TokenTypePCIB
)

type PreviousTransactionIdType string

//goland:noinspection GoUnusedConst
const (
	PreviousTransactionIdTypeFirstAuth   PreviousTransactionIdType = "firstAuth"
	PreviousTransactionIdTypeLastCapture PreviousTransactionIdType = "lastCapture"
)

type PreviousTransactionIdSource string

//goland:noinspection GoUnusedConst
const (
	PreviousTransactionIdSourceConnector PreviousTransactionIdSource = "connector"
	PreviousTransactionIdSourceNetwork   PreviousTransactionIdSource = "network"
)

type DeviceType string

const (
	DeviceTypeInvalid  = "invalid"
	DeviceTypeMobile   = "mobile"
	DeviceTypeTablet   = "tablet"
	DeviceTypeDesktop  = "desktop"
	DeviceTypeWearable = "wearable"
	DeviceTypeConsole  = "console"
	DeviceTypeVehicle  = "vehicle"
)

func (d DeviceType) ToInt() int {
	switch d {
	case DeviceTypeInvalid:
		return 0
	case DeviceTypeMobile:
		return 1
	case DeviceTypeTablet:
		return 2
	case DeviceTypeDesktop:
		return 3
	case DeviceTypeWearable:
		return 4
	case DeviceTypeConsole:
		return 5
	case DeviceTypeVehicle:
		return 6
	}

	return 0
}

type MethodType string

//goland:noinspection GoUnusedConst
const (
	MethodTypeCard      MethodType = "card"
	MethodTypeGooglePay MethodType = "googlePay"
	MethodTypeApplePay  MethodType = "applePay"
)

type SCAChallengePreference string

//goland:noinspection GoUnusedConst
const (
	SCAChallengePreferenceNone    SCAChallengePreference = "none"
	SCAChallengePreferenceRequest SCAChallengePreference = "request"
	SCAChallengePreferenceMandate SCAChallengePreference = "mandate"
)

type ThreeRI struct {
	PriorAuthData      string // information about how previous 3DS transaction was done
	PriorAuthMethod    ThreeDSAuthMethod
	PriorAuthTimestamp time.Time
	PriorACSTransID    string // ACS Transaction ID for a prior authenticated transaction
}

type ThreeDSAuthMethod string

//goland:noinspection GoUnusedConst
const (
	ThreeDSAuthMethodFrictionless ThreeDSAuthMethod = "frictionless"
	ThreeDSAuthMethodChallenged   ThreeDSAuthMethod = "challenged"
	ThreeDSAuthMethodAVS          ThreeDSAuthMethod = "avs"
	ThreeDSAuthMethodOther        ThreeDSAuthMethod = "other"
)
