# invoice generator

## Summary

A simple api to create invoices in pdf format. Based on fpdf package.

## Dependencies

- [fpdf](https://www.github.com/go-pdf/fpdf)
- [go-qrcode](https://www.github.com/skip2/go-qrcode)

## Installation

```bash
go get -u github.com/buffos/go-aade/pdfinvoice
```

## Usage

First, We need to set to environment variables to use the api. 
 - `FONT_PATH`: The path to the fonts directory
 - `LOGO_FILE_PATH`: the path to the logo file (including the file name and extension)


We create a document that holds the required data
```go
doc, err := NewInvoice(nil)
```

If we pass nil as the options struct, the default options are used.
You can view and create custom options. The options struct is defined in the options.go file.

Then we add data to the document.

```go

doc.LogoFileName = os.Getenv("LOGO_FILE_PATH")
doc.SetIssuer(&Contact{
		Name:            "Κωνσταντίνος Παπαδόπουλος",
		AFM:             "090000000",
		WorkDescription: "Τίποτα Απολύτως",
		Address: &Address{
			Address:    "Εγνατίας 12",
			Address2:   "",
			City:       "Θεσσαλονίκη",
			PostalCode: "12345",
		},
		DOY:            "Α΄ ΘΕΣΣΑΛΟΝΙΚΗΣ",
		AdditionalInfo: []string{"ΓΕΜΗ: 231506418987897 · email: support@test.gr \nΤηλ: 2210000111 · Κιν: 697777777"},
	})
	doc.SetCustomer(&Contact{
		Name:            "Παπαδόπουλος Ανδρέας",
		AFM:             "090000000",
		WorkDescription: "Παροχή Υπηρεσιών",
		Address: &Address{
			Address:    "Τσιμισκή 12",
			Address2:   "",
			City:       "Θεσσαλονίκη",
			PostalCode: "12345",
		},
		DOY: "Β΄ Θεσσαλονίκης",
	})

doc.Type = "Τιμολόγιο Παροχής Υπηρεσιών"
doc.Series = "ΤΠΥ-Α"
doc.Number = "1"
doc.Date = "01/01/2021"
doc.Mark = "4000000012351351"
doc.PaymentMethod = "Μετρητά"
doc.Notes = `Η παραγγελία θα εκτελεσθεί μέχρι τις 24 Ιουνίου 2021. Αριθμός συναλλαγής 40909345`
doc.QRCodeString = "https://www.example.gr"
```

Now we need to add the invoice entries for each item. For example:

```go
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή υπηρεσιών κατασκευής ιστοσελίδας",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
```

Finally we need to create the document. It is as simple as possible

```go
err = doc.CreateAndSave("test.pdf", SimpleA4Invoice{})
```

The second argument is a struct that implements the `InvoiceCreator` interface.
This interface has three functions that 'draw' the header, footer and middle part of the invoice using the data we provided before.

We can create different implementations of this interface that create different invoices.

If we need to get a buffer, instead of a file, to upload it to a service, for example, S3, we can use

```go
buffer, err := doc.CreateAndBuffer(SimpleA4Invoice{})
```

where buffer is a `*bytes.Buffer` pointer.