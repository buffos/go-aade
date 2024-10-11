package mydata

import (
	"net/http"
	"testing"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/stretchr/testify/require"
)

func TestRequestTransmittedDocs(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
		Mark: "1",
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		docs.Print()
	}
}

func TestRequestTransmittedDocsPastDays(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocsPastDays(1)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		docs.Print()
	}
}

func TestRequestTransmittedDocsWithMark(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocWithMark(invoiceWithKnownMark)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		docs.Print()
	}
}
