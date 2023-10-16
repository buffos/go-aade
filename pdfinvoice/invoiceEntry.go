package pdfinvoice

import "fmt"

// InvoiceEntry represents an invoice entry. Theses are the data we need to create an invoice entry row
type InvoiceEntry struct {
	Description       string  `json:"description,omitempty" validate:"required,min=1,max=41"`
	Quantity          float64 `json:"quantity,omitempty" validate:"required"`
	UnitOfMeasurement string  `json:"unit_of_measurement,omitempty"`
	NetPrice          float64 `json:"net_price,omitempty" validate:"required"`
	Discount          float64 `json:"discount,omitempty"`
	FinalPrice        float64 `json:"final_price,omitempty"`
	TaxPercent        float64 `json:"tax,omitempty"`
	TaxAmount         float64 `json:"tax_amount,omitempty"`
}

// TaxEntry represents the taxes of an invoice
type TaxEntry struct {
	NetAmount        float64 `json:"net_amount,omitempty" validate:"required"`
	TaxAmount        float64 `json:"tax_amount,omitempty" validate:"required"`
	FinalAmount      float64 `json:"final_amount,omitempty" validate:"required"`
	WithHoldingTaxes float64 `json:"with_holding_taxes,omitempty" validate:"required"`
	Deductions       float64 `json:"deductions,omitempty" validate:"required"`
	StampDuty        float64 `json:"stamp_duty,omitempty" validate:"required"`
	Fees             float64 `json:"fees,omitempty" validate:"required"`
	OtherTaxes       float64 `json:"other_taxes,omitempty" validate:"required"`
}

// InvoiceDetails represents the details of an invoice
type InvoiceDetails struct {
	Entries []*InvoiceEntry `json:"entries,omitempty" validate:"required"`
}

func (inv *InvoiceDetails) AddEntry(entry *InvoiceEntry) {
	inv.Entries = append(inv.Entries, entry)
}

func (inv *InvoiceDetails) CalculateTotals() (net, tax, total, discount float64) {
	for _, entry := range inv.Entries {
		net += entry.NetPrice
		discount += entry.Discount
		tax += entry.TaxAmount
	}
	net -= discount
	total = net + tax
	return
}

// CalculateTotals calculates the totals of the invoice and appends it to the Taxes struct
func (doc *Document) CalculateTotals() {
	net, tax, total, discount := doc.InvoiceDetails.CalculateTotals()
	doc.Taxes.NetAmount = net
	doc.Taxes.TaxAmount = tax
	doc.Taxes.FinalAmount = total
	doc.TotalDiscount = discount
}

// AppendInvoiceDetailsHeader appends the invoice details to the document. This is the header of the table we create
func (doc *Document) AppendInvoiceDetailsHeader(y float64) float64 {
	return doc.AppendTableRow(
		BaseMargin, y, []float64{60, 17, 10, 19, 20, 25, 13, 22}, 10, doc.Options.Layout.InvoiceColumnGap,
		doc.Options.DarkBgColor, doc.Options.LightTextColor,
		[]string{
			doc.Options.TextItemsNameTitle,
			doc.Options.TextItemsQuantityTitle,
			doc.Options.TextItemsMeasurementUnit,
			doc.Options.TextItemsUnitCostTitle,
			doc.Options.TextItemsDiscountTitle,
			doc.Options.TextItemsTotalHTTitle,
			doc.Options.TextItemsTaxTitle,
			doc.Options.TextItemsTotalTTCTitle,
		},
		SmallTextFontSize, "0")
}

// AppendInvoiceDetail appends an invoice detail to the document. This is a row of the table we create
func (doc *Document) AppendInvoiceDetail(y float64, entry *InvoiceEntry) float64 {
	if entry == nil {
		return y
	}
	return doc.AppendTableRow(
		BaseMargin, y, []float64{60, 17, 10, 19, 20, 25, 13, 22}, 10, doc.Options.Layout.InvoiceColumnGap,
		doc.Options.LightBgColor, doc.Options.DarkTextColor,
		[]string{
			entry.Description,
			fmt.Sprintf("%.2f", entry.Quantity),
			entry.UnitOfMeasurement,
			fmt.Sprintf("%.2f", entry.NetPrice),
			fmt.Sprintf("%.2f", entry.Discount),
			fmt.Sprintf("%.2f", entry.NetPrice-entry.Discount),
			fmt.Sprintf("%.2f %%", entry.TaxPercent),
			fmt.Sprintf("%.2f", entry.NetPrice-entry.Discount+entry.TaxAmount),
		},
		SmallTextFontSize, "0")
}
