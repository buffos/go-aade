package pdfinvoice

import (
	"fmt"
	"os"
)

// Contact represents contact information.
type Contact struct {
	Name            string   `json:"name" validate:"required,min=1,max=256"`
	AFM             string   `json:"afm" validate:"required,min=9,max=9"`
	WorkDescription string   `json:"work_description" validate:"required,min=1,max=256"`
	Address         *Address `json:"address" validate:"required,min=1,max=256"`
	DOY             string   `json:"doy" validate:"required,min=1,max=256"`
	AdditionalInfo  []string `json:"additional_info,omitempty"`
}

// AppendLogoToDocument appends the logo to the document at position x, y.
// The logo is offset by logoOffsetX, and the height is also defined in options by Layout.LogoHeight.
func (doc *Document) AppendLogoToDocument(x, y float64) {
	// region 1.handle logo
	doc.pdf.SetXY(x, y)
	if doc.LogoFileName != "" {
		logoBytes, err := os.ReadFile(doc.LogoFileName)
		if err != nil {
			return // we do not add the image
		}
		doc.AppendImage(
			x+doc.Options.Layout.LogoOffsetX, y,
			0, doc.Options.Layout.LogoHeight,
			logoBytes, doc.LogoFileName)
	}
	// endregion
}

// AppendIssuerToDoc appends the issuer's contact to the document on the specified position.
// the logo, if provided, is added to the left of the contact.
func (doc *Document) AppendIssuerToDoc(x, y float64) float64 {
	doc.AppendLogoToDocument(x, y)
	doc.pdf.SetX(doc.Options.Layout.IssuerPX)
	doc.pdf.SetFontSize(doc.Options.Layout.IssuerFontSize)
	line := fmt.Sprintf("%s\n  %s\n %s\n Α.Φ.Μ.: %s  Δ.Ο.Υ.: %s",
		doc.Issuer.Name, doc.Issuer.WorkDescription, doc.Issuer.Address.ToStringSingleLine(), doc.Issuer.AFM, doc.Issuer.DOY)
	doc.pdf.MultiCell(100, 5, doc.EncodeString(line), "0", "C", false)

	if doc.Issuer.AdditionalInfo != nil {
		doc.pdf.SetFontSize(SmallTextFontSize)
		doc.pdf.SetXY(doc.Options.Layout.IssuerPX, doc.pdf.GetY())

		for _, line := range doc.Issuer.AdditionalInfo {
			doc.pdf.SetXY(doc.Options.Layout.IssuerPX, doc.pdf.GetY())
			doc.pdf.MultiCell(100, 5, doc.EncodeString(line), "0", "C", false)
		}
		// return to the left of the page.
		doc.pdf.SetXY(x, doc.pdf.GetY())
		doc.pdf.SetFontSize(BaseTextFontSize)
	}
	return y
}

// AppendCounterPartToDoc appends the counterpart's contact to the document on the top left of the page.
func (doc *Document) AppendCounterPartToDoc(x, y float64) float64 {
	doc.pdf.SetXY(x, y)
	fontsizeLabel := doc.Options.Layout.CustomerLabelFontSize
	fontsizeValue := doc.Options.Layout.CustomerValueFontSize

	doc.CreateLabeledLines(
		x, doc.Options.Layout.CustomerLabelWidth, doc.Options.Layout.CustomerValueWidth,
		doc.Options.Layout.CustomerLineHeight, fontsizeLabel, fontsizeValue,
		doc.Options.DarkTextColor, doc.Options.BaseTextColor, doc.Options.DefaultBgColor,
		[]string{"Επωνυμία:", "Επάγγελμα", "Α.Φ.Μ.:", "Δ.Ο.Υ.:", "Διεύθυνση:"},
		[]string{doc.CounterPart.Name, doc.CounterPart.WorkDescription, doc.CounterPart.AFM, doc.CounterPart.DOY, doc.CounterPart.Address.String()}, "L")

	return doc.pdf.GetY()
}
