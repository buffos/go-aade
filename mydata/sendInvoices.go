package mydata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

type SendInvoiceParams struct {
	AadeUserID string `json:"aade-user-id"`
}

// SendInvoices sends the given invoices to the myDATA API
func (c *Client) SendInvoices(invoices *mydataInvoices.InvoicesDoc) (int, *mydataInvoices.ResponseDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()

	err := invoices.Validate()
	// returning if one of the invoices is invalid
	if err != nil && c.onInvalid == ErrorOnInvalid {
		return InternalErrorCode, nil, errors.Join(ErrorInvalidInvoices, err)
	}
	// filtering out invalid invoices if this is the required action
	if err != nil && c.onInvalid == FilterOnInvalid {
		invoices = invoices.FilterInvalid()
		// check if no invoices are left. if so, return
		if invoices.Invoices == nil || len(invoices.Invoices) == 0 {
			return 200, nil, nil
		}
	}

	// we are here if the default action was pass through or if the invalid invoices were filtered out

	body, err := xml.Marshal(invoices)
	//fmt.Println(string(body))

	if err != nil {
		return InternalErrorCode, nil, ErrorXMLMarshal
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getURL(URLSendInvoices, nil), bytes.NewBuffer(body))
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}

	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}

	//b, _ := c.responseToString(response)
	//fmt.Printf("%s\n", b)

	result, err := ParseXMLResponse[mydataInvoices.ResponseDoc](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	if result.HasErrors() {
		return response.StatusCode, result, result.Errors("send invoices")
	}
	return response.StatusCode, result, nil
}
