package mydataInvoices

import (
	"encoding/xml"
	"errors"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type InvoicesDoc struct {
	Xmlns             string     `xml:"xmlns,attr"`
	XmlnsXsi          string     `xml:"xmlns:xsi,attr"`
	XmlnsICLS         string     `xml:"xmlns:icls,attr"`
	XmlnsECLS         string     `xml:"xmlns:ecls,attr"`
	XsiSchemaLocation string     `xml:"xsi:schemaLocation,attr"`
	Invoices          []*Invoice `xml:"invoice"`
}

func (i *InvoicesDoc) Print() {
	for _, invoice := range i.Invoices {
		invoice.Print()
	}
}

func (i *InvoicesDoc) ShouldXML() ([]byte, error) {
	return xml.Marshal(i)
}

func (i *InvoicesDoc) MustXML() []byte {
	body, err := xml.Marshal(i)
	if err != nil {
		panic(err)
	}
	return body
}

//goland:noinspection GoUnusedExportedFunction
func NewInvoices() *InvoicesDoc {
	return &InvoicesDoc{
		Xmlns:             mydatavalues.Xmlns,
		XmlnsXsi:          mydatavalues.XmlnsXsi,
		XmlnsICLS:         mydatavalues.XmlnsICLS,
		XmlnsECLS:         mydatavalues.XmlnsECLS,
		XsiSchemaLocation: mydatavalues.SchemaLocation,
		Invoices:          make([]*Invoice, 0),
	}
}

// region InvoicesDoc

// AddInvoice adds an invoice to the invoices array
func (i *InvoicesDoc) AddInvoice(invoice *Invoice) *InvoicesDoc {
	i.Invoices = append(i.Invoices, invoice)
	return i
}

// Validate validates the invoices doc. Checks if the invoices array is not empty and validates each invoice
func (i *InvoicesDoc) Validate() error {
	if i.Invoices == nil || len(i.Invoices) == 0 {
		return errors.New("invoices are required")
	}
	for _, invoice := range i.Invoices {
		if err := invoice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// FilterInvalid returns a new invoices doc with only the valid invoices
func (i *InvoicesDoc) FilterInvalid() *InvoicesDoc {
	filtered := NewInvoices()
	for _, invoice := range i.Invoices {
		if invoice.Validate() == nil {
			filtered.Invoices = append(filtered.Invoices, invoice)
		}
	}
	return filtered
}

// endregion
