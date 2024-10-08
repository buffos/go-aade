package mydataInvoices

import (
	"errors"
	"time"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type InvoiceHeaderType struct {
	Series                    string                                 `xml:"series"`                    // * Σειρά παραστατικού (max 50 chars)
	Aa                        string                                 `xml:"aa"`                        // * ΑΑ παραστατικού (max 50 chars)
	IssueDate                 string                                 `xml:"issueDate"`                 // * Ημερομηνία έκδοσης
	InvoiceType               *mydatavalues.InvoiceType              `xml:"invoiceType"`               // * Είδος παραστατικού
	VatPaymentSuspension      *bool                                  `xml:"vatPaymentSuspension"`      // Αναστολή καταβολής ΦΠΑ
	Currency                  *string                                `xml:"currency"`                  // Νόμισμα. Ο κωδικός νομίσματος προέρχεται από το πρότυπο ISO 4217.
	ExchangeRate              *float64                               `xml:"exchangeRate"`              // Ισοτιμία. Συμπληρώνεται μόνο αν το νόμισμα δεν έχει τιμή EUR.
	SelfPricing               *bool                                  `xml:"selfPricing"`               // Αυτοτιμολόγηση
	CorrelatedInvoices        []*uint                                `xml:"correlatedInvoices"`        // Συσχετιζόμενα παραστατικά
	DispatchDate              *string                                `xml:"dispatchDate"`              // Ημερομηνία αποστολής
	DispatchTime              *string                                `xml:"dispatchTime"`              // Ώρα αποστολής
	VehicleNumber             *string                                `xml:"vehicleNumber"`             // Αριθμός οχήματος
	MovePurpose               *mydatavalues.InvoicePurposeOfMovement `xml:"movePurpose"`               // Σκοπός διακίνησης 1-8
	FuelInvoice               *bool                                  `xml:"fuelInvoice"`               // Παραστατικό καυσίμων
	SpecialInvoiceCategory    *mydatavalues.InvoiceSpecialCategory   `xml:"specialInvoiceCategory"`    // Ειδική κατηγορία παραστατικού 1-4
	InvoiceVariationType      *mydatavalues.InvoiceVariationType     `xml:"invoiceVariationType"`      // Τύπος απόκλισης παραστατικού 1-4
	OtherCorrelatedEntities   []*EntityType                          `xml:"otherCorrelatedEntities"`   // Λοιπές συσχετιζόμενες οντότητες
	OtherDeliveryNoteHeader   *OtherDeliveryNoteHeaderType           `xml:"otherDeliveryNoteHeader"`   // Λοιπές συσχετιζόμενες οντότητες
	IsDeliveryNote            *bool                                  `xml:"isDeliveryNote"`            // Ορίζει αν το παραστατικό είναι και δελτίο αποστολής και θα πρέπει να αποσταλούν επιπλέον στοιχεία διακίνησης.
	OtherMovePurposeTitle     *string                                `xml:"otherMovePurposeTitle"`     // Άλλος σκοπός διακίνησης Αποδεκτό μόνο για την περίπτωση που movePurpose = 19
	ThirdPartyCollection      *bool                                  `xml:"thirdPartyCollection"`      // Συλλογή από τρίτο πρόσωπο Αποδεκτό μόνο για παραστατικά τύπου 8.4 και 8.5
	MultipleConnectedMarks    *[]uint64                              `xml:"multipleConnectedMarks"`    // Πολλαπλές σημειώσεις συνδεδεμένες με το παραστατικό (Δεν είναι αποδεκτό για τα παραστατικά των τύπων 1.6, 2.4 και 5.1)
	TableAA                   *string                                `xml:"tableAA"`                   // AA ΤΡΑΠΕΖΙOY (για Δελτία Παραγγελίας Εστίασης). Μόνο για παραστατικά τύπου 8.6 και μέγιστο μήκος 50 χαρακτήρων
	TotalCancelDeliveryOrders *bool                                  `xml:"totalCancelDeliveryOrders"` // Ένδειξη συνολικής αναίρεσης Δελτίων Παραγγελίας (Αποδεκτό μόνο για παραστατικά τύπου 8.6)
}

//goland:noinspection GoUnusedExportedFunction
func NewInvoiceHeader(series string, aa string, issueDate time.Time, invType mydatavalues.InvoiceType) *InvoiceHeaderType {
	return &InvoiceHeaderType{
		Series:      series,
		Aa:          aa,
		IssueDate:   issueDate.Format(time.DateOnly),
		InvoiceType: &invType,
	}
}

// region InvoiceHeaderType

// Validate validates the invoice header
func (i *InvoiceHeaderType) validate() error {
	if i.Series == "" {
		return errors.New("invoice series is required")
	}
	if i.Aa == "" {
		return errors.New("invoice aa is required")
	}
	if i.IssueDate == "" {
		return errors.New("invoice issue date is required")
	}
	if *i.InvoiceType == "" {
		return errors.New("invoice type is required")
	}
	return nil
}

// SetVatPaymentSuspension sets the invoice header's vat payment suspension.
func (i *InvoiceHeaderType) SetVatPaymentSuspension(value bool) *InvoiceHeaderType {
	i.VatPaymentSuspension = &value
	return i
}

// SetCurrency sets the invoice header's currency.
func (i *InvoiceHeaderType) SetCurrency(currency string) *InvoiceHeaderType {
	i.Currency = &currency
	return i
}

// SetExchangeRate sets the invoice header's exchange rate.
func (i *InvoiceHeaderType) SetExchangeRate(rate float64) *InvoiceHeaderType {
	i.ExchangeRate = &rate
	return i
}

// SetSelfPricing sets the invoice header's self-pricing.
func (i *InvoiceHeaderType) SetSelfPricing(value bool) *InvoiceHeaderType {
	i.SelfPricing = &value
	return i
}

// SetDispatchDate sets the invoice header's dispatch date.
func (i *InvoiceHeaderType) SetDispatchDate(date time.Time) *InvoiceHeaderType {
	dispatchDate := date.Format(time.DateOnly)
	i.DispatchDate = &dispatchDate
	return i
}

// SetDispatchTime sets the invoice header's dispatch time.
func (i *InvoiceHeaderType) SetDispatchTime(time time.Time) *InvoiceHeaderType {
	dispatchTime := time.Format("15:04:05")
	i.DispatchTime = &dispatchTime
	return i
}

// SetVehicleNumber sets the invoice header's vehicle number.
func (i *InvoiceHeaderType) SetVehicleNumber(vehicle string) *InvoiceHeaderType {
	i.VehicleNumber = &vehicle
	return i
}

// SetMovePurpose sets the invoice header's move purpose.
func (i *InvoiceHeaderType) SetMovePurpose(value mydatavalues.InvoicePurposeOfMovement) *InvoiceHeaderType {
	i.MovePurpose = &value
	return i
}

// SetFuelInvoice sets the invoice header's fuel invoice.
func (i *InvoiceHeaderType) SetFuelInvoice(value bool) *InvoiceHeaderType {
	i.FuelInvoice = &value
	return i
}

// SetInvoiceSpecialCategory sets the invoice header's special category.
func (i *InvoiceHeaderType) SetInvoiceSpecialCategory(value mydatavalues.InvoiceSpecialCategory) *InvoiceHeaderType {
	i.SpecialInvoiceCategory = &value
	return i
}

// SetInvoiceVariationType sets the invoice header's variation type.
func (i *InvoiceHeaderType) SetInvoiceVariationType(value mydatavalues.InvoiceVariationType) *InvoiceHeaderType {
	i.InvoiceVariationType = &value
	return i
}

// SetOtherCorrelatedEntity sets the invoice header's other correlated entities.
func (i *InvoiceHeaderType) SetOtherCorrelatedEntity(category mydatavalues.EntityCategory, vat string, country string, branch uint64) *InvoiceHeaderType {
	if i.OtherCorrelatedEntities == nil {
		i.OtherCorrelatedEntities = make([]*EntityType, 0)
	}
	i.OtherCorrelatedEntities = append(i.OtherCorrelatedEntities, &EntityType{
		Type: &category,
		entityData: &PartyType{
			VatNumber: &vat,
			Country:   &country,
			Branch:    &branch,
		},
	})
	return i
}

// endregion
