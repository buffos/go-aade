package mydata

import (
	"testing"
	"time"

	"github.com/buffos/go-aade/mydata/mydataInvoices"
)

const (
	iterationFmt = "iteration %d"
)

func TestMyIncomeIterator(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	params := mydataInvoices.RequestMyIncomeParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	}
	iter := NewIncomeIterator(c, params)
	for i, res := range iter.Iterate() {
		t.Logf(iterationFmt, i)
		res.Print()
	}
}

func TestVatInfoIterator(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	params := mydataInvoices.RequestVatInfoParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	}
	iter := NewVatInfoIterator(c, params)
	for i, res := range iter.Iterate() {
		t.Logf(iterationFmt, i)
		res.Print()
	}
}

func TestMyExpensesIterator(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	params := mydataInvoices.RequestMyExpensesParams{
		DateFrom: time.Now().UTC().AddDate(0, -12, 0),
		DateTo:   time.Now().UTC(),
	}
	iter := NewExpensesIterator(c, params)
	for i, res := range iter.Iterate() {
		t.Logf(iterationFmt, i)
		res.Print()
	}
}

func TestTransmittedDocsIterator(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	params := mydataInvoices.RequestDocsParams{
		Mark:     "0",
		DateFrom: time.Now().UTC().AddDate(0, -3, 0),
		DateTo:   time.Now().UTC(),
	}
	iter := NewTransmittedDocsIterator(c, params)
	for i, res := range iter.Iterate() {
		t.Logf(iterationFmt, i)
		res.Print()
	}
}

func TestRequestDocsIterator(t *testing.T) {
	c := NewClient(userID, subscriptionKey, 30, false)
	params := mydataInvoices.RequestDocsParams{
		Mark:     "0",
		DateFrom: time.Now().UTC().AddDate(0, -3, 0),
		DateTo:   time.Now().UTC(),
	}
	iter := NewDocsIterator(c, params)
	for i, res := range iter.Iterate() {
		t.Logf(iterationFmt, i)
		res.Print()
	}
}
