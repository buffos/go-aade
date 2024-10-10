package viesService

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const (
	// europeanViesEndpoint is the current default vies endpoint
	VIESEndpoint = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"
)

const (
	checkVatTemplate = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
<Body>
	<checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
		<countryCode>%s</countryCode>
		<vatNumber>%s</vatNumber>
	</checkVat>
</Body>
</Envelope>`
)

var validCountryCodes = []string{"AT", "BE", "BG", "CY", "CZ", "DE", "DK", "EE", "EL", "ES", "FI", "FR", "GB", "HR", "HU", "IE", "IT", "LT", "LU", "LV", "MT", "NL", "PL", "PT", "RO", "SE", "SI", "SK", "XI"}

// parseXML parses the vat service http response and returns an XMLBody or an error
func parseXML(r *http.Response) (*XMLBody, error) {

	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	xmlResp := XMLResponse{}
	err = xml.Unmarshal(responseBody, &xmlResp)
	if err != nil {
		return nil, err
	}

	return &xmlResp.Body, nil
}

func prepareVatRequest(vat string) (string, error) {
	// the minimum length of vat is 4. 2 for country code and 2 for vat number
	// the maximum length of vat is 15. 2 for country code and 13 for vat number
	if len(vat) < 4 || len(vat) > 15 {
		return "", ErrorVatTooShort
	}

	// check if the country code is valid (also the caller should not be in the same country)
	countryCode := vat[:2]
	if !slices.Contains(validCountryCodes, countryCode) {
		return "", ErrorInvalidCountryCode
	}

	vatNumber := vat[2:]

	return fmt.Sprintf(checkVatTemplate, countryCode, vatNumber), nil
}

func GetViesVatInfo(vat string) (VATInfo, error) {
	body, err := prepareVatRequest(vat)
	if err != nil {
		return VATInfo{}, err
	}

	req, err := http.NewRequest(http.MethodPost, VIESEndpoint, strings.NewReader(body))
	if err != nil {
		return VATInfo{}, err
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("SOAPAction", fmt.Sprintf("urn:%s", "checkVatService"))
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return VATInfo{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body %s\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return VATInfo{}, fmt.Errorf("HTTP Status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	xmlBody, err := parseXML(resp)
	if err != nil {
		return VATInfo{}, err
	}

	err = xmlBody.error()
	if err != nil {
		return VATInfo{}, err
	}

	return xmlBody.VATInfo, nil
}
