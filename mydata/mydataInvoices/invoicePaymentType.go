package mydataInvoices

import (
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type PaymentMethods struct {
	PaymentMethodDetails []PaymentMethodDetailsType `xml:"paymentMethodDetails"`
}

func (p *PaymentMethods) Print() {
	for _, detail := range p.PaymentMethodDetails {
		detail.Print()
	}
}

type PaymentMethodDetailsType struct {
	Type               *mydatavalues.InvoicePaymentType `xml:"type"`               // Ελάχιστη τιμή = 1 Μέγιστη τιμή = 5
	Amount             *float64                         `xml:"amount"`             // Ελάχιστη τιμή 0. Δεκαδικά ψηφία 2
	PaymentMethodInfo  string                           `xml:"paymentMethodInfo"`  // Πληροφορίες (πχ Αριθμός Τραπέζης)
	TipAmount          *float64                         `xml:"tipAmount"`          // Ποσό φιλοδωρήματος.Ελάχιστη τιμή 0.Δεκαδικά ψηφία 2
	Tid                *string                          `xml:"tid"`                // Αναγνωριστικό πληρωμής.Διαβιβάζεται στην περίπτωση πληρωμών με type = 7 (POS)
	TransactionId      *string                          `xml:"transactionId"`      // Μοναδική ταυτότητα πληρωμής.Διαβιβάζεται στην περίπτωση πληρωμών με type = 7 (POS)
	ProvidersSignature *ProviderSignatureType           `xml:"ProvidersSignature"` // Υπογραφή πληρωμής παρόχου.Διαβιβάζεται στην περίπτωση πληρωμών με type = 7 (POS).Μέσω παρόχου
	ECRToken           *ECRTokenType                    `xml:"ECRToken"`           // Κρυπτογραφημένος κωδικός πληρωμής. Για πληρωμές με type = 7 (POS).Μέσω ERP.
}

type ProviderSignatureType struct {
	SigningAuthor *string `xml:"SigningAuthor"` // Αριθμός Απόφασης έγκρισης ΥΠΑΗΕΣ	Παρόχου (μέγιστο 20 χαρακτήρες)
	Signature     *string `xml:"Signature"`     // Υπογραφή παρόχου (μέγιστο 20 χαρακτήρες).Λεπτομέρειες στην υπ’ αριθμ.Α.1155/09-10-2023 απόφαση ΦΕΚ 5992 Β΄/13.10.2023
}

func (p *ProviderSignatureType) Print() {
	fmt.Println("Αριθμός Απόφασης έγκρισης ΥΠΑΗΕΣ Παρόχου:", *p.SigningAuthor)
	fmt.Println("Υπογραφή παρόχου:", *p.Signature)
}

type ECRTokenType struct {
	SigningAuthor *string `xml:"SigningAuthor"` // ECR id: Αριθμός μητρώου του φορολογικού μηχανισμού
	Signature     *string `xml:"Signature"`     // Υπογραφή παρόχου (μέγιστο 20 χαρακτήρες).Λεπτομέρειες στην υπ’ αριθμ.Α.1155/09-10-2023 απόφαση ΦΕΚ 5992 Β΄/13.10.2023
}

func (p *ECRTokenType) Print() {
	fmt.Println("ECR id:", *p.SigningAuthor)
	fmt.Println("Υπογραφή παρόχου:", *p.Signature)
}

func (p *PaymentMethodDetailsType) Print() {
	if p.Type != nil {
		fmt.Println("Τύπος Πληρωμής:", p.Type.String())
	}
	if p.Amount != nil {
		fmt.Println("Ποσό Πληρωμής:", *p.Amount)
	}
	fmt.Println("Πληροφορίες Πληρωμής:", p.PaymentMethodInfo)
	if p.TipAmount != nil {
		fmt.Println("Ποσό Φιλοδωρήματος:", *p.TipAmount)
	}
	if p.Tid != nil {
		fmt.Println("Αναγνωριστικό Πληρωμής:", *p.Tid)
	}
	if p.TransactionId != nil {
		fmt.Println("Μοναδική Ταυτότητα Πληρωμής:", *p.TransactionId)
	}
	if p.ProvidersSignature != nil {
		p.ProvidersSignature.Print()
	}
	if p.ECRToken != nil {
		p.ECRToken.Print()
	}
}
