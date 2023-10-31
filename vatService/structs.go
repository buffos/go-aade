package vatService

import (
	"encoding/xml"
	"fmt"
)

// XMLResponse is where we parse an http response
type XMLResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    XMLBody  `xml:"Body"`
}

// XMLBody is the body of a response
type XMLBody struct {
	VATInfo VATInfo    `xml:"rgWsPublic2AfmMethodResponse>result>rg_ws_public2_result_rtType"`
	Version *string    `xml:"rgWsPublic2VersionInfoResponse>result"`
	Error   *ErrorInfo `xml:"Fault" json:"error,omitempty"`
}

func (b *XMLBody) error() error {

	if b.Error == nil {
		return nil
	}

	if b.Error.Code == "" {
		return nil
	}

	return Error(b.Error.Message)
}

// ErrorInfo holds error info
type ErrorInfo struct {
	Code    string `xml:"Code>Value" json:"code"`
	Message string `xml:"Reason>Text" json:"message"`
}

type VATInfo struct {
	CallSeqID  int            `xml:"call_seq_id" json:"call_seq_id"`
	CalledBy   VATCalledBy    `xml:"afm_called_by_rec" json:"called_by"`
	Result     VATResult      `xml:"basic_rec"  json:"result"`
	Activities []FirmActivity `xml:"firm_act_tab>item" json:"activities"`
	Error      *ErrorVATInfo  `xml:"error_rec" json:"error,omitempty"`
}

// error extracts the error transmitted from the service if present.
func (v *VATInfo) error() error {

	if v.Error == nil {
		return nil
	}

	if v.Error.Code == "" {
		return nil
	}

	vatError, ok := ServiceErrors[v.Error.Code]
	if !ok {
		return ErrGeneral
	}
	return vatError
}

// GetInvoiceData returns the data needed for an invoice
func (v *VATInfo) GetInvoiceData() (onomasia, doy, afm, address, zip, area string) {
	return v.Result.Onomasia, v.Result.DOYDescription, v.Result.AFM,
		fmt.Sprintf("%s %s", v.Result.PostalAddress, v.Result.PostalAddressNo),
		v.Result.PostalZipCode, v.Result.PostalAreaDescription
}

// GetMainActivity returns the main activity of the entity
func (v *VATInfo) GetMainActivity() string {
	for _, a := range v.Activities {
		if a.Kind == 1 {
			return a.Description
		}
	}
	return ""
}

// ErrorVATInfo holds error info
type ErrorVATInfo struct {
	Code    string `xml:"error_code" json:"code"`
	Message string `xml:"error_descr" json:"message"`
}

// VATCalledBy is the data relative who did the search.
type VATCalledBy struct {
	TokenUsername       string `xml:"token_username" json:"username"`
	TokenAFM            string `xml:"token_afm" json:"vat"`
	TokenAFMFullName    string `xml:"token_afm_fullname" json:"vat_fullname"`
	AFMCalledBy         string `xml:"afm_called_by" json:"called_by"`
	AFMCalledByFullName string `xml:"afm_called_by_fullname" json:"vat_called_by_fullname"`
	AsOnDate            string `xml:"as_on_date" json:"as_on_date"`
}

// VATResult is the data relative to an entity's VAT search
type VATResult struct {
	AFM                         string `xml:"afm" json:"afm"`                                              // ΑΦΜ
	DOY                         string `xml:"doy" json:"doy"`                                              // ΚΩΔΙΚΟΣ ΔΟΥ
	DOYDescription              string `xml:"doy_descr" json:"doy_description"`                            // ΠΕΡΙΓΡΑΦΗ ΔΟΥ
	InitialFlagDescription      string `xml:"i_ni_flag_descr" json:"initial_flag_description"`             // ΦΠ /ΜΗ ΦΠ
	DeactivationFlag            string `xml:"deactivation_flag" json:"deactivation_flag"`                  // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ:1=ΕΝΕΡΓΟΣ ΑΦΜ 2=ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	DeactivationFlagDescription string `xml:"deactivation_flag_desc" json:"deactivation_flag_description"` // ΕΝΔΕΙΞΗ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ(ΠΕΡΙΓΡΑΦΗ): ΕΝΕΡΓΟΣ ΑΦΜ ΑΠΕΝΕΡΓΟΠΟΙΗΜΕΝΟΣ ΑΦΜ
	FirmFlagDescription         string `xml:"firm_flag_descr" json:"firm_flag_description"`                // ΤΙΜΕΣ: ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΜΗ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ, ΠΡΩΗΝ ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ
	Onomasia                    string `xml:"onomasia" json:"onomasia"`                                    // ΕΠΩΝΥΜΙΑ
	CommercialTitle             string `xml:"commer_title" json:"commercial_title"`                        // ΤΙΤΛΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	LegalStatusDescription      string `xml:"legal_status_descr" json:"legal_status_descr"`                // ΠΕΡΙΓΡΑΦΗ ΜΟΡΦΗΣ ΜΗ Φ.Π.
	PostalAddress               string `xml:"postal_address" json:"postal_address"`                        // ΟΔΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAddressNo             string `xml:"postal_address_no" json:"postal_address_no"`                  // ΑΡΙΘΜΟΣ ΕΠΙΧΕΙΡΗΣΗΣ
	PostalZipCode               string `xml:"postal_zip_code" json:"postal_zip_code"`                      // ΤΑΧ. ΚΩΔ. ΕΠΙΧΕΙΡΗΣΗΣ
	PostalAreaDescription       string `xml:"postal_area_description" json:"postal_area_description"`      // ΠΕΡΙΟΧΗ ΕΠΙΧΕΙΡΗΣΗΣ
	RegistrationDate            string `xml:"regist_date" json:"registration_date"`                        // ΗΜ/ΝΙΑ ΕΝΑΡΞΗΣ
	StopDate                    string `xml:"stop_date" json:"stop_date"`                                  // ΗΜ/ΝΙΑ ΔΙΑΚΟΠΗΣ
	NormalVATSystemFlag         string `xml:"normal_vat_system_flag" json:"normal_vat_system_flag"`
}

type FirmActivity struct {
	Code            int    `xml:"firm_act_code" json:"code"`                   // ΚΩΔΙΚΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	Description     string `xml:"firm_act_descr" json:"description"`           // ΠΕΡΙΓΡΑΦΗ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ
	Kind            int    `xml:"firm_act_kind" json:"kind"`                   // ΕΙΔΟΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: 1=ΚΥΡΙΑ, 2=ΔΕΥΤΕΡΕΥΟΥΣΑ, 3=ΛΟΙΠΗ, 4=ΒΟΗΘΗΤΙΚΗ
	KindDescription string `xml:"firm_act_kind_descr" json:"kind_description"` // ΠΕΡΙΓΡΑΦΗ ΕΙΔΟΥΣ ΔΡΑΣΤΗΡΙΟΤΗΤΑΣ: ΚΥΡΙΑ, ΔΕΥΤΕΡΕΥΟΥΣΑ, ΛΟΙΠΗ, ΒΟΗΘΗΤΙΚΗ
}

const (
	// Endpoint is the url for WSDL service
	Endpoint = "https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2"
)

// String returns a string representation of a VATInfo
func (v *VATInfo) String() string {
	var s string

	s += fmt.Sprintf("afm:%s\n", v.Result.AFM)
	s += fmt.Sprintf("doy:%s\n", v.Result.DOY)
	s += fmt.Sprintf("doy_descr:%s\n", v.Result.DOYDescription)
	s += fmt.Sprintf("i_ni_flag_descr:%s\n", v.Result.InitialFlagDescription)
	s += fmt.Sprintf("deactivation_flag:%s\n", v.Result.DeactivationFlag)
	s += fmt.Sprintf("deactivation_flag_desc:%s\n", v.Result.DeactivationFlagDescription)
	s += fmt.Sprintf("firm_flag_descr:%s\n", v.Result.FirmFlagDescription)
	s += fmt.Sprintf("onomasia:%s\n", v.Result.Onomasia)
	s += fmt.Sprintf("commer_title:%s\n", v.Result.CommercialTitle)
	s += fmt.Sprintf("legal_status_descr:%s\n", v.Result.LegalStatusDescription)
	s += fmt.Sprintf("postal_address:%s\n", v.Result.PostalAddress)
	s += fmt.Sprintf("postal_address_no:%s\n", v.Result.PostalAddressNo)
	s += fmt.Sprintf("postal_zip_code:%s\n", v.Result.PostalZipCode)
	s += fmt.Sprintf("postal_area_description:%s\n", v.Result.PostalAreaDescription)
	s += fmt.Sprintf("regist_date:%s\n", v.Result.RegistrationDate)
	s += fmt.Sprintf("stop_date:%s\n", v.Result.StopDate)
	s += fmt.Sprintf("normal_vat_system_flag:%s\n", v.Result.NormalVATSystemFlag)

	s += fmt.Sprintf("ACTIVITIES:--------------------\n")
	for k, v := range v.Activities {
		s += fmt.Sprintf("ACTIVITY #%d\n", k)
		s += fmt.Sprintf("FirmActCode: %d\n", v.Code)
		s += fmt.Sprintf("FirmActDescr: %s\n", v.Description)
		s += fmt.Sprintf("FirmActKind: %d\n", v.Kind)
		s += fmt.Sprintf("FirmActKindDescr: %s\n", v.KindDescription)
	}

	//s += fmt.Sprintf("Error Description: %s\n", a.Error.Message)
	//s += fmt.Sprintf("ErrorCode: %s\n", a.Error.Code)
	s += fmt.Sprintf("Error: %v", v.Error)

	return s
}
