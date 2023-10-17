package mydata

const (
	develHost = "mydataapidev.aade.gr"
	prodHost  = "mydatapi.aade.grrrrrr"

	URLSendInvoices               = "/SendInvoices"
	URLRequestMyIncome            = "/RequestMyIncome"
	URLRequestMyExpenses          = "/RequestMyExpenses"
	URLRequestDocs                = "/RequestDocs"
	URLRequestTransmittedDocs     = "/RequestTransmittedDocs"
	URLSendIncomeClassification   = "/SendIncomeClassification"
	URLSendExpensesClassification = "/SendExpensesClassification"
	URLCancelInvoice              = "/CancelInvoice"

	InternalErrorCode = -500
)

type OnInvalidAction int

const (
	ErrorOnInvalid OnInvalidAction = iota
	PassThroughOnInvalid
	FilterOnInvalid
)
