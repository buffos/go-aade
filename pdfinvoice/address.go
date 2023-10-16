package pdfinvoice

// Address is a struct that holds the address of a company or a person
type Address struct {
	Address    string `json:"address,omitempty" validate:"required"`
	Address2   string `json:"address_2,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
}

// String output address as string.
// Line break is added for new lines
func (a *Address) String() string {
	addrString := a.Address

	if len(a.Address2) > 0 {
		addrString += "\n"
		addrString += a.Address2
	}

	if len(a.PostalCode) > 0 {
		addrString += "\n"
		addrString += a.PostalCode
	} else {
		addrString += "\n"
	}

	if len(a.City) > 0 {
		addrString += " "
		addrString += a.City
	}

	if len(a.Country) > 0 {
		addrString += "\n"
		addrString += a.Country
	}

	return addrString
}

func (a *Address) ToStringSingleLine() string {
	addrString := a.Address

	if len(a.City) > 0 {
		addrString += " "
		addrString += a.City
	}

	if len(a.PostalCode) > 0 {
		addrString += ",ΤΚ "
		addrString += a.PostalCode
	}

	return addrString
}
