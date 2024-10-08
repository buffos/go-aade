package mydataInvoices

import (
	"errors"
	"fmt"
	"time"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type InvoiceRowType struct {
	LineNumber                   *uint                                    `xml:"lineNumber"`                   // * ΑΑ γραμμής
	RecType                      *mydatavalues.InvoiceLineType            `xml:"recType"`                      // Είδος Γραμμής
	FuelCode                     *mydatavalues.FuelCode                   `xml:"fuelCode"`                     // Κωδικός Καυσίμου
	Quantity                     *float64                                 `xml:"quantity"`                     // Ποσότητα
	MeasurementUnit              *mydatavalues.InvoiceMeasurementUnit     `xml:"measurementUnit"`              // Είδος Ποσότητας
	InvoiceDetailType            *mydatavalues.InvoiceDetailType          `xml:"invoiceDetailType"`            // Επισήμανση
	NetValue                     *float64                                 `xml:"netValue"`                     // * Καθαρή αξία
	VatCategory                  *mydatavalues.InvoiceVATCategory         `xml:"vatCategory"`                  // * Κατηγορία ΦΠΑ
	VatAmount                    *float64                                 `xml:"vatAmount"`                    // * Ποσό ΦΠΑ
	VatExemptionCategory         *mydatavalues.VATExceptionReasonType     `xml:"vatExemptionCategory"`         // Κατηγορία Αιτίας Απαλλαγής ΦΠΑ
	Dienergia                    *ShipType                                `xml:"dienergia"`                    // ΠΟΛ 1177/2018	Αρ. 27
	DiscountOption               *bool                                    `xml:"discountOption"`               // Δικαίωμα έκπτωσης
	WithheldAmount               *float64                                 `xml:"withheldAmount"`               // Ποσό παρακράτησης φόρου
	WithheldPercentCategory      *mydatavalues.WithholdingTaxCategoryType `xml:"withheldPercentCategory"`      // Κατηγορία ποσοστού παρακράτησης φόρου
	StampDutyAmount              *float64                                 `xml:"stampDutyAmount"`              // Ποσό Χαρτοσήμου
	StampDutyPercentCategory     *mydatavalues.PaperStampCategoryType     `xml:"stampDutyPercentCategory"`     // Κατηγορία ποσοστού χαρτοσήμου
	FeesAmount                   *float64                                 `xml:"feesAmount"`                   // Ποσό Τελών
	FeesPercentCategory          *mydatavalues.FeeCategoriesType          `xml:"feesPercentCategory"`          // Κατηγορία ποσοστού τελών
	OtherTaxesPercentCategory    *mydatavalues.MiscTaxCategoryType        `xml:"otherTaxesPercentCategory"`    // Κατηγορία ποσοστού λοιπών φόρων
	OtherTaxesAmount             *float64                                 `xml:"otherTaxesAmount"`             // Ποσό Λοιπών Φόρων
	DeductionsAmount             *float64                                 `xml:"deductionsAmount"`             // Ποσό Κρατήσεων
	LineComments                 *string                                  `xml:"lineComments"`                 // Σχόλια γραμμής
	IncomeClassification         []*IncomeClassificationType              `xml:"incomeClassification"`         // Χαρακτηρισμοί Εσόδων
	ExpensesClassification       []*ExpensesClassificationType            `xml:"expensesClassification"`       // Χαρακτηρισμοί Εξόδων
	Quantity15                   *float64                                 `xml:"quantity15"`                   // Ποσότητα Θερμοκρασίας 15 βαθμών (για παραστατικά καυσίμων από παρόχους)
	ItemDescr                    *string                                  `xml:"itemDescr"`                    // Περιγραφή Είδους (max 300 chars).Μόνο για την ειδική κατηγορία tax free
	TaricNo                      *string                                  `xml:"TaricNo"`                      // Κωδικός ΤΑΡΙΚ.
	ItemCode                     *string                                  `xml:"itemCode"`                     // Κωδικός Είδους.
	OtherMeasurementUnitQuantity *int                                     `xml:"otherMeasurementUnitQuantity"` // Ποσότητα Εναλλακτικής Μονάδας Μέτρησης.Μόνο όταν measurementUnit = 7
	OtherMeasurementUnitTitle    *string                                  `xml:"otherMeasurementUnitTitle"`    // Τίτλος Εναλλακτικής Μονάδας Μέτρησης.Μόνο όταν measurementUnit = 7
	NotVat195                    *bool                                    `xml:"notVat195"`                    // Ένδειξη μη συμμετοχής στο ΦΠΑ (έσοδα - έξοδα). Αποδεκτό μόνο για παραστατικά τύπων μεταξύ 1.1 - 11.5
}

func (i *InvoiceRowType) Print() {
	if i.LineNumber != nil {
		fmt.Println("ΑΑ γραμμής:", *i.LineNumber)
	}
	if i.RecType != nil {
		fmt.Println("Είδος Γραμμής:", i.RecType.String())
	}
	if i.FuelCode != nil {
		fmt.Println("Κωδικός Καυσίμου:", i.FuelCode.String())
	}
	if i.Quantity != nil {
		fmt.Println("Ποσότητα:", *i.Quantity)
	}
	if i.MeasurementUnit != nil {
		fmt.Println("Είδος Ποσότητας:", i.MeasurementUnit.String())
	}
	if i.InvoiceDetailType != nil {
		fmt.Println("Επισήμανση:", i.InvoiceDetailType.String())
	}
	if i.NetValue != nil {
		fmt.Println("Καθαρή Αξία:", *i.NetValue)
	}
	if i.VatCategory != nil {
		fmt.Println("Κατηγορία ΦΠΑ:", i.VatCategory.String())
	}
	if i.VatAmount != nil {
		fmt.Println("Ποσό ΦΠΑ:", *i.VatAmount)
	}
	if i.VatExemptionCategory != nil {
		fmt.Println("Κατηγορία Αιτίας Απαλλαγής ΦΠΑ:", i.VatExemptionCategory.String())
	}
	if i.Dienergia != nil {
		fmt.Println("Διενέργεια:", i.Dienergia)
	}
	if i.DiscountOption != nil {
		fmt.Println("Δικαίωμα έκπτωσης:", *i.DiscountOption)
	}
	if i.WithheldAmount != nil {
		fmt.Println("Ποσό παρακράτησης φόρου:", *i.WithheldAmount)
	}
	if i.WithheldPercentCategory != nil {
		fmt.Println("Κατηγορία ποσοστού παρακράτησης φόρου:", i.WithheldPercentCategory.String())
	}
	if i.StampDutyAmount != nil {
		fmt.Println("Ποσό Χαρτοσήμου:", *i.StampDutyAmount)
	}
	if i.StampDutyPercentCategory != nil {
		fmt.Println("Κατηγορία ποσοστού χαρτοσήμου:", i.StampDutyPercentCategory.String())
	}
	if i.FeesAmount != nil {
		fmt.Println("Ποσό Τελών:", *i.FeesAmount)
	}
	if i.FeesPercentCategory != nil {
		fmt.Println("Κατηγορία ποσοστού τελών:", i.FeesPercentCategory.String())
	}
	if i.OtherTaxesAmount != nil {
		fmt.Println("Ποσό Λοιπών Φόρων:", *i.OtherTaxesAmount)
	}
	if i.OtherTaxesPercentCategory != nil {
		fmt.Println("Κατηγορία ποσοστού λοιπών φόρων:", i.OtherTaxesPercentCategory.String())
	}
	if i.DeductionsAmount != nil {
		fmt.Println("Ποσό Κρατήσεων:", *i.DeductionsAmount)
	}
	if i.LineComments != nil {
		fmt.Println("Σχόλια γραμμής:", *i.LineComments)
	}
	for _, income := range i.IncomeClassification {
		income.Print()
	}
	for _, expense := range i.ExpensesClassification {
		expense.Print()
	}
	if i.Quantity15 != nil {
		fmt.Println("Ποσότητα Θερμοκρασίας 15 βαθμών:", *i.Quantity15)
	}
	if i.ItemDescr != nil {
		fmt.Println("Περιγραφή Είδους:", *i.ItemDescr)
	}
	if i.TaricNo != nil {
		fmt.Println("Κωδικός ΤΑΡΙΚ:", *i.TaricNo)
	}
	if i.ItemCode != nil {
		fmt.Println("Κωδικός Είδους:", *i.ItemCode)
	}
	if i.OtherMeasurementUnitQuantity != nil {
		fmt.Println("Ποσότητα Εναλλακτικής Μονάδας Μέτρησης:", *i.OtherMeasurementUnitQuantity)
	}
	if i.OtherMeasurementUnitTitle != nil {
		fmt.Println("Τίτλος Εναλλακτικής Μονάδας Μέτρησης:", *i.OtherMeasurementUnitTitle)
	}
	if i.NotVat195 != nil {
		fmt.Println("Ένδειξη μη συμμετοχής στο ΦΠΑ:", *i.NotVat195)
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewInvoiceRow(netValue float64, vatCat mydatavalues.InvoiceVATCategory) *InvoiceRowType {
	var vatPercent float64
	switch vatCat {
	case mydatavalues.InvoiceVAT24Percent:
		vatPercent = 0.24
	case mydatavalues.InvoiceVAT13Percent:
		vatPercent = 0.13
	case mydatavalues.InvoiceVAT6Percent:
		vatPercent = 0.06
	case mydatavalues.InvoiceVAT17Percent:
		vatPercent = 0.17
	case mydatavalues.InvoiceVAT9Percent:
		vatPercent = 0.09
	case mydatavalues.InvoiceVAT4Percent:
		vatPercent = 0.04
	case mydatavalues.InvoiceVAT0Percent:
		vatPercent = 0
	case mydatavalues.InvoiceVATExempt:
		vatPercent = 0
	case mydatavalues.InvoiceVAT3PercentArticle31:
		vatPercent = 0.03
	case mydatavalues.InvoiceVAT4PercentArticle31:
		vatPercent = 0.04
	}
	vatAmount := roundToMoney(netValue * vatPercent)
	i := &InvoiceRowType{}
	i.LineNumber = nil
	i.NetValue = &netValue
	i.VatCategory = &vatCat
	i.VatAmount = &vatAmount
	return i
}

// region InvoiceRowType

// Validate validates the invoice row
func (i *InvoiceRowType) Validate() error {
	if *i.LineNumber == 0 {
		return errors.New("invoice row line number is required")
	}
	if *i.NetValue == 0 {
		return errors.New("invoice row net value is required")
	}
	if *i.VatCategory == 0 {
		return errors.New("invoice row vat category is required")
	}
	if *i.VatAmount == 0 && !i.IsVATExempt() {
		return errors.New("invoice row vat amount is required")
	}

	return nil
}

func (i *InvoiceRowType) IsVATExempt() bool {
	if i.VatExemptionCategory != nil && *i.VatExemptionCategory != 0 {
		return true
	}
	if *i.VatCategory == mydatavalues.InvoiceVATExempt || *i.VatCategory == mydatavalues.InvoiceVAT0Percent {
		return true
	}
	return false
}

// SetRecType sets the invoice row's recType. What type of invoice we have
func (i *InvoiceRowType) SetRecType(recType mydatavalues.InvoiceLineType) *InvoiceRowType {
	i.RecType = &recType
	return i
}

// SetFuelCode sets the invoice row's fuel code
func (i *InvoiceRowType) SetFuelCode(code mydatavalues.FuelCode) *InvoiceRowType {
	i.FuelCode = &code
	return i
}

// SetQuantity sets the invoice row's quantity.
func (i *InvoiceRowType) SetQuantity(quantity float64) *InvoiceRowType {
	i.Quantity = &quantity
	return i
}

// SetMeasurementUnit sets the invoice row's measurement unit.
func (i *InvoiceRowType) SetMeasurementUnit(unit mydatavalues.InvoiceMeasurementUnit) *InvoiceRowType {
	i.MeasurementUnit = &unit
	return i
}

// SetInvoiceDetailType sets the invoice row's detail type.
func (i *InvoiceRowType) SetInvoiceDetailType(detailType mydatavalues.InvoiceDetailType) *InvoiceRowType {
	i.InvoiceDetailType = &detailType
	return i
}

// SetVatExemptionCategory sets the invoice row's vat exemption category.
func (i *InvoiceRowType) SetVatExemptionCategory(category mydatavalues.VATExceptionReasonType) *InvoiceRowType {
	i.VatExemptionCategory = &category
	return i
}

// SetDienergia sets the invoice row's dienergia.
func (i *InvoiceRowType) SetDienergia(appID string, appDate time.Time, doy string, shipID string) *InvoiceRowType {
	i.Dienergia = &ShipType{
		ApplicationID:   appID,
		ApplicationDate: appDate.Format(time.DateOnly),
		Doy:             doy,
		ShipID:          shipID,
	}
	return i
}

// SetDiscountOption sets the invoice row's discount option.
func (i *InvoiceRowType) SetDiscountOption(discount bool) *InvoiceRowType {
	i.DiscountOption = &discount
	return i
}

// SetWithheldAmount sets the invoice row's withheld amount and category.
func (i *InvoiceRowType) SetWithheldAmount(amount float64, category mydatavalues.WithholdingTaxCategoryType) *InvoiceRowType {
	i.WithheldAmount = &amount
	i.WithheldPercentCategory = &category
	return i
}

// SetStampAmount sets the invoice row's stamp amount and category.
func (i *InvoiceRowType) SetStampAmount(amount float64, category mydatavalues.PaperStampCategoryType) *InvoiceRowType {
	i.StampDutyAmount = &amount
	i.StampDutyPercentCategory = &category
	return i
}

// SetFeesAmount sets the invoice row's fee amount and category.
func (i *InvoiceRowType) SetFeesAmount(amount float64, category mydatavalues.FeeCategoriesType) *InvoiceRowType {
	i.FeesAmount = &amount
	i.FeesPercentCategory = &category
	return i
}

// SetOtherTaxesAmount sets the invoice row's other taxes amount and category.
func (i *InvoiceRowType) SetOtherTaxesAmount(amount float64, category mydatavalues.MiscTaxCategoryType) *InvoiceRowType {
	i.OtherTaxesAmount = &amount
	i.OtherTaxesPercentCategory = &category
	return i
}

// SetDeductionsAmount sets the invoice row's.
func (i *InvoiceRowType) SetDeductionsAmount(amount float64) *InvoiceRowType {
	i.DeductionsAmount = &amount
	return i
}

// SetComment sets the invoice row's comment.
func (i *InvoiceRowType) SetComment(comment string) *InvoiceRowType {
	i.LineComments = &comment
	return i
}

// SetItemDescription sets the invoice row's item description.
func (i *InvoiceRowType) SetItemDescription(description string) *InvoiceRowType {
	i.ItemDescr = &description
	return i
}

// AddIncomeClassification adds an income classification to the invoice row's income classification array.
func (i *InvoiceRowType) AddIncomeClassification(clType mydatavalues.IncomeClassificationTypeStringType, clCategory mydatavalues.IncomeClassificationCategoryStringType, amount float64) *InvoiceRowType {
	if i.IncomeClassification == nil {
		i.IncomeClassification = make([]*IncomeClassificationType, 0)
	}
	i.IncomeClassification = append(i.IncomeClassification, NewIncomeClassification(clType, clCategory, amount, byte(len(i.IncomeClassification)+1)))
	return i
}

// AddE3ExpenseClassification adds an expense classification to the invoice row's expense classification array. This does not classify VAT values
func (i *InvoiceRowType) AddE3ExpenseClassification(clType mydatavalues.ExpenseClassificationTypeStringType, clCategory mydatavalues.ExpensesClassificationCategoryStringType, amount float64) *InvoiceRowType {
	if i.ExpensesClassification == nil {
		i.ExpensesClassification = make([]*ExpensesClassificationType, 0)
	}
	i.ExpensesClassification = append(i.ExpensesClassification, NewExpenseClassification(clType, clCategory, 0, 0, amount, byte(len(i.ExpensesClassification)+1)))

	return i
}

// AddVATExpenseClassification adds an expense classification to the invoice row's expense classification array. This does not classify VAT values
func (i *InvoiceRowType) AddVATExpenseClassification(
	clType mydatavalues.ExpenseClassificationTypeStringType,
	clCategory mydatavalues.ExpensesClassificationCategoryStringType,
	vatCategory mydatavalues.InvoiceVATCategory,
	vatExemptionCategory mydatavalues.VATExceptionReasonType,
	amount float64) *InvoiceRowType {
	if i.ExpensesClassification == nil {
		i.ExpensesClassification = make([]*ExpensesClassificationType, 0)
	}
	i.ExpensesClassification = append(i.ExpensesClassification,
		NewExpenseClassification(clType, clCategory, vatCategory, vatExemptionCategory, amount, byte(len(i.ExpensesClassification)+1)),
	)
	return i
}

// endregion
