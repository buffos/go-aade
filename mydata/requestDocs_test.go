package mydata

import (
	"net/http"
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/stretchr/testify/require"
)

//Mark: "400001917182008",

func TestRequestDocs(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestDocs(mydataInvoices.RequestDocsParams{
		Mark:     "1",
		DateFrom: time.Now().AddDate(0, -2, 0), // last 3 months
		DateTo:   time.Now(),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	if docs != nil {
		docs.Print()
	}
}
