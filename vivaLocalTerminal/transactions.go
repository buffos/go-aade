package vivaLocalTerminal

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type SaleToAcquirerData struct {
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
}

type ApplicationInfo struct {
	ExternalPlatform *ExternalPlatform `json:"externalPlatform,omitempty"`
}

type ExternalPlatform struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Integrator string `json:"integrator"`
}

func NewSalesToAcquirerDataString(name, version, integrator string) string {
	s := SaleToAcquirerData{
		ApplicationInfo: &ApplicationInfo{
			ExternalPlatform: &ExternalPlatform{
				Name:       name,
				Version:    version,
				Integrator: integrator,
			},
		},
	}

	json, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(json)
}

type AadeProviderID int32

const (
	VivaDemo                    AadeProviderID = 999
	SoftOne                     AadeProviderID = 101
	EnterSoft                   AadeProviderID = 102
	Impact                      AadeProviderID = 103
	EpsilonNet                  AadeProviderID = 104
	NovusConceptus              AadeProviderID = 105
	CloudServicesIKE            AadeProviderID = 106
	PrimerSoftware              AadeProviderID = 107
	ILYDA                       AadeProviderID = 108
	Prosvasis                   AadeProviderID = 109
	MAT                         AadeProviderID = 110
	Simply                      AadeProviderID = 111
	Kappa                       AadeProviderID = 112
	SBZ                         AadeProviderID = 113
	OneSys                      AadeProviderID = 114
	ORIAN                       AadeProviderID = 115
	TESAE                       AadeProviderID = 116
	Karpodinis                  AadeProviderID = 117
	CloudT                      AadeProviderID = 118
	ParoxosLiseonPliroforikisAE AadeProviderID = 119
	FHMAS                       AadeProviderID = 800
)

type AadeProviderSignatureData struct {
	ReceiptUID        string
	Mark              string
	DateTimeSignature int
	Amount            int
	AmountWithoutVat  int
	VatPercentage     int
	SumAmountWithVat  int
	Tid               string
}

func (a *AadeProviderSignatureData) String() string {
	// join with ';' as delimiter
	s := strings.Join([]string{
		a.ReceiptUID,
		a.Mark,
		strconv.Itoa(a.DateTimeSignature),
		strconv.Itoa(a.Amount),
		strconv.Itoa(a.AmountWithoutVat),
		strconv.Itoa(a.VatPercentage),
		strconv.Itoa(a.SumAmountWithVat),
		a.Tid,
	}, ";")
	return s
}

type SalesRequest struct {
	SessionID                 uuid.UUID       `json:"sessionId"`
	Amount                    int64           `json:"amount"`
	MerchantReference         *string         `json:"merchantReference,omitempty"`         // default "some-reference"
	CustomerTrns              *string         `json:"customerTrns,omitempty"`              // default "some-reference"
	Preauth                   *bool           `json:"preauth,omitempty"`                   // default false
	TipAmount                 *int32          `json:"tipAmount,omitempty"`                 // default 0
	ShowTransactionResult     *bool           `json:"showTransactionResult,omitempty"`     // default true
	ShowReceipt               *bool           `json:"showReceipt,omitempty"`               // default true
	CurrencyCode              *string         `json:"currencyCode,omitempty"`              // ISO 4217 numeric currency code. default 978 (EUR)
	SaleToAcquirerData        *string         `json:"saleToAcquirerData,omitempty"`        // JSON converted to base64 string containing additional metadata information.
	MaxInstallments           *int32          `json:"maxInstalments,omitempty"`            // default 1
	AadeProviderID            *AadeProviderID `json:"aadeProviderId,omitempty"`            // default 999
	AadeProviderSignatureData *string         `json:"aadeProviderSignatureData,omitempty"` // default ""
	AadeProviderSignature     *string         `json:"aadeProviderSignature,omitempty"`     // default "". The fields of providerSignatureFields encrypted using a public key and the ECDSA.
	AadePreloaded             *bool           `json:"aadePreloaded,omitempty"`             // default false.
	AadePreloadedDuration     *int32          `json:"aadePreloadedDuration,omitempty"`     // default 24hours. The duration of the preloaded transaction in hours.
}

type SalesResponse struct {
	State       string `json:"state"`       // PROCESSING, BUSY_ERROR, SERVER_ERROR
	SessionID   string `json:"sessionId"`   // The sessionId of the transaction
	SessionType string `json:"sessionType"` // Possible value SALE specifies the sale transaction.
}

type SaleRefundRequest struct {
	SessionID                 string          `json:"sessionId"`
	Amount                    int64           `json:"amount"`
	TransactionID             string          `json:"transactionId"`
	OrderCode                 *int64          `json:"orderCode,omitempty"`
	ShortOrderCode            *int64          `json:"shortOrderCode,omitempty"`
	MerchantReference         *string         `json:"merchantReference,omitempty"`
	CustomerTrns              *string         `json:"customerTrns,omitempty"`
	CurrencyCode              *string         `json:"currencyCode,omitempty"`
	ShowTransactionResult     *bool           `json:"showTransactionResult,omitempty"`
	ShowReceipt               *bool           `json:"showReceipt,omitempty"`
	TxnDateFrom               *string         `json:"txnDateFrom,omitempty"`
	TxnDateTo                 *string         `json:"txnDateTo,omitempty"`
	AadeProviderID            *AadeProviderID `json:"aadeProviderId,omitempty"`
	AadeProviderSignatureData *string         `json:"aadeProviderSignatureData,omitempty"`
	AadeProviderSignature     *string         `json:"aadeProviderSignature,omitempty"`
}

type SaleRefundResponse struct {
	State       string `json:"state"`       // PROCESSING, BUSY_ERROR, SERVER_ERROR
	SessionID   string `json:"sessionId"`   // The sessionId of the transaction
	SessionType string `json:"sessionType"` // Possible value CANCEL specifies the sale transaction.
}

// #region GetSessionResponse
type GetSessionResponse struct {
	State       string             `json:"state"`       // SUCCESS, FAILURE, PROCESSING, SESSION_NOT_FOUND_ERROR, SERVER_ERROR
	SessionType string             `json:"sessionType"` // SALE, CANCEL
	SessionID   string             `json:"sessionId"`   // The sessionId of the transaction
	PayloadData SessionPayloadData `json:"payloadData"`
}

type SessionPayloadData struct {
	SaleResponse SessionSalesResponse `json:"saleResponse"`
}

type NbgProgramDetails struct {
	PacketNo          string `json:"packetNo"`          // The packet number associated with this transaction.
	TransactionNo     string `json:"transactionNo"`     // The unique ID assigned to this transaction.
	PointsCollected   int32  `json:"pointsCollected"`   // The points earned from this transaction as a reward for the cardholder.
	PointsRedeemed    int32  `json:"pointsRedeemed"`    // The points redeemed by the cardholder.
	PointsPrevBalance int32  `json:"pointsPrevBalance"` // The points balance before the completion of this transaction.
	PointsNewBalance  int32  `json:"pointsNewBalance"`  // The points balance after the completion of this transaction.
	ExtraMessage      string `json:"extraMessage"`      // The message provided by the National Bank of Greece (NBG).
}

type LoyaltyInfo struct {
	MerchantId        string            `json:"merchantId"`        // The merchant ID assigned by the other bank.
	TerminalId        string            `json:"terminalId"`        // The terminal id that has been assigned by the other bank
	PaymentAmount     int32             `json:"paymentAmount"`     // The payment amount
	RedemptionAmount  int32             `json:"redemptionAmount"`  // The amount the cardholder chooses to redeem.
	FinalAmount       int32             `json:"finalAmount"`       // The final amount after the redemption
	ProgramId         int32             `json:"programId"`         // The loyalty program id. Possible values are:1 = Eurobank επιστροφή 2 = NBG Go4more 3 = ALPHA Bonus
	LogoUrl           string            `json:"logoUrl"`           // The loyalty program logo provided by the bank that manages the loyalty program. It can be printed on the receipt if the merchant chooses to do so.
	NbgProgramDetails NbgProgramDetails `json:"nbgProgramDetails"` // The details of the NBG loyalty program.
}

type SessionSalesResponse struct {
	TransactionID              string      `json:"transactionId"`              // Transaction identification value. Returns only in online mode.
	Tid                        string      `json:"tid"`                        // The ID of the terminal, as returned in the transaction response (normally the same as terminalId). Note: when a sale request cannot be processed, the tid will be 'null'
	VerificationMethod         string      `json:"verificationMethod"`         // The verification method used
	RetrievalReferenceNumber   int32       `json:"retrievalReferenceNumber"`   // Unique transaction identification
	TransactionDate            string      `json:"transactionDate"`            // Transaction date time in ISO 8601 format
	AuthorizationId            string      `json:"authorizationId"`            // Authorization Id response from the authorizing institution
	PanEntryMode               string      `json:"panEntryMode"`               // Indicates the method used for PAN entry to initiate a transaction
	Aid                        string      `json:"aid"`                        // Application identifier
	TipAmount                  int32       `json:"tipAmount"`                  // Tip amount
	Amount                     int64       `json:"amount"`                     // Total amount authorized, including tip amount
	IsAbortRequested           bool        `json:"isAbortRequested"`           // Default: false
	ApplicationLabel           string      `json:"applicationLabel"`           // Selected application label, VISA / AMEX etc
	PrimaryAccountNumberMasked string      `json:"primaryAccountNumberMasked"` // Masked primary account number
	ReferenceNumber            int64       `json:"referenceNumber"`            // STAN number
	OrderCode                  string      `json:"orderCode"`                  // Order code of completed sale transaction
	ShortOrderCode             string      `json:"shortOrderCode"`             // 10-digit integer
	Installments               int32       `json:"installments"`               // Number of card installments
	EventId                    int32       `json:"eventId"`                    // Used for transferring system state information, such as Error IDs.
	Message                    string      `json:"message"`                    // Description of status
	IsSuccess                  bool        `json:"isSuccess"`                  // Indicates successful authorization result
	TransactionTypeId          int32       `json:"transactionTypeId"`          // Transaction type: 5 - sale / 4 - refund
	MerchantApproved           bool        `json:"merchantApproved"`           // This field is returned only if the terminal is in offline mode. It indicates that the transaction is completed in offline mode.
	LoyaltyInfo                LoyaltyInfo `json:"loyaltyInfo"`                // The loyalty information for the transaction.
	AadeTransactionId          string      `json:"aadeTransactionId"`          // The transaction id for the Aade transaction.
}

// #endregion GetSessionResponse

type AbortTransactionRequest struct {
	SessionID string `json:"sessionId"`
}

type AbortTransactionResponse struct {
	State       string `json:"state"`       // PROCESSING, SESSION_NOT_FOUND_ERROR
	SessionID   string `json:"sessionId"`   // The sessionId of the transaction
	SessionType string `json:"sessionType"` // SALE, CANCEL
}
