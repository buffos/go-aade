package mydata

import (
	"net/http"
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/stretchr/testify/require"
)

func TestRequestMyIncome(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestMyIncome(mydataInvoices.RequestMyIncomeParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
	require.NotEmpty(t, docs.Xmlns, "books xmlns should not be empty")

	if docs != nil {
		docs.Print()
	}
}

func TestRequestMyExpenses(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestMyExpenses(mydataInvoices.RequestMyExpensesParams{
		DateFrom: time.Now().UTC().AddDate(0, -3, 0),
		DateTo:   time.Now().UTC(),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
	require.NotEmpty(t, docs.Xmlns, "books xmlns should not be empty")

	if docs != nil {
		docs.Print()
	}
}
