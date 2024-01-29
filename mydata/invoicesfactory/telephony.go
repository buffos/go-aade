package invoicesfactory

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
)

func ClassifyTelephonyBillInvoice(mark uint64, entityVatNumber string, amount float64, vat float64) *mydataInvoices.ExpensesClassificationsDoc {
	doc := mydataInvoices.NewExpensesClassificationDoc()
	cl := doc.NewInvoiceClassificationForMark(mark, entityVatNumber)
	cl.AddE3ClassificationDetail(mydatavalues.E3_585_013, mydatavalues.ECategory2_4, amount)
	cl.AddVatClassificationDetail(
		mydatavalues.VAT_361,
		mydatavalues.ECategory2_4,
		mydatavalues.InvoiceVAT24Percent,
		mydatavalues.VATExceptionReasonType(0),
		vat)
	return doc
}
