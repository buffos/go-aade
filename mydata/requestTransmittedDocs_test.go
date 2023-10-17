package mydata

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestRequestTransmittedDocs(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
		Mark: "1",
	})

	spew.Dump(docs)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
}

func TestRequestTransmittedDocsPastDays(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocsPastDays(1)

	spew.Dump(docs)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
}

func TestRequestTransmittedDocsWithMark(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestTransmittedDocWithMark(invoiceWithKnownMark)
	spew.Dump(docs)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
}
