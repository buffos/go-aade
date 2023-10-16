package pdfinvoice

type InvoiceCreator interface {
	InvoiceHeader(doc *Document)
	InvoiceFooter(doc *Document, endOfDoc bool)
	InvoiceMiddle(doc *Document)
}

type SimpleA4Invoice struct{}

func (s SimpleA4Invoice) InvoiceHeader(doc *Document) {
	// region append issuer header
	x, y, _, _ := doc.pdf.GetMargins()
	issuerBottomY := doc.AppendIssuerToDoc(x, y)
	doc.SetXY(BaseMargin, issuerBottomY)
	// endregion

	// region appends invoice header details
	currentY := doc.AppendTableRow(
		BaseMargin, 50, []float64{188}, 10, 1,
		doc.Options.DarkBgColor, doc.Options.LightTextColor,
		[]string{doc.Type}, LargeTextFontSize, "0")

	currentY = doc.AppendTableRow(
		BaseMargin, currentY, []float64{40, 25, 54, 66}, 10, 1,
		doc.Options.LightBgColor, doc.Options.DarkTextColor,
		[]string{
			doc.GetSeries(),
			doc.GetAA(),
			doc.GetDate(),
			doc.GetMark(),
		}, BaseTextFontSize, "0")
	currentY += LineHeight
	// endregion

	// region appends customer information
	if doc.CounterPart != nil {
		_ = doc.AppendCounterPartToDoc(BaseMargin, currentY)
		doc.SetXY(BaseMargin, currentY)
	}
	// endregion

	// region payment methods
	doc.CreateLabeledLines(
		doc.Options.Layout.CustomerColumnTwoX, doc.Options.Layout.CustomerLabelWidth, doc.Options.Layout.CustomerValueWidth,
		doc.Options.Layout.CustomerLineHeight, doc.Options.Layout.CustomerLabelFontSize, doc.Options.Layout.CustomerValueFontSize,
		doc.Options.DarkTextColor, doc.Options.BaseTextColor, doc.Options.DefaultBgColor,
		[]string{doc.Options.TextPaymentTermTitle + ":"},
		[]string{doc.PaymentMethod}, "L")

	// endregion

	// region qrcode
	doc.AppendQRCode(doc.Options.Layout.QRCodeX, doc.Options.Layout.QRCodeY,
		doc.Options.Layout.QRCodeWidth, doc.Options.Layout.QRCodeWidth, doc.QRCodeString)
	// endregion
}
func (s SimpleA4Invoice) InvoiceFooter(doc *Document, endOfDoc bool) {
	doc.AppendTaxes(endOfDoc)
	doc.AppendNotes()
	doc.AppendTotals(endOfDoc)
	doc.AppendFootnote("-")
}
func (s SimpleA4Invoice) InvoiceMiddle(doc *Document) {
	// region appends invoice details
	currentY := 134.0
	currentY = doc.AppendInvoiceDetailsHeader(currentY)
	if doc.InvoiceDetails != nil && len(doc.InvoiceDetails.Entries) > 0 {
		for index, entry := range doc.InvoiceDetails.Entries {
			currentY = doc.AppendInvoiceDetail(currentY, entry) // handle new page
			currentY += doc.Options.Layout.InvoiceRowOffset
			remainingItems := len(doc.InvoiceDetails.Entries) - index + 1
			if currentY > doc.Options.Layout.MaxAllowedDetailsHeight && remainingItems > 2 {
				doc.pdf.AddPage()
				currentY = 134.0
				currentY = doc.AppendInvoiceDetailsHeader(currentY)
			}
		}
	}

	// endregion

}

var _ InvoiceCreator = SimpleA4Invoice{} // compile time check
