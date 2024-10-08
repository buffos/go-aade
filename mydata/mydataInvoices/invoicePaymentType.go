package mydataInvoices

import "github.com/buffos/go-aade/mydata/mydatavalues"

type PaymentMethods struct {
	PaymentMethodDetails []PaymentMethodDetailsType `xml:"paymentMethodDetails"`
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

type ECRTokenType struct {
	SigningAuthor *string `xml:"SigningAuthor"` // ECR id: Αριθμός μητρώου του φορολογικού μηχανισμού
	Signature     *string `xml:"Signature"`     // Υπογραφή παρόχου (μέγιστο 20 χαρακτήρες).Λεπτομέρειες στην υπ’ αριθμ.Α.1155/09-10-2023 απόφαση ΦΕΚ 5992 Β΄/13.10.2023
}
