package pdfinvoice

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func setupDoc(doc *Document) {
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
	doc.UID = "GR090000000"
	doc.PaymentMethod = "Μετρητά"
	doc.Notes = `Η παραγγελία θα εκτελεσθεί μέχρι τις 24 Ιουνίου 2021. Αριθμός συναλλαγής 40909345`
	doc.QRCodeString = "https://www.example.gr"

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
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          1,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          2,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          2,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          2,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})
	doc.InvoiceDetails.AddEntry(&InvoiceEntry{
		Description:       "Παροχή συντήρηση ιστοσελίδας για ένα έτος η και για δύο έτη ανάλογα.",
		Quantity:          2,
		UnitOfMeasurement: "",
		NetPrice:          100,
		Discount:          0,
		FinalPrice:        0,
		TaxPercent:        24,
		TaxAmount:         24,
	})

	doc.CalculateTotals()

	doc.Taxes.WithHoldingTaxes = 20
	doc.Taxes.Deductions = 0
	doc.Taxes.StampDuty = 0
	doc.Taxes.Fees = 0
	doc.Taxes.OtherTaxes = 0
}

func TestNewDocument(t *testing.T) {
	_ = os.Setenv("LOGO_FILE_PATH", "../assets/logos/logoLight.png")
	_ = os.Setenv("FONT_PATH", "../assets/fonts")
	doc, err := NewInvoice(nil)
	require.NoError(t, err)

	setupDoc(doc)

	err = doc.CreateAndSave("test.pdf", SimpleA4Invoice{})
	require.NoError(t, err)

}

func TestNewBufferedDocument(t *testing.T) {
	doc, err := NewInvoice(nil)
	require.NoError(t, err)

	setupDoc(doc)
	result, err := doc.CreateAndBuffer(SimpleA4Invoice{})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
