package mydataInvoices

type Address struct {
	Street     string `xml:"street"`
	Number     string `xml:"number"`
	PostalCode string `xml:"postalCode"`
	City       string `xml:"city"`
}
