package invoicesfactory

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"time"
)

var location, _ = time.LoadLocation("Europe/Athens")

// region create invoice

type SalesInvoiceParams struct {
	Series         string
	AA             string
	IssueDate      time.Time
	InvType        mydatavalues.InvoiceType
	Currency       string
	IssuerVat      string
	IssuerCountry  string
	IssuerBranch   uint64
	CounterVat     string
	CounterCountry string
	CounterBranch  uint64
	PaymentMethod  mydatavalues.InvoicePaymentType
	PaymentAmount  float64
	PaymentInfo    string
}

func CreateSalesInvoice(params *SalesInvoiceParams) *mydataInvoices.Invoice {
	invoice := mydataInvoices.NewInvoice(params.Series, params.AA, params.IssueDate, params.InvType)
	invoice.SetIssuer(params.IssuerVat, params.IssuerCountry, params.IssuerBranch)
	invoice.SetCounterPart(params.CounterVat, params.CounterCountry, params.CounterBranch)
	invoice.SetPaymentMethod(params.PaymentMethod, params.PaymentAmount, params.PaymentInfo)

	invoice.InvoiceHeader.SetCurrency(params.Currency)

	return invoice
}

type SalesInvoiceWithMovementParams struct {
	DispatchTimestamp time.Time
	VehicleNumber     string
	Movement          mydatavalues.InvoicePurposeOfMovement
}

func CreateSalesInvoiceWithMovement(sales *SalesInvoiceParams, params *SalesInvoiceWithMovementParams) *mydataInvoices.Invoice {
	invoice := CreateSalesInvoice(sales)
	invoice.InvoiceHeader.
		SetDispatchDate(params.DispatchTimestamp.In(location)).
		SetDispatchTime(params.DispatchTimestamp.In(location)).
		SetVehicleNumber(params.VehicleNumber).
		SetMovePurpose(params.Movement)

	return invoice
}

type PayrollInvoiceParams struct {
	Series        string
	AA            string
	IssueDate     time.Time
	Currency      string
	IssuerVat     string
	IssuerCountry string
	IssuerBranch  uint64
	PaymentMethod mydatavalues.InvoicePaymentType
	PaymentAmount float64
	PaymentInfo   string
}

func CreatePayrollInvoice(params *PayrollInvoiceParams) *mydataInvoices.Invoice {
	invoice := mydataInvoices.NewInvoice(params.Series, params.AA, params.IssueDate, mydatavalues.InvoiceTypePayroll)
	invoice.SetIssuer(params.IssuerVat, params.IssuerCountry, params.IssuerBranch)
	invoice.SetPaymentMethod(params.PaymentMethod, params.PaymentAmount, params.PaymentInfo)

	invoice.InvoiceHeader.SetCurrency(params.Currency)

	return invoice
}

type RetailReceiptParams struct {
	Series        string
	AA            string
	IssueDate     time.Time
	InvType       mydatavalues.InvoiceType
	Currency      string
	IssuerVat     string
	IssuerCountry string
	IssuerBranch  uint64
	PaymentMethod mydatavalues.InvoicePaymentType
	PaymentAmount float64
	PaymentInfo   string
}

func CreateRetailReceipt(params *RetailReceiptParams) *mydataInvoices.Invoice {
	invoice := mydataInvoices.NewInvoice(params.Series, params.AA, params.IssueDate, params.InvType)
	invoice.SetIssuer(params.IssuerVat, params.IssuerCountry, params.IssuerBranch)
	invoice.SetPaymentMethod(params.PaymentMethod, params.PaymentAmount, params.PaymentInfo)

	invoice.InvoiceHeader.SetCurrency(params.Currency)

	return invoice
}

// endregion
