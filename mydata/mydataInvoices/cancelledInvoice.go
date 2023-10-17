package mydataInvoices

type CancelledInvoice struct {
	InvoiceMark      *uint64 `xml:"invoiceMark"`
	CancellationMark *uint64 `xml:"cancellationMark"`
	CancellationDate string  `xml:"cancellationDate"`
}

type CancelledInvoicesDoc struct {
	CancelledInvoices []*CancelledInvoice `xml:"cancelledInvoice"`
}
