package mydata

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestRequestMyIncome(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestMyIncome(mydataInvoices.RequestMyIncomeParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	})

	spew.Dump(docs)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	require.NotEmpty(t, docs.Xmlns, "books xmlns should not be empty")
}

func TestRequestMyExpenses(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestMyExpenses(mydataInvoices.RequestMyExpensesParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	})

	spew.Dump(docs)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	require.NotEmpty(t, docs.Xmlns, "books xmlns should not be empty")
}
