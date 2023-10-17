package mydata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"
)

// SendExpensesClassification sends the expense classification of invoices submitted to AADE.
// If postMethod is false, then we change invoice classification per line number.
// If postMethod is true, then we change invoice classification per invoice mark. One for E3 and one for VAT.
func (c *Client) SendExpensesClassification(changes *mydataInvoices.ExpensesClassificationsDoc, postMethodPerInvoice bool) (int, *mydataInvoices.ResponseDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()

	body, err := xml.Marshal(changes)
	//fmt.Println(string(body))

	if err != nil {
		return InternalErrorCode, nil, ErrorXMLMarshal
	}

	var queryArgs map[string]string
	if postMethodPerInvoice {
		queryArgs = map[string]string{
			"postPerInvoice": "true",
		}
	} else {
		queryArgs = nil
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getURL(URLSendExpensesClassification, queryArgs), bytes.NewBuffer(body))
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestSendIncomeClassification
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
