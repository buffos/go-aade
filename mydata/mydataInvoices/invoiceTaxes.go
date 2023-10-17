package mydataInvoices

import "github.com/buffos/go-aade/mydata/mydatavalues"

type TaxTotalsType struct {
	Taxes []*Taxes `xml:"taxes"`
}

type Taxes struct {
	TaxType         *mydatavalues.TaxType `xml:"taxType"`         // Είδος Φόρου
	TaxCategory     *uint                 `xml:"taxCategory"`     // Κατηγορία Φόρου
	UnderlyingValue *float64              `xml:"underlyingValue"` // Υποκείμενη Αξία (την αξία στην οποία εφαρμόζεται ο συγκεκριμένος φόρος)
	TaxAmount       *float64              `xml:"taxAmount"`       // Ποσό Φόρου
	ID              *byte                 `xml:"id"`              // Αύξων αριθμός γραμμής
}
