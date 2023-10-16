package pdfinvoice

import (
	"bytes"
	"github.com/go-pdf/fpdf"
)

// Create creates a new pdf document using the provided header, middle and footer functions that define the document's content.
func (doc *Document) Create(creator InvoiceCreator) (*fpdf.Fpdf, error) {
	documentEnd := false

	// region sets auto header and footer
	doc.pdf.SetFont(doc.Options.Font, "", BaseTextFontSize) // Load font
	doc.SetTextColor(doc.Options.BaseTextColor)
	doc.pdf.SetAutoPageBreak(false, 0)
	doc.pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
	doc.pdf.SetHeaderFuncMode(func() { creator.InvoiceHeader(doc) }, false)
	doc.pdf.SetFooterFunc(func() { creator.InvoiceFooter(doc, documentEnd) })
	// endregion

	doc.pdf.AddPage() // Add first page
	creator.InvoiceMiddle(doc)
	documentEnd = true

	// region append js to auto-print if AutoPrint == true
	if doc.Options.AutoPrint {
		doc.pdf.SetJavascript("print(true);")
	}
	// endregion
	return doc.pdf, nil
}

// CreateAndSave creates a new pdf document using the provided header, middle and footer functions that define the document's content and saves it to the specified filename.
func (doc *Document) CreateAndSave(filename string, creator InvoiceCreator) error {
	pdf, err := doc.Create(creator)
	if err != nil {
		return err
	}

	return pdf.OutputFileAndClose(filename)
}

// CreateAndBuffer creates a new pdf document using the provided header, middle and footer functions that define the document's content and returns a buffer with the pdf content.
func (doc *Document) CreateAndBuffer(creator InvoiceCreator) (*bytes.Buffer, error) {
	pdf, err := doc.Create(creator)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
