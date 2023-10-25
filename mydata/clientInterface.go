package mydata

import "github.com/buffos/go-aade/mydata/mydataInvoices"

type ConnectionClient interface {
	SetOnInvalidAction(action OnInvalidAction)
	SendInvoices(invoices *mydataInvoices.InvoicesDoc) (int, *mydataInvoices.ResponseDoc, error)
	RequestDocs(params mydataInvoices.RequestDocsParams) (int, *mydataInvoices.RequestedDoc, error)
	RequestDocsPastDays(days int) (int, *mydataInvoices.RequestedDoc, error)
	RequestDocWithMark(mark uint) (int, *mydataInvoices.RequestedDoc, error)
	RequestMyIncome(params mydataInvoices.RequestMyIncomeParams) (int, *mydataInvoices.RequestedBookInfo, error)
	RequestMyExpenses(params mydataInvoices.RequestMyExpensesParams) (int, *mydataInvoices.RequestedBookInfo, error)
	RequestMyIncomeLastDays(days int) (int, *mydataInvoices.RequestedBookInfo, error)
	RequestMyExpensesLastDays(days int) (int, *mydataInvoices.RequestedBookInfo, error)
	RequestTransmittedDocs(params mydataInvoices.RequestDocsParams) (int, *mydataInvoices.RequestedDoc, error)
	RequestTransmittedDocsPastDays(days int) (int, *mydataInvoices.RequestedDoc, error)
	RequestTransmittedDocWithMark(mark uint) (int, *mydataInvoices.RequestedDoc, error)
	RequestTransmittedDocBetweenMarks(lowMark, higMark uint) (int, *mydataInvoices.RequestedDoc, error)
	SendIncomeClassification(changes *mydataInvoices.IncomeClassificationsDoc) (int, *mydataInvoices.ResponseDoc, error)
	SendExpensesClassification(changes *mydataInvoices.ExpensesClassificationsDoc, postMethodPerInvoice bool) (int, *mydataInvoices.ResponseDoc, error)
	CancelInvoice(params mydataInvoices.CancelInvoiceParams) (int, *mydataInvoices.ResponseDoc, error)
}
