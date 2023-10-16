package pdfinvoice

import (
	"bytes"
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/skip2/go-qrcode"
	"image"
	"os"
)

// Document represents an invoice document.
type Document struct {
	pdf *fpdf.Fpdf

	Options       Options `json:"options,omitempty"`
	Type          string  `json:"type,omitempty" validate:"required"`
	Date          string  `json:"date,omitempty"`
	Mark          string  `json:"mark,omitempty" validate:"required"`
	Number        string  `json:"number,omitempty" validate:"required"`
	Notes         string  `json:"notes,omitempty" validate:"min_length=1,max_length=256"`
	Series        string  `json:"series,omitempty" validate:"required"`
	PaymentMethod string  `json:"paymentMethod,omitempty" validate:"required"`
	UID           string  `json:"uid,omit"`
	QRCodeString  string  `json:"qr_code_string,omitempty"`

	InvoiceDetails *InvoiceDetails `json:"invoices_details,omitempty" validate:"required"`
	Taxes          *TaxEntry       `json:"taxes,omitempty" validate:"required"`
	TotalDiscount  float64         `json:"total_discount,omitempty" validate:"required"`

	LogoFileName string   `json:"logo_file_name,omitempty"`
	Issuer       *Contact `json:"company,omitempty" validate:"required"`
	CounterPart  *Contact `json:"customer,omitempty"`
}

// NewInvoice creates a new invoice document. If the option struct is nil, default options are used.
func NewInvoice(options *Options) (*Document, error) {
	if options == nil {
		options = &defaultOptions
	}
	// create a new pdf object
	fontPath := os.Getenv(options.FontPathEnvName)
	newPdf := fpdf.New(options.Orientation, options.DocumentUnits, options.DocumentSize, fontPath)

	// set default options if none provided
	if options == nil {
		options = &Options{}
		*options = defaultOptions
	}

	// set unicode translator
	//options.UnicodeTranslateFunc = newPdf.UnicodeTranslatorFromDescriptor("")

	newPdf.AddUTF8Font(options.Font, "", options.FontFileName)
	newPdf.AddUTF8Font(options.BoldFont, "B", options.BoldFontFileName)

	invoiceDetails := &InvoiceDetails{Entries: make([]*InvoiceEntry, 0)}

	return &Document{
		pdf:            newPdf,
		Options:        *options,
		Type:           "invoice",
		InvoiceDetails: invoiceDetails,
		Taxes:          &TaxEntry{},
	}, nil
}

// AddPage adds a new page to the underlying pdf document
func (doc *Document) AddPage() {
	doc.pdf.AddPage()
}

// AppendBoxToDocument appends a box to the document at position x, y, with width w, and height h.
// The box is rounded with radius r.
// The color is specified by the color array [r, g, b] and is the draw color of the box.
// The styleStr is the style of the box. It can be "D" for dashed, "F" for filled, "DF" for dashed and filled.
func (doc *Document) AppendBoxToDocument(x, y, w, h, r float64, color []int, styleStr string) {
	doc.pdf.SetXY(x, y)
	doc.SetDrawColor(color)
	doc.pdf.RoundedRect(x-2, y-2, w, h, r, "1234", styleStr)
}

// AppendFootnote appends the footer to the document
// note is the text to be appended to the page number.
// if the note is equal to, "-", then it includes the mark and the uid of the document.
func (doc *Document) AppendFootnote(note string) {
	if note == "-" {
		note = fmt.Sprintf("%s %s", doc.GetMark(), doc.GetUID())
	}
	doc.pdf.SetFont(doc.Options.Font, "", SmallTextFontSize)
	doc.pdf.SetXY(doc.Options.Layout.FooterX, doc.Options.Layout.FooterY)
	footNote := fmt.Sprintf("%s Σελίδα: %d", note, doc.pdf.PageNo())
	doc.pdf.CellFormat(0, 0, footNote, "", 0, "C", false, 0, "")
}

// AppendImage appends an image to the document at position x, y, with width w, and height h.
// The image is specified by the imageBytes, and the name is the name to store the image internally.
func (doc *Document) AppendImage(x, y, w, h float64, imageBytes []byte, name string) {
	doc.pdf.SetXY(x, y)
	// Create filename to embed image
	fileName := doc.EncodeString(name)
	// Create reader from logo bytes
	format := readImageFormat(imageBytes)
	ioReader := bytes.NewReader(imageBytes)
	imageInfo := doc.pdf.RegisterImageOptionsReader(fileName, fpdf.ImageOptions{
		ImageType: format,
	}, ioReader)
	if imageInfo != nil {
		var imageOpt fpdf.ImageOptions
		imageOpt.ImageType = format
		doc.pdf.ImageOptions(
			fileName, x, y,
			w, h, false, imageOpt, 0, "")
	}
}

// AppendNotes appends the notes to the document. Its box with some "remarks".
func (doc *Document) AppendNotes() {
	doc.AppendBoxToDocument(
		doc.Options.Layout.NotesX, doc.Options.Layout.NotesY, doc.Options.Layout.NotesWidth, doc.Options.Layout.NotesHeight, 1,
		doc.Options.DarkBgColor, "")

	notes := fmt.Sprintf("%s: \n%s", doc.Options.TextItemNotesTitle, doc.Notes)
	doc.pdf.SetXY(doc.Options.Layout.NotesX, doc.Options.Layout.NotesY)
	doc.pdf.MultiCell(doc.Options.Layout.NotesWidth-BaseMargin, LineHeight, doc.EncodeString(notes), "0", "L", false)
}

// AppendQRCode appends the QR code to the document at position x, y, with width w, and height h.
func (doc *Document) AppendQRCode(x, y, w, h float64, qrCodeString string) {
	qrCodeBytes, err := qrcode.Encode(qrCodeString, qrcode.Medium, int(w))
	if err != nil {
		return
	}
	doc.AppendImage(x, y, w, h, qrCodeBytes, "QRCode")
}

// AppendTableRow appends a table row to the document at position x, y, with width w, and height h.
// The gap is the gap between the cells.
// The color is specified by the color array [r, g, b] and is the fill color of the row.
// The textColor is specified by the textColor array [r, g, b] and is the text color of the row.
// The values are the values of the cells.
// The fontSize is the font size of the cells.
// borderStr specifies how the cell border will be drawn. An empty string
// indicates no border, "1" indicates a full border, and one or more of "L",
// "T", "R" and "B" indicate the left, top, right and bottom sides of the
// border.
func (doc *Document) AppendTableRow(
	x float64, y float64, w []float64, h float64, gap float64,
	color []int, textColor []int, values []string, fontSize float64, borderStr string) float64 {
	doc.pdf.SetXY(x, y) // set the current position
	doc.pdf.SetFontSize(fontSize)
	doc.SetFillColor(color)
	doc.SetTextColor(textColor)

	if len(w) != len(values) {
		panic("length of w and values must be the same")
	}
	for i, value := range values {
		doc.pdf.SetX(doc.pdf.GetX() + gap)
		doc.pdf.CellFormat(w[i], h, doc.EncodeString(value), borderStr, 0, "C", true, 0, "")
	}
	return doc.pdf.GetY() + h
}

// AppendTaxes appends the taxes section to the document
// The position is specified in Options.Layout struct by TaxesX and TaxesY.
// The row height is specified in Options.Layout struct by TaxesRowHeight.
// The gap is set at 0.5.
// The fill color is the Options.DarkBgColor, and the text color is the Options.LightTextColor
// It displays empty values if we are not at the end of the document, on invoices that span multiple pages.
// This is controlled by endOfDoc.
func (doc *Document) AppendTaxes(endOfDoc bool) {
	currentY := doc.AppendTableRow(
		doc.Options.Layout.TaxesX, doc.Options.Layout.TaxesY, []float64{31, 31, 31, 31, 31, 31}, doc.Options.Layout.TaxesRowHeight, 0.5,
		doc.Options.DarkBgColor, doc.Options.LightTextColor,
		[]string{
			doc.Options.TextItemsTotalTTCTitle,
			doc.Options.TextWithHoldingTaxes,
			doc.Options.TextDeductions,
			doc.Options.TextStampTaxes,
			doc.Options.TextFeesTaxes,
			doc.Options.TextMiscTaxes,
		}, doc.Options.Layout.TaxesFontSize, "0")
	if !endOfDoc {
		currentY = doc.AppendTableRow(
			doc.Options.Layout.TaxesX, currentY, []float64{31, 31, 31, 31, 31, 31}, doc.Options.Layout.TaxesRowHeight, 0.5,
			doc.Options.LightBgColor, doc.Options.DarkTextColor,
			[]string{
				fmt.Sprintf("-"),
				fmt.Sprintf("-"),
				fmt.Sprintf("-"),
				fmt.Sprintf("-"),
				fmt.Sprintf("-"),
				fmt.Sprintf("-"),
			}, doc.Options.Layout.TaxesFontSize, "0")
	} else {
		currentY = doc.AppendTableRow(
			BaseMargin, currentY, []float64{31, 31, 31, 31, 31, 31}, 10, 0.5,
			doc.Options.LightBgColor, doc.Options.DarkTextColor,
			[]string{
				fmt.Sprintf("%.2f", doc.Taxes.FinalAmount),
				fmt.Sprintf("%.2f", doc.Taxes.WithHoldingTaxes),
				fmt.Sprintf("%.2f", doc.Taxes.Deductions),
				fmt.Sprintf("%.2f", doc.Taxes.StampDuty),
				fmt.Sprintf("%.2f", doc.Taxes.Fees),
				fmt.Sprintf("%.2f", doc.Taxes.OtherTaxes),
			}, SmallTextFontSize, "0")
	}
}

// AppendTotals appends the totals section to the document
// Those are the totals of the document and should be displayed at the footer of the document.
// The position is specified in Options.Layout struct by TotalsX and TotalsY.
// If endOfDoc is false, the values displayed are -.
// The labels for the taxes have the Options.AccentColors
// The values have the Options.DarkTextColor
// The background color is the Options.LightBgColor
func (doc *Document) AppendTotals(endOfDoc bool) {
	doc.pdf.SetXY(doc.Options.Layout.TotalsX, doc.Options.Layout.TotalsY)

	var taxesValues []string
	if !endOfDoc {
		taxesValues = []string{
			"-", "-", "-", "-", "-", "-",
		}
	} else {
		otherTaxes := doc.Taxes.OtherTaxes + doc.Taxes.Fees + doc.Taxes.StampDuty
		withholdingTaxes := doc.Taxes.WithHoldingTaxes + doc.Taxes.Deductions
		taxesValues = []string{
			fmt.Sprintf("%.2f%s", doc.Taxes.FinalAmount, doc.Options.CurrencySymbol),
			fmt.Sprintf("%.2f%s", doc.TotalDiscount, doc.Options.CurrencySymbol),
			fmt.Sprintf("%.2f%s", doc.Taxes.TaxAmount, doc.Options.CurrencySymbol),
			fmt.Sprintf("%.2f%s", otherTaxes, doc.Options.CurrencySymbol),
			fmt.Sprintf("%.2f%s", withholdingTaxes, doc.Options.CurrencySymbol),
			fmt.Sprintf("%.2f%s", doc.Taxes.FinalAmount+otherTaxes-withholdingTaxes, doc.Options.CurrencySymbol),
		}
	}

	doc.CreateLabeledLines(
		doc.Options.Layout.TotalsX, doc.Options.Layout.TotalsLabelWidth, doc.Options.Layout.TotalsValueWidth,
		doc.Options.Layout.TotalsLineHeight, doc.Options.Layout.TotalsFontLabelSize, doc.Options.Layout.TotalsFontValueSize,
		doc.Options.AccentColor, doc.Options.DarkTextColor, doc.Options.LightBgColor,
		[]string{
			doc.Options.TextTotalTotal,
			doc.Options.TextTotalDiscounted,
			doc.Options.TextTotalVatTax,
			doc.Options.TextTotalVariousTaxes,
			doc.Options.TextTotalWithHoldingTaxes,
			doc.Options.TextTotalWithTax,
		}, taxesValues, "R")
}

// CreateLabeledLines creates labeled lines with label: value
// The function uses the current y position of the document.
// The x position is specified by the x parameter.
// The wLabel is the width of the label, and wValue is the width of the value.
// The height is the height of the line, and the fontSizeLabel and fontSizeValue are the font sizes of the label and value respectively.
// The colorLabel is the color of the label, and the colorValue is the color of the value.
func (doc *Document) CreateLabeledLines(
	x float64, wLabel float64, wValue float64, height float64, fontSizeLabel float64, fontSizeValue float64,
	colorLabel []int, colorValue []int, bgColor []int, labels []string, values []string, valueAlign string) {

	for i, label := range labels {
		doc.pdf.SetX(x)
		doc.SetFillColor(bgColor)

		doc.SetTextColor(colorLabel)
		doc.pdf.SetFont(doc.Options.BoldFont, "B", fontSizeLabel)
		doc.pdf.CellFormat(wLabel, height, doc.EncodeString(label), "0", 0, "L", true, 0, "")

		doc.SetTextColor(colorValue)
		doc.pdf.SetFont(doc.Options.Font, "", fontSizeValue)
		doc.pdf.MultiCell(wValue, height, doc.EncodeString(values[i]), "0", valueAlign, true)
	}

	// restore values for next calls
	doc.pdf.SetX(x)
	doc.pdf.SetFont(doc.Options.Font, "", BaseTextFontSize)
	doc.SetTextColor(doc.Options.BaseTextColor)
}

// EncodeString encodes the string using doc.Options.UnicodeTranslateFunc
// This function can be used if we use a unicode translation function, which I could not make it work with greek encodings.
// So I just return the string as is.
func (doc *Document) EncodeString(str string) string {
	//return d.Options.UnicodeTranslateFunc(str)
	return str
}

// GetPDF returns the underlying fpdf object
func (doc *Document) GetPDF() *fpdf.Fpdf {
	return doc.pdf
}

// GetSeries returns the series of the document as a labeled string
func (doc *Document) GetSeries() string {
	return fmt.Sprintf("Σειρά: %s", doc.Series)
}

// GetAA returns the number of the document as a labeled string
func (doc *Document) GetAA() string {
	return fmt.Sprintf("Α.Α.: %s", doc.Number)
}

// GetDate returns the date of the document as a labeled string
func (doc *Document) GetDate() string {
	return fmt.Sprintf("Ημερομηνία: %s", doc.Date)
}

// GetMark returns the mark of the document as a labeled string
func (doc *Document) GetMark() string {
	return fmt.Sprintf("MARK: %s", doc.Mark)
}

// GetUID returns the uid of the document as a labeled string
func (doc *Document) GetUID() string {
	return fmt.Sprintf("UID: %s", doc.UID)
}

// SetIssuer sets the issuer's contact
func (doc *Document) SetIssuer(issuer *Contact) {
	doc.Issuer = issuer
}

// SetCustomer set the customer contact
func (doc *Document) SetCustomer(customer *Contact) {
	doc.CounterPart = customer
}

// SetDrawColor sets the color of the draw
func (doc *Document) SetDrawColor(color []int) {
	doc.pdf.SetDrawColor(color[0], color[1], color[2])
}

// SetFillColor sets the color of the fill
func (doc *Document) SetFillColor(color []int) {
	doc.pdf.SetFillColor(color[0], color[1], color[2])
}

// SetTextColor sets the color of the text
func (doc *Document) SetTextColor(color []int) {
	doc.pdf.SetTextColor(color[0], color[1], color[2])
}

// SetXY sets the current position
func (doc *Document) SetXY(x, y float64) {
	doc.pdf.SetXY(x, y)
}

// SetFont sets the font of the document
func (doc *Document) SetFont(font string, style string, size float64) {
	doc.pdf.SetFont(font, style, size)
}

// readImageFormat reads the format of the image
func readImageFormat(imageBytes []byte) string {
	ioReader := bytes.NewReader(imageBytes)
	_, format, _ := image.DecodeConfig(ioReader)
	return format
}
