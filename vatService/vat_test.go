package vatService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestVersion(t *testing.T) {
	c := NewClient("", "")
	version, err := c.Version()
	require.NoError(t, err)
	require.NotEmpty(t, version)
}

func TestInvalids(t *testing.T) {

	// some invalid input to test returned service errors
	inputs := []map[string]string{
		{
			"vat":      "1234567890",
			"username": "someUser",
			"password": "somePass",
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		{
			"vat":      "104807035",
			"username": "someUser",
			"password": "somePass",
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		{
			"vat":      "1234567890",
			"username": "promedialab", // valid user but,
			"password": "123456",      // wrong pass
			"error":    "RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED",
		},
		//{
		//	"vat":      "1234567890",
		//	"username": "complete to run", // put valid credentials here
		//	"password": "123456", // for this test to pass
		//	"error":    "RG_WS_PUBLIC_WRONG_AFM",
		//},
	}
	for k, v := range inputs {
		t.Logf("testing case #%d, vat:%s, user:%s, pass:%s", k, v["vat"], v["username"], v["password"])
		c := NewClient(v["username"], v["password"])
		i, err := c.GetVATInfo("", v["vat"])
		require.Error(t, err, "error getting VAT info", err.Error())
		require.Nil(t, i, "VAT info should be nil")
	}

}

func TestGetVatInfo(t *testing.T) {
	t.Skip("skipping test")
	demo := "000000000"
	// demo for VAT number
	// replace username and password with the ones you got from
	// https://www1.aade.gr/sgsisapps/tokenservices/protected/displayConsole.htm
	// username and password will be read from env
	c := NewClient("", "")
	i, err := c.GetVATInfo("", demo)
	require.NoError(t, err, "error getting VAT info")
	require.NotEmpty(t, i.String())
	js, err := json.Marshal(i)
	require.NoError(t, err, "error marshaling json")
	spew.Dump(string(js))
}

func TestParseVatInfo(t *testing.T) {
	// a mock http response
	body := fmt.Sprintf(`<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
   <env:Header/>
   <env:Body>
      <srvc:rgWsPublic2AfmMethodResponse xmlns="http://rgwspublic2/RgWsPublic2" xmlns:srvc="http://rgwspublic2/RgWsPublic2Service">
         <srvc:result>
            <rg_ws_public2_result_rtType>
               <call_seq_id>862701698</call_seq_id>
               <error_rec>
                  <error_code xsi:nil="true"/>
                  <error_descr xsi:nil="true"/>
               </error_rec>
               <afm_called_by_rec>
                  <token_username>XXXXXXXXXXXXXXXXXXXXXXXXXXXX</token_username>
                  <token_afm>XXXXXXXXXX</token_afm>
                  <token_afm_fullname>FIF*** USE*** του TAX***</token_afm_fullname>
                  <afm_called_by>XXXXXXXXXX</afm_called_by>
                  <afm_called_by_fullname>FIF*** USE*** του TAX***</afm_called_by_fullname>
                  <as_on_date>2016-07-01</as_on_date>
               </afm_called_by_rec>
               <basic_rec>
                  <afm>090165560</afm>
                  <doy>1104</doy>
                  <doy_descr>Δ΄ ΑΘΗΝΩΝ</doy_descr>
                  <i_ni_flag_descr>ΜΗ ΦΠ</i_ni_flag_descr>
                  <deactivation_flag>1</deactivation_flag>
                  <deactivation_flag_descr>ΕΝΕΡΓΟΣ ΑΦΜ</deactivation_flag_descr>
                  <firm_flag_descr>ΕΠΙΤΗΔΕΥΜΑΤΙΑΣ</firm_flag_descr>
                  <onomasia>ΥΠΟΥΡΓΕΙΟ ΟΙΚΟΝΟΜΙΚΩΝ ΓΕΝΙΚΗ   ΔIEYΘΥΝΣΗ  Δ ΚΗΣΥΠ ΞΗΣ Δ ΟΙΚ</onomasia>
                  <commer_title xsi:nil="true"/>
                  <legal_status_descr>ΔΗΜΟΣΙΑ ΥΠΗΡΕΣΙΑ</legal_status_descr>
                  <postal_address>Κ ΣΕΡΒΙΑΣ</postal_address>
                  <postal_address_no>10</postal_address_no>
                  <postal_zip_code>10110</postal_zip_code>
                  <postal_area_description>ΑΘΗΝΑ</postal_area_description>
                  <regist_date>1993-02-08</regist_date>
                  <stop_date xsi:nil="true"/>
                  <normal_vat_system_flag>N</normal_vat_system_flag>
               </basic_rec>
               <firm_act_tab>
                  <item>
                     <firm_act_code>84111000</firm_act_code>
                     <firm_act_descr>ΓΕΝΙΚΕΣ ΔΗΜΟΣΙΕΣ ΥΠΗΡΕΣΙΕΣ</firm_act_descr>
                     <firm_act_kind>1</firm_act_kind>
                     <firm_act_kind_descr>ΚΥΡΙΑ</firm_act_kind_descr>
                  </item>
               </firm_act_tab>
            </rg_ws_public2_result_rtType>
         </srvc:result>
      </srvc:rgWsPublic2AfmMethodResponse>
   </env:Body>
</env:Envelope>`)

	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header),
	}

	i, err := parseXML(r)
	spew.Dump(i)
	require.NoError(t, err, "error parsing xml")

	onomasia, doy, afm, address, zip, area := i.VATInfo.GetInvoiceData()
	require.Equal(t, "Δ΄ ΑΘΗΝΩΝ", doy)
	require.Equal(t, "ΥΠΟΥΡΓΕΙΟ ΟΙΚΟΝΟΜΙΚΩΝ ΓΕΝΙΚΗ   ΔIEYΘΥΝΣΗ  Δ ΚΗΣΥΠ ΞΗΣ Δ ΟΙΚ", onomasia)
	require.Equal(t, "090165560", afm)
	require.Equal(t, "Κ ΣΕΡΒΙΑΣ 10", address)
	require.Equal(t, "10110", zip)
	require.Equal(t, "ΑΘΗΝΑ", area)

	mainActivity := i.VATInfo.GetMainActivity()
	require.Equal(t, "ΓΕΝΙΚΕΣ ΔΗΜΟΣΙΕΣ ΥΠΗΡΕΣΙΕΣ", mainActivity)
}

func TestParseVersion(t *testing.T) {
	// a mock http response
	body := fmt.Sprintf(`
<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
   <env:Header/>
   <env:Body>
      <srvc:rgWsPublic2VersionInfoResponse xmlns="http://rgwspublic2/RgWsPublic2" xmlns:srvc="http://rgwspublic2/RgWsPublic2Service">
         <srvc:result>Διαδικτυακή Υπηρεσία Α.Α.Δ.Ε. "Βασικά στοιχεία μητρώου για νομικά πρόσωπα, νομικές οντότητες, και φυσικά πρόσωπα με εισόδημα από επιχειρηματική δραστηριότητα» με όριο κλήσεων και ταυτοποίηση χρήστη. Release: 4.0.0, 01/07/2018, Copyright Α.Α.Δ.Ε. 2018.</srvc:result>
      </srvc:rgWsPublic2VersionInfoResponse>
   </env:Body>
</env:Envelope>`)

	r := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header),
	}

	v, err := parseXML(r)
	require.NoError(t, err, "error parsing xml")
	require.Equal(t, *v.Version, "Διαδικτυακή Υπηρεσία Α.Α.Δ.Ε. \"Βασικά στοιχεία μητρώου για νομικά πρόσωπα, νομικές οντότητες, και φυσικά πρόσωπα με εισόδημα από επιχειρηματική δραστηριότητα» με όριο κλήσεων και ταυτοποίηση χρήστη. Release: 4.0.0, 01/07/2018, Copyright Α.Α.Δ.Ε. 2018.")
}
