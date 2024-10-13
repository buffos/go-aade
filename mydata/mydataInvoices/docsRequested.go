package mydataInvoices

import "encoding/xml"

type RequestedDoc struct {
	XMLName                   xml.Name                    `xml:"http://www.aade.gr/myDATA/invoice/v1.0 RequestedDoc"`
	XmlnsICLS                 string                      `xml:"xmlns icls,attr"`
	XmlnsECLS                 string                      `xml:"xmlns ecls,attr"`
	ContinuationToken         *ContinuationToken          `xml:"continuationToken"`
	InvoicesDoc               *InvoicesDoc                `xml:"invoicesDoc"`
	CancelledInvoicesDoc      *CancelledInvoicesDoc       `xml:"cancelledInvoicesDoc"`
	IncomeClassificationDoc   *IncomeClassificationsDoc   `xml:"incomeClassificationsDoc"`
	ExpensesClassificationDoc *ExpensesClassificationsDoc `xml:"expensesClassificationDoc"`
	PaymentMethodsDoc         *PaymentMethodsDoc          `xml:"paymentMethodsDoc"`
}

func (d *RequestedDoc) GetNextPartitionData() (string, string) {
	if d.ContinuationToken == nil {
		return "", ""
	}
	return d.ContinuationToken.NextPartitionKey, d.ContinuationToken.NextRowKey
}

func (d *RequestedDoc) Print() {
	if d.ContinuationToken != nil {
		d.ContinuationToken.Print()
	}
	if d.InvoicesDoc != nil {
		d.InvoicesDoc.Print()
	}
	if d.CancelledInvoicesDoc != nil {
		d.CancelledInvoicesDoc.Print()
	}
	if d.IncomeClassificationDoc != nil {
		d.IncomeClassificationDoc.Print()
	}
	if d.ExpensesClassificationDoc != nil {
		d.ExpensesClassificationDoc.Print()
	}
	if d.PaymentMethodsDoc != nil {
		d.PaymentMethodsDoc.Print()
	}
}
