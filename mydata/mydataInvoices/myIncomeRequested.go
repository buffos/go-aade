package mydataInvoices

import "github.com/buffos/go-aade/mydata/mydatavalues"

type BookInfo struct {
	ContinuationToken struct {
		NextPartitionKey string `xml:"nextPartitionKey"`
		NextRowKey       string `xml:"nextRowKey"`
	} `xml:"continuationToken"` // Στοιχείο για την τμηματική λήψη	αποτελεσμάτων
	IssueDate         string                         `xml:"issueDate"`         // Ημερομηνία έκδοσης
	InvType           mydatavalues.InvoiceType       `xml:"invType"`           // Τύπος Παραστατικού
	SelfPricing       bool                           `xml:"selfPricing"`       // Αυτοτιμολόγηση
	InvoiceDetailType mydatavalues.InvoiceDetailType `xml:"invoiceDetailType"` // Επισήμανση
	NetValue          float64                        `xml:"netValue"`          // Καθαρή αξία
	VatAmount         float64                        `xml:"vatAmount"`         // Ποσό ΦΠΑ
	WithheldAmount    float64                        `xml:"withheldAmount"`    // Ποσό Παρακράτησης Φόρου
	OtherTaxesAmount  float64                        `xml:"otherTaxesAmount"`  // Ποσό Λοιπών Φόρων
	StampDutyAmount   float64                        `xml:"stampDutyAmount"`   // Ποσό Χαρτοσήμου
	FeesAmount        float64                        `xml:"feesAmount"`        // Ποσό Τελών
	DeductionsAmount  float64                        `xml:"deductionsAmount"`  // Ποσό Κρατήσεων
	ThirdPartyAmount  float64                        `xml:"thirdPartyAmount"`  // Ποσό Περί Τρίτων
	GrossValue        float64                        `xml:"grossValue"`        // Συνολική Αξία
	Count             int                            `xml:"count"`             // * Πλήθος Παραστατικών
	MinMark           string                         `xml:"minMark"`           // Ελάχιστο ΜΑΡΚ πλήθους
	MaxMark           string                         `xml:"maxMark"`           // Μέγιστο ΜΑΡΚ πλήθους
}

type RequestedBookInfo struct {
	Xmlns     string     `xml:"xmlns,attr"`
	XmlnsICLS string     `xml:"xmlns:icls,attr"`
	XmlnsECLS string     `xml:"xmlns:ecls,attr"`
	BooksInfo []BookInfo `xml:"bookInfo"`
}
