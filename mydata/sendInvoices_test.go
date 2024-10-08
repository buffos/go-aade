package mydata

import (
	"fmt"
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/invoicesfactory"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestSendInvoice(t *testing.T) {
	t.Skip("skipping test")
	testCases := []struct {
		caseName         string
		createInvoices   func() *mydataInvoices.InvoicesDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Error No Invoices",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				return mydataInvoices.NewInvoices()
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       InternalErrorCode,
			wantErr:          true,
		},
		{
			caseName: "Case 2: Error Empty Header",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				invoice := mydataInvoices.Invoice{}
				invoices.Invoices = append(invoices.Invoices, &invoice)
				return invoices
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       InternalErrorCode,
			wantErr:          true,
		},
		{
			caseName: "Case 3: Error No Invoice Details",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				// create new invoice
				invoice := mydataInvoices.NewInvoice("A", "1", time.Now(), mydatavalues.InvoiceTypeSales)
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       InternalErrorCode,
			wantErr:          true,
		},
		{
			caseName: "Case 4: Error 101 - Invalid XML",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				// create new invoice
				invoice := mydataInvoices.NewInvoice("A", "1", time.Now(), mydatavalues.InvoiceTypeSales)
				// create and add invoice row
				invoiceRow := mydataInvoices.NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent)
				invoiceRow.SetItemDescription("") // empty item results to invalid xml
				invoice.AddInvoiceRow(invoiceRow)
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				require.Len(t, d.Response, 1)
				require.NotEmpty(t, d.Response[0].Errors.Error)
				require.Equal(t, "101", d.Response[0].Errors.Error[0].Code)
				fmt.Println(d.Response[0].Errors.Error)
			},
			wantedCode: 200,
			wantErr:    true,
		},
		{
			caseName: "Case 4: Create a sales invoice",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				params := invoicesfactory.SalesInvoiceParams{
					Series:         "A",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeServices,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoice(&params)
				// create and add invoice row
				invoiceRow := mydataInvoices.NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent)
				invoiceRow.AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
				require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
				require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
				require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
				require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
				require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
				require.Empty(t, d.Response[0].Errors.Error)
				//require.NotEmpty(t, d.Response[0].Errors.Error)
				//require.Equal(t, "101", d.Response[0].Errors.Error[0].Code)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 5: Create a self invoice (example 1)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				params := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      testVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     authorizedVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoice(&params)
				invoice.InvoiceHeader.SetSelfPricing(true)
				// create and add invoice row
				invoiceRow := mydataInvoices.NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent)
				invoiceRow.AddE3ExpenseClassification(mydatavalues.E3_102_001, mydatavalues.ECategory2_1, 1)
				invoiceRow.AddE3ExpenseClassification(mydatavalues.VAT_361, "", 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 6: Create sales invoice with taxes declared in tax totals (example 2)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				params := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoice(&params)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_2, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// create and add invoice row 2
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// now add taxes totals
				invoice.AddTaxTotals(mydatavalues.TaxTypeWithHoldingTax, 2, 1, 0.2, 1)
				invoice.AddTaxTotals(mydatavalues.TaxTypeDeductions, 1, 1, 0.1, 2)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 7: Create sales invoice with withholding taxes declared in invoice details row (example 4)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				params := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoice(&params)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_2, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// create and add invoice row 2
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					SetWithheldAmount(0.2, mydatavalues.WithHoldingTaxRights).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 8: Create products invoice with taxes declared in invoice details row and movement of products (example 3)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				salesParams := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				moveParams := invoicesfactory.SalesInvoiceWithMovementParams{
					DispatchTimestamp: time.Date(2021, 10, 1, 0, 0, 0, 0, location),
					VehicleNumber:     "ΑΑΑ-1234",
					Movement:          mydatavalues.MovePurposeSales,
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoiceWithMovement(&salesParams, &moveParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_2, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// create and add invoice row 2
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					SetWithheldAmount(0.2, mydatavalues.WithHoldingTaxRights).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 9: Create sales invoice on behalf of (example 6)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				salesParams := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now().In(location),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				moveParams := invoicesfactory.SalesInvoiceWithMovementParams{
					DispatchTimestamp: time.Date(2021, 10, 1, 0, 0, 0, 0, location),
					VehicleNumber:     "ΑΑΑ-1234",
					Movement:          mydatavalues.MovePurposeSales,
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoiceWithMovement(&salesParams, &moveParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_881_001, mydatavalues.ICategory1_7, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 10: Create sales invoice on behalf of, with payment (εκκαθάριση - example 7)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				salesParams := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now().In(location),
					InvType:        mydatavalues.InvoiceTypeSalesOnBehalfOfPayment,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoice(&salesParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetInvoiceDetailType(mydatavalues.InvoiceDetailSalesClearanceOfThirdParties).
					AddE3ExpenseClassification("", mydatavalues.ECategory2_9, 1)
				invoice.AddInvoiceRow(invoiceRow)
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetInvoiceDetailType(mydatavalues.InvoiceDetailPaymentFromThirdPartySales).
					SetWithheldAmount(0.2, mydatavalues.WithHoldingTaxRights).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 10: Payroll Invoice (example 13)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				payrollParams := invoicesfactory.PayrollInvoiceParams{
					Series:        "ΤΠΥ",
					AA:            "1",
					IssueDate:     time.Now().In(location),
					Currency:      "EUR",
					IssuerVat:     authorizedVat,
					IssuerCountry: "GR",
					IssuerBranch:  0,
					PaymentMethod: mydatavalues.Cash,
					PaymentAmount: 1.0,
					PaymentInfo:   "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreatePayrollInvoice(&payrollParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVATExempt).
					AddE3ExpenseClassification(mydatavalues.E3_581_001, mydatavalues.ECategory2_6, 1) // Μικτές Αποδοχές (Αμοιβές και Παροχές προσωπικού)
				invoice.AddInvoiceRow(invoiceRow)
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVATExempt).
					AddE3ExpenseClassification(mydatavalues.E3_581_002, mydatavalues.ECategory2_6, 1) // Εργοδοτικές Εισφορές (Αμοιβές και Παροχές προσωπικού)
				invoice.AddInvoiceRow(invoiceRow)

				invoice.AddTaxTotals(mydatavalues.TaxTypeWithHoldingTax, uint(mydatavalues.WithHoldingTaxFMY), 0, 0.2, 0)                     // ΦΜΥ
				invoice.AddTaxTotals(mydatavalues.TaxTypeWithHoldingTax, uint(mydatavalues.WithHoldingTaxSolidarityCompensations), 0, 0.2, 0) // Φόρος Αλληλεγγύης
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 11: Receipt Services",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				receiptParams := invoicesfactory.RetailReceiptParams{
					Series:        "ΑΠΥ",
					AA:            "1",
					IssueDate:     time.Now().In(location),
					InvType:       mydatavalues.RetailServiceReceipt,
					Currency:      "EUR",
					IssuerVat:     authorizedVat,
					IssuerCountry: "GR",
					IssuerBranch:  0,
					PaymentMethod: mydatavalues.Cash,
					PaymentAmount: 1.0,
					PaymentInfo:   "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateRetailReceipt(&receiptParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_003, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)

				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 11: Send 3 Receipts at the same time.",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				receiptParams := invoicesfactory.RetailReceiptParams{
					Series:        "ΑΠΥ",
					AA:            "1",
					IssueDate:     time.Now().In(location),
					InvType:       mydatavalues.RetailServiceReceipt,
					Currency:      "EUR",
					IssuerVat:     authorizedVat,
					IssuerCountry: "GR",
					IssuerBranch:  0,
					PaymentMethod: mydatavalues.Cash,
					PaymentAmount: 1.0,
					PaymentInfo:   "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateRetailReceipt(&receiptParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_003, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)

				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				invoices.AddInvoice(invoice)
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				//require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 11: Receipt Goods",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				receiptParams := invoicesfactory.RetailReceiptParams{
					Series:        "ΑΠΥ",
					AA:            "1",
					IssueDate:     time.Now().In(location),
					InvType:       mydatavalues.RetailReceipt,
					Currency:      "EUR",
					IssuerVat:     authorizedVat,
					IssuerCountry: "GR",
					IssuerBranch:  0,
					PaymentMethod: mydatavalues.Cash,
					PaymentAmount: 1.0,
					PaymentInfo:   "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateRetailReceipt(&receiptParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_003, mydatavalues.ICategory1_1, 1)
				invoice.AddInvoiceRow(invoiceRow)

				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 11: Send 3 Receipts at the same time. One with errors",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				receiptParams := invoicesfactory.RetailReceiptParams{
					Series:        "ΑΠΥ",
					AA:            "1",
					IssueDate:     time.Now().In(location),
					InvType:       mydatavalues.RetailServiceReceipt,
					Currency:      "EUR",
					IssuerVat:     authorizedVat,
					IssuerCountry: "GR",
					IssuerBranch:  0,
					PaymentMethod: mydatavalues.Cash,
					PaymentAmount: 1.0,
					PaymentInfo:   "paypal transaction id 1234567890",
				}
				// create new invoice
				invoice := invoicesfactory.CreateRetailReceipt(&receiptParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_003, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)

				invoice.CalculateSummary()

				invoice2 := invoicesfactory.CreateRetailReceipt(&receiptParams)
				invoiceRow2 := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_003, mydatavalues.ICategory1_7, 1)
				invoice2.AddInvoiceRow(invoiceRow2)
				invoice2.CalculateSummary()

				// add invoice to invoices
				invoices.AddInvoice(invoice)
				invoices.AddInvoice(invoice2)
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				spew.Dump(d.Response)
				require.NotEmpty(t, d.Response)
				require.Len(t, d.Response, 3)
				require.NotEqual(t, uint64(0), d.Response[0].InvoiceMark)
				require.Equal(t, uint64(0), d.Response[1].InvoiceMark)
				require.NotEqual(t, uint64(0), d.Response[2].InvoiceMark)

			},
			wantedCode: 200,
			wantErr:    true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			invoices := tc.createInvoices()
			code, docs, err := c.SendInvoices(invoices)
			tc.validateResponse(docs)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendInvoiceFilterInvalid(t *testing.T) {
	testCases := []struct {
		caseName         string
		createInvoices   func() *mydataInvoices.InvoicesDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 2: Error Empty Header",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				invoice := mydataInvoices.Invoice{}
				invoices.Invoices = append(invoices.Invoices, &invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.Nil(t, d)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 3: Error No Invoice Details",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				// create new invoice
				invoice := mydataInvoices.NewInvoice("A", "1", time.Now(), mydatavalues.InvoiceTypeSales)
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.Nil(t, d)
			},
			wantedCode: 200,
			wantErr:    false,
		},
		{
			caseName: "Case 3: Error 101 - Invalid XML",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				// create new invoice
				invoice := mydataInvoices.NewInvoice("A", "1", time.Now(), mydatavalues.InvoiceTypeSales)
				// create and add invoice row
				invoiceRow := mydataInvoices.NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent)
				invoiceRow.SetItemDescription("") // empty item results to invalid xml
				invoice.AddInvoiceRow(invoiceRow)
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				require.Len(t, d.Response, 1)
				require.NotEmpty(t, d.Response[0].Errors.Error)
				require.Equal(t, "101", d.Response[0].Errors.Error[0].Code)
				fmt.Println(d.Response[0].Errors.Error)
			},
			wantedCode: 200,
			wantErr:    true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)
	c.SetOnInvalidAction(FilterOnInvalid)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			invoices := tc.createInvoices()
			code, docs, err := c.SendInvoices(invoices)
			tc.validateResponse(docs)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendInvoicePassThroughInvalid(t *testing.T) {
	testCases := []struct {
		caseName         string
		createInvoices   func() *mydataInvoices.InvoicesDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Error Empty Header",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				invoice := mydataInvoices.Invoice{}
				invoices.Invoices = append(invoices.Invoices, &invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.Len(t, d.Response, 1)
			},
			wantedCode: 200,
			wantErr:    true,
		},
		{
			caseName: "Case 2: Error No Invoice Details",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				// create new invoice
				invoice := mydataInvoices.NewInvoice("A", "1", time.Now(), mydatavalues.InvoiceTypeSales)
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.Len(t, d.Response, 1)
			},
			wantedCode: 200,
			wantErr:    true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)
	c.SetOnInvalidAction(PassThroughOnInvalid)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			invoices := tc.createInvoices()
			code, docs, err := c.SendInvoices(invoices)
			tc.validateResponse(docs)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(docs)
			spew.Dump(err)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendInvoiceSmallSuite(t *testing.T) {
	testCases := []struct {
		caseName         string
		createInvoices   func() *mydataInvoices.InvoicesDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Error No Invoices",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				return mydataInvoices.NewInvoices()
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       InternalErrorCode,
			wantErr:          true,
		},
		{
			caseName: "Case 2: Error Empty Header",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				invoice := mydataInvoices.Invoice{}
				invoices.Invoices = append(invoices.Invoices, &invoice)
				return invoices
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       InternalErrorCode,
			wantErr:          true,
		},
		{
			caseName: "Case 8: Create products invoice with taxes declared in invoice details row and movement of products (example 3)",
			createInvoices: func() *mydataInvoices.InvoicesDoc {
				invoices := mydataInvoices.NewInvoices()
				salesParams := invoicesfactory.SalesInvoiceParams{
					Series:         "ΤΠΥ",
					AA:             "1",
					IssueDate:      time.Now(),
					InvType:        mydatavalues.InvoiceTypeSales,
					Currency:       "EUR",
					IssuerVat:      authorizedVat,
					IssuerCountry:  "GR",
					IssuerBranch:   0,
					CounterVat:     testVat,
					CounterCountry: "GR",
					CounterBranch:  0,
					PaymentMethod:  mydatavalues.Cash,
					PaymentAmount:  1.0,
					PaymentInfo:    "paypal transaction id 1234567890",
				}
				moveParams := invoicesfactory.SalesInvoiceWithMovementParams{
					DispatchTimestamp: time.Date(2021, 10, 1, 0, 0, 0, 0, location),
					VehicleNumber:     "ΑΑΑ-1234",
					Movement:          mydatavalues.MovePurposeSales,
				}
				// create new invoice
				invoice := invoicesfactory.CreateSalesInvoiceWithMovement(&salesParams, &moveParams)
				// create and add invoice row 1
				invoiceRow := mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_2, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// create and add invoice row 2
				invoiceRow = mydataInvoices.
					NewInvoiceRow(1.0, mydatavalues.InvoiceVAT24Percent).
					SetDiscountOption(true).
					SetWithheldAmount(0.2, mydatavalues.WithHoldingTaxRights).
					AddIncomeClassification(mydatavalues.IE3_561_001, mydatavalues.ICategory1_3, 1)
				invoice.AddInvoiceRow(invoiceRow)
				// make summary for invoice
				invoice.CalculateSummary()
				// add invoice to invoices
				invoices.AddInvoice(invoice)
				return invoices
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.NotEmpty(t, d.Response)
				fmt.Println(d.Response[0].Errors.Error)
				require.Len(t, d.Response, 1)
				if d.Response[0].StatusCode == "Success" {
					require.Equal(t, "Success", d.Response[0].StatusCode, "StatusCode should be Success")
					require.NotEmpty(t, d.Response[0].InvoiceUID, "InvoiceUID should not be empty")
					require.NotEmpty(t, d.Response[0].InvoiceMark, "InvoiceMark should not be empty")
					require.Equal(t, uint64(0), d.Response[0].ClassificationMark, "ClassificationMark should be 0")
					require.Empty(t, d.Response[0].AuthenticationCode, "AuthenticationCode should be empty")
					require.Equal(t, uint64(0), d.Response[0].CancellationMark, "CancellationMark should be 0")
					require.Empty(t, d.Response[0].Errors.Error)
				} else {
					require.Equal(t, "201", d.Response[0].Errors.Error[0].Code, "if we fail it must be because we do not have an authorized vat to execute")
				}
				spew.Dump(d.Response)
			},
			wantedCode: 200,
			wantErr:    false,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			invoices := tc.createInvoices()
			code, docs, err := c.SendInvoices(invoices)
			tc.validateResponse(docs)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
		})
		time.Sleep(2 * time.Second)
	}
}
