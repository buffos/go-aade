package mydata

import (
	"context"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"
	"time"
)

// RequestMyIncome requests the invoices that have income characterization for the user for the given period
func (c *Client) RequestMyIncome(params mydataInvoices.RequestMyIncomeParams) (int, *mydataInvoices.RequestedBookInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getURL(URLRequestMyIncome, queryArgs), nil)
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}
	result, err := ParseXMLResponse[mydataInvoices.RequestedBookInfo](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
}

// RequestMyExpenses requests the invoices that have expense characterization for the user for the given period
func (c *Client) RequestMyExpenses(params mydataInvoices.RequestMyExpensesParams) (int, *mydataInvoices.RequestedBookInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()
	queryArgs, err := params.ToMap()
	if err != nil {
		return InternalErrorCode, nil, ErrorQueryURLCreation
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getURL(URLRequestMyExpenses, queryArgs), nil)
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}
	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}
	result, err := ParseXMLResponse[mydataInvoices.RequestedBookInfo](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	return response.StatusCode, result, nil
}

// RequestMyIncomeLastDays requests the invoices that have income characterization for the user for the last days
func (c *Client) RequestMyIncomeLastDays(days int) (int, *mydataInvoices.RequestedBookInfo, error) {
	return c.RequestMyIncome(mydataInvoices.RequestMyIncomeParams{
		DateFrom: time.Now().AddDate(0, 0, -days),
		DateTo:   time.Now(),
	})
}

// RequestMyExpensesLastDays requests the invoices that have expense characterization for the user for the last days
func (c *Client) RequestMyExpensesLastDays(days int) (int, *mydataInvoices.RequestedBookInfo, error) {
	return c.RequestMyExpenses(mydataInvoices.RequestMyExpensesParams{
		DateFrom: time.Now().AddDate(0, 0, -days),
		DateTo:   time.Now(),
	})
}
