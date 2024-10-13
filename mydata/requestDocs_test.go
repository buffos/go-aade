package mydata

import (
	"fmt"
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
		Mark:     "0",
		DateFrom: time.Now().AddDate(0, -16, 0), // last 3 months
		DateTo:   time.Now(),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code, "status code should be 200")

	fmt.Printf("Πλήθος παραστατικών: %d\n", len(docs.InvoicesDoc.Invoices))
	for _, invoice := range docs.InvoicesDoc.Invoices {
		fmt.Printf("Από ΑΦΜ: %s στις %s\n", *invoice.Issuer.VatNumber, invoice.InvoiceHeader.IssueDate)
	}
	fmt.Printf("Αναλυτικά:\n\n\n")
	if docs != nil {
		docs.Print()
	}
}
