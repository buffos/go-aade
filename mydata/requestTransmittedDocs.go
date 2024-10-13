package mydata

import (
	"strconv"
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

// RequestTransmittedDocs gets the invoices from myDATA that the user has submitted.
func (c *Client) RequestTransmittedDocs(params mydataInvoices.RequestDocsParams) (int, *mydataInvoices.RequestedDoc, error) {
	return Requester[*mydataInvoices.RequestDocsParams, mydataInvoices.RequestedDoc](c, &params, URLRequestTransmittedDocs)
}

// RequestTransmittedDocsPastDays gets the invoices from myDATA for the past days
func (c *Client) RequestTransmittedDocsPastDays(days int) (int, *mydataInvoices.RequestedDoc, error) {
	return c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
		Mark:     "1",
		DateFrom: time.Now().AddDate(0, 0, -days),
		DateTo:   time.Now(),
	})
}

// RequestTransmittedDocWithMark gets the invoices from myDATA with a specific mark
func (c *Client) RequestTransmittedDocWithMark(mark uint) (int, *mydataInvoices.RequestedDoc, error) {
	return c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
		Mark:    strconv.Itoa(int(mark - 1)),
		MaxMark: strconv.Itoa(int(mark)),
	})
}

// RequestTransmittedDocBetweenMarks gets the invoices from myDATA between marks
func (c *Client) RequestTransmittedDocBetweenMarks(lowMark, highMark uint) (int, *mydataInvoices.RequestedDoc, error) {
	return c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
		Mark:    strconv.Itoa(int(lowMark - 1)),
		MaxMark: strconv.Itoa(int(highMark)),
	})
}
