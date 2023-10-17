package mydataInvoices

import "github.com/buffos/go-aade/mydata/mydatavalues"

type InvoiceSummaryType struct {
	TotalNetValue          float64                       `xml:"totalNetValue"`          // * Σύνολο Καθαρής Αξίας
	TotalVatAmount         float64                       `xml:"totalVatAmount"`         // * Σύνολο ΦΠΑ
	TotalWithheldAmount    float64                       `xml:"totalWithheldAmount"`    // * Σύνολο Παρακράτησης Φόρου
	TotalFeesAmount        float64                       `xml:"totalFeesAmount"`        // * Σύνολο Τελών
	TotalStampDutyAmount   float64                       `xml:"totalStampDutyAmount"`   // * Σύνολο Χαρτοσήμου
	TotalOtherTaxesAmount  float64                       `xml:"totalOtherTaxesAmount"`  // * Σύνολο Λοιπών Φόρων
	TotalDeductionsAmount  float64                       `xml:"totalDeductionsAmount"`  // * Σύνολο Κρατήσεων
	TotalGrossValue        float64                       `xml:"totalGrossValue"`        // * Σύνολο Αξίας
	IncomeClassification   []*IncomeClassificationType   `xml:"incomeClassification"`   // Χαρακτηρισμοί Εσόδων
	ExpensesClassification []*ExpensesClassificationType `xml:"expensesClassification"` // Χαρακτηρισμοί Εξόδων
}

// region InvoiceSummaryType functions

// calculate calculates the summary values from the rows
func (i *InvoiceSummaryType) calculate(rows []*InvoiceRowType) {

	/* region first set all values to zero */

	*i = InvoiceSummaryType{}

	// endregion

	// region adds the values from the rows to the summary
	for _, row := range rows {
		// if recType is VATEnd or MiscTaxesWithVAT, then netValue contains the VAT value.
		if (row.NetValue != nil && row.RecType != nil) && (*row.RecType == mydatavalues.VATEnd || *row.RecType == mydatavalues.MiscTaxesWithVAT) {
			i.TotalVatAmount += *row.NetValue
			// all other values should be zero for that type of rows
			continue
		}
		if (row.NetValue != nil && row.RecType != nil) && (*row.RecType != mydatavalues.VATEnd && *row.RecType != mydatavalues.MiscTaxesWithVAT) {
			i.TotalNetValue += *row.NetValue
		}
		if row.NetValue != nil && row.RecType == nil {
			i.TotalNetValue += *row.NetValue
		}
		if row.VatAmount != nil {
			i.TotalVatAmount += *row.VatAmount
		}
		if row.WithheldAmount != nil {
			i.TotalWithheldAmount += *row.WithheldAmount
		}
		if row.FeesAmount != nil {
			i.TotalFeesAmount += *row.FeesAmount
		}
		if row.StampDutyAmount != nil {
			i.TotalStampDutyAmount += *row.StampDutyAmount
		}
		if row.OtherTaxesAmount != nil {
			i.TotalOtherTaxesAmount += *row.OtherTaxesAmount
		}
		if row.DeductionsAmount != nil {
			i.TotalDeductionsAmount += *row.DeductionsAmount
		}
	}
	// endregion

	// region calculate gross value
	i.TotalGrossValue = i.TotalNetValue + i.TotalVatAmount - i.TotalWithheldAmount +
		i.TotalFeesAmount + i.TotalStampDutyAmount + i.TotalOtherTaxesAmount - i.TotalDeductionsAmount
	// endregion

	// region completes income's classification array

	// for every row's income classification array,
	// if the classification type is not present in the summary income classification, add it
	// else add the amount to the existing classification
	for _, row := range rows {
		if row.IncomeClassification != nil {
			for _, cl := range row.IncomeClassification {
				if i.IncomeClassification == nil {
					i.IncomeClassification = make([]*IncomeClassificationType, 0)
				}
				if !i.incomeClassificationExists(cl) {
					i.IncomeClassification = append(i.IncomeClassification, cl)
				} else {
					i.addIncomeClassificationAmount(cl)
				}
			}
		}

		if row.ExpensesClassification != nil {
			for _, cl := range row.ExpensesClassification {
				if i.ExpensesClassification == nil {
					i.ExpensesClassification = make([]*ExpensesClassificationType, 0)
				}
				if !i.expensesClassificationExists(cl) {
					i.ExpensesClassification = append(i.ExpensesClassification, cl)
				} else {
					i.addExpensesClassificationAmount(cl)
				}
			}
		}
	}

	// endregion
}

// incomeClassificationExists returns true if the classification type and category exist in the summary's income classification array
func (i *InvoiceSummaryType) incomeClassificationExists(cl *IncomeClassificationType) bool {
	for _, ic := range i.IncomeClassification {
		if ic.ClassificationType == cl.ClassificationType && ic.ClassificationCategory == cl.ClassificationCategory {
			return true
		}
	}
	return false
}

// addIncomeClassificationAmount adds the amount to the existing classification
func (i *InvoiceSummaryType) addIncomeClassificationAmount(cl *IncomeClassificationType) {
	for _, ic := range i.IncomeClassification {
		if ic.ClassificationType == cl.ClassificationType && ic.ClassificationCategory == cl.ClassificationCategory {
			ic.Amount += cl.Amount
		}
	}
}

// expensesClassificationExists returns true if the classification type and category exist in the summary's expenses classification array
func (i *InvoiceSummaryType) expensesClassificationExists(cl *ExpensesClassificationType) bool {
	for _, ic := range i.ExpensesClassification {
		if ic.ClassificationType == cl.ClassificationType && ic.ClassificationCategory == cl.ClassificationCategory {
			return true
		}
	}
	return false
}

// addExpensesClassificationAmount adds the amount to the existing classification
func (i *InvoiceSummaryType) addExpensesClassificationAmount(cl *ExpensesClassificationType) {
	for _, ic := range i.ExpensesClassification {
		if ic.ClassificationType == cl.ClassificationType && ic.ClassificationCategory == cl.ClassificationCategory {
			ic.Amount += cl.Amount
		}
	}
}

// endregion
