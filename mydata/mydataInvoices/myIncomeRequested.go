package mydataInvoices

import (
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type BookInfo struct {
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

func (b *BookInfo) Print() {

	if b.IssueDate != "" {
		fmt.Println("Ημερομηνία έκδοσης:", b.IssueDate)
	}
	if b.InvType != "" {
		fmt.Println("Τύπος Παραστατικού:", b.InvType)
	}
	if b.SelfPricing {
		fmt.Println("Αυτοτιμολόγηση:", b.SelfPricing)
	}
	if b.InvoiceDetailType != 0 {
		fmt.Println("Επισήμανση:", b.InvoiceDetailType.String())
	}
	if b.NetValue != 0 {
		fmt.Println("Καθαρή αξία:", b.NetValue)
	}
	if b.VatAmount != 0 {
		fmt.Println("Ποσό ΦΠΑ:", b.VatAmount)
	}
	if b.WithheldAmount != 0 {
		fmt.Println("Ποσό Παρακράτησης Φόρου:", b.WithheldAmount)
	}
	if b.OtherTaxesAmount != 0 {
		fmt.Println("Ποσό Λοιπών Φόρων:", b.OtherTaxesAmount)
	}
	if b.StampDutyAmount != 0 {
		fmt.Println("Ποσό Χαρτοσήμου:", b.StampDutyAmount)
	}
	if b.FeesAmount != 0 {
		fmt.Println("Ποσό Τελών:", b.FeesAmount)
	}
	if b.DeductionsAmount != 0 {
		fmt.Println("Ποσό Κρατήσεων:", b.DeductionsAmount)
	}
	if b.ThirdPartyAmount != 0 {
		fmt.Println("Ποσό Περί Τρίτων:", b.ThirdPartyAmount)
	}
	if b.GrossValue != 0 {
		fmt.Println("Συνολική Αξία:", b.GrossValue)
	}
	if b.Count != 0 {
		fmt.Println("Πλήθος Παραστατικών:", b.Count)
	}
	if b.MinMark != "" {
		fmt.Println("Ελάχιστο ΜΑΡΚ πλήθους:", b.MinMark)
	}
	if b.MaxMark != "" {
		fmt.Println("Μέγιστο ΜΑΡΚ πλήθους:", b.MaxMark)
	}
}

type RequestedBookInfo struct {
	Xmlns             string             `xml:"xmlns,attr"`
	XmlnsICLS         string             `xml:"xmlns:icls,attr"`
	XmlnsECLS         string             `xml:"xmlns:ecls,attr"`
	ContinuationToken *ContinuationToken `xml:"continuationToken"` // Στοιχείο για την τμηματική λήψη	αποτελεσμάτων
	CounterVatNumber  *string            `xml:"counterVatNumber"`  // ΑΦΜ αντισυμβαλλόμενου
	BooksInfo         []BookInfo         `xml:"bookInfo"`
}

func (b *RequestedBookInfo) Print() {
	fmt.Println("Καταχωρήσεις βιβλίου εσόδων - εξόδων:")
	if b.ContinuationToken != nil {
		b.ContinuationToken.Print()
	}
	if b.CounterVatNumber != nil {
		fmt.Println("ΑΦΜ αντισυμβαλλόμενου:", *b.CounterVatNumber)
	}

	for _, book := range b.BooksInfo {
		book.Print()
		fmt.Println()
	}
}
func (b *RequestedBookInfo) GetNextPartitionData() (string, string) {
	if b.ContinuationToken == nil {
		return "", ""
	}
	return b.ContinuationToken.NextPartitionKey, b.ContinuationToken.NextRowKey
}
