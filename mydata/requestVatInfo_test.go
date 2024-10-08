package mydata

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/stretchr/testify/require"
)

func TestRequestVatInfo(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	code, docs, err := c.RequestVatInfo(mydataInvoices.RequestVatInfoParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")
	require.NotEmpty(t, docs.Xmlns, "docs xmlns should not be empty")

	// output the number of docs returned
	fmt.Printf("Number of docs returned: %d\n", len(docs.VatInfo))
	// output if there is a continuation token
	docs.Print()
}
