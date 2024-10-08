package mydataInvoices

import (
	"fmt"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type EntityType struct {
	Type       *mydatavalues.EntityCategory `xml:"type"` // Κατηγορία Οντότητας
	entityData *PartyType                   // Στοιχεία Οντότητας
}

func (e *EntityType) Print() {
	if e.Type != nil {
		fmt.Println("Κατηγορία Οντότητας:", e.Type.String())
	}
	if e.entityData != nil {
		e.entityData.Print()
	}
}

type PartyType struct {
	VatNumber         *string  `xml:"vatNumber"`         // * ΑΦΜ
	Country           *string  `xml:"country"`           // * Κωδικός Χώρας (Ελλάδα = GR)
	Branch            *uint64  `xml:"branch"`            // * Αριθμός Υποκαταστήματος. Ελάχιστη τιμή = 0 που είναι η έδρα του εκδότη.
	Name              *string  `xml:"name"`              // Επωνυμία. Τα στοιχεία Επωνυμία και Διεύθυνση δε γίνονται αποδεκτά στην περίπτωση που αφορούν οντότητα εντός Ελλάδας
	Address           *Address `xml:"address"`           // Διεύθυνση
	DocumentIdNo      *string  `xml:"documentIdNo"`      // Αριθμός επίσημου εγγράφου (max 100 chars). Έγκυρο μόνο στην περίπτωση που ο τύπος εγγράφου είναι ο 4
	SupplyAccountNo   *string  `xml:"supplyAccountNo"`   // Αριθμός Παροχής Ηλεκτρικού Ρεύματος (max 100 chars). Έγκυρα μόνο στην περίπτωση παραστατικών καυσίμων
	CountryDocumentId *string  `xml:"countryDocumentId"` // Κωδικός χώρας έκδοσης επίσημου εγγράφου
}

func (p *PartyType) Print() {
	if p.VatNumber != nil {
		fmt.Println("ΑΦΜ:", *p.VatNumber)
	}
	if p.Country != nil {
		fmt.Println("Χώρα:", *p.Country)
	}
	if p.Branch != nil {
		fmt.Println("Αριθμός Υποκαταστήματος:", *p.Branch)
	}
	if p.SupplyAccountNo != nil {
		fmt.Println("Αριθμός Παροχής Ηλεκτρικού Ρεύματος:", *p.SupplyAccountNo)
	}
	if p.CountryDocumentId != nil {
		fmt.Println("Χώρα έκδοσης επίσημου εγγράφου:", *p.CountryDocumentId)
	}
	if p.Name != nil {
		fmt.Println("Επωνυμία:", *p.Name)
	}
	if p.Address != nil {
		p.Address.Print()
	}
	if p.DocumentIdNo != nil {
		fmt.Println("Αριθμός επίσημου εγγράφου:", *p.DocumentIdNo)
	}
}
