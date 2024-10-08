package mydataInvoices

import "fmt"

type Address struct {
	Street     string `xml:"street"`
	Number     string `xml:"number"`
	PostalCode string `xml:"postalCode"`
	City       string `xml:"city"`
}

func (a *Address) Print() {
	fmt.Println("Διεύθυνση:", a.Street, a.Number, a.PostalCode, a.City)
}
