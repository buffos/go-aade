package mydataInvoices

import (
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type IncomeClassificationsDoc struct {
	Xmlns                       string                         `xml:"xmlns,attr"`
	SchemaLocation              string                         `xml:"xsi:schemaLocation,attr"`
	XmlnsXsi                    string                         `xml:"xmlns:xsi,attr"`
	XmlnsIcls                   string                         `xml:"xmlns:icls,attr"`
	IncomeInvoiceClassification []*IncomeInvoiceClassification `xml:"incomeInvoiceClassification"` // Χαρακτηρισμοί Εσόδων Πρότυπων Παραστατικών ΑΑΔΕ
}

func (i *IncomeClassificationsDoc) Print() {
	fmt.Println("IncomeClassificationsDoc Length:", len(i.IncomeInvoiceClassification))
	for _, invoiceClassification := range i.IncomeInvoiceClassification {
		invoiceClassification.Print()
	}
}

type IncomeInvoiceClassification struct {
	InvoiceMark                         uint64                                 `xml:"invoiceMark"`                         // Μοναδικός Αριθμός Καταχώρησης Παραστατικού
	ClassificationMark                  *uint64                                `xml:"classificationMark"`                  // Αποδεικτικό Λήψης Χαρακτηρισμού Εσόδων. Συμπληρώνεται από την Υπηρεσία
	EntityVatNumber                     *string                                `xml:"entityVatNumber"`                     // ΑΦΜ Οντότητας Αναφοράς (μόνο όταν κληθεί από λογιστή ή εκπρόσωπο)
	TransactionMode                     *int                                   `xml:"transactionMode"`                     // Αιτιολογία Συναλλαγής. 1:Reject 2: Deviation.
	InvoicesIncomeClassificationDetails []*InvoicesIncomeClassificationDetails `xml:"invoicesIncomeClassificationDetails"` // Στοιχεία Χαρακτηρισμού Εσόδων
}

func (i *IncomeInvoiceClassification) Print() {
	fmt.Printf("Κατηγορία Εσόδων Τιμολογίου: %d", i.InvoiceMark)
	if i.ClassificationMark != nil {
		fmt.Println("ClassificationMark:", *i.ClassificationMark)
	}
	if i.EntityVatNumber != nil {
		fmt.Println("EntityVatNumber:", *i.EntityVatNumber)
	}
	if i.TransactionMode != nil {
		if *i.TransactionMode == 1 {
			fmt.Println("TransactionMode: Reject")
		} else if *i.TransactionMode == 2 {
			fmt.Println("TransactionMode: Deviation")
		} else {
			fmt.Println("TransactionMode: Unknown")
		}
	}
	for _, detail := range i.InvoicesIncomeClassificationDetails {
		detail.Print()
	}
}

type InvoicesIncomeClassificationDetails struct {
	LineNumber                     int                         `xml:"lineNumber"`                     // Αριθμός Γραμμής Παραστατικού
	IncomeClassificationDetailData []*IncomeClassificationType `xml:"incomeClassificationDetailData"` // Στοιχεία Χαρακτηρισμού Εσόδων
}

func (i *InvoicesIncomeClassificationDetails) Print() {
	fmt.Println("LineNumber:", i.LineNumber)
	for _, detail := range i.IncomeClassificationDetailData {
		detail.Print()
	}
}

func NewIncomeClassificationDoc() *IncomeClassificationsDoc {
	return &IncomeClassificationsDoc{
		Xmlns:                       "https://www.aade.gr/myDATA/incomeClassificaton/v1.0",
		XmlnsXsi:                    "http://www.w3.org/2001/XMLSchema-instance",
		SchemaLocation:              "https://www.aade.gr/myDATA/incomeClassificaton/v1.0 schema.xsd",
		XmlnsIcls:                   mydatavalues.XmlnsICLS,
		IncomeInvoiceClassification: make([]*IncomeInvoiceClassification, 0),
	}
}

func (d *IncomeClassificationsDoc) RejectClassification(mark uint64, entityVatNumber string) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}
	transactionMode := 1
	d.IncomeInvoiceClassification = append(d.IncomeInvoiceClassification, &IncomeInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		TransactionMode: &transactionMode,
	})
}

func (d *IncomeClassificationsDoc) DeviateClassification(mark uint64, entityVatNumber string) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}
	transactionMode := 2
	d.IncomeInvoiceClassification = append(d.IncomeInvoiceClassification, &IncomeInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		TransactionMode: &transactionMode,
	})
}

// EditLineNumberDetail adds an income classification to an invoice with a given mark for a specific line number.
func (d *IncomeClassificationsDoc) EditLineNumberDetail(
	mark uint64, entityVatNumber string, lineNumber int,
	clType mydatavalues.IncomeClassificationTypeStringType,
	clCategory mydatavalues.IncomeClassificationCategoryStringType,
	amount float64, id byte) {
	var entityVatNumberPointer *string
	if entityVatNumber == "" {
		entityVatNumberPointer = nil
	} else {
		entityVatNumberPointer = &entityVatNumber
	}

	// search d.IncomeInvoiceClassification to see if the invoice mark already exists
	for _, invoiceClassification := range d.IncomeInvoiceClassification {
		if invoiceClassification.InvoiceMark == mark {
			// search invoiceClassification.InvoicesExpensesClassificationDetails to see if the line number already exists
			for _, classificationDetails := range invoiceClassification.InvoicesIncomeClassificationDetails {
				if classificationDetails.LineNumber == lineNumber {
					// append to the existing classification
					classificationDetails.IncomeClassificationDetailData = append(
						classificationDetails.IncomeClassificationDetailData,
						NewIncomeClassification(clType, clCategory, amount, id))
					return
				}
			}
			// add a new classification and a new array of classifications with the given line number (mark exists but not line number)
			invoiceClassification.InvoicesIncomeClassificationDetails = append(
				invoiceClassification.InvoicesIncomeClassificationDetails, &InvoicesIncomeClassificationDetails{
					LineNumber:                     lineNumber,
					IncomeClassificationDetailData: []*IncomeClassificationType{NewIncomeClassification(clType, clCategory, amount, id)},
				})
			return
		}
	}
	d.IncomeInvoiceClassification = append(d.IncomeInvoiceClassification, &IncomeInvoiceClassification{
		InvoiceMark:     mark,
		EntityVatNumber: entityVatNumberPointer,
		InvoicesIncomeClassificationDetails: []*InvoicesIncomeClassificationDetails{
			{
				LineNumber:                     lineNumber,
				IncomeClassificationDetailData: []*IncomeClassificationType{NewIncomeClassification(clType, clCategory, amount, id)},
			},
		},
	})
}
