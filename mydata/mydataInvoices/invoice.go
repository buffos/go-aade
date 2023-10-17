package mydataInvoices

import (
	"errors"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"time"
)

type Invoice struct {
	UID                   *string                               `xml:"uid"`                   // Αναγνωριστικό Παραστατικού (από υπηρεσία)
	Mark                  *uint64                               `xml:"mark"`                  // Μοναδικός Αριθμός Καταχώρησης Παραστατικού (από υπηρεσία)
	CancelledByMark       *uint64                               `xml:"cancelledByMark"`       // Μοναδικός Αριθμός Καταχώρησης Ακυρωτικού (από υπηρεσία)
	AuthenticationCode    *string                               `xml:"authenticationCode"`    // Συμβολοσειρά Αυθεντικοποίησης (από υπηρεσία)
	TransmissionFailure   *mydatavalues.TransmissionFailureType `xml:"transmissionFailure"`   // Αδυναμία επικοινωνίας παρόχου ή διαβίβασης δεδομένων
	Issuer                *PartyType                            `xml:"issuer"`                // Στοιχεία Εκδότη Παραστατικού
	Counterpart           *PartyType                            `xml:"counterpart"`           // Στοιχεία Λήπτη Παραστατικού
	InvoiceHeader         *InvoiceHeaderType                    `xml:"invoiceHeader"`         // * Επικεφαλίδα Παραστατικού
	PaymentMethods        *PaymentMethods                       `xml:"paymentMethods"`        // Στοιχεία Πληρωμής
	InvoiceDetails        []*InvoiceRowType                     `xml:"invoiceDetails"`        // * Γραμμές Παραστατικού
	TaxesTotals           *TaxTotalsType                        `xml:"taxesTotals"`           // Σύνολα Φόρων
	InvoiceSummary        *InvoiceSummaryType                   `xml:"invoiceSummary"`        // * Περίληψη Παραστατικού
	QrCodeUrl             *string                               `xml:"qrCodeUrl"`             // Κωδικοποιημένο αλφαριθμητικό για να χρησιμοποιηθεί από τα προγράμματα για τη δημιουργία QR Code τύπου Url (από υπηρεσία)
	OtherTransportDetails *TransportDetailType                  `xml:"otherTransportDetails"` // Λοιπές λεπτομέρειες διακίνησης
}

//goland:noinspection GoUnusedExportedFunction
func NewInvoice(series string, aa string, issueDate time.Time, invType mydatavalues.InvoiceType) *Invoice {
	return &Invoice{
		InvoiceHeader:  NewInvoiceHeader(series, aa, issueDate, invType),
		InvoiceDetails: make([]*InvoiceRowType, 0),
		InvoiceSummary: &InvoiceSummaryType{},
	}
}

// region Invoice

// Validate validates the invoice
func (i *Invoice) Validate() error {
	if i.InvoiceHeader == nil {
		return errors.New("invoice header is required")
	}
	if i.InvoiceDetails == nil || len(i.InvoiceDetails) == 0 {
		return errors.New("invoice details are required")
	}
	if i.InvoiceSummary == nil {
		return errors.New("invoice summary is required")
	}
	// validate header
	if err := i.InvoiceHeader.validate(); err != nil {
		return err
	}
	// validate each of the rows
	for _, detail := range i.InvoiceDetails {
		if err := detail.Validate(); err != nil {
			return err
		}
	}
	// validate income classifications
	for _, row := range i.InvoiceSummary.IncomeClassification {
		if !AllowedIncomeCharacterisations(i.InvoiceHeader.InvoiceType, row.ClassificationType, row.ClassificationCategory) {
			return errors.New("invalid income classification")
		}
	}

	// validate summary
	return nil
}

func (i *Invoice) ValidateRow() error {
	if i.InvoiceHeader.InvoiceType == nil {
		return errors.New("invoice type is nil or empty")
	}
	return nil
}

// AddInvoiceRow adds an invoice row to the invoice's rows array
func (i *Invoice) AddInvoiceRow(row *InvoiceRowType) *Invoice {
	currentNumberOfRows := uint(len(i.InvoiceDetails))
	currentNumberOfRows++                 // rows start from 1
	row.LineNumber = &currentNumberOfRows // set the line number for the invoice row
	i.InvoiceDetails = append(i.InvoiceDetails, row)
	return i
}

// CalculateSummary calculates the summary values from the rows and TaxTotals
func (i *Invoice) CalculateSummary() {
	// first calculate the summary from the rows
	i.InvoiceSummary.calculate(i.InvoiceDetails)

	// region adds values from the taxesTotals if they exist
	if i.TaxesTotals != nil && len(i.TaxesTotals.Taxes) > 0 {
		for _, tax := range i.TaxesTotals.Taxes {
			switch *tax.TaxType {
			case mydatavalues.TaxTypeWithHoldingTax:
				i.InvoiceSummary.TotalWithheldAmount += *tax.TaxAmount
				i.InvoiceSummary.TotalGrossValue -= *tax.TaxAmount
			case mydatavalues.TaxTypeFees:
				i.InvoiceSummary.TotalFeesAmount += *tax.TaxAmount
				i.InvoiceSummary.TotalGrossValue += *tax.TaxAmount
			case mydatavalues.TaxTypeMiscellaneous:
				i.InvoiceSummary.TotalOtherTaxesAmount += *tax.TaxAmount
				i.InvoiceSummary.TotalGrossValue += *tax.TaxAmount
			case mydatavalues.TaxTypeStamp:
				i.InvoiceSummary.TotalStampDutyAmount += *tax.TaxAmount
				i.InvoiceSummary.TotalGrossValue += *tax.TaxAmount
			case mydatavalues.TaxTypeDeductions:
				i.InvoiceSummary.TotalDeductionsAmount += *tax.TaxAmount
				i.InvoiceSummary.TotalGrossValue -= *tax.TaxAmount
			}
		}
	}
	// endregion

	// region round values to two digits
	i.InvoiceSummary.TotalNetValue = roundToMoney(i.InvoiceSummary.TotalNetValue)
	i.InvoiceSummary.TotalVatAmount = roundToMoney(i.InvoiceSummary.TotalVatAmount)
	i.InvoiceSummary.TotalWithheldAmount = roundToMoney(i.InvoiceSummary.TotalWithheldAmount)
	i.InvoiceSummary.TotalFeesAmount = roundToMoney(i.InvoiceSummary.TotalFeesAmount)
	i.InvoiceSummary.TotalStampDutyAmount = roundToMoney(i.InvoiceSummary.TotalStampDutyAmount)
	i.InvoiceSummary.TotalOtherTaxesAmount = roundToMoney(i.InvoiceSummary.TotalOtherTaxesAmount)
	i.InvoiceSummary.TotalDeductionsAmount = roundToMoney(i.InvoiceSummary.TotalDeductionsAmount)
	// endregion

	// region recalculate gross value just in case rounding values changed it
	i.InvoiceSummary.TotalGrossValue = i.InvoiceSummary.TotalNetValue + i.InvoiceSummary.TotalVatAmount - i.InvoiceSummary.TotalWithheldAmount +
		i.InvoiceSummary.TotalFeesAmount + i.InvoiceSummary.TotalStampDutyAmount + i.InvoiceSummary.TotalOtherTaxesAmount - i.InvoiceSummary.TotalDeductionsAmount

	// round gross value to two digits
	i.InvoiceSummary.TotalGrossValue = roundToMoney(i.InvoiceSummary.TotalGrossValue)
	// endregion
}

// SetCounterPart sets the invoice's counterpart
func (i *Invoice) SetCounterPart(vat string, country string, branch uint64) *Invoice {
	i.Counterpart = &PartyType{
		VatNumber: &vat,
		Country:   &country,
		Branch:    &branch,
	}
	return i
}

// SetIssuer sets the invoice's issuer
func (i *Invoice) SetIssuer(vat string, country string, branch uint64) *Invoice {
	i.Issuer = &PartyType{
		VatNumber: &vat,
		Country:   &country,
		Branch:    &branch,
	}
	return i
}

// SetPaymentMethod sets the invoice's payment method
func (i *Invoice) SetPaymentMethod(method mydatavalues.InvoicePaymentType, amount float64, paymentInfo string) *Invoice {
	paymentDetail := &PaymentMethodDetailsType{
		Type:              &method, // values 1 - 7 are valid
		Amount:            &amount,
		PaymentMethodInfo: paymentInfo,
	}
	i.PaymentMethods = &PaymentMethods{
		PaymentMethodDetails: paymentDetail,
	}

	return i
}

// AddTaxTotals adds a tax total to the invoice's tax totals array
func (i *Invoice) AddTaxTotals(taxType mydatavalues.TaxType, taxCategory uint, underlyingValue float64, amount float64, id byte) *Invoice {
	if i.TaxesTotals == nil {
		i.TaxesTotals = &TaxTotalsType{
			Taxes: make([]*Taxes, 0),
		}
	}

	// if values are zero, omit them from the struct
	underlyingValuePtr := &underlyingValue
	idPtr := &id
	if underlyingValue == 0 {
		underlyingValuePtr = nil
	}
	if id == 0 {
		idPtr = nil
	}

	taxEntry := &Taxes{
		TaxType:         &taxType,
		TaxCategory:     &taxCategory,
		UnderlyingValue: underlyingValuePtr,
		TaxAmount:       &amount,
		ID:              idPtr,
	}
	i.TaxesTotals.Taxes = append(i.TaxesTotals.Taxes, taxEntry)
	return i
}

// endregion
