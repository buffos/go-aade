package mydata

import (
	"net/http"
	"testing"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

//Mark: "400001917182008",

func TestCancelNonExistentInvoice(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.CancelInvoice(mydataInvoices.CancelInvoiceParams{
		Mark: "400001917182008",
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		spew.Dump(docs)
	}
}

func TestCancelInvoice(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.CancelInvoice(mydataInvoices.CancelInvoiceParams{
		Mark: "400001917778937",
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		spew.Dump(docs)
	}
}
