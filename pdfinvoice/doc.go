/*
Package pdfinvoice is a library for generating PDF invoices.

# Defining the required environment variables

The following environment variables are required:

	`FONT_PATH` - the path to the font files to be used
	`LOGO_FILE_PATH` - the path to the logo file to be used.

In the default options struct, we are expecting the following font files:

	`FONT_PATH`/Roboto-Regular.ttf
	`FONT_PATH`/Roboto-Bold.ttf

# Creating the document

	doc, err := NewInvoice(nil)

By passing nil as the first argument, the default configuration will be used.

# Filling the document with the required information

  - Issuer
  - Customer
  - Type
  - Number
  - Date
  - Series
  - Mark
  - PaymentMethod
  - Notes
  - QrCodeString

# Adding invoice entries

Using the doc.InvoiceDetails.AddEntry method, you can add invoice entries. Those are the items sold or
services provided.

# Calculate Totals and Taxes

  - call doc.CalculateTotals() to calculate the totals

and fill all Taxes entries

	doc.Taxes.WithHoldingTaxes =
	doc.Taxes.Deductions =
	doc.Taxes.StampDuty =
	doc.Taxes.Fees =
	doc.Taxes.OtherTaxes =

# Generate the PDF

After providing all the required data we need to create the actual document. Each document has a header, a middle body and a footer.
Three corresponding functions are responsible to place on the pdf document the data provided before.
We need to provide any struct or data type, that implements the InvoiceCreator interface, which provides those three functions.

We have already one implementation of this interface, the simpleA4Invoice.

We create the document with the following code:

	err := doc.CreateAndSave("test.pdf", SimpleA4Invoice{})
*/
package pdfinvoice
