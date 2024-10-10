package viesService

import (
	"encoding/xml"
	"fmt"
)

// ValidationVAT Response message for valid responses from VIES
type VATInfo struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	RequestDate string `xml:"requestDate"`
	Valid       bool   `xml:"valid"`
	Name        string `xml:"name"`
	Address     string `xml:"address"`
}

// ValidationErrorResponse message for invalid responses from VIES
type ErrorInfo struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

// body is the body of a response
type XMLBody struct {
	XMLName xml.Name
	VATInfo VATInfo   `xml:"checkVatResponse"`
	Error   ErrorInfo `xml:"Fault"`
}

func (b *XMLBody) error() error {
	if b.Error.FaultCode == "" {
		return nil
	}
	return fmt.Errorf("%s: %s", b.Error.FaultCode, b.Error.FaultString)
}

// XMLResponse is the response from the vies service
type XMLResponse struct {
	XMLName xml.Name
	Body    XMLBody
}
