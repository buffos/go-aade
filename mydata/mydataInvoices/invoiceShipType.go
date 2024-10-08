package mydataInvoices

import "fmt"

type ShipType struct {
	ApplicationID   string `xml:"applicationId"`   // Αριθμός Δήλωσης Διενέργειας Δραστηριότητας
	ApplicationDate string `xml:"applicationDate"` // Ημερομηνία Δήλωσης
	Doy             string `xml:"doy"`             // ΔΟΥ Δήλωσης
	ShipID          string `xml:"shipID"`          // Στοιχεία Πλοίου
}

func (s *ShipType) Print() {
	fmt.Println("Στοιχεία Πλοίου:")
	fmt.Println("Αριθμός Δήλωσης Διενέργειας Δραστηριότητας:", s.ApplicationID)
	fmt.Println("Ημερομηνία Δήλωσης:", s.ApplicationDate)
	fmt.Println("ΔΟΥ Δήλωσης:", s.Doy)
	fmt.Println("Στοιχεία Πλοίου:", s.ShipID)
}
