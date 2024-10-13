package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/invoicesfactory"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"github.com/stretchr/testify/require"
)

func TestCreateSalesInvoice(t *testing.T) {
	invoices := mydataInvoices.NewInvoices()
	params := &invoicesfactory.SalesInvoiceParams{
		Series:         "B",
		AA:             "23",
		IssueDate:      time.Now(),
		InvType:        mydatavalues.InvoiceTypeSales,
		Currency:       "EUR",
		IssuerVat:      "1234567890",
		IssuerCountry:  "GR",
		IssuerBranch:   1,
		CounterVat:     "0987654321",
		CounterCountry: "GR",
		CounterBranch:  0,
		PaymentMethod:  mydatavalues.Cash,
		PaymentAmount:  483.6,
		PaymentInfo:    "Payment info",
	}
	moveParams := &invoicesfactory.SalesInvoiceWithMovementParams{
		DispatchTimestamp: time.Now(),
		VehicleNumber:     "1234567890",
		Movement:          mydatavalues.MovePurposeSales,
	}
	invoice := invoicesfactory.CreateSalesInvoiceWithMovement(params, moveParams)
	invoiceRow := mydataInvoices.NewInvoiceRow(390, mydatavalues.InvoiceVAT24Percent)
	invoiceRow.AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_1, 390)
	invoice.AddInvoiceRow(invoiceRow)
	invoice.CalculateSummary()
	invoices.AddInvoice(invoice)

	body := invoices.MustXML()
	fmt.Println(string(body))
	require.NotEmpty(t, body)
}
