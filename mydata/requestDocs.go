package mydata

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"

	"strconv"
	"time"
)

// RequestDocs gets the requested documents, that others have submitted from myDATA
// we get back the response status code and the requested documents
func (c *Client) RequestDocs(params mydataInvoices.RequestDocsParams) (int, *mydataInvoices.RequestedDoc, error) {
	return Requester[*mydataInvoices.RequestDocsParams, mydataInvoices.RequestedDoc](c, &params, URLRequestDocs)
}

// RequestDocsPastDays gets the invoices from myDATA for the past days
func (c *Client) RequestDocsPastDays(days int) (int, *mydataInvoices.RequestedDoc, error) {
	return c.RequestDocs(mydataInvoices.RequestDocsParams{
		Mark:     "1",
		DateFrom: time.Now().AddDate(0, 0, -days),
		DateTo:   time.Now(),
	})
}

// RequestDocWithMark gets the invoices from myDATA with a specific mark
func (c *Client) RequestDocWithMark(mark uint) (int, *mydataInvoices.RequestedDoc, error) {
	return c.RequestDocs(mydataInvoices.RequestDocsParams{
		Mark:    strconv.Itoa(int(mark - 1)),
		MaxMark: strconv.Itoa(int(mark)),
	})
}
