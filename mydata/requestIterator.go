package mydata

import (
	"iter"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

type WithNextDataSetter interface {
	SetNextPartitionData(partitionKey, rowKey string)
}

type WithNextDataGetter interface {
	GetNextPartitionData() (string, string)
}

type IteratorParams interface {
	WithNextDataSetter
	ToMap
}

// type for a function that takes a WithNextDataSetter and returns a WithNextDataGetter
type RequestIteratorFunc func(WithNextDataSetter) (int, WithNextDataGetter, error)

type RequestIterator[P IteratorParams, R WithNextDataGetter] struct {
	c       *Client
	params  P
	arg     R
	urlPath string
}

func NewRequestIterator[P IteratorParams, R WithNextDataGetter](c *Client, params P, arg R, urlPath string) RequestIterator[P, R] {
	return RequestIterator[P, R]{c: c, params: params, arg: arg, urlPath: urlPath}
}

func NewIncomeIterator(c *Client, params mydataInvoices.RequestMyIncomeParams) RequestIterator[*mydataInvoices.RequestMyIncomeParams, *mydataInvoices.RequestedBookInfo] {
	return NewRequestIterator(c, &params, &mydataInvoices.RequestedBookInfo{}, URLRequestMyIncome)
}

func NewVatInfoIterator(c *Client, params mydataInvoices.RequestVatInfoParams) RequestIterator[*mydataInvoices.RequestVatInfoParams, *mydataInvoices.RequestedVatInfoType] {
	return NewRequestIterator(c, &params, &mydataInvoices.RequestedVatInfoType{}, URLRequestVatInfo)
}

func NewExpensesIterator(c *Client, params mydataInvoices.RequestMyExpensesParams) RequestIterator[*mydataInvoices.RequestMyExpensesParams, *mydataInvoices.RequestedBookInfo] {
	return NewRequestIterator(c, &params, &mydataInvoices.RequestedBookInfo{}, URLRequestMyExpenses)
}

func NewTransmittedDocsIterator(c *Client, params mydataInvoices.RequestDocsParams) RequestIterator[*mydataInvoices.RequestDocsParams, *mydataInvoices.RequestedDoc] {
	return NewRequestIterator(c, &params, &mydataInvoices.RequestedDoc{}, URLRequestTransmittedDocs)
}

func NewDocsIterator(c *Client, params mydataInvoices.RequestDocsParams) RequestIterator[*mydataInvoices.RequestDocsParams, *mydataInvoices.RequestedDoc] {
	return NewRequestIterator(c, &params, &mydataInvoices.RequestedDoc{}, URLRequestDocs)
}

func (ri RequestIterator[P, R]) Iterate() func(func(int, R) bool) {
	index := 0
	return func(yield func(int, R) bool) {
		for {
			_, arg, err := Requester[P, R](ri.c, ri.params, ri.urlPath)
			if err != nil {
				return
			}
			ri.arg = *arg
			if !yield(index, ri.arg) {
				return
			}
			index++
			// get the next partition and row keys
			partitionKey, rowKey := ri.arg.GetNextPartitionData()
			if partitionKey == "" && rowKey == "" {
				return
			}
			ri.params.SetNextPartitionData(partitionKey, rowKey)
		}
	}
}

func (ri RequestIterator[P, R]) Pull() (func() (int, R, bool), func()) {
	return iter.Pull2(ri.Iterate())
}
