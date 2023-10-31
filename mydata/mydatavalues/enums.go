package mydatavalues

import "math"

const (
	Xmlns     = "http://www.aade.gr/myDATA/invoice/v1.0"
	XmlnsICLS = "https://www.aade.gr/myDATA/incomeClassificaton/v1.0"
	XmlnsECLS = "https://www.aade.gr/myDATA/expensesClassificaton/v1.0"
)

// region Transmission Failure

type TransmissionFailureType byte

//goland:noinspection GoUnusedConst
const (
	FailureConnectionToProvider       TransmissionFailureType = iota + 1 // Αποτυχία σύνδεσης με τον Πάροχο
	FailureConnectionProviderToMyData                                    // Αποτυχία σύνδεσης παρόχου με το myDATA
	FailureConnectionToMyDataFromERP                                     // Αποτυχία σύνδεσης με το myDATA από το ERP
)

// endregion

// region Τρόποι πληρωμής

type InvoicePaymentType uint

//goland:noinspection GoUnusedConst
const (
	BankAccountLocal   InvoicePaymentType = iota + 1 // Επαγγελματικός Λογαριασμός Πληρωμών Ημεδαπής
	BankAccountForeign                               // Επαγγελματικός Λογαριασμός Πληρωμών Αλλοδαπής
	Cash                                             // Μετρητά
	Cheque                                           // Επιταγή
	Credit                                           // Επί πιστώσει
	WebBanking                                       // Ηλεκτρονική Τραπεζική Πληρωμή
	POS                                              // Point of Sale
)

// endregion

// region Είδος Ποσότητας

type InvoiceMeasurementUnit uint

//goland:noinspection GoUnusedConst
const (
	Pieces    InvoiceMeasurementUnit = iota + 1 // Τεμάχια
	Kilograms                                   // Κιλά
	Liters                                      // Λίτρα
)

// endregion

// region Σκοπός Διακίνησης

type InvoicePurposeOfMovement uint

//goland:noinspection GoUnusedConst
const (
	MovePurposeSales         InvoicePurposeOfMovement = iota + 1 // Πώληση
	MovePurposeSalesOnBehalf                                     // Πώληση εκ μέρους τρίτων
	MovePurposeSample                                            // Δειγματισμός

)

// endregion

// region Επισήμανση Παραστατικού

type InvoiceDetailType uint

//goland:noinspection GoUnusedConst
const (
	InvoiceDetailSalesClearanceOfThirdParties InvoiceDetailType = iota + 1 // Εκκαθάριση Πωλήσεων Τρίτων
	InvoiceDetailPaymentFromThirdPartySales                                // Αμοιβή από Πωλήσεις Τρίτων
)

// endregion

// region Είδος Γραμμής

type InvoiceLineType uint

//goland:noinspection GoUnusedConst
const (
	SpecialLineOfWithholdingTaxes InvoiceLineType = iota + 1 // Ειδική Γραμμή Κρατήσεων Φόρου
	VATEnd                                                   // Γραμμή Τέλους με Φ.Π.Α.
	MiscTaxesWithVAT                                         // Γραμμή Λοιπών Φόρων με Φ.Π.Α.
	SpecialLinePaperStamp                                    // Ειδική Γραμμή Χαρτοσήμου
	SpecialLineDeductions                                    // Ειδική Γραμμή Κρατήσεων
	GiftCode                                                 // Δωροεπιταγή
	NegativeSignature                                        // Αρνητικό πρόσημο αξιών
)

// endregion

// region Τύπος Απόκλισης Παραστατικού

type InvoiceVariationType uint

//goland:noinspection GoUnusedConst
const (
	MissingInvoiceFromReceiver   InvoiceVariationType = iota + 1 // Διαβίβαση δεδομένων από παραλήπτη λόγω παράλειψης του εκδότη
	MissingInvoiceAgreedIssuer                                   // Διαβίβαση δεδομένων από εκδότη που συμφωνεί με την επισήμανση του λήπτη για παράλειψη
	DeviationInvoiceFromReceiver                                 // Διαβίβαση δεδομένων από παραλήπτη λόγω απόκλισης του εκδότη
	DeviationInvoiceAgreedIssuer                                 // Διαβίβαση δεδομένων από εκδότη που συμφωνεί με την επισήμανση του λήπτη για απόκλιση
)

// endregion

// region Ειδική Κατηγορία Παραστατικού

type InvoiceSpecialCategory uint

//goland:noinspection GoUnusedConst
const (
	Subsidy                InvoiceSpecialCategory = iota + 1 // Επιδοτήσεις – Επιχορηγήσεις
	HotelIncomeRoomCharges                                   // Έσοδα Ξενοδοχείων – Χρεώσεις Δωματίων
	AccountingEntry                                          // Λογιστική Εγγραφή
	TaxFree                                                  // Χωρίς Φόρο. Έγκυρη τιμή μόνο για	διαβίβαση μέσω erp ή έκδοση μέσω παρόχου ή timologio
)

// endregion

// region Είδος Φόρου

type TaxType byte

//goland:noinspection GoUnusedConst
const (
	TaxTypeWithHoldingTax TaxType = iota + 1 // Παρακρατούμενος Φόρος
	TaxTypeFees                              // Τέλη
	TaxTypeMiscellaneous                     // Λοιποί Φόροι
	TaxTypeStamp                             // Χαρτόσημο
	TaxTypeDeductions                        // Κρατήσεις
)

// endregion

// region Τύπος Παραστατικού

type InvoiceType string

// Τιμολόγια Πώλησης

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeSales                  InvoiceType = "1.1" // Τιμολόγιο Πώλησης
	InvoiceTypeSalesInsideEU          InvoiceType = "1.2" // Τιμολόγιο Πώλησης / Ενδοκοινοτικές Παραδόσεις
	InvoiceTypeSalesOutsideEU         InvoiceType = "1.3" // Τιμολόγιο Πώλησης / Παραδόσεις Τρίτων Χωρών
	InvoiceTypeSalesOnBehalfOf        InvoiceType = "1.4" // Τιμολόγιο Πώλησης εκ μέρους τρίτων
	InvoiceTypeSalesOnBehalfOfPayment InvoiceType = "1.5" // Τιμολόγιο Πώλησης / Εκκαθάριση Πωλήσεων Τρίτων - Αμοιβή από Πωλήσεις Τρίτων
	InvoiceTypeSalesComplementary     InvoiceType = "1.6" // Τιμολόγιο Πώλησης / Συμπληρωματικό Παραστατικό
)

// Τιμολόγια Παροχής Υπηρεσιών

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeServices             InvoiceType = "2.1" // Τιμολόγιο Παροχής Υπηρεσιών
	InvoiceTypeServicesInsideEU     InvoiceType = "2.2" // Τιμολόγιο Παροχής Υπηρεσιών / Ενδοκοινοτικές Παροχές Υπηρεσιών
	InvoiceTypeServicesOutsideEU    InvoiceType = "2.3" // Τιμολόγιο Παροχής Υπηρεσιών / Παροχές Υπηρεσιών Τρίτων Χωρών
	InvoiceTypeServiceComplementary InvoiceType = "2.4" // Τιμολόγιο Παροχής Υπηρεσιών / Συμπληρωματικό Παραστατικό
)

// Τίτλος Κτήσης

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeOwnershipTitleNoObligationIssuer InvoiceType = "3.1" // Τίτλος Κτήσης (μή υπόχρεος εκδότης)
	InvoiceTypeOwnershipTitleRefuseByIssuer     InvoiceType = "3.2" // Τίτλος Κτήσης (αρνητική δήλωση εκδότη)
)

// Πιστωτικό Τιμολόγιο

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeCreditSalesWithReference    InvoiceType = "5.1" // Πιστωτικό Τιμολόγιο / Συσχετιζόμενο
	InvoiceTypeCreditSalesWithoutReference InvoiceType = "5.2" // Πιστωτικό Τιμολόγιο / Μη Συσχετιζόμενο
)

// Στοιχείο Αυτοπαράδοσης - Ιδιοχρησιμοποίησης

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeSelfDelivery InvoiceType = "6.1" // Στοιχείο Αυτοπαράδοσης
	InvoiceTypeSelfUsage    InvoiceType = "6.2" // Στοιχείο Ιδιοχρησιμοποίησης
)

//goland:noinspection GoUnusedConst
const InvoiceTypeContractIncome InvoiceType = "7.1" // Συμβόλαιο Έσοδο

// Ειδικό Στοιχείο (Έσοδο) – Απόδειξη Είσπραξης

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeRentIncome                InvoiceType = "8.1" // Ενοίκια - Έσοδο
	InvoiceTypeReceiptOfAccommodationTax InvoiceType = "8.2" // Απόδειξη Είσπραξης Φόρου Διαμονής
)

// Παραστατικά Λιανικής

//goland:noinspection GoUnusedConst
const (
	RetailReceipt           InvoiceType = "11.1" // Απόδειξη Λιανικής
	RetailServiceReceipt    InvoiceType = "11.2" // Απόδειξη Παροχής Υπηρεσιών
	RetailSimplifiedInvoice InvoiceType = "11.3" // Απλοποιημένο Τιμολόγιο
	RetailCreditReceipt     InvoiceType = "11.4" // Πιστωτικό Στοιχείο Λιανικής
	RetailReceiptOnBehalfOf InvoiceType = "11.5" // Απόδειξη Λιανικής εκ μέρους τρίτων
)

// Λήψη Παραστατικών Λιανικής

//goland:noinspection GoUnusedConst
const (
	InvoiceTypeReceiveRetailReceipt  InvoiceType = "13.1" // Έξοδα - Αγορές Λιανικών Συναλλαγών ημεδαπής / αλλοδαπής
	InvoiceTypeReceiveServiceReceipt InvoiceType = "13.2" // Παροχή Λιανικών Συναλλαγών ημεδαπής / αλλοδαπής
	InvoiceTypeSharedExpenses        InvoiceType = "13.3" // Έξοδα - Κοινόχρηστα
	InvoiceTypeSubscriptions         InvoiceType = "13.4" // Έξοδα - Συνδρομές
	InvoiceTypeReceiptAsIsDynamic    InvoiceType = "13.5" // Παραστατικά Οντότητας ως αναγράφονται από την ίδια (δυναμικό)
	InvoiceTypeReceiveCreditReceipt  InvoiceType = "13.6" // Πιστωτικό Στοιχείο Λιανικής
)

// Παραστατικά Εξαιρουμένων Οντοτήτων ημεδαπής / αλλοδαπής

//goland:noinspection GoUnusedConst
const (
	ExemptInvoiceGoodsInsideEU     InvoiceType = "14.1" // Τιμολόγιο / Ενδοκοινοτικές Αποκτήσεις
	ExemptInvoiceGoodsOutsideEU    InvoiceType = "14.2" // Τιμολόγιο / Αποκτήσεις Εμπορευμάτων από Τρίτες Χώρες
	ExemptInvoiceServicesInsideEU  InvoiceType = "14.3" // Τιμολόγιο / Ενδοκοινοτικές Παροχές Υπηρεσιών
	ExemptInvoiceServicesOutsideEU InvoiceType = "14.4" // Τιμολόγιο / Παροχές Υπηρεσιών από Τρίτες Χώρες
	ExemptInsuranceServices        InvoiceType = "14.5" // ΕΦΚΑ και λοιποί ασφαλιστικοί οργανισμοί
	ExemptInvoiceAsIsDynamic       InvoiceType = "14.6" // Παραστατικά Οντότητας ως αναγράφονται από την ίδια (δυναμικό)
	ExemptionInvoiceCredit         InvoiceType = "14.7" // Πιστωτικό ημεδαπής / αλλοδαπής
)

//goland:noinspection GoUnusedConst
const InvoiceTypeContractExpense InvoiceType = "15.1" // Συμβόλαιο Έξοδο
//goland:noinspection GoUnusedConst
const InvoiceTypeRentExpense InvoiceType = "16.1" // Ενοίκια - Έξοδο

// Εγγραφές Οντότητας

//goland:noinspection GoUnusedConst
const (
	InvoiceTypePayroll                               InvoiceType = "17.1" // Μισθοδοσία
	InvoiceTypeAmortization                          InvoiceType = "17.2" // Αποσβέσεις
	InvoiceTypeMiscIncomeRegistrationsAccountingBase InvoiceType = "17.3" // Λοιπές Εγγραφές Τακτοποίησης Εσόδων - Λογιστική Βάση
	InvoiceTypeMiscIncomeRegistrationsTaxBase        InvoiceType = "17.4" // Λοιπές Εγγραφές Τακτοποίησης Εσόδων - Φορολογική Βάση
	InvoiceTypeMiscExpenseRegistrationAccountingBase InvoiceType = "17.5" // Λοιπές Εγγραφές Τακτοποίησης Εξόδων - Λογιστική Βάση
	InvoiceTypeMiscExpenseRegistrationTaxBase        InvoiceType = "17.6" // Λοιπές Εγγραφές Τακτοποίησης Εξόδων - Φορολογική Βάση
)

// endregion

// region Κατηγορία Αιτίας Εξαίρεσης ΦΠΑ

type VATExceptionReasonType uint

//goland:noinspection GoUnusedConst
const (
	Article2And3         VATExceptionReasonType = iota + 1 // Χωρίς ΦΠΑ – άρθρο 2 και 3 του	Κώδικα ΦΠΑ
	Article5                                               // Χωρίς ΦΠΑ – άρθρο 5 του Κώδικα ΦΠΑ
	Article13                                              // Χωρίς ΦΠΑ – άρθρο 13 του Κώδικα ΦΠΑ
	Article14                                              // Χωρίς ΦΠΑ – άρθρο 14 του Κώδικα ΦΠΑ
	Article16                                              // Χωρίς ΦΠΑ – άρθρο 16 του Κώδικα ΦΠΑ
	Article19                                              // Χωρίς ΦΠΑ – άρθρο 19 του Κώδικα ΦΠΑ
	Article22                                              // Χωρίς ΦΠΑ – άρθρο 22 του Κώδικα ΦΠΑ
	Article24                                              // Χωρίς ΦΠΑ – άρθρο 24 του Κώδικα ΦΠΑ
	Article25                                              // Χωρίς ΦΠΑ – άρθρο 25 του Κώδικα ΦΠΑ
	Article26                                              // Χωρίς ΦΠΑ – άρθρο 26 του Κώδικα ΦΠΑ
	Article27                                              // Χωρίς ΦΠΑ – άρθρο 27 του Κώδικα ΦΠΑ
	Article27ForShips                                      // Χωρίς ΦΠΑ – άρθρο 27 του Κώδικα ΦΠΑ για πλοία ανοιχτής θαλάσσης
	Article271cForShips                                    // Χωρίς ΦΠΑ – άρθρο 27 παρ. 1 εδ γ του Κώδικα ΦΠΑ για πλοία
	Article28                                              // Χωρίς ΦΠΑ – άρθρο 28 του Κώδικα ΦΠΑ
	Article39                                              // Χωρίς ΦΠΑ – άρθρο 39 του Κώδικα ΦΠΑ
	Article39a                                             // Χωρίς ΦΠΑ – άρθρο 39α του Κώδικα ΦΠΑ
	Article40                                              // Χωρίς ΦΠΑ – άρθρο 40 του Κώδικα ΦΠΑ
	Article41                                              // Χωρίς ΦΠΑ – άρθρο 41 του Κώδικα ΦΠΑ
	Article47                                              // Χωρίς ΦΠΑ – άρθρο 47 του Κώδικα ΦΠΑ
	Article43VatIncluded                                   // ΦΠΑ εμπεριεχόμενος - άρθρο 43 του	Κώδικα ΦΠΑ
	Article44VatIncluded                                   // ΦΠΑ εμπεριεχόμενος - άρθρο 44 του	Κώδικα ΦΠΑ
	Article45VatIncluded                                   // ΦΠΑ εμπεριεχόμενος - άρθρο 45 του	Κώδικα ΦΠΑ
	Article46VatIncluded                                   // ΦΠΑ εμπεριεχόμενος - άρθρο 46 του	Κώδικα ΦΠΑ
	Article6                                               // Χωρίς ΦΠΑ – άρθρο 6 του Κώδικα ΦΠΑ
	Pol1029of1995                                          // Χωρίς ΦΠΑ - ΠΟΛ.1029/1995
	Pol1167Of2015                                          // Χωρίς ΦΠΑ - ΠΟΛ.1167/2015
	MiscException                                          // Λοιπές Εξαιρέσεις
	Article24CaseBPar1                                     // Χωρίς ΦΠΑ – άρθρο 24 περίπτωση β παράγραφος 1 του Κώδικα ΦΠΑ
	Article47b                                             // Χωρίς ΦΠΑ – άρθρο 47β, του Κώδικα ΦΠΑ (OSS μη ενωσιακό καθεστώς)
	Article47c                                             // Χωρίς ΦΠΑ – άρθρο 47γ, του Κώδικα ΦΠΑ (OSS ενωσιακό καθεστώς)
	Article47d                                             // Χωρίς ΦΠΑ – άρθρο 47δ, του Κώδικα ΦΠΑ (OSS ενωσιακό καθεστώς)
)

// endregion

// region Κατηγορία Παρακρατούμενων Φόρων

type WithholdingTaxCategoryType uint

//goland:noinspection GoUnusedConst
const (
	WithHoldingTaxInterest                       WithholdingTaxCategoryType = iota + 1 // Τόκοι 15%
	WithHoldingTaxRights                                                               // Δικαιώματα 20%
	WithHoldingTaxManagementAdvisoryFees                                               // Αμοιβές Συμβουλών Διοίκησης 20%
	WithHoldingTaxConstruction                                                         // Τεχνικά Έργα 3%
	WithHoldingTaxGasAndTobacco                                                        // Υγρά καύσιμα και προϊόντα καπνοβιομηχανίας 1%
	WithHoldingTaxVariousGoods                                                         // Λοιπά Αγαθά 4%
	WithHoldingTaxServices                                                             // Παροχή Υπηρεσιών 8%
	WithHoldingTaxArchitectsDesignFees                                                 // Προκαταβλητέος Φόρος Αρχιτεκτόνων και Μηχανικών επί Συμβατικών Αμοιβών, για Εκπόνηση Μελετών και Σχεδίων 4%
	WithHoldingTaxArchitectsOtherFees                                                  // Προκαταβλητέος Φόρος Αρχιτεκτόνων και Μηχανικών επί Συμβατικών Αμοιβών, που αφορούν οποιασδήποτε άλλης φύσης έργα 10%
	WithHoldingTaxLayerFees                                                            // Προκαταβλητέος Φόρος στις Αμοιβές Δικηγόρων 15%
	WithHoldingTaxFMY                                                                  // Παρακράτηση Φόρου Μισθωτών Υπηρεσιών παρ. 1 αρ. 15 ν. 4172/2013
	WithHoldingTaxFMYCase2MerchantMarineOfficers                                       // Παρακράτηση Φόρου Μισθωτών Υπηρεσιών παρ. 2 αρ. 15 ν. 4172/2013 Αξιωματικών Εμπορικού Ναυτικού
	WithHoldingTaxFMYJuniorCrewNave                                                    // Παρακράτηση Φόρου Μισθωτών Υπηρεσιών παρ. 2 αρ. 15 ν. 4172/2013 Κατώτερο Πλήρωμα Εμπορικού Ναυτικού
	WithHoldingTaxSolidarityCompensations                                              // Παρακράτηση Ειδικής Εισφοράς Αλληλεγγύης
	WithHoldingTaxCompensationTaxFromTermination                                       // Παρακράτηση Φόρου Αποζημίωσης λόγω Διακοπής Σχέσης Εργασίας
	WithHoldingTaxNoDoubleTaxing                                                       // Παρακρατήσεις συναλλαγών αλλοδαπής βάσει συμβάσεων αποφυγής διπλής φορολογίας
	WithHoldingTaxMiscWithholdingTaxes                                                 // Λοιπές Παρακρατήσεις
	WithHoldingTaxTaxDividends                                                         // Φόρος Μερισμάτων 5%
)

// endregion

// region Κατηγορία Λοιπών Φόρων

type MiscTaxCategoryType uint

//goland:noinspection GoUnusedConst
const (
	Fire15Percent                        MiscTaxCategoryType = iota + 1 // Ασφάλιστρα κλάδου πυρός 20% - 15%
	Fire5Percent                                                        // Ασφάλιστρα κλάδου πυρός 20% - 5%
	LifeInsurance                                                       // Ασφάλιστρα κλάδου ζωής 4%
	VariousInsurance                                                    // Ασφάλιστρα λοιπών κλάδων 15%
	ExemptedInsuranceTax                                                // Aπαλλασσόμενα φόρου ασφαλίστρων 0%
	Hotel1OR2Stars                                                      // Φόρος Διαμονής Ξενοδοχείων 1 ή 2 αστέρων 0,5 ευρώ
	Hotel3Stars                                                         // Φόρος Διαμονής Ξενοδοχείων 3 αστέρων 1,5 ευρώ
	Hotel4Stars                                                         // Φόρος Διαμονής Ξενοδοχείων 4 αστέρων 3 ευρώ
	Hotel5Stars                                                         // Φόρος Διαμονής Ξενοδοχείων 5 αστέρων 4 ευρώ
	RoomsToLet                                                          // Φόρος Διαμονής Ενοικιαζομένων Δωματίων 0,5 ευρώ
	SpecialTvAdsTax                                                     // Ειδικός Φόρος Τηλεοπτικών Διαφημίσεων 5%
	LuxuryTax10PercentFromThirdCountries                                // Πολυτελή Είδη 10% από Τρίτες Χώρες
	LuxuryTax10PercentFromGreece                                        // Πολυτελή Είδη 10%, προ ΦΠΑ από Ελλάδα
	CasinoTickets                                                       // Δικαίωμα του Δημοσίου στα εισιτήρια των καζίνο 80%
	FireInsuranceTaxes20Percent                                         // Ασφάλιστρα κλάδου πυρός 20%
	CustomsTaxes                                                        // Δασμοί
	VariousTaxes                                                        // Λοιποί Φόροι
	ChargesOnVariousTaxes                                               // Επιβαρύνσεις Λοιπών Φόρων
	EFK                                                                 // Ειδικός Φόρος Κατανάλωσης
)

// endregion

// region Κατηγορία Συντελεστή Χαρτοσήμου

type PaperStampCategoryType uint

//goland:noinspection GoUnusedConst
const (
	Coefficient12Percent PaperStampCategoryType = iota + 1 // Συντελεστής 1.2%
	Coefficient24Percent                                   // Συντελεστής 2.4%
	Coefficient36Percent                                   // Συντελεστής 3.6%
	VariousCases                                           // Λοιπές Περιπτώσεις
)

// endregion

// region Κατηγορία Τελών

type FeeCategoriesType uint

//goland:noinspection GoUnusedConst
const (
	MonthlyBillsUpto50Euros    FeeCategoriesType = iota + 1 // Για μηνιαίο λογαριασμό μέχρι και 50 ευρώ 12%
	MonthlyBillsUpto100Euros                                // Για μηνιαίο λογαριασμό από 50,01 μέχρι και 100 ευρώ 15%
	MonthlyBillsUpto150Euros                                // Για μηνιαίο λογαριασμό από 100,01 μέχρι και 150 ευρώ 18%
	MonthlyBillAbove150                                     // Για μηνιαίο λογαριασμό από 150,01 ευρώ και άνω 20%
	CartMobilePhone12Percent                                // Τέλος καρτοκινητής επί της αξίας του χρόνου ομιλίας (12%)
	SubscriptionTV                                          // Τέλος στη συνδρομητική τηλεόραση 10%
	Telephony                                               // Τέλος συνδρομητών σταθερής τηλεφωνίας 5%
	PlasticBags                                             // Περιβαλλοντικό Τέλος & πλαστικής σακούλας 0,07 ευρώ ανά τεμάχιο
	Dakoktonia2Percent                                      // Εισφορά δακοκτονίας 2%
	MiscFees                                                // Λοιπά τέλη
	FeesOnMiscTaxes                                         // Τέλη Λοιπών Φόρων
	Dakoktonia                                              // Εισφορά δακοκτονίας
	MonthlyBill                                             // Για μηνιαίο λογαριασμό κάθε σύνδεσης (10%)
	CartMobilePhone10Percent                                // Τέλος καρτοκινητής επί της αξίας του χρόνου ομιλίας (2%)
	CartMobilePhoneForYoung                                 // Τέλος κινητής και καρτοκινητής για φυσικά πρόσωπα ηλικίας 15 έως	και 29 ετών (0%)
	PlasticFee                                              // Εισφορά προστασίας περιβάλλοντος πλαστικών προϊόντων 0,04 λεπτά ανά τεμάχιο
	RecyclingFee                                            // Τέλος ανακύκλωσης 0,08 λεπτά ανά τεμάχιο
	TransientAccommodationFees                              // Τέλος διαμονής παρεπιδημούντων
	FeeOnRestaurants                                        // Tέλος επί των ακαθάριστων εσόδων των εστιατορίων και συναφών καταστημάτων
	FeeOnEntertainmentServices                              // Τέλος επί των ακαθάριστων εσόδων των κέντρων διασκέδασης
	FeeOnCasinoProfits                                      // Τέλος επί των ακαθάριστων εσόδων των καζίνο
	FeeOnVariousProfits                                     // Τέλος επί των ακαθάριστων εσόδων των λοιπών επιχειρήσεων
)

// endregion

// region Κατηγορία Οντότητας

type EntityCategory uint

//goland:noinspection GoUnusedConst
const (
	TaxRepresentative EntityCategory = iota + 1 // Φορολογικός Εκπρόσωπος
	Mediator                                    // Διαμεσολαβητής
)

// endregion

// region Κατηγορία ΦΠΑ

type InvoiceVATCategory uint

const (
	InvoiceVAT24Percent InvoiceVATCategory = iota + 1 // Κανονικό Τιμολόγιο 24%
	InvoiceVAT13Percent                               // Κανονικό Τιμολόγιο 13%
	InvoiceVAT6Percent                                // Κανονικό Τιμολόγιο 6%
	InvoiceVAT17Percent                               // Μειωμένο Τιμολόγιο 17%
	InvoiceVAT9Percent                                // Μειωμένο Τιμολόγιο 9%
	InvoiceVAT4Percent                                // Μειωμένο Τιμολόγιο 4%
	InvoiceVAT0Percent                                // Μηδενικό Τιμολόγιο 0%
	InvoiceVATExempt                                  // Εγγραφές χωρίς ΦΠΑ (πχ Μισθοδοσία, Αποσβέσεις)
)

func (i InvoiceVATCategory) CalculateVAT(amount float64) float64 {
	vatAmount := float64(0)

	switch i {
	case InvoiceVAT24Percent:
		vatAmount = amount * 0.24
	case InvoiceVAT13Percent:
		vatAmount = amount * 0.13
	case InvoiceVAT6Percent:
		vatAmount = amount * 0.6
	case InvoiceVAT0Percent:
		vatAmount = 0
	case InvoiceVATExempt:
		vatAmount = 0
	case InvoiceVAT17Percent:
		vatAmount = amount * 0.17
	case InvoiceVAT9Percent:
		vatAmount = amount * 0.9
	case InvoiceVAT4Percent:
		vatAmount = amount * 0.4
	default:
		vatAmount = 0
	}
	// round money down to cents
	return math.Round(vatAmount*100) / 100
}

// endregion

// region Κωδικοί Καυσίμων

type FuelCode uint

//goland:noinspection GoUnusedConst
const (
	Benzine95RON           FuelCode = iota + 10 // Βενζίνη 95 RON
	Benzine95RONPlus                            // Βενζίνη 95 RON Plus
	Benzine100RON                               // Βενζίνη 100 RON
	BenzineLRPG                                 // Βενζίνη LRP
	BenzineAirplane                             // Βενζίνη Αεροπλάνου
	BenzineSpecialAirplane                      // Ειδικό καύσιμο αεριωθουμένων
)

//goland:noinspection GoUnusedConst
const (
	Diesel     FuelCode = iota + 20 // Diesel
	DieselPlus                      // Diesel Plus
)

//goland:noinspection GoUnusedConst
const (
	DieselHeatNN      FuelCode = iota + 30 // Diesel Heat nn
	DieselHeatPremium                      // Diesel Heat premium
	DieselLight                            // Diesel Light
	DieselOtherUses                        // Diesel άλλων χρήσεων
	DieselMarine                           // Diesel Ναυτιλίας
	Kerosene                               // Κηροζίνη JP1
	KeroseneOtherUses                      // Κηροζίνη άλλων χρήσεων
	Mazout                                 // Μαζούτ
	MazoutMarine                           // Μαζούτ Ναυτιλίας
)

// endregion

// region Κωδικός Τύπου Χαρακτηρισμού Εσόδων

type IncomeClassificationTypeStringType string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	IE_106  IncomeClassificationTypeStringType = "E3_106" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Εμπορεύματα
	IE3_205 IncomeClassificationTypeStringType = "E3_205" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά
	IE3_210 IncomeClassificationTypeStringType = "E3_210" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη
	IE3_305 IncomeClassificationTypeStringType = "E3_305" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις – Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά
	IE3_310 IncomeClassificationTypeStringType = "E3_310" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη
	IE3_318 IncomeClassificationTypeStringType = "E3_318" // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Έξοδα παραγωγής

	IE3_561_001 IncomeClassificationTypeStringType = "E3_561_001" // Πωλήσεις αγαθών και υπηρεσιών Χονδρικές - Επιτηδευματιών
	IE3_561_002 IncomeClassificationTypeStringType = "E3_561_002" // Πωλήσεις αγαθών και υπηρεσιών Χονδρικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)
	IE3_561_003 IncomeClassificationTypeStringType = "E3_561_003" // Πωλήσεις αγαθών και υπηρεσιών Λιανικές - Ιδιωτική Πελατεία
	IE3_561_004 IncomeClassificationTypeStringType = "E3_561_004" // Πωλήσεις αγαθών και υπηρεσιών Λιανικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)
	IE3_561_005 IncomeClassificationTypeStringType = "E3_561_005" // Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Ενδοκοινοτικές
	IE3_561_006 IncomeClassificationTypeStringType = "E3_561_006" // Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Τρίτες Χώρες
	IE3_561_007 IncomeClassificationTypeStringType = "E3_561_007" // Πωλήσεις αγαθών και υπηρεσιών Λοιπά

	IE3_562 IncomeClassificationTypeStringType = "E3_562" // Λοιπά συνήθη έσοδα
	IE3_563 IncomeClassificationTypeStringType = "E3_563" // Πιστωτικοί τόκοι και συναφή έσοδα
	IE3_564 IncomeClassificationTypeStringType = "E3_564" // Πιστωτικές συναλλαγματικές διαφορές
	IE3_565 IncomeClassificationTypeStringType = "E3_565" // Έσοδα συμμετοχών
	IE3_566 IncomeClassificationTypeStringType = "E3_566" // Κέρδη από διάθεση μη κυκλοφορούντων περιουσιακών στοιχείων
	IE3_567 IncomeClassificationTypeStringType = "E3_567" // Κέρδη από αναστροφή προβλέψεων και απομειώσεων
	IE3_568 IncomeClassificationTypeStringType = "E3_568" // Κέρδη από επιμέτρηση στην εύλογη αξία
	IE3_570 IncomeClassificationTypeStringType = "E3_570" // Ασυνήθη έσοδα και κέρδη
	IE3_595 IncomeClassificationTypeStringType = "E3_595" // Έξοδα σε ιδιοπαραγωγή
	IE3_596 IncomeClassificationTypeStringType = "E3_596" // Επιδοτήσεις - Επιχορηγήσεις
	IE3_597 IncomeClassificationTypeStringType = "E3_597" // Επιδοτήσεις - Επιχορηγήσεις για επενδυτικούς σκοπούς - κάλυψη δαπανών

	IE3_880_001 IncomeClassificationTypeStringType = "E3_880_001" // Πωλήσεις Παγίων Χονδρικές
	IE3_880_002 IncomeClassificationTypeStringType = "E3_880_002" // Πωλήσεις Παγίων Λιανικές
	IE3_880_003 IncomeClassificationTypeStringType = "E3_880_003" //  Πωλήσεις Παγίων Εξωτερικού Ενδοκοινοτικές
	IE3_880_004 IncomeClassificationTypeStringType = "E3_880_004" // Πωλήσεις Παγίων Εξωτερικού Τρίτες Χώρες
	IE3_881_001 IncomeClassificationTypeStringType = "E3_881_001" // Πωλήσεις για λογ/σμο Τρίτων Χονδρικές
	IE3_881_002 IncomeClassificationTypeStringType = "E3_881_002" //  Πωλήσεις για λογ/σμο Τρίτων Λιανικές
	IE3_881_003 IncomeClassificationTypeStringType = "E3_881_003" // Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Ενδοκοινοτικές
	IE3_881_004 IncomeClassificationTypeStringType = "E3_881_004" // Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Τρίτες Χώρες
	IE3_598_001 IncomeClassificationTypeStringType = "E3_598_001" // Πωλήσεις αγαθών που υπάγονται σε ΕΦΚ
	IE3_598_003 IncomeClassificationTypeStringType = "E3_598_003" //  Πωλήσεις για λογαριασμό αγροτών μέσω αγροτικού συνεταιρισμού κ.λ.π
)

// endregion

// region Κωδικός Κατηγορίας Χαρακτηρισμού Εσόδων

type IncomeClassificationCategoryStringType string

//goland:noinspection GoUnusedConst, GoSnakeCaseUsage
const (
	ICategory1_1  IncomeClassificationCategoryStringType = "category1_1"  //  Έσοδα από Πώληση Εμπορευμάτων (+) / (-)
	ICategory1_2  IncomeClassificationCategoryStringType = "category1_2"  // Έσοδα από Πώληση Προϊόντων (+) / (-)
	ICategory1_3  IncomeClassificationCategoryStringType = "category1_3"  // Έσοδα από Παροχή Υπηρεσιών (+) / (-)
	ICategory1_4  IncomeClassificationCategoryStringType = "category1_4"  // Έσοδα από Πώληση Παγίων (+) / (-)
	ICategory1_5  IncomeClassificationCategoryStringType = "category1_5"  // Λοιπά Έσοδα/ Κέρδη (+) / (-)
	ICategory1_6  IncomeClassificationCategoryStringType = "category1_6"  // Αυτοπαραδόσεις / Ιδιοχρησιμοποιήσεις (+) / (-)
	ICategory1_7  IncomeClassificationCategoryStringType = "category1_7"  // Έσοδα για λ/σμο τρίτων (+) / (-)
	ICategory1_8  IncomeClassificationCategoryStringType = "category1_8"  // Έσοδα προηγούμενων χρήσεων (+)/ (-)
	ICategory1_9  IncomeClassificationCategoryStringType = "category1_9"  // Έσοδα επομένων χρήσεων (+) / (-)
	ICategory1_10 IncomeClassificationCategoryStringType = "category1_10" // Λοιπές Εγγραφές Τακτοποίησης Εσόδων (+) / (-)
	ICategory1_95 IncomeClassificationCategoryStringType = "category1_95" // Λοιπά Πληροφοριακά Στοιχεία Εσόδων (+) / (-)
)

// endregion

// region Κωδικός Τύπου Χαρακτηρισμού Εξόδων

type ExpenseClassificationTypeStringType string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	E3_101     ExpenseClassificationTypeStringType = "E3_101"     // Εμπορεύματα έναρξης
	E3_102_001 ExpenseClassificationTypeStringType = "E3_102_001" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)/Χονδρικές
	E3_102_002 ExpenseClassificationTypeStringType = "E3_102_002" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)/Λιανικές
	E3_102_003 ExpenseClassificationTypeStringType = "E3_102_003" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)/Αγαθών του άρθρου 39α παρ.5 του Κώδικα Φ.Π.Α. (ν.2859/2000)
	E3_102_004 ExpenseClassificationTypeStringType = "E3_102_004" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)/Εξωτερικού Ενδοκοινοτικές
	E3_102_005 ExpenseClassificationTypeStringType = "E3_102_005" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)/Εξωτερικού Τρίτες Χώρες
	E3_102_006 ExpenseClassificationTypeStringType = "E3_102_006" // Αγορές εμπορευμάτων χρήσης (καθαρό ποσό)Λοιπά
	E3_104     ExpenseClassificationTypeStringType = "E3_104"     // Εμπορεύματα λήξης
	E3_201     ExpenseClassificationTypeStringType = "E3_201"     // Πρώτες ύλες και υλικά έναρξης/Παραγωγή
	E3_202_001 ExpenseClassificationTypeStringType = "E3_202_001" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Χονδρικές
	E3_202_002 ExpenseClassificationTypeStringType = "E3_202_002" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Λιανικές
	E3_202_003 ExpenseClassificationTypeStringType = "E3_202_003" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Εξωτερικού Ενδοκοινοτικές
	E3_202_004 ExpenseClassificationTypeStringType = "E3_202_004" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Εξωτερικού Τρίτες Χώρες
	E3_202_005 ExpenseClassificationTypeStringType = "E3_202_005" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Λοιπά
	E3_204     ExpenseClassificationTypeStringType = "E3_204"     // Αποθέματα λήξης πρώτων υλών και υλικών/Παραγωγή
	E3_207     ExpenseClassificationTypeStringType = "E3_207"     // Προϊόντα και παραγωγή σε εξέλιξη έναρξης/Παραγωγή
	E3_209     ExpenseClassificationTypeStringType = "E3_209"     // Προϊόντα και παραγωγή σε εξέλιξη λήξης/Παραγωγή
	E3_301     ExpenseClassificationTypeStringType = "E3_301"     // Πρώτες ύλες και υλικά έναρξης/Αγροτική
	E3_302_001 ExpenseClassificationTypeStringType = "E3_302_001" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Χονδρικές
	E3_302_002 ExpenseClassificationTypeStringType = "E3_302_002" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Λιανικές
	E3_302_003 ExpenseClassificationTypeStringType = "E3_302_003" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Εξωτερικού Ενδοκοινοτικές
	E3_302_004 ExpenseClassificationTypeStringType = "E3_302_004" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Εξωτερικού Τρίτες Χώρες
	E3_302_005 ExpenseClassificationTypeStringType = "E3_302_005" // Αγορές πρώτων υλών και υλικών χρήσης (καθαρό ποσό)/Λοιπά
	E3_304     ExpenseClassificationTypeStringType = "E3_304"     // Αποθέματα λήξης πρώτων υλών και υλικών/Αγροτική
	E3_307     ExpenseClassificationTypeStringType = "E3_307"     // Προϊόντα και παραγωγή σε εξέλιξη έναρξης/Αγροτική
	E3_309     ExpenseClassificationTypeStringType = "E3_309"     // Προϊόντα και παραγωγή σε εξέλιξη λήξης/Αγροτική
	E3_312     ExpenseClassificationTypeStringType = "E3_312"     // Αποθέματα έναρξης (ζώων - φυτών)
	E3_313_001 ExpenseClassificationTypeStringType = "E3_313_001" // Αγορές ζώων - φυτών (καθαρό ποσό)/Χονδρικές
	E3_313_002 ExpenseClassificationTypeStringType = "E3_313_002" // Αγορές ζώων - φυτών (καθαρό ποσό)/Λιανικές
	E3_313_003 ExpenseClassificationTypeStringType = "E3_313_003" // Αγορές ζώων - φυτών (καθαρό ποσό)/Εξωτερικού Ενδοκοινοτικές
	E3_313_004 ExpenseClassificationTypeStringType = "E3_313_004" // Αγορές ζώων - φυτών (καθαρό ποσό)/Εξωτερικού Τρίτες Χώρες
	E3_313_005 ExpenseClassificationTypeStringType = "E3_313_005" // Αγορές ζώων - φυτών (καθαρό ποσό)/Λοιπά
	E3_315     ExpenseClassificationTypeStringType = "E3_315"     // Αποθέματα τέλους (ζώων - φυτών)/Αγροτική
	E3_581_001 ExpenseClassificationTypeStringType = "E3_581_001" // Παροχές σε εργαζόμενους/Μικτές αποδοχές
	E3_581_002 ExpenseClassificationTypeStringType = "E3_581_002" // Παροχές σε εργαζόμενους/Εργοδοτικές εισφορές
	E3_581_003 ExpenseClassificationTypeStringType = "E3_581_003" // Παροχές σε εργαζόμενους/Λοιπές παροχές
	E3_582     ExpenseClassificationTypeStringType = "E3_582"     // Ζημιές επιμέτρησης περιουσιακών στοιχείων
	E3_583     ExpenseClassificationTypeStringType = "E3_583"     // Χρεωστικές συναλλαγματικές διαφορές
	E3_584     ExpenseClassificationTypeStringType = "E3_584"     // Ζημιές από διάθεση-απόσυρση μη κυκλοφορούντων περιουσιακών στοιχείων
	E3_585_001 ExpenseClassificationTypeStringType = "E3_585_001" // Προμήθειες διαχείρισης ημεδαπής - αλλοδαπής (management fees)
	E3_585_002 ExpenseClassificationTypeStringType = "E3_585_002" // Δαπάνες από συνδεδεμένες επιχειρήσεις
	E3_585_003 ExpenseClassificationTypeStringType = "E3_585_003" // Δαπάνες από μη συνεργαζόμενα κράτη ή από κράτη με προνομιακό φορολογικό καθεστώς
	E3_585_004 ExpenseClassificationTypeStringType = "E3_585_004" // Δαπάνες για ενημερωτικές ημερίδες
	E3_585_005 ExpenseClassificationTypeStringType = "E3_585_005" // Έξοδα υποδοχής και φιλοξενίας
	E3_585_006 ExpenseClassificationTypeStringType = "E3_585_006" // Έξοδα ταξιδιών
	E3_585_007 ExpenseClassificationTypeStringType = "E3_585_007" // Ασφαλιστικές Εισφορές Αυτοαπασχολούμενων
	E3_585_008 ExpenseClassificationTypeStringType = "E3_585_008" // Έξοδα και προμήθειες παραγγελιοδόχου για λογαριασμό αγροτών
	E3_585_009 ExpenseClassificationTypeStringType = "E3_585_009" // Λοιπές Αμοιβές για υπηρεσίες ημεδαπής
	E3_585_010 ExpenseClassificationTypeStringType = "E3_585_010" // Λοιπές Αμοιβές για υπηρεσίες αλλοδαπής
	E3_585_011 ExpenseClassificationTypeStringType = "E3_585_011" // Ενέργεια
	E3_585_012 ExpenseClassificationTypeStringType = "E3_585_012" // Ύδρευση
	E3_585_013 ExpenseClassificationTypeStringType = "E3_585_013" // Τηλεπικοινωνίες
	E3_585_014 ExpenseClassificationTypeStringType = "E3_585_014" // Ενοίκια
	E3_585_015 ExpenseClassificationTypeStringType = "E3_585_015" // Διαφήμιση και προβολή
	E3_585_016 ExpenseClassificationTypeStringType = "E3_585_016" // Λοιπά έξοδα
	E3_586     ExpenseClassificationTypeStringType = "E3_586"     // Χρεωστικοί τόκοι και συναφή έξοδα
	E3_587     ExpenseClassificationTypeStringType = "E3_587"     // Αποσβέσεις
	E3_588     ExpenseClassificationTypeStringType = "E3_588"     // Ασυνήθη έξοδα, ζημιές και πρόστιμα
	E3_589     ExpenseClassificationTypeStringType = "E3_589"     // Προβλέψεις (εκτός από προβλέψεις για το προσωπικό)
	E3_882_001 ExpenseClassificationTypeStringType = "E3_882_001" // Αγορές ενσώματων παγίων χρήσης/Χονδρικές
	E3_882_002 ExpenseClassificationTypeStringType = "E3_882_002" // Αγορές ενσώματων παγίων χρήσης/Λιανικές
	E3_882_003 ExpenseClassificationTypeStringType = "E3_882_003" // Αγορές ενσώματων παγίων χρήσης/Εξωτερικού Ενδοκοινοτικές
	E3_882_004 ExpenseClassificationTypeStringType = "E3_882_004" // Αγορές ενσώματων παγίων χρήσης/Εξωτερικού Τρίτες Χώρες
	E3_883_001 ExpenseClassificationTypeStringType = "E3_883_001" // Αγορές μη ενσώματων παγίων χρήσης/Χονδρικές
	E3_883_002 ExpenseClassificationTypeStringType = "E3_883_002" // Αγορές μη ενσώματων παγίων χρήσης/Λιανικές
	E3_883_003 ExpenseClassificationTypeStringType = "E3_883_003" // Αγορές μη ενσώματων παγίων χρήσης/Εξωτερικού Ενδοκοινοτικές
	E3_883_004 ExpenseClassificationTypeStringType = "E3_883_004" // Αγορές μη ενσώματων παγίων χρήσης/Εξωτερικού Τρίτες Χώρες
	VAT_361    ExpenseClassificationTypeStringType = "VAT_361"    // Αγορές & δαπάνες στο εσωτερικό της χώρας
	VAT_362    ExpenseClassificationTypeStringType = "VAT_362"    // Αγορές & εισαγωγές επενδ. Αγαθών (πάγια)
	VAT_363    ExpenseClassificationTypeStringType = "VAT_363"    // Λοιπές εισαγωγές εκτός επενδ. Αγαθών (πάγια)
	VAT_364    ExpenseClassificationTypeStringType = "VAT_364"    // Ενδοκοινοτικές αποκτήσεις αγαθών
	VAT_365    ExpenseClassificationTypeStringType = "VAT_365"    // Ενδοκοινοτικές λήψεις υπηρεσιών άρθρο 14.2.α
	VAT_366    ExpenseClassificationTypeStringType = "VAT_366"    // Λοιπές πράξεις λήπτη
	E3_103     ExpenseClassificationTypeStringType = "E3_103"     // Απομείωση εμπορευμάτων
	E3_203     ExpenseClassificationTypeStringType = "E3_203"     // Απομείωση πρώτων υλών και υλικών
	E3_303     ExpenseClassificationTypeStringType = "E3_303"     // Απομείωση πρώτων υλών και υλικών
	E3_208     ExpenseClassificationTypeStringType = "E3_208"     // Απομείωση προϊόντων και παραγωγής σε εξέλιξη
	E3_308     ExpenseClassificationTypeStringType = "E3_308"     // Απομείωση προϊόντων και παραγωγής σε εξέλιξη
	E3_314     ExpenseClassificationTypeStringType = "E3_314"     // Απομείωση ζώων - φυτών – εμπορευμάτων
	E3_106     ExpenseClassificationTypeStringType = "E3_106"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_205     ExpenseClassificationTypeStringType = "E3_205"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_305     ExpenseClassificationTypeStringType = "E3_305"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_210     ExpenseClassificationTypeStringType = "E3_210"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_310     ExpenseClassificationTypeStringType = "E3_310"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_318     ExpenseClassificationTypeStringType = "E3_318"     // Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων
	E3_598_002 ExpenseClassificationTypeStringType = "E3_598_002" // Αγορές αγαθών που υπάγονται σε ΕΦΚΑ
)

// endregion

// region Κωδικός Κατηγορίας Χαρακτηρισμού Εξόδων

type ExpensesClassificationCategoryStringType string

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	ECategory2_1  ExpensesClassificationCategoryStringType = "category2_1"  // Αγορές Εμπορευμάτων (-) / (+)
	ECategory2_2  ExpensesClassificationCategoryStringType = "category2_2"  // Αγορές Α'-Β' Υλών (-) / (+)
	ECategory2_3  ExpensesClassificationCategoryStringType = "category2_3"  // Λήψη Υπηρεσιών (-) / (+)
	ECategory2_4  ExpensesClassificationCategoryStringType = "category2_4"  // Γενικά Έξοδα με δικαίωμα έκπτωσης ΦΠΑ (-) / (+)
	ECategory2_5  ExpensesClassificationCategoryStringType = "category2_5"  // Γενικά Έξοδα χωρίς δικαίωμα έκπτωσης ΦΠΑ (-) / (+)
	ECategory2_6  ExpensesClassificationCategoryStringType = "category2_6"  // Αμοιβές και Παροχές προσωπικού (-) / (+)
	ECategory2_7  ExpensesClassificationCategoryStringType = "category2_7"  // Αγορές Παγίων (-) / (+)
	ECategory2_8  ExpensesClassificationCategoryStringType = "category2_8"  // Αποσβέσεις Παγίων (-) / (+)
	ECategory2_9  ExpensesClassificationCategoryStringType = "category2_9"  // Έξοδα για λ/σμο τρίτων (-) / (+)
	ECategory2_10 ExpensesClassificationCategoryStringType = "category2_10" // Έξοδα προηγούμενων χρήσεων (-) / (+)
	ECategory2_11 ExpensesClassificationCategoryStringType = "category2_11" // Έξοδα επομένων χρήσεων (-) / (+)
	ECategory2_12 ExpensesClassificationCategoryStringType = "category2_12" // Λοιπές Εγγραφές Τακτοποίησης Εξόδων (-) / (+)
	ECategory2_13 ExpensesClassificationCategoryStringType = "category2_13" // Αποθέματα Έναρξης Περιόδου (-) / (+)
	ECategory2_14 ExpensesClassificationCategoryStringType = "category2_14" // Αποθέματα Λήξης Περιόδου (-) / (+)
	ECategory2_95 ExpensesClassificationCategoryStringType = "category2_95" // Λοιπά Πληροφοριακά Στοιχεία Εξόδων (-) / (+)
)

// endregion
