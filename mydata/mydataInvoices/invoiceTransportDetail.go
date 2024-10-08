package mydataInvoices

import "fmt"

type TransportDetailType struct {
	VehicleNumber string `xml:"vehicleNumber"` // * Αριθμός οχήματος (max 50 chars)
}

func (t *TransportDetailType) Print() {
	fmt.Println("VehicleNumber:", t.VehicleNumber)
}
