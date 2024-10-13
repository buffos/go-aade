package mydata

import (
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

// RequestMyIncome requests the invoices that have income characterization for the user for the given period
func (c *Client) RequestMyIncome(params mydataInvoices.RequestMyIncomeParams) (int, *mydataInvoices.RequestedBookInfo, error) {
	return Requester[*mydataInvoices.RequestMyIncomeParams, mydataInvoices.RequestedBookInfo](c, &params, URLRequestMyIncome)
}

// RequestMyExpenses requests the invoices that have expense characterization for the user for the given period
func (c *Client) RequestMyExpenses(params mydataInvoices.RequestMyExpensesParams) (int, *mydataInvoices.RequestedBookInfo, error) {
	return Requester[*mydataInvoices.RequestMyExpensesParams, mydataInvoices.RequestedBookInfo](c, &params, URLRequestMyExpenses)
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
