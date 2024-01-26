package mydataInvoices

type PaymentMethodType struct {
	InvoiceMark          *int64                    `xml:"invoiceMark"`
	PaymentMethodMark    *int64                    `xml:"paymentMethodMark"`    // Μοναδικός Αριθμός Καταχώρησης Τρόπου Πληρωμής
	EntityVatNumber      *string                   `xml:"entityVatNumber"`      // ΑΦΜ Οντότητας
	PaymentMethodDetails *PaymentMethodDetailsType `xml:"paymentMethodDetails"` // Στοιχεία Τρόπου Πληρωμής
}

type PaymentMethodsDoc struct {
	PaymentMethods []*PaymentMethodType `xml:"paymentMethods"`
}

// 1. Το πεδίο paymentMethodMark συμπληρώνεται από την υπηρεσία
// 2. Όταν η μέθοδος καλείται από τρίτο πρόσωπο (πχ πάροχος), ο ΑΦΜ αναφοράς αποστέλλεται μέσω του πεδίου entityVatNumber
// διαφορετικά το εν λόγω πεδίο παραμένει κενό.
// 3. Κατά τη χρήση της μεθόδου, τουλάχιστον ένα αντικείμενο PaymentMethodDetailType ανά παραστατικό πρέπει να είναι τύπου POS
// 4. Το σύνολο των ποσών amount ανά αντικείμενο PaymentMethodType πρέπει να ισούται με το totalValue του παραστατικού στο οποίο αντιστοιχεί το invoiceMark
