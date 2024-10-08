package mydataInvoices

import "fmt"

type CancelledInvoice struct {
	InvoiceMark      *uint64 `xml:"invoiceMark"`
	CancellationMark *uint64 `xml:"cancellationMark"`
	CancellationDate string  `xml:"cancellationDate"`
}

func (c *CancelledInvoice) Print() {
	fmt.Println("Τιμολόγιο Ακύρωσης:")
	if c.InvoiceMark != nil {
		fmt.Println("InvoiceMark:", *c.InvoiceMark)
	}
	if c.CancellationMark != nil {
		fmt.Println("CancellationMark:", *c.CancellationMark)
	}
	fmt.Println("CancellationDate:", c.CancellationDate)
}

type CancelledInvoicesDoc struct {
	CancelledInvoices []*CancelledInvoice `xml:"cancelledInvoice"`
}

func (c *CancelledInvoicesDoc) Print() {
	fmt.Println("Cancelled invoices Length:", len(c.CancelledInvoices))
	for _, invoice := range c.CancelledInvoices {
		invoice.Print()
	}
}
