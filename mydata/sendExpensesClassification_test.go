package mydata

import (
	"fmt"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSendExpensesClassificationRejectPostPerLine(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.ExpensesClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Reject 1",
			createRequest: func() *mydataInvoices.ExpensesClassificationsDoc {
				doc := mydataInvoices.NewExpensesClassificationDoc()
				doc.RejectClassification(400001906829007, "")
				return doc
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          false,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			request := tc.createRequest()
			code, doc, err := c.SendExpensesClassification(request, false)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendExpensesClassificationDeviatePostPerLine(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.ExpensesClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Deviate 1",
			createRequest: func() *mydataInvoices.ExpensesClassificationsDoc {
				doc := mydataInvoices.NewExpensesClassificationDoc()
				doc.DeviateClassification(400001906829007, "")
				return doc
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          false,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			request := tc.createRequest()
			code, doc, err := c.SendExpensesClassification(request, false)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
			spew.Dump(err)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendExpensesClassificationEditDetailPostPerLine(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.ExpensesClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Edit one invoice (own invoice - forbidden)",
			createRequest: func() *mydataInvoices.ExpensesClassificationsDoc {
				doc := mydataInvoices.NewExpensesClassificationDoc()
				doc.EditLineNumberDetail(400001906829007, "", 1,
					mydatavalues.E3_585_009, mydatavalues.ECategory2_3, 5, 1)
				return doc
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          false,
		},
		{
			caseName: "Case 2: Edit two invoices (own invoices - both forbidden)",
			createRequest: func() *mydataInvoices.ExpensesClassificationsDoc {
				doc := mydataInvoices.NewExpensesClassificationDoc()
				doc.EditLineNumberDetail(400001906829007, "", 1,
					mydatavalues.E3_585_009, mydatavalues.ECategory2_3, 5, 1)
				doc.EditLineNumberDetail(400001906829007, "", 1,
					mydatavalues.E3_585_009, mydatavalues.ECategory2_3, 5, 1)
				return doc
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {
				require.Equal(t, 2, len(d.Response))
			},
			wantedCode: 200,
			wantErr:    true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			request := tc.createRequest()
			code, doc, err := c.SendExpensesClassification(request, false)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
			spew.Dump(err)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSimple(t *testing.T) {
	doc := mydataInvoices.NewExpensesClassificationDoc()
	doc.EditLineNumberDetail(400001906829007, "", 1,
		mydatavalues.E3_585_009, mydatavalues.ECategory2_3, 5, 1)

	c := NewClient(userID, subscriptionKey, 30, false)
	_, result, err := c.SendExpensesClassification(doc, false)
	require.Greater(t, result.Response[0].ClassificationMark, uint64(0))
	require.NoError(t, err)
	fmt.Printf("classification mark: %d\n", result.Response[0].ClassificationMark)
}

func TestSimpleNewWay(t *testing.T) {
	doc := mydataInvoices.NewExpensesClassificationDoc()
	inv1 := doc.NewInvoiceClassificationForMark(400001906829007, "")
	inv1.AddE3ClassificationDetail(mydatavalues.E3_102_001, mydatavalues.ECategory2_1, 5, 1)
	//inv1.AddVatClassificationDetail(mydatavalues.InvoiceVAT24Percent, mydatavalues.VATExceptionReasonType(0), 5, 1.2, 2)

	c := NewClient(userID, subscriptionKey, 30, false)
	_, result, err := c.SendExpensesClassification(doc, true)
	spew.Dump(result)
	spew.Dump(err)
}

//TODO: fix errors with the new method. It seems that the new method does not work.
