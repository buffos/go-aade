package mydata

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrorRequestCreation    = Error("error creating request")
	ErrorGettingResponse    = Error("error getting response")
	ErrorXMLMarshal         = Error("error marshaling xml")
	ErrorXMLParsingResponse = Error("error parsing xml response")
	ErrorQueryURLCreation   = Error("error creating query url")

	ErrorInvalidInvoices                 = Error("invalid invoices")
	ErrorRequestSendIncomeClassification = Error("error sending income classification")
)
