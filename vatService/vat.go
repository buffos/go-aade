package vatService

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	username string
	password string
}

// NewClient creates a new client to call the greek VAT service
func NewClient(username, password string) *Client {
	if username == "" {
		username = os.Getenv("GSIS_VAT_USERNAME")
	}
	if password == "" {
		password = os.Getenv("GSIS_VAT_PASSWORD")
	}
	return &Client{
		username: username,
		password: password,
	}
}

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

// Version makes a call at the service and returns the version of the service
func (c *Client) Version() (*string, error) {
	body := `<?xml version="1.0" encoding="UTF-8"?>
		<soap:Envelope 
			xmlns:soap="http://www.w3.org/2003/05/soap-envelope" 
			xmlns:rgw="http://rgwspublic2/RgWsPublic2Service">
			<soap:Header/>
			<soap:Body>
	   			<rgw:rgWsPublic2VersionInfo/>
			</soap:Body>
 		</soap:Envelope>`
	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(body))
	if err != nil {
		return nil, ErrCannotReachService
	}
	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body %s\n", err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, ErrCannotReachService
	}

	xmlBody, err := parseXML(resp)
	if err != nil {
		return nil, err
	}

	err = xmlBody.error()
	if err != nil {
		return nil, err
	}
	xmlBody.Error = nil // to correct parser creating an object
	return xmlBody.Version, nil
}

// GetVATInfo associated with a VAT number
// accepts a called by VAT and a called for VAT, username and password
// returns AFMData or an error
func (c *Client) GetVATInfo(calledBy, calledFor string) (*VATInfo, error) {

	// vat numbers must be between 9 and 12 chars
	if len(calledFor) < 9 || len(calledFor) > 12 {
		return nil, ErrInvalidVAT
	}
	// the first one (calledBy) can be empty
	if calledBy != "" {
		if len(calledBy) < 9 || len(calledBy) > 12 {
			return nil, ErrInvalidVAT
		}
	}
	// same for username/password
	if len(c.username) < 6 || len(c.password) < 6 {
		return nil, ErrInvalidCredentials
	}

	body := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
		<env:Envelope 
			xmlns:env="http://www.w3.org/2003/05/soap-envelope" 
			xmlns:ns1="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
			xmlns:ns2="http://rgwspublic2/RgWsPublic2Service" 
			xmlns:ns3="http://rgwspublic2/RgWsPublic2">
			<env:Header>
	   			<ns1:Security>
		  			<ns1:UsernameToken>
			 			<ns1:Username>%v</ns1:Username>
			 			<ns1:Password>%v</ns1:Password>
		  			</ns1:UsernameToken>
	   			</ns1:Security>
			</env:Header>
			<env:Body>
	   			<ns2:rgWsPublic2AfmMethod>
		  			<ns2:INPUT_REC>
						<ns3:afm_called_by>%v</ns3:afm_called_by>
			 			<ns3:afm_called_for>%v</ns3:afm_called_for>
		  			</ns2:INPUT_REC>
	   			</ns2:rgWsPublic2AfmMethod>
			</env:Body>
 		</env:Envelope>`, c.username, c.password, calledBy, calledFor)

	req, err := http.NewRequest("POST", Endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Set("Content-Type", "application/soap+xml")
	header.Set("Connection", "keep-alive")
	header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body %s\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status: %d, error: %s", resp.StatusCode, resp.Status)
	}

	xmlBody, err := parseXML(resp)
	if err != nil {
		return nil, err
	}

	err = xmlBody.VATInfo.error()
	if err != nil {
		return nil, err
	}

	// to correct parser creating an object
	xmlBody.VATInfo.Error = nil
	return &xmlBody.VATInfo, nil
}
