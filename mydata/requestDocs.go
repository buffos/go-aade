package mydata

import (
	"context"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"

	"strconv"
	"time"
)

// RequestDocs gets the requested documents, that others have submitted from myDATA
// we get back the response status code and the requested documents
func (c *Client) RequestDocs(params mydataInvoices.RequestDocsParams) (int, *mydataInvoices.RequestedDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getURL(URLRequestDocs, queryArgs), nil)
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}

	//b, _ := c.responseToString(response)
	//fmt.Println(b)

	result, err := ParseXMLResponse[mydataInvoices.RequestedDoc](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
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
