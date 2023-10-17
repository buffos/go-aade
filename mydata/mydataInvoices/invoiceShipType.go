package mydataInvoices

type ShipType struct {
	ApplicationID   string `xml:"applicationId"`   // Αριθμός Δήλωσης Διενέργειας Δραστηριότητας
	ApplicationDate string `xml:"applicationDate"` // Ημερομηνία Δήλωσης
	Doy             string `xml:"doy"`             // ΔΟΥ Δήλωσης
	ShipID          string `xml:"shipID"`          // Στοιχεία Πλοίου
}
