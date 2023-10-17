package mydataInvoices

import "github.com/buffos/go-aade/mydata/mydatavalues"

type PaymentMethods struct {
	PaymentMethodDetails *PaymentMethodDetailsType `xml:"paymentMethodDetails"`
}

type PaymentMethodDetailsType struct {
	Type              *mydatavalues.InvoicePaymentType `xml:"type"`
	Amount            *float64                         `xml:"amount"`
	PaymentMethodInfo string                           `xml:"paymentMethodInfo"` // Πληροφορίες (πχ Αριθμός Τραπέζης)
}
