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
}
