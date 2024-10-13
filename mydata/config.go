package mydata

const (
	developmentHost = "mydataapidev.aade.gr"
	productionHost  = "mydatapi.aade.gr/myDATA"

	URLSendInvoices               = "/SendInvoices"
	URLRequestMyIncome            = "/RequestMyIncome"
	URLRequestMyExpenses          = "/RequestMyExpenses"
	URLRequestDocs                = "/RequestDocs"
	URLRequestTransmittedDocs     = "/RequestTransmittedDocs"
	URLRequestVatInfo             = "/RequestVatInfo"
	URLSendIncomeClassification   = "/SendIncomeClassification"
	URLSendExpensesClassification = "/SendExpensesClassification"
	URLCancelInvoice              = "/CancelInvoice"
	URLSendPaymentsMethod         = "/SendPaymentsMethod"

	InternalErrorCode = -500
)

type OnInvalidAction int

const (
	ErrorOnInvalid OnInvalidAction = iota
	PassThroughOnInvalid
	FilterOnInvalid
)
