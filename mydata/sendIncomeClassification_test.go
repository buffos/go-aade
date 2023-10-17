package mydata

import (
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSendIncomeClassificationReject(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.IncomeClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Reject 1",
			createRequest: func() *mydataInvoices.IncomeClassificationsDoc {
				doc := mydataInvoices.NewIncomeClassificationDoc()
				doc.RejectClassification(400001917670105, "")
				return doc
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			request := tc.createRequest()
			code, doc, err := c.SendIncomeClassification(request)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendIncomeClassificationDeviate(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.IncomeClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Deviate 1",
			createRequest: func() *mydataInvoices.IncomeClassificationsDoc {
				doc := mydataInvoices.NewIncomeClassificationDoc()
				doc.DeviateClassification(400001917670105, "")
				return doc
			},
			validateResponse: func(*mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          true,
		},
	}
	c := NewClient(userID, subscriptionKey, 30, false)

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			request := tc.createRequest()
			code, doc, err := c.SendIncomeClassification(request)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
			spew.Dump(err)
		})
		time.Sleep(2 * time.Second)
	}
}

func TestSendIncomeClassificationEditDetail(t *testing.T) {
	testCases := []struct {
		caseName         string
		createRequest    func() *mydataInvoices.IncomeClassificationsDoc
		validateResponse func(*mydataInvoices.ResponseDoc)
		wantedCode       int
		wantErr          bool
	}{
		{
			caseName: "Case 1: Edit one invoice (own invoice - forbidden)",
			createRequest: func() *mydataInvoices.IncomeClassificationsDoc {
				doc := mydataInvoices.NewIncomeClassificationDoc()
				doc.EditLineNumberDetail(400001917670105, "", 1,
					mydatavalues.IE3_561_002, mydatavalues.ICategory1_1, 1, 1)
				return doc
			},
			validateResponse: func(d *mydataInvoices.ResponseDoc) {},
			wantedCode:       200,
			wantErr:          true,
		},
		{
			caseName: "Case 2: Edit two invoices (own invoices - both forbidden)",
			createRequest: func() *mydataInvoices.IncomeClassificationsDoc {
				doc := mydataInvoices.NewIncomeClassificationDoc()
				doc.EditLineNumberDetail(400001917670105, "", 1,
					mydatavalues.IE3_561_002, mydatavalues.ICategory1_1, 1, 0)
				doc.EditLineNumberDetail(400001917670105, "", 1,
					mydatavalues.IE3_561_002, mydatavalues.ICategory1_1, 1, 0)
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
			code, doc, err := c.SendIncomeClassification(request)
			tc.validateResponse(doc)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantedCode, code)
			spew.Dump(doc)
			spew.Dump(err)
		})
		time.Sleep(2 * time.Second)
	}
}
