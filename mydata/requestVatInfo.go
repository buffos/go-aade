package mydata

import (
	"context"
	"errors"
	"net/http"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

func (c *Client) RequestVatInfo(params mydataInvoices.RequestVatInfoParams) (int, *mydataInvoices.RequestedVatInfoType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getURL(URLRequestVatInfo, queryArgs), nil)
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}

	result, err := ParseXMLResponse[mydataInvoices.RequestedVatInfoType](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
}
