package mydataInvoices

import (
	"testing"

	"github.com/buffos/go-aade/mydata/mydatavalues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExpenseClassificationSuite struct {
	suite.Suite
}

func TestExpenseClassificationSuite(t *testing.T) {
	suite.Run(t, new(ExpenseClassificationSuite))
}

func (suite *ExpenseClassificationSuite) TestNewExpensesClassificationDoc() {
	t := suite.T()
	v := NewExpensesClassificationDoc()
	require.NotNil(t, v)
	require.Len(t, v.ExpensesInvoiceClassification, 0)
}

func (suite *ExpenseClassificationSuite) TestRejectClassification() {
	t := suite.T()
	v := NewExpensesClassificationDoc()
	v.RejectClassification(123456789, "")
	require.Len(t, v.ExpensesInvoiceClassification, 1)
	require.Equal(t, uint64(123456789), v.ExpensesInvoiceClassification[0].InvoiceMark)
	require.Equal(t, 1, *v.ExpensesInvoiceClassification[0].TransactionMode)

	v.RejectClassification(987654321, "123456789")
	require.Len(t, v.ExpensesInvoiceClassification, 2)
	require.Equal(t, uint64(987654321), v.ExpensesInvoiceClassification[1].InvoiceMark)
	require.Equal(t, 1, *v.ExpensesInvoiceClassification[0].TransactionMode)
}

func (suite *ExpenseClassificationSuite) TestDeviateClassification() {
	t := suite.T()
	v := NewExpensesClassificationDoc()
	v.DeviateClassification(123456789, "")
	require.Len(t, v.ExpensesInvoiceClassification, 1)
	require.Equal(t, uint64(123456789), v.ExpensesInvoiceClassification[0].InvoiceMark)
	require.Equal(t, 2, *v.ExpensesInvoiceClassification[0].TransactionMode)

	v.DeviateClassification(987654321, "123456789")
	require.Len(t, v.ExpensesInvoiceClassification, 2)
	require.Equal(t, uint64(987654321), v.ExpensesInvoiceClassification[1].InvoiceMark)
	require.Equal(t, 2, *v.ExpensesInvoiceClassification[0].TransactionMode)
}

func (suite *ExpenseClassificationSuite) TestEditLineNumberDetail() {
	t := suite.T()
	v := NewExpensesClassificationDoc()

	// nesting of arrays goes invoiceEntry->lineEntry->classificationEntry

	v.EditLineNumberDetail(123456789, "", 1,
		mydatavalues.E3_102_004, mydatavalues.ECategory2_2, 100)
	require.Len(t, v.ExpensesInvoiceClassification, 1)
	require.Equal(t, uint64(123456789), v.ExpensesInvoiceClassification[0].InvoiceMark)
	require.Nil(t, v.ExpensesInvoiceClassification[0].TransactionMode)
	require.Equal(t, 1, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails))          // only one line number
	require.Equal(t, 1, v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].LineNumber) // the first line number
	require.Equal(t, mydatavalues.E3_102_004, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].ClassificationType)
	require.Equal(t, mydatavalues.ECategory2_2, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].ClassificationCategory)
	require.Equal(t, 100.0, v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].Amount)

	// add another line number for the same invoice
	v.EditLineNumberDetail(123456789, "", 2,
		mydatavalues.E3_102_003, mydatavalues.ECategory2_1, 50)

	require.Len(t, v.ExpensesInvoiceClassification, 1)                                                          // still one invoice
	require.Nil(t, v.ExpensesInvoiceClassification[0].TransactionMode)                                          // still using the old transaction mode
	require.Equal(t, 2, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails))          // we now have two line numbers
	require.Equal(t, 2, v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[1].LineNumber) // the second line number
	// now verifies the details of that second line number
	require.Equal(t, mydatavalues.E3_102_003, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[1].ExpensesClassificationDetailData[0].ClassificationType)
	require.Equal(t, mydatavalues.ECategory2_1, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[1].ExpensesClassificationDetailData[0].ClassificationCategory)
	require.Equal(t, 50.0, v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[1].ExpensesClassificationDetailData[0].Amount)

	// add another characterization to the same invoice for LineNumber 1
	v.EditLineNumberDetail(123456789, "", 1,
		mydatavalues.E3_102_005, mydatavalues.ECategory2_5, 10)
	require.Len(t, v.ExpensesInvoiceClassification, 1)                                                                                                                             // still one invoice
	require.Nil(t, v.ExpensesInvoiceClassification[0].TransactionMode)                                                                                                             // still using the old transaction mode
	require.Equal(t, 2, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails))                                                                             // we have two line numbers
	require.Equal(t, 2, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData))                                         // select the first line and the details of that line.
	require.Equal(t, 1, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[1].ExpensesClassificationDetailData))                                         // the other lineNumber has still one entry.
	require.Equal(t, mydatavalues.E3_102_005, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[1].ClassificationType) // now access the second characterization
	require.Equal(t, mydatavalues.ECategory2_5, *v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[1].ClassificationCategory)
	require.Equal(t, 10.0, v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[1].Amount)

	// finally, add a new invoice with a new mark
	v.EditLineNumberDetail(987654321, "", 1,
		mydatavalues.E3_102_005, mydatavalues.ECategory2_5, 10)
	require.Len(t, v.ExpensesInvoiceClassification, 2)                                                                                     // we now have two invoices
	require.Nil(t, v.ExpensesInvoiceClassification[1].TransactionMode)                                                                     // the new invoice uses the same mode too
	require.Equal(t, 1, len(v.ExpensesInvoiceClassification[1].InvoicesExpensesClassificationDetails))                                     // the second invoice has one line number
	require.Equal(t, 2, len(v.ExpensesInvoiceClassification[0].InvoicesExpensesClassificationDetails))                                     // the first invoice has two line numbers
	require.Equal(t, 1, len(v.ExpensesInvoiceClassification[1].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData)) // the first line entry, of the second invoice has one characterization
	// check those characterizations
	require.Equal(t, mydatavalues.E3_102_005, *v.ExpensesInvoiceClassification[1].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].ClassificationType) // now access the second characterization
	require.Equal(t, mydatavalues.ECategory2_5, *v.ExpensesInvoiceClassification[1].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].ClassificationCategory)
	require.Equal(t, 10.0, v.ExpensesInvoiceClassification[1].InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData[0].Amount)
	if v != nil {
		v.Print()
	}
}

func (suite *ExpenseClassificationSuite) TestInvoiceExpensesClassificationNewWay() {
	t := suite.T()
	v := NewExpensesClassificationDoc()
	inv1 := v.NewInvoiceClassificationForMark(123456789, "")
	inv2 := v.NewInvoiceClassificationForMark(453455589, "")
	require.Len(t, v.ExpensesInvoiceClassification, 2)
	require.Equal(t, inv1.InvoiceMark, v.ExpensesInvoiceClassification[0].InvoiceMark)
	require.Equal(t, inv2.InvoiceMark, v.ExpensesInvoiceClassification[1].InvoiceMark)

	require.Equal(t, 0, len(inv1.InvoicesExpensesClassificationDetails))
	require.Equal(t, 0, len(inv1.InvoicesExpensesClassificationDetails))

	inv1.AddE3ClassificationDetail(mydatavalues.E3_102_004, mydatavalues.ECategory2_2, 100)
	inv1.AddE3ClassificationDetail(mydatavalues.E3_102_003, mydatavalues.ECategory2_1, 110)
	inv1.AddVatClassificationDetail(mydatavalues.VAT_361, mydatavalues.ECategory2_2, mydatavalues.InvoiceVAT24Percent, mydatavalues.VATExceptionReasonType(0), 120)
	inv1.AddVatClassificationDetail(mydatavalues.VAT_361, mydatavalues.ECategory2_2, mydatavalues.InvoiceVAT0Percent, mydatavalues.Article39a, 130)

	require.Equal(t, 4, len(inv1.InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData))
	require.Equal(t, 1, inv1.InvoicesExpensesClassificationDetails[0].LineNumber) // line number is zero for the invoice level

	inv1Entries := inv1.InvoicesExpensesClassificationDetails[0].ExpensesClassificationDetailData
	// classification type (E3)
	require.Equal(t, mydatavalues.E3_102_004, *inv1Entries[0].ClassificationType)
	require.Equal(t, mydatavalues.E3_102_003, *inv1Entries[1].ClassificationType)
	require.Equal(t, mydatavalues.VAT_361, *inv1Entries[2].ClassificationType)
	require.Equal(t, mydatavalues.VAT_361, *inv1Entries[3].ClassificationType)
	// vat amount testing
	require.Equal(t, (*float64)(nil), inv1Entries[0].VatAmount)
	require.Equal(t, (*float64)(nil), inv1Entries[1].VatAmount)
	require.Equal(t, 28.8, *inv1Entries[2].VatAmount)
	require.Equal(t, float64(0), *inv1Entries[3].VatAmount)

	// classification category
	require.Equal(t, mydatavalues.ECategory2_2, *inv1Entries[0].ClassificationCategory)
	require.Equal(t, mydatavalues.ECategory2_1, *inv1Entries[1].ClassificationCategory)
	require.Equal(t, mydatavalues.ECategory2_2, *inv1Entries[2].ClassificationCategory)
	require.Equal(t, mydatavalues.ECategory2_2, *inv1Entries[3].ClassificationCategory)

	// amount
	require.Equal(t, 100.0, inv1Entries[0].Amount)
	require.Equal(t, 110.0, inv1Entries[1].Amount)
	require.Equal(t, 120.0, inv1Entries[2].Amount)
	require.Equal(t, 130.0, inv1Entries[3].Amount)

	// vat amount
	require.Equal(t, (*float64)(nil), inv1Entries[0].VatAmount)
	require.Equal(t, (*float64)(nil), inv1Entries[1].VatAmount)
	require.Equal(t, 28.8, *inv1Entries[2].VatAmount)
	require.Equal(t, 0.0, *inv1Entries[3].VatAmount)

	// vat category
	require.Equal(t, (*mydatavalues.InvoiceVATCategory)(nil), inv1Entries[0].VatCategory)
	require.Equal(t, (*mydatavalues.InvoiceVATCategory)(nil), inv1Entries[1].VatCategory)
	require.Equal(t, mydatavalues.InvoiceVAT24Percent, *inv1Entries[2].VatCategory)
	require.Equal(t, mydatavalues.InvoiceVAT0Percent, *inv1Entries[3].VatCategory)

	// vat exception reason
	require.Equal(t, (*mydatavalues.VATExceptionReasonType)(nil), inv1Entries[0].VatExemptionCategory)
	require.Equal(t, (*mydatavalues.VATExceptionReasonType)(nil), inv1Entries[1].VatExemptionCategory)
	require.Equal(t, (*mydatavalues.VATExceptionReasonType)(nil), inv1Entries[2].VatExemptionCategory)
	require.Equal(t, mydatavalues.Article39a, *inv1Entries[3].VatExemptionCategory)

	if inv1 != nil {
		inv1.Print()
	}
}

func (suite *ExpenseClassificationSuite) TestValidateAgainstInvoice() {
	//TODO: add tests
}
