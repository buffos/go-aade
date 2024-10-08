package mydataInvoices

import (
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

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

func (t *TaxTotalsType) Print() {
	fmt.Println("Φόροι:")
	for _, tax := range t.Taxes {
		if tax.TaxType == nil {
			continue
		}
		fmt.Println("Είδος Φόρου:", tax.TaxType.String())
		if tax.TaxCategory != nil {
			fmt.Println("Κατηγορία Φόρου:", *tax.TaxCategory)
		}
		if tax.UnderlyingValue != nil {
			fmt.Println("Υποκείμενη Αξία:", *tax.UnderlyingValue)
		}
		if tax.TaxAmount != nil {
			fmt.Println("Ποσό Φόρου:", *tax.TaxAmount)
		}
		if tax.ID != nil {
			fmt.Println("Αύξων αριθμός γραμμής:", *tax.ID)
		}
	}
}
