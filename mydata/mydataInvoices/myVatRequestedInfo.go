package mydataInvoices

import (
	"encoding/xml"
	"fmt"
	"time"
)

type RequestedVatInfoType struct {
	Xmlns             string                  `xml:"xmlns,attr"`
	ContinuationToken *ContinuationToken      `xml:"continuationToken"`
	VatInfo           []*InvoiceVatDetailType `xml:"VatInfo"`
}

func (r *RequestedVatInfoType) Print() {
	if r.ContinuationToken != nil {
		r.ContinuationToken.Print()
	}

	for _, vat := range r.VatInfo {
		vat.Print()
	}
}

func (r *RequestedVatInfoType) GetNextPartitionData() (string, string) {
	if r.ContinuationToken == nil {
		return "", ""
	}
	return r.ContinuationToken.NextPartitionKey, r.ContinuationToken.NextRowKey
}

type InvoiceVatDetailType struct {
	Mark               *string    `xml:"Mark"`
	IsCancelled        *bool      `xml:"IsCancelled"`
	IssueDate          *time.Time `xml:"IssueDate"`
	Vat301             *float64   `xml:"Vat301"`
	Vat302             *float64   `xml:"Vat302"`
	Vat303             *float64   `xml:"Vat303"`
	Vat304             *float64   `xml:"Vat304"`
	Vat305             *float64   `xml:"Vat305"`
	Vat306             *float64   `xml:"Vat306"`
	Vat331             *float64   `xml:"Vat331"`
	Vat332             *float64   `xml:"Vat332"`
	Vat333             *float64   `xml:"Vat333"`
	Vat334             *float64   `xml:"Vat334"`
	Vat335             *float64   `xml:"Vat335"`
	Vat336             *float64   `xml:"Vat336"`
	Vat361             *float64   `xml:"Vat361"`
	Vat362             *float64   `xml:"Vat362"`
	Vat363             *float64   `xml:"Vat363"`
	Vat364             *float64   `xml:"Vat364"`
	Vat365             *float64   `xml:"Vat365"`
	Vat366             *float64   `xml:"Vat366"`
	Vat381             *float64   `xml:"Vat381"`
	Vat382             *float64   `xml:"Vat382"`
	Vat383             *float64   `xml:"Vat383"`
	Vat384             *float64   `xml:"Vat384"`
	Vat385             *float64   `xml:"Vat385"`
	Vat386             *float64   `xml:"Vat386"`
	Vat342             *float64   `xml:"Vat342"`
	Vat345             *float64   `xml:"Vat345"`
	Vat348             *float64   `xml:"Vat348"`
	Vat349             *float64   `xml:"Vat349"`
	Vat310             *float64   `xml:"Vat310"`
	Vat402             *float64   `xml:"Vat402"`
	Vat407             *float64   `xml:"Vat407"`
	Vat411             *float64   `xml:"Vat411"`
	Vat423             *float64   `xml:"Vat423"`
	Vat422             *float64   `xml:"Vat422"`
	VatUnclassified361 *float64   `xml:"VatUnclassified361"`
	VatUnclassified381 *float64   `xml:"VatUnclassified381"`
}

func (i *InvoiceVatDetailType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias InvoiceVatDetailType
	aux := &struct {
		IssueDate string `xml:"IssueDate"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	// Parse the IssueDate from string to time.Time
	parsedDate, err := time.Parse("2006-01-02T15:04:05", aux.IssueDate)
	if err != nil {
		return err
	}
	i.IssueDate = &parsedDate
	return nil
}

func (i *InvoiceVatDetailType) VatKeysWithValues() map[string]string {
	taxes := make(map[string]string)
	if i.Vat301 != nil {
		taxes["Vat301"] = fmt.Sprintf("%f", *i.Vat301)
	}
	if i.Vat302 != nil {
		taxes["Vat302"] = fmt.Sprintf("%f", *i.Vat302)
	}
	if i.Vat303 != nil {
		taxes["Vat303"] = fmt.Sprintf("%f", *i.Vat303)
	}
	if i.Vat304 != nil {
		taxes["Vat304"] = fmt.Sprintf("%f", *i.Vat304)
	}
	if i.Vat305 != nil {
		taxes["Vat305"] = fmt.Sprintf("%f", *i.Vat305)
	}
	if i.Vat306 != nil {
		taxes["Vat306"] = fmt.Sprintf("%f", *i.Vat306)
	}
	if i.Vat331 != nil {
		taxes["Vat331"] = fmt.Sprintf("%f", *i.Vat331)
	}
	if i.Vat332 != nil {
		taxes["Vat332"] = fmt.Sprintf("%f", *i.Vat332)
	}
	if i.Vat333 != nil {
		taxes["Vat333"] = fmt.Sprintf("%f", *i.Vat333)
	}
	if i.Vat334 != nil {
		taxes["Vat334"] = fmt.Sprintf("%f", *i.Vat334)
	}
	if i.Vat335 != nil {
		taxes["Vat335"] = fmt.Sprintf("%f", *i.Vat335)
	}
	if i.Vat336 != nil {
		taxes["Vat336"] = fmt.Sprintf("%f", *i.Vat336)
	}
	if i.Vat361 != nil {
		taxes["Vat361"] = fmt.Sprintf("%f", *i.Vat361)
	}
	if i.Vat362 != nil {
		taxes["Vat362"] = fmt.Sprintf("%f", *i.Vat362)
	}
	if i.Vat363 != nil {
		taxes["Vat363"] = fmt.Sprintf("%f", *i.Vat363)
	}
	if i.Vat364 != nil {
		taxes["Vat364"] = fmt.Sprintf("%f", *i.Vat364)
	}
	if i.Vat365 != nil {
		taxes["Vat365"] = fmt.Sprintf("%f", *i.Vat365)
	}
	if i.Vat366 != nil {
		taxes["Vat366"] = fmt.Sprintf("%f", *i.Vat366)
	}
	if i.Vat381 != nil {
		taxes["Vat381"] = fmt.Sprintf("%f", *i.Vat381)
	}
	if i.Vat382 != nil {
		taxes["Vat382"] = fmt.Sprintf("%f", *i.Vat382)
	}
	if i.Vat383 != nil {
		taxes["Vat383"] = fmt.Sprintf("%f", *i.Vat383)
	}
	if i.Vat384 != nil {
		taxes["Vat384"] = fmt.Sprintf("%f", *i.Vat384)
	}
	if i.Vat385 != nil {
		taxes["Vat385"] = fmt.Sprintf("%f", *i.Vat385)
	}
	if i.Vat386 != nil {
		taxes["Vat386"] = fmt.Sprintf("%f", *i.Vat386)
	}
	if i.Vat342 != nil {
		taxes["Vat342"] = fmt.Sprintf("%f", *i.Vat342)
	}
	if i.Vat345 != nil {
		taxes["Vat345"] = fmt.Sprintf("%f", *i.Vat345)
	}
	if i.Vat348 != nil {
		taxes["Vat348"] = fmt.Sprintf("%f", *i.Vat348)
	}
	if i.Vat349 != nil {
		taxes["Vat349"] = fmt.Sprintf("%f", *i.Vat349)
	}
	if i.Vat310 != nil {
		taxes["Vat310"] = fmt.Sprintf("%f", *i.Vat310)
	}
	if i.Vat402 != nil {
		taxes["Vat402"] = fmt.Sprintf("%f", *i.Vat402)
	}
	if i.Vat407 != nil {
		taxes["Vat407"] = fmt.Sprintf("%f", *i.Vat407)
	}
	if i.Vat411 != nil {
		taxes["Vat411"] = fmt.Sprintf("%f", *i.Vat411)
	}
	if i.Vat423 != nil {
		taxes["Vat423"] = fmt.Sprintf("%f", *i.Vat423)
	}
	if i.Vat422 != nil {
		taxes["Vat422"] = fmt.Sprintf("%f", *i.Vat422)
	}
	if i.VatUnclassified361 != nil {
		taxes["VatUnclassified361"] = fmt.Sprintf("%f", *i.VatUnclassified361)
	}
	if i.VatUnclassified381 != nil {
		taxes["VatUnclassified381"] = fmt.Sprintf("%f", *i.VatUnclassified381)
	}
	return taxes
}

// PrintInfo prints the vat info to the console
func (i *InvoiceVatDetailType) Print() {
	if i.Mark != nil {
		fmt.Printf("Mark: %s\n", *i.Mark)
	}
	if i.IsCancelled != nil {
		fmt.Printf("IsCancelled: %t\n", *i.IsCancelled)
	}
	if i.IssueDate != nil {
		fmt.Printf("IssueDate: %s\n", i.IssueDate.Format("2006-01-02T15:04:05"))
	}
	fmt.Printf("Vat keys with values: %v\n", i.VatKeysWithValues())
}
