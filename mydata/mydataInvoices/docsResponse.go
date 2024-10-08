package mydataInvoices

import (
	"errors"
	"fmt"
)

type ResponseDoc struct {
	Response []*Response `xml:"response"`
}

func (rDoc ResponseDoc) Print() {
	for _, resp := range rDoc.Response {
		resp.Print()
	}
}

func (rDoc ResponseDoc) String() string {
	str := ""
	for i, resp := range rDoc.Response {
		str += fmt.Sprintf("Response %d:\n", i)
		str += fmt.Sprintf("%s\n", resp)
	}
	return str
}

type Response struct {
	StatusCode         string `xml:"statusCode"`
	Errors             Errors `xml:"errors"`
	Index              uint   `xml:"index"`
	InvoiceUID         string `xml:"invoiceUid"`
	InvoiceMark        uint64 `xml:"invoiceMark"`
	QrCodeUrl          string `xml:"qrUrl"`
	ClassificationMark uint64 `xml:"classificationMark"`
	AuthenticationCode string `xml:"authenticationCode"`
	CancellationMark   uint64 `xml:"cancellationMark"`
}

func (r Response) String() string {
	return fmt.Sprintf(
		"StatusCode:%s\n"+
			"Index:%d\n"+
			"InvoiceUID:%s\n"+
			"InvoiceMark:%d\n"+
			"QrCodeUrl:%s\n"+
			"ClassificationMark:%d\n"+
			"AuthenticationCode:%s\n"+
			"CancellationMark:%d",
		r.StatusCode, r.Index, r.InvoiceUID, r.InvoiceMark, r.QrCodeUrl, r.ClassificationMark, r.AuthenticationCode, r.CancellationMark)
}

func (r Response) Print() {
	fmt.Printf(r.String())
}

type Error struct {
	Message string `xml:"message"`
	Code    string `xml:"code"`
}

// Error implements the error interface and returns a human-readable representation of the given error.
func (e Error) Error() string {
	return fmt.Sprintf("Code:%s - Message:%s", e.Code, e.Message)
}

type Errors struct {
	Error []Error `xml:"error"`
}

// HasErrors returns true if the response contains at least one errors, else false.
func (rDoc ResponseDoc) HasErrors() bool {
	for _, resp := range rDoc.Response {
		if len(resp.Errors.Error) > 0 {
			return true
		}
	}

	return false
}

// Errors returns all errors of the response doc wrapped with a given prefix.
// If no errors are found, then it returns nil.
func (rDoc ResponseDoc) Errors(prefix string) error {
	if !rDoc.HasErrors() {
		return nil
	}
	errs := errors.New("response error")
	for _, resp := range rDoc.Response {
		for _, respErr := range resp.Errors.Error {
			respErr.Message = fmt.Sprintf("%s: %s", prefix, respErr.Message)
			errs = errors.Join(errs, respErr)
		}
	}
	return errs
}
