package mydata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"net/http"
)

func (c *Client) SendPaymentsMethod(paymentMethods *mydataInvoices.PaymentMethodsDoc) (int, *mydataInvoices.ResponseDoc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()

	body, err := xml.Marshal(paymentMethods)
	if err != nil {
		return InternalErrorCode, nil, ErrorXMLMarshal
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getURL(URLSendPaymentsMethod, nil), bytes.NewBuffer(body))
	if err != nil {
		return InternalErrorCode, nil, ErrorRequestCreation
	}

	c.authorize(request)
	response, err := c.client.Do(request)
	if err != nil {
		return InternalErrorCode, nil, ErrorGettingResponse
	}
	result, err := ParseXMLResponse[mydataInvoices.ResponseDoc](response)
	if err != nil {
		return InternalErrorCode, nil, errors.Join(ErrorXMLParsingResponse, err)
	}
	if result.HasErrors() {
		return response.StatusCode, result, result.Errors("send payment methods")
	}
	return response.StatusCode, result, nil
}
