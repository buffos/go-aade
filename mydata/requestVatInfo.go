package mydata

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

func (c *Client) RequestVatInfo(params mydataInvoices.RequestVatInfoParams) (int, *mydataInvoices.RequestedVatInfoType, error) {
	return Requester[*mydataInvoices.RequestVatInfoParams, mydataInvoices.RequestedVatInfoType](c, &params, URLRequestVatInfo)
}
