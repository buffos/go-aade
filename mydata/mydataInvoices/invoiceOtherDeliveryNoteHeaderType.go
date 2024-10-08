package mydataInvoices

import "fmt"

type OtherDeliveryNoteHeaderType struct {
	LoadingAddress         *Address `xml:"loadingAddress"`         // Διεύθυνση φόρτωσης συμπληρώνεται για παραστατικά που είναι δελτία αποστολής ή τιμολόγιο και δελτίο αποστολής
	DeliveryAddress        *Address `xml:"deliveryAddress"`        // Διεύθυνση παράδοσης συμπληρώνεται για παραστατικά που είναι δελτία αποστολής ή τιμολόγιο και δελτίο αποστολής
	StartShippingBranch    *int64   `xml:"startShippingBranch"`    // Αριθμός Υποκαταστήματος Έναρξης Μεταφοράς
	CompleteShippingBranch *int64   `xml:"completeShippingBranch"` // Αριθμός Υποκαταστήματος Ολοκλήρωσης Μεταφοράς
}

func (o *OtherDeliveryNoteHeaderType) Print() {
	if o.LoadingAddress != nil {
		o.LoadingAddress.Print()
	}
	if o.DeliveryAddress != nil {
		o.DeliveryAddress.Print()
	}
	if o.StartShippingBranch != nil {
		fmt.Println("Αριθμός Υποκαταστήματος Έναρξης Μεταφοράς:", *o.StartShippingBranch)
	}
	if o.CompleteShippingBranch != nil {
		fmt.Println("Αριθμός Υποκαταστήματος Ολοκλήρωσης Μεταφοράς:", *o.CompleteShippingBranch)
	}
}
