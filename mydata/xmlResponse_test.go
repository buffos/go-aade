package mydata

import (
	"bytes"
	"encoding/xml"
	"github.com/buffos/go-aade/mydata/mydataInvoices"
	"github.com/stretchr/testify/require"
	"testing"
)

var xmlResponse = `<RequestedDoc xmlns:icls="https://www.aade.gr/myDATA/incomeClassificaton/v1.0" xmlns:ecls="https://www.aade.gr/myDATA/expensesClassificaton/v1.0" xmlns="http://www.aade.gr/myDATA/invoice/v1.0">
  <invoicesDoc>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000290597</aa>
        <issueDate>2022-11-19</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000395264</aa>
        <issueDate>2022-11-20</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000505417</aa>
        <issueDate>2022-11-22</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000358630</aa>
        <issueDate>2022-12-12</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>12.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>12</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>12</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>12</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000024688</aa>
        <issueDate>2023-01-12</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>1.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>1</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>1</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>1</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000386621</aa>
        <issueDate>2022-12-14</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000572355</aa>
        <issueDate>2022-12-19</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000080918</aa>
        <issueDate>2023-01-17</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000158032</aa>
        <issueDate>2023-01-23</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>24.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>24</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>24</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>24</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000328065</aa>
        <issueDate>2023-01-26</issueDate>
        <invoiceType>2.1</invoiceType>

        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <cancelledByMark>400001901103528</cancelledByMark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000619468</aa>
        <issueDate>2023-02-15</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000619468</aa>
        <issueDate>2023-02-15</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000793242</aa>
        <issueDate>2023-02-20</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1000966043</aa>
        <issueDate>2023-03-02</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>6.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>6</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>6</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>6</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1001121118</aa>
        <issueDate>2023-03-14</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1001290812</aa>
        <issueDate>2023-03-29</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1001476966</aa>
        <issueDate>2023-04-12</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1001559736</aa>
        <issueDate>2023-04-18</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1002010345</aa>
        <issueDate>2023-05-15</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1002205807</aa>
        <issueDate>2023-05-22</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1002205807</aa>
        <issueDate>2023-05-22</issueDate>
        <invoiceType>2.1</invoiceType>

        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1002658312</aa>
        <issueDate>2023-06-13</issueDate>
        <invoiceType>2.1</invoiceType>

        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>0.30</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>0.3</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>0.3</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>0.3</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </counterpart>
      <invoiceHeader>
        <series>ΕΑ</series>
        <aa>1002822656</aa>
        <issueDate>2023-06-19</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>1</type>
          <amount>5.00</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>5</netValue>
        <vatCategory>7</vatCategory>
        <vatAmount>0</vatAmount>
        <vatExemptionCategory>7</vatExemptionCategory>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>5</totalNetValue>
        <totalVatAmount>0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>5</totalGrossValue>
      </invoiceSummary>
    </invoice>
    <invoice>
      <uid>CE8</uid>
      <mark>40000</mark>
      <issuer>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
      </issuer>
      <counterpart>
         <vatNumber>11111111111</vatNumber>
        <country>GR</country>
        <branch>0</branch>
        <address>
          <postalCode>56626</postalCode>
          <city>THESSALONIKI</city>
        </address>
      </counterpart>
      <invoiceHeader>
        <series>test_Α</series>
        <aa>8909</aa>
        <issueDate>2023-02-16</issueDate>
        <invoiceType>2.1</invoiceType>
        <vatPaymentSuspension>false</vatPaymentSuspension>
        <currency>EUR</currency>
      </invoiceHeader>
      <paymentMethods>
        <paymentMethodDetails>
          <type>3</type>
          <amount>62.02</amount>
        </paymentMethodDetails>
      </paymentMethods>
      <invoiceDetails>
        <lineNumber>1</lineNumber>
        <netValue>50.02</netValue>
        <vatCategory>1</vatCategory>
        <vatAmount>12.0</vatAmount>
      </invoiceDetails>
      <invoiceSummary>
        <totalNetValue>50.02</totalNetValue>
        <totalVatAmount>12.0</totalVatAmount>
        <totalWithheldAmount>0</totalWithheldAmount>
        <totalFeesAmount>0</totalFeesAmount>
        <totalStampDutyAmount>0</totalStampDutyAmount>
        <totalOtherTaxesAmount>0</totalOtherTaxesAmount>
        <totalDeductionsAmount>0</totalDeductionsAmount>
        <totalGrossValue>62.02</totalGrossValue>
      </invoiceSummary>
    </invoice>
  </invoicesDoc>
</RequestedDoc>`

func TestResponse(t *testing.T) {
	var result mydataInvoices.RequestedDoc
	err := xml.NewDecoder(bytes.NewReader([]byte(xmlResponse))).Decode(&result)
	require.NoError(t, err)
	require.Equal(t, "11111111111", *result.InvoicesDoc.Invoices[0].Issuer.VatNumber)
}
