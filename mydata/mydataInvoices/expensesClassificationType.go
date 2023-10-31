package mydataInvoices

import (
	"encoding/xml"
	"errors"
	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type ExpensesClassificationType struct {
	ClassificationType     *mydatavalues.ExpenseClassificationTypeStringType      `xml:"classificationType"`     // Κωδικός Χαρακτηρισμού
	ClassificationCategory *mydatavalues.ExpensesClassificationCategoryStringType `xml:"classificationCategory"` // Κατηγορία Χαρακτηρισμού
	Amount                 float64                                                `xml:"amount"`                 // * Ποσό
	VatAmount              *float64                                               `xml:"vatAmount"`              // Ποσό ΦΠΑ
	VatCategory            *mydatavalues.InvoiceVATCategory                       `xml:"vatCategory"`            // Κατηγορία ΦΠΑ
	VatExemptionCategory   *mydatavalues.VATExceptionReasonType                   `xml:"vatExemptionCategory"`   // Κατηγορία Αιτίας Απαλλαγής ΦΠΑ
	ID                     *byte                                                  `xml:"id"`                     // Αύξων αριθμός χαρακτηρισμού
}

// NewExpenseClassification returns a new ExpensesClassificationType for E3 classifications
func NewExpenseClassification(
	clType mydatavalues.ExpenseClassificationTypeStringType,
	clCategory mydatavalues.ExpensesClassificationCategoryStringType,
	vatCategory mydatavalues.InvoiceVATCategory,
	vatExemptCategory mydatavalues.VATExceptionReasonType,
	amount float64,
	id byte) *ExpensesClassificationType {
	// region initialize pointers
	var clPtr *mydatavalues.ExpenseClassificationTypeStringType
	var catPtr *mydatavalues.ExpensesClassificationCategoryStringType
	var vatCatPtr *mydatavalues.InvoiceVATCategory
	var vatExemptCatPtr *mydatavalues.VATExceptionReasonType
	vatAmountPtr := (*float64)(nil)
	idPtr := (*byte)(nil)
	if clType != "" {
		clPtr = &clType
	}
	if clCategory != "" {
		catPtr = &clCategory
	}
	if vatCategory != 0 {
		vatCatPtr = &vatCategory
	}
	if vatExemptCategory != 0 {
		vatExemptCatPtr = &vatExemptCategory
	}
	if id != 0 {
		idPtr = &id
	}
	// vat amount is calculated only if vatCategory and vatExemptCategory are not 0
	if vatCategory != 0 {
		vatAmount := vatCategory.CalculateVAT(amount)
		vatAmountPtr = &vatAmount
	}
	// endregion

	return &ExpensesClassificationType{
		ClassificationType:     clPtr,
		ClassificationCategory: catPtr,
		Amount:                 amount,
		VatAmount:              vatAmountPtr,
		VatCategory:            vatCatPtr,
		VatExemptionCategory:   vatExemptCatPtr,
		ID:                     idPtr,
	}
}

// MarshalXML marshals the invoice's ExpensesClassificationType to xml
func (cl *ExpensesClassificationType) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type expensesClassificationDoc struct {
		ClassificationType     *string  `xml:"ecls:classificationType"`
		ClassificationCategory *string  `xml:"ecls:classificationCategory"`
		Amount                 float64  `xml:"ecls:amount"`
		VatAmount              *float64 `xml:"ecls:vatAmount"`
		VatCategory            *uint    `xml:"ecls:vatCategory"`
		VatExemptionCategory   *uint    `xml:"ecls:vatExemptionCategory"`
		ID                     *byte    `xml:"ecls:id"`
	}
	err := enc.EncodeElement(expensesClassificationDoc{
		ClassificationType:     (*string)(cl.ClassificationType),
		ClassificationCategory: (*string)(cl.ClassificationCategory),
		Amount:                 cl.Amount,
		VatAmount:              cl.VatAmount,
		VatCategory:            (*uint)(cl.VatCategory),
		VatExemptionCategory:   (*uint)(cl.VatExemptionCategory),
		ID:                     cl.ID,
	}, start)

	if err != nil {
		return errors.New("error marshaling expenses classification")
	}

	return err
}

func (cl *ExpensesClassificationType) SetVatAmount(vatAmount float64) *ExpensesClassificationType {
	cl.VatAmount = &vatAmount
	return cl
}

func (cl *ExpensesClassificationType) SetVatCategory(vatCategory mydatavalues.InvoiceVATCategory) *ExpensesClassificationType {
	cl.VatCategory = &vatCategory
	return cl
}

func (cl *ExpensesClassificationType) SetVatExemptionCategory(vatExemptionCategory mydatavalues.VATExceptionReasonType) *ExpensesClassificationType {
	cl.VatExemptionCategory = &vatExemptionCategory
	return cl
}
