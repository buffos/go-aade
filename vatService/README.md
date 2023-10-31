# vatService

A client library for the Greek VAT service written in Go.
This is based on a fork of [github.com/kamilakis/rgwspublic](https://github.com/kamilakis/rgwspublic) adding some minor functionality.

## Installation

```bash
go get -u github.com/buffos/go-aade/vatservice
```

## Dependencies

This library has no dependencies. 
The test library uses 
 - [github.com/stretchr/testify](https://github.com/stretchr/testify) for assertions and
 - [github.com/davecgh/go-spew/spew](https://github.com/davecgh/go-spew/spew) for pretty printing.

## Usage

```go
package main

import (
    "fmt"
    "github.com/buffos/go-aade/vatservice"
	
	func main() {
		// Create a new client
		client := vatservice.NewClient("", "") // username, password
		// if left empty, the client loads the credentials from the environment variables GSIS_VAT_USERNAME and GSIS_VAT_PASSWORD
		vatCallingTheService := "" // the vat calling the service can be left empty.
		vatToCheck := "012314564" // the vat
		info, err := client.GetVATInfo(vatCallingTheService, vatToCheck) //
		
		if err != nil {
			fmt.Println(err)
		}
		
		// we can extract the basic info needed
		onomasia, doy, afm, address, zip, area := info.GetInvoiceData()
		fmt.Println(onomasia, doy, afm, address, zip, area)
		// also, we can get the main activity
		fmt.Println(info.GetMainActivity())
    }
)

```

The service translates the error codes to human-readable messages. You can also check if the error is of type `vatservice.Error` 
and get the error code and message.

```go
	var e *vatService.Error
	errors.As(err, e)
```

## References

- [Αναζήτηση Βασικών Στοιχείων Μητρώου Επιχειρήσεων](https://www.aade.gr/epiheiriseis/forologikes-ypiresies/mitroo/anazitisi-basikon-stoiheion-mitrooy-epiheiriseon)

- [Οδηγίες για προγραμματιστές](http://www.logistis-forotexnis.gr/ckfinder/userfiles/files/%CE%9F%CE%B4%CE%B7%CE%B3%CE%AF%CE%B5%CF%82%20%CE%B3%CE%B9%CE%B1%20%CE%A0%CF%81%CE%BF%CE%B3%CF%81%CE%B1%CE%BC%CE%BC%CE%B1%CF%84%CE%B9%CF%83%CF%84%CE%AD%CF%82.pdf)

Με τη χρήση αυτής της υπηρεσίας, τα νομικά πρόσωπα, οι νομικές οντότητες, 
και τα φυσικά πρόσωπα με εισόδημα από επιχειρηματική δραστηριότητα μπορούν να αναζητήσουν βασικές πληροφορίες,
προκειμένου να διακριβώσουν τη φορολογική ή την επαγγελματική υπόσταση άλλων νομικών προσώπων ή 
νομικών οντοτήτων ή φορολογουμένων/φυσικών προσώπων που ασκούν επιχειρηματική δραστηριότητα.

H υπηρεσία στοχεύει στην καταπολέμηση της έκδοσης πλαστών και εικονικών στοιχείων από 
φορολογικά ανύπαρκτα πρόσωπα και παρέχεται κατόπιν γνωμοδότησης της Α.Π.Δ.Π.Χ. (με αριθ. 1/2011) 
και σύμφωνα με το άρθρο 20 του ν. 3842/2010 Αποκατάσταση φορολογικής δικαιοσύνης, 
αντιμετώπιση της φοροδιαφυγής και άλλες διατάξεις και το εδάφιο (ι) του άρθρου 17 του ν. 4174/2013 
περί Κώδικα Φορολογικής Διαδικασίας.

### Βήμα - βήμα

1. [x] Εγγραφή στην [υπηρεσία](https://www1.aade.gr/webtax/wspublicreg/faces/pages/wspublicreg/menu.xhtml) κάνοντας χρήση των κωδικών TAXISnet.
2. [x] Απόκτηση ειδικών κωδικών πρόσβασης μέσω της εφαρμογής [Διαχείριση Ειδικών Κωδικών](https://www1.aade.gr/sgsisapps/tokenservices/protected/displayConsole.htm).
3. [x] Χρήση ένος προγράμματος της αρεσκείας σας για την [κλήση της υπηρεσίας](https://www.aade.gr/sites/default/files/2018-07/AadeWebServiceRgWsPublicV401Client.zip).


### Τα βασικά χαρακτηριστικά της υπηρεσίας είναι:

* Η υπηρεσία μπορεί να αξιοποιηθεί απ’ όλους τους πιστοποιημένους χρήστες του TAXISnet.
* Υπάρχει μηνιαίο όριο κλήσεων της υπηρεσίας.
* Ο ΑΦΜ τα στοιχεία του οποίου αναζητούνται, ενημερώνεται με ειδική ειδοποίηση, για το ΑΦΜ / ονοματεπώνυμο που έκανε την αναζήτηση.
* Μέσω της οθόνης εγγραφής στην υπηρεσία μπορεί κάποιος να εξουσιοδοτήσει ένα τρίτο ΑΦΜ να καλεί την υπηρεσία γι’ αυτόν.

**Τα WSDL / ENDPOINT / XSD της αναβαθμισμένης υπηρεσίας είναι:**

* WSDL		: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2?WSDL
* ENDPOINT	: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2
* XSD		: https://www1.gsis.gr/wsaade/RgWsPublic2/RgWsPublic2?xsd=1

Πρόκειται για Soap JAX-WS 2.0 Web Service (έκδοσης SOAP 1.2).

    Για να καλέσει ένας σταθμός εργασίας την υπηρεσία απαιτείται δικτυακή πρόσβαση στο www1.gsis.gr και στο port 443.

    Εφόσον γίνει χρήση Java, απαιτείται χρήση Java 1.8 ή μεταγενέστερη λόγω της χρήσης του πρωτοκόλλου επικοινωνίας TLS1.2.

    Περιλαμβάνονται:

    a) παραδείγματα κλήσης (Request XML / Response XML) του Web Service,
    b) ένα SoapUI project για να γίνει import στο SoapUI. Προτείνεται χρήση SoapUI Version 5.4.0 ή μεταγενέστερη λόγω της Java 1.8 ( https://www.soapui.org/downloads/latest-release.html ).
    c) Τύπος δεδομένων, μέγεθος και τιμές επιστρεφόμενων στοιχείων

### Συχνές ερωτήσεις - απαντήσεις

- [x] [Αναζήτηση Βασικών Στοιχείων Μητρώου Επιχειρήσεων (έκδοση 1.6., 18/07/2023)](https://www.aade.gr/sites/default/files/2023-07/FAQS_anazitisi_vasikwn_stx_mitrwou_epix.pdf)
- [x] [Ειδικοί Κωδικοί Πρόσβασης](https://www.aade.gr/sites/default/files/2018-07/eidikoi_kwdikoi_FAQs.pdf)

### Βοηθητικό Υλικό

- [x] [Υλικό τεκμηρίωσης για προγραμματιστές (έκδοση 1.1., 18/07/2023)](https://www.aade.gr/sites/default/files/2023-07/RgWsPublic2DevelopersInfoV1.1.zip)
- [x] [Οδηγίες Χρήσης της Υπηρεσίας](https://www.aade.gr/sites/default/files/2018-07/RgWsPublic2OroiXrisisV4.0.pdf)