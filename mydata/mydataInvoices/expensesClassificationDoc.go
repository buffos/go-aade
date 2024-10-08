package mydataInvoices

import (
	"errors"
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type ExpensesClassificationsDoc struct {
	Xmlns     string `xml:"xmlns,attr"`
	XmlnsEcls string `xml:"xmlns:ecls,attr"`
	//XmlnsXsi                      string                           `xml:"xmlns:xsi,attr"`
	//SchemaLocation string `xml:"xsi:schemaLocation,attr"`
	ExpensesInvoiceClassification []*ExpensesInvoiceClassification `xml:"expensesInvoiceClassification"` // Χαρακτηρισμοί Εξόδων Πρότυπων Παραστατικών ΑΑΔΕ
}

func (d *ExpensesClassificationsDoc) Print() {
	for _, classification := range d.ExpensesInvoiceClassification {
		classification.Print()
	}
}

type ExpensesInvoiceClassification struct {
	InvoiceMark                           uint64                                   `xml:"invoiceMark"`                           // Μοναδικός Αριθμός Καταχώρησης Παραστατικού
	ClassificationMark                    *uint64                                  `xml:"classificationMark"`                    // Αποδεικτικό Λήψης Χαρακτηρισμού Εσόδων. Συμπληρώνεται από την Υπηρεσία
	EntityVatNumber                       *string                                  `xml:"entityVatNumber"`                       // ΑΦΜ Οντότητας Αναφοράς (μόνο όταν κληθεί από λογιστή ή εκπρόσωπο)
	TransactionMode                       *int                                     `xml:"transactionMode"`                       // Αιτιολογία Συναλλαγής. 1:Reject 2: Deviation.
	InvoicesExpensesClassificationDetails []*InvoicesExpensesClassificationDetails `xml:"invoicesExpensesClassificationDetails"` // Στοιχεία Χαρακτηρισμού Εξόδων
	ClassificationPostMode                *byte                                    `xml:"classificationPostMode"`                // Μέθοδος Υποβολής Χαρακτηρισμού (1: Σημαίνει οτι ο χαρακτηρισμός αφορά ολόκληρο το παραστατικό 0: Ο χαρακτηρισμός αφορά μια γραμμή μόνο)
}

func (d *ExpensesInvoiceClassification) Print() {
	fmt.Printf("Κατηγορία Εξόδων Τιμολογίου: %d", d.InvoiceMark)
	if d.ClassificationMark != nil {
		fmt.Println("Αποδεικτικό Λήψης Χαρακτηρισμού Εσόδων:", *d.ClassificationMark)
	}
	if d.EntityVatNumber != nil {
		fmt.Println("ΑΦΜ Οντότητας Αναφοράς:", *d.EntityVatNumber)
	}
	if d.TransactionMode != nil {
		fmt.Println("Αιτιολογία Συναλλαγής:", *d.TransactionMode)
	}
	for _, classificationDetails := range d.InvoicesExpensesClassificationDetails {
		classificationDetails.Print()
	}
}

type InvoicesExpensesClassificationDetails struct {
	LineNumber                       int                           `xml:"lineNumber"`                       // Αριθμός Γραμμής Παραστατικού
	ExpensesClassificationDetailData []*ExpensesClassificationType `xml:"expensesClassificationDetailData"` // Στοιχεία Χαρακτηρισμού Εσόδων
}

func (d *InvoicesExpensesClassificationDetails) Print() {
	fmt.Println("Αριθμός Γραμμής Παραστατικού:", d.LineNumber)
	for _, classification := range d.ExpensesClassificationDetailData {
		classification.Print()
	}
}

func NewExpensesClassificationDoc() *ExpensesClassificationsDoc {
	return &ExpensesClassificationsDoc{
		Xmlns:     "https://www.aade.gr/myDATA/expensesClassificaton/v1.0",
		XmlnsEcls: mydatavalues.XmlnsECLS,
		//XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		//SchemaLocation: "https://www.aade.gr/myDATA/expensesClassificaton/v1.0 schema.xsd",
		ExpensesInvoiceClassification: make([]*ExpensesInvoiceClassification, 0),
	}
}

func (d *ExpensesClassificationsDoc) RejectClassification(mark uint64, entityVatNumber string) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}
	transactionMode := 1
	d.ExpensesInvoiceClassification = append(d.ExpensesInvoiceClassification, &ExpensesInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		TransactionMode: &transactionMode,
	})
}

func (d *ExpensesClassificationsDoc) DeviateClassification(mark uint64, entityVatNumber string) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}
	transactionMode := 2
	d.ExpensesInvoiceClassification = append(d.ExpensesInvoiceClassification, &ExpensesInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		TransactionMode: &transactionMode,
	})
}

// EditLineNumberDetail adds an expense classification to an invoice with a given mark for a specific line number.
// This is the old way of adding classifications to an invoice. It is kept for backwards compatibility.
// It should be used with postPerInvoice to false
func (d *ExpensesClassificationsDoc) EditLineNumberDetail(
	mark uint64, entityVatNumber string, lineNumber int,
	clType mydatavalues.ExpenseClassificationTypeStringType,
	clCategory mydatavalues.ExpensesClassificationCategoryStringType,
	amount float64) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}

	// search d.ExpensesInvoiceClassification to see if the invoice mark already exists
	for _, invoiceClassification := range d.ExpensesInvoiceClassification {
		if invoiceClassification.InvoiceMark == mark {
			// search invoiceClassification.InvoicesExpensesClassificationDetails to see if the line number already exists
			for _, classificationDetails := range invoiceClassification.InvoicesExpensesClassificationDetails {
				if classificationDetails.LineNumber == lineNumber {
					// append to the existing classification
					id := byte(len(classificationDetails.ExpensesClassificationDetailData) + 1)
					classificationDetails.ExpensesClassificationDetailData = append(
						classificationDetails.ExpensesClassificationDetailData,
						NewExpenseClassification(clType, clCategory, mydatavalues.InvoiceVATCategory(0), mydatavalues.VATExceptionReasonType(0), amount, id))
					return
				}
			}
			// add a new classification and a new array of classifications with the given line number
			// (mark exists but not line number)
			invoiceClassification.InvoicesExpensesClassificationDetails = append(invoiceClassification.InvoicesExpensesClassificationDetails, &InvoicesExpensesClassificationDetails{
				LineNumber: lineNumber,
				ExpensesClassificationDetailData: []*ExpensesClassificationType{
					NewExpenseClassification(clType, clCategory, mydatavalues.InvoiceVATCategory(0), mydatavalues.VATExceptionReasonType(0), amount, byte(1)),
				},
			})
			return
		}
	}

	// if we reach this point, then the invoice mark does not exist. add a new invoice classification
	d.ExpensesInvoiceClassification = append(d.ExpensesInvoiceClassification, &ExpensesInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		InvoicesExpensesClassificationDetails: []*InvoicesExpensesClassificationDetails{
			{
				LineNumber: lineNumber,
				ExpensesClassificationDetailData: []*ExpensesClassificationType{
					NewExpenseClassification(clType, clCategory, mydatavalues.InvoiceVATCategory(0), mydatavalues.VATExceptionReasonType(0), amount, byte(1)),
				},
			},
		},
	})
}

// NewInvoiceClassificationForMark adds an expense classification to an invoice with a given mark.
// We should append at this slice, classifications for E3 and VAT. At least one for each.
func (d *ExpensesClassificationsDoc) NewInvoiceClassificationForMark(mark uint64, entityVatNumber string) *ExpensesInvoiceClassification {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}

	newInvoiceClassification := &ExpensesInvoiceClassification{
		InvoiceMark:                           mark,
		EntityVatNumber:                       entityVatNumberPointer,
		InvoicesExpensesClassificationDetails: make([]*InvoicesExpensesClassificationDetails, 0),
	}
	d.ExpensesInvoiceClassification = append(d.ExpensesInvoiceClassification, newInvoiceClassification)
	return newInvoiceClassification
}

// AddE3ClassificationDetail adds an E3 expense classification to an invoice with a given mark.
// returns a pointer to the ExpensesInvoiceClassification for chaining. this way we can add multiple e3 and vat classifications.
func (d *ExpensesInvoiceClassification) AddE3ClassificationDetail(
	clType mydatavalues.ExpenseClassificationTypeStringType,
	clCategory mydatavalues.ExpensesClassificationCategoryStringType,
	amount float64) *ExpensesInvoiceClassification {
	// search d.InvoicesExpensesClassificationDetails to see if the LineNumber zero already exists
	for _, classificationDetails := range d.InvoicesExpensesClassificationDetails {
		if classificationDetails.LineNumber == 1 {
			// append to the existing classification
			id := byte(len(classificationDetails.ExpensesClassificationDetailData) + 1)
			classificationDetails.ExpensesClassificationDetailData = append(
				classificationDetails.ExpensesClassificationDetailData,
				NewExpenseClassification(clType, clCategory, mydatavalues.InvoiceVATCategory(0), mydatavalues.VATExceptionReasonType(0), amount, id))
			return d
		}
	}
	d.InvoicesExpensesClassificationDetails = append(d.InvoicesExpensesClassificationDetails, &InvoicesExpensesClassificationDetails{
		LineNumber: 1,
		ExpensesClassificationDetailData: []*ExpensesClassificationType{
			NewExpenseClassification(clType, clCategory, mydatavalues.InvoiceVATCategory(0), mydatavalues.VATExceptionReasonType(0), amount, byte(1)), // id is one
			// since this is the first entry in the ExpensesClassificationDetailData array.
		},
	})
	return d
}

func (d *ExpensesInvoiceClassification) AddVatClassificationDetail(
	clType mydatavalues.ExpenseClassificationTypeStringType,
	clCategory mydatavalues.ExpensesClassificationCategoryStringType,
	vatCategory mydatavalues.InvoiceVATCategory,
	vatExemptionCategory mydatavalues.VATExceptionReasonType,
	amount float64) *ExpensesInvoiceClassification {
	if vatCategory == mydatavalues.InvoiceVATExempt {
		// Τα ποσά των γραμμών του παραστατικού με τιμή vatCategory 8 (Άνευ ΦΠΑ) δεν
		//χαρακτηρίζονται εδώ
		return d
	}
	if vatCategory == mydatavalues.InvoiceVAT0Percent && vatExemptionCategory != mydatavalues.Article39a {
		// Τα ποσά των γραμμών του παραστατικού με τιμή vatCategory 7 (0%) δε
		// χαρακτηρίζονται εδώ, με εξαίρεση όταν το vatExemptionCategory έχει τιμή 16 (άρθρο
		// 39α του Κώδικα ΦΠΑ)
		return d
	}
	// search d.InvoicesExpensesClassificationDetails to see if the LineNumber zero already exists
	for _, classificationDetails := range d.InvoicesExpensesClassificationDetails {
		if classificationDetails.LineNumber == 1 {
			// append to the existing classification
			id := byte(len(classificationDetails.ExpensesClassificationDetailData) + 1)
			classificationDetails.ExpensesClassificationDetailData = append(
				classificationDetails.ExpensesClassificationDetailData,
				NewExpenseClassification(clType, clCategory, vatCategory, vatExemptionCategory, amount, id))
			return d
		}
	}
	d.InvoicesExpensesClassificationDetails = append(d.InvoicesExpensesClassificationDetails, &InvoicesExpensesClassificationDetails{
		LineNumber: 1,
		ExpensesClassificationDetailData: []*ExpensesClassificationType{
			NewExpenseClassification(clType, clCategory, vatCategory, vatExemptionCategory, amount, byte(1)), // id is one
			// since this is the first entry in the ExpensesClassificationDetailData array.
		},
	})
	return d
}

// ValidateAgainstInvoice validates the classification doc against an invoice.
// requirements can be viewed in https://www.aade.gr/sites/default/files/2023-07/SendExpensesClassificationPostPerInvoiceGuidelines.pdf
func (d *ExpensesClassificationsDoc) ValidateAgainstInvoice(v *Invoice) error {
	// first look if the invoice mark exists in the classification doc
	for _, classification := range d.ExpensesInvoiceClassification {
		if classification.InvoiceMark == *v.Mark {
			//TODO: check requirements
			return nil
		}
	}
	return errors.New("invoice mark not found in classification doc")
}

func (d *ExpensesInvoiceClassification) CalculateClassificationTotals() (float64, float64) {
	var e3Total float64
	var vatTotal float64
	for _, classificationDetails := range d.InvoicesExpensesClassificationDetails {
		for _, classification := range classificationDetails.ExpensesClassificationDetailData {
			if classification.ClassificationType != nil {
				e3Total += classification.Amount
			} else {
				vatTotal += classification.Amount
			}
		}
	}
	return e3Total, vatTotal
}
