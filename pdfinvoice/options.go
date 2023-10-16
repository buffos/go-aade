package pdfinvoice

// UnicodeTranslateFunc ...
type UnicodeTranslateFunc func(string) string

// Layout of the document
type Layout struct {
	CustomerColumnTwoX      float64 `json:"customer_column_two_x,omitempty" validate:"required"`      // x position of customer column two
	CustomerLabelFontSize   float64 `json:"customer_label_font_size,omitempty" validate:"required"`   // font size of customer label
	CustomerValueFontSize   float64 `json:"customer_value_font_size,omitempty" validate:"required"`   // font size of customer value
	CustomerLabelWidth      float64 `json:"customer_label_width,omitempty" validate:"required"`       // width of customer label
	CustomerValueWidth      float64 `json:"customer_value_width,omitempty" validate:"required"`       // width of customer value
	CustomerLineHeight      float64 `json:"customer_line_height,omitempty" validate:"required"`       // height of customer line
	IssuerPX                float64 `json:"issuer_px,omitempty" validate:"required"`                  // x position of issuer contact
	IssuerPY                float64 `json:"issuer_py,omitempty" validate:"required"`                  // y position of issuer contact
	IssuerFontSize          float64 `json:"issuer_font_size,omitempty" validate:"required"`           // font size of issuer contact
	FooterX                 float64 `json:"footer_x,omitempty" validate:"required"`                   // x position of footer
	FooterY                 float64 `json:"footer_y,omitempty" validate:"required"`                   // y position of footer
	FooterFontSize          float64 `json:"footer_font_size,omitempty" validate:"required"`           // font size of footer
	InvoiceRowOffset        float64 `json:"invoice_row_offset,omitempty" validate:"required"`         // offset of invoice row
	InvoiceColumnGap        float64 `json:"invoice_column_gap,omitempty" validate:"required"`         // gap between invoice columns
	LogoOffsetX             float64 `json:"logo_offset,omitempty" validate:"required"`                // offsetX of logo
	LogoHeight              float64 `json:"logo_height,omitempty" validate:"required"`                // height of logo
	MaxAllowedDetailsHeight float64 `json:"max_allowed_details_height,omitempty" validate:"required"` // max allowed height of details
	NotesX                  float64 `json:"notes_x,omitempty" validate:"required"`                    // x position of notes
	NotesY                  float64 `json:"notes_y,omitempty" validate:"required"`                    // y position of notes
	NotesWidth              float64 `json:"notes_width,omitempty" validate:"required"`                // width of notes
	NotesHeight             float64 `json:"notes_height,omitempty" validate:"required"`               // height of notes
	QRCodeX                 float64 `json:"qrcode_x,omitempty" validate:"required"`                   // x position of qrcode
	QRCodeY                 float64 `json:"qrcode_y,omitempty" validate:"required"`                   // y position of qrcode
	QRCodeWidth             float64 `json:"qrcode_width,omitempty" validate:"required"`               // width of qrcode in mm
	QRCodeSize              float64 `json:"qrcode_size,omitempty" validate:"required"`                // size of qrcode in pixels
	TaxesX                  float64 `json:"taxes_x,omitempty" validate:"required"`                    // x position of taxes
	TaxesY                  float64 `json:"taxes_y,omitempty" validate:"required"`                    // y position of taxes
	TaxesRowHeight          float64 `json:"taxes_row_height,omitempty" validate:"required"`           // height of taxes row
	TaxesFontSize           float64 `json:"taxes_font_size,omitempty" validate:"required"`            // font size of taxes
	TotalsX                 float64 `json:"totals_x,omitempty" validate:"required"`                   // x position of totals
	TotalsY                 float64 `json:"totals_y,omitempty" validate:"required"`                   // y position of totals
	TotalsFontLabelSize     float64 `json:"totals_label_font_size,omitempty" validate:"required"`     // font size of totals
	TotalsFontValueSize     float64 `json:"totals_value_font_size,omitempty" validate:"required"`     // font size of totals
	TotalsLabelWidth        float64 `json:"totals_label_width,omitempty" validate:"required"`         // width of "totals" label
	TotalsValueWidth        float64 `json:"totals_value_width,omitempty" validate:"required"`         // width of "totals" value
	TotalsLineHeight        float64 `json:"totals_line_height,omitempty" validate:"required"`         // height of totals line
}

// Options for Document
type Options struct {
	Layout Layout `json:"layout,omitempty"` // layout of the document

	AutoPrint bool `json:"auto_print,omitempty"` // print without the dialogue box?

	Orientation   string `json:"orientation,omitempty"`    // orientation of the document
	DocumentUnits string `json:"document_units,omitempty"` // units of the document
	DocumentSize  string `json:"document_size,omitempty"`  // size of the document

	CurrencySymbol string `json:"currency_symbol,omitempty"`

	TextTypeInvoice      string `json:"text_type_invoice,omitempty"`
	TextTypeQuotation    string `json:"text_type_quotation,omitempty"`
	TextTypeDeliveryNote string `json:"text_type_delivery_note,omitempty"`

	TextRefTitle         string `json:"text_ref_title,omitempty"`
	TextVersionTitle     string `json:"text_version_title,omitempty"`
	TextDateTitle        string `json:"text_date_title,omitempty"`
	TextPaymentTermTitle string `json:"text_payment_term_title,omitempty"`

	TextItemsNameTitle       string `json:"text_items_name_title,omitempty"`
	TextItemsUnitCostTitle   string `json:"text_items_unit_cost_title,omitempty"`
	TextItemsMeasurementUnit string `json:"text_items_measurement_unit,omitempty"`
	TextItemsQuantityTitle   string `json:"text_items_quantity_title,omitempty"`
	TextItemsTotalHTTitle    string `json:"text_items_total_ht_title,omitempty"`
	TextItemsTaxTitle        string `json:"text_items_tax_title,omitempty"`
	TextItemsDiscountTitle   string `json:"text_items_discount_title,omitempty"`
	TextItemsTotalTTCTitle   string `json:"text_items_total_ttc_title,omitempty"`

	TextItemNotesTitle string `json:"text_item_notes_title,omitempty"`

	TextTotalTotal            string `json:"text_total_total,omitempty"`
	TextTotalDiscounted       string `json:"text_total_discounted,omitempty"`
	TextTotalVatTax           string `json:"text_total_vat_tax,omitempty"`
	TextTotalVariousTaxes     string `json:"text_total_various_taxes,omitempty"`
	TextTotalWithHoldingTaxes string `json:"text_total_with_holding_taxes,omitempty"`
	TextTotalWithTax          string `json:"text_total_with_tax,omitempty"`

	TextWithHoldingTaxes string `json:"text_with_holding_taxes,omitempty"`
	TextMiscTaxes        string `json:"text_misc_taxes,omitempty"`
	TextDeductions       string `json:"text_deductions,omitempty"`
	TextFeesTaxes        string `json:"text_fees_taxes,omitempty"`
	TextStampTaxes       string `json:"text_stamp_taxes,omitempty"`

	BaseTextColor  []int `json:"base_text_color,omitempty"`
	LightTextColor []int `json:"grey_text_color,omitempty"`
	LightBgColor   []int `json:"grey_bg_color,omitempty"`
	DarkBgColor    []int `json:"dark_bg_color,omitempty"`
	DarkTextColor  []int `json:"dark_text_color,omitempty"`
	DefaultBgColor []int `json:"default_bg_color,omitempty"`
	AccentColor    []int `json:"accent_color,omitempty"`

	FontPathEnvName  string
	Font             string
	FontFileName     string
	BoldFont         string
	BoldFontFileName string

	UnicodeTranslateFunc UnicodeTranslateFunc
}

var defaultLayout = Layout{
	IssuerPX:       A4PageWidth / 2,
	IssuerPY:       BaseMarginTop,
	IssuerFontSize: SmallTextFontSize,

	LogoOffsetX: 5,
	LogoHeight:  15,

	CustomerColumnTwoX:    -90,
	CustomerLabelFontSize: MediumTextFontSize,
	CustomerValueFontSize: MediumTextFontSize,
	CustomerLabelWidth:    35,
	CustomerValueWidth:    70,
	CustomerLineHeight:    LineHeight,

	FooterX:        BaseMargin,
	FooterY:        -BaseMargin,
	FooterFontSize: SmallTextFontSize,

	InvoiceRowOffset: 0.3,
	InvoiceColumnGap: 0.5,

	MaxAllowedDetailsHeight: 220,

	NotesX:      BaseMargin + 2,
	NotesY:      252,
	NotesWidth:  A4PageWidth/2 - 2*BaseMargin,
	NotesHeight: 30,

	QRCodeX:     -70,
	QRCodeY:     90,
	QRCodeWidth: 35,
	QRCodeSize:  256,

	TaxesX:         BaseMargin,
	TaxesY:         227,
	TaxesRowHeight: 10,
	TaxesFontSize:  SmallTextFontSize,

	TotalsX:             118,
	TotalsY:             250,
	TotalsLabelWidth:    46,
	TotalsValueWidth:    35,
	TotalsFontLabelSize: MediumTextFontSize,
	TotalsFontValueSize: MediumTextFontSize,
	TotalsLineHeight:    5,
}

var defaultOptions = Options{
	Layout: defaultLayout,

	AutoPrint: false,

	Orientation:   "P",
	DocumentUnits: "mm",
	DocumentSize:  "A4",

	CurrencySymbol: "€ ",

	TextTypeInvoice:      "Τιμολόγιο Πώλησης",
	TextTypeQuotation:    "Προσφορά",
	TextTypeDeliveryNote: "Δελτίο Αποστολής",

	TextRefTitle:         "Αρ. Παραστατικού",
	TextVersionTitle:     "Σειρά",
	TextDateTitle:        "Ημερομηνία",
	TextPaymentTermTitle: "Τρόπος Πληρωμής",

	TextItemsNameTitle:       "Προϊόν-Υπηρεσία",
	TextItemsQuantityTitle:   "Ποσότητα",
	TextItemsMeasurementUnit: "Μ.Μ",
	TextItemsUnitCostTitle:   "Αρχ. Τιμή (€)",
	TextItemsDiscountTitle:   "Έκπτωση (€)",
	TextItemsTaxTitle:        "ΦΠΑ%",
	TextItemsTotalHTTitle:    "Καθαρή Αξία (€)",
	TextItemsTotalTTCTitle:   "Τελική Αξία (€)",

	TextItemNotesTitle: "Παρατηρήσεις",

	TextTotalTotal:            "Συνολική Αξία",
	TextTotalDiscounted:       "Έκπτωση",
	TextTotalVatTax:           "ΦΠΑ",
	TextTotalVariousTaxes:     "Επιπλέον Φόροι (+)",
	TextTotalWithHoldingTaxes: "Παρακρατούμενοι Φόροι (-)",
	TextTotalWithTax:          "Πληρωτέο Ποσό",

	TextWithHoldingTaxes: "Παρακρατούμενοι (-)",
	TextDeductions:       "Κρατήσεις (-)",
	TextMiscTaxes:        "Λοιποί Φόροι (+)",
	TextFeesTaxes:        "Τέλη (+)",
	TextStampTaxes:       "Χαρτόσημο (+)",

	BaseTextColor:  []int{35, 35, 35},
	LightTextColor: []int{255, 255, 255},
	LightBgColor:   []int{232, 232, 232},
	DarkBgColor:    []int{5, 46, 102},
	DarkTextColor:  []int{35, 35, 35},
	DefaultBgColor: []int{255, 255, 255},
	AccentColor:    []int{0, 0, 153},

	FontPathEnvName:  "FONT_PATH",
	Font:             "Roboto",
	FontFileName:     "Roboto-Regular.ttf",
	BoldFont:         "Roboto",
	BoldFontFileName: "Roboto-Bold.ttf",
}
