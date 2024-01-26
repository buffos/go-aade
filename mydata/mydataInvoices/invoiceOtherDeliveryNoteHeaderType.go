package mydataInvoices

type OtherDeliveryNoteHeaderType struct {
	LoadingAddress         *Address `xml:"loadingAddress"`         // Διεύθυνση φόρτωσης συμπληρώνεται για παραστατικά που είναι δελτία αποστολής ή τιμολόγιο και δελτίο αποστολής
	DeliveryAddress        *Address `xml:"deliveryAddress"`        // Διεύθυνση παράδοσης συμπληρώνεται για παραστατικά που είναι δελτία αποστολής ή τιμολόγιο και δελτίο αποστολής
	StartShippingBranch    *int64   `xml:"startShippingBranch"`    // Αριθμός Υποκαταστήματος Έναρξης Μεταφοράς
	CompleteShippingBranch *int64   `xml:"completeShippingBranch"` // Αριθμός Υποκαταστήματος Ολοκλήρωσης Μεταφοράς
}
