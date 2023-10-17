package mydata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"
)

// SendIncomeClassification sends the income classification of invoices already submitted to AADE
func (c *Client) SendIncomeClassification(changes *mydataInvoices.IncomeClassificationsDoc) (int, *mydataInvoices.ResponseDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()

	body, err := xml.Marshal(changes)
	//fmt.Println(string(body))

	if err != nil {
		return InternalErrorCode, nil, ErrorXMLMarshal
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getURL(URLSendIncomeClassification, nil), bytes.NewBuffer(body))
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
