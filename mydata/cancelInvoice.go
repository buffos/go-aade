package mydata

import (
	"context"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"
)

func (c *Client) CancelInvoice(params mydataInvoices.CancelInvoiceParams) (int, *mydataInvoices.ResponseDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getURL(URLCancelInvoice, queryArgs), nil)

	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}

	result, err := ParseXMLResponse[mydataInvoices.ResponseDoc](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
}
