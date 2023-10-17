package mydataInvoices

import (
	"encoding/xml"
	"errors"
	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type IncomeClassificationType struct {
	ClassificationType     *mydatavalues.IncomeClassificationTypeStringType     `xml:"classificationType"`     // Κωδικός Χαρακτηρισμού
	ClassificationCategory *mydatavalues.IncomeClassificationCategoryStringType `xml:"classificationCategory"` // Κατηγορία Χαρακτηρισμού
	Amount                 float64                                              `xml:"amount"`                 // * Ποσό
	ID                     *byte                                                `xml:"id"`                     // Αύξων αριθμός χαρακτηρισμού
}

//goland:noinspection GoUnusedExportedFunction
func NewIncomeClassification(clType mydatavalues.IncomeClassificationTypeStringType, clCategory mydatavalues.IncomeClassificationCategoryStringType, amount float64, id byte) *IncomeClassificationType {
	var clPtr *mydatavalues.IncomeClassificationTypeStringType
	var catPtr *mydatavalues.IncomeClassificationCategoryStringType
	if clType != "" {
		clPtr = &clType
	}
	if clCategory != "" {
		catPtr = &clCategory
	}

	return &IncomeClassificationType{
		ClassificationType:     clPtr,
		ClassificationCategory: catPtr,
		Amount:                 amount,
		ID:                     &id,
	}
}

// MarshalXML marshals the invoice's IncomeClassificationType to xml
func (cl IncomeClassificationType) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type incomeClassificationDoc struct {
		ClassificationType     *string `xml:"icls:classificationType"`
		ClassificationCategory *string `xml:"icls:classificationCategory"`
		Amount                 float64 `xml:"icls:amount"`
		ID                     *byte   `xml:"icls:id"`
	}
	err := enc.EncodeElement(incomeClassificationDoc{
		ClassificationType:     (*string)(cl.ClassificationType),
		ClassificationCategory: (*string)(cl.ClassificationCategory),
		Amount:                 cl.Amount,
		ID:                     cl.ID,
	}, start)

	if err != nil {
		return errors.New("error marshaling income classification")
	}

	return err
}
