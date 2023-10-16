package vatService

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrCannotReachService = Error("δεν μπορέσαμε να συνδεθούμε με την υπηρεσία")
	ErrInvalidVAT         = Error("λανθασμένη μορφή ΑΦΜ")
	ErrInvalidCredentials = Error("το username και το password δεν μπορεί να είναι μικρότερα από 6 χαρακτήρες")
	ErrGeneral            = Error("Γενικό σφάλμα")
)

var ServiceErrors = map[string]Error{
	"RG_WS_PUBLIC_AFM_CALLED_BY_BLOCKED":             Error("Ο χρήστης που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."),
	"RG_WS_PUBLIC_AFM_CALLED_BY_NOT_FOUND":           Error("Ο Α.Φ.Μ. για τον οποίο γίνεται η κλήση δε βρέθηκε στους έγκυρους Α.Φ.Μ του Μητρώου TAXIS."),
	"RG_WS_PUBLIC_EPIT_NF":                           Error("O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν ανήκει και δεν ανήκε ποτέ σε νομικό πρόσωπο, νομική οντότητα, ή φυσικό πρόσωπο με εισόδημα από επιχειρηματική δραστηριότητα."),
	"RG_WS_PUBLIC_FAILURES_TOLERATED_EXCEEDED":       Error("Υπέρβαση μέγιστου επιτρεπτού ορίου πρόσφατων αποτυχημένων κλήσεων. Προσπαθήστε εκ νέου σε μερικές ώρες."),
	"RG_WS_PUBLIC_MAX_DAILY_USERNAME_CALLS_EXCEEDED": Error("Υπέρβαση μέγιστου επιτρεπτού ορίου ημερήσιων κλήσεων ανά χρήστη (ανεξαρτήτως εξουσιοδοτήσεων)."),
	"RG_WS_PUBLIC_MONTHLY_LIMIT_EXCEEDED":            Error("Υπέρβαση του Μέγιστου Επιτρεπτού Μηνιαίου Ορίου Κλήσεων."),
	"RG_WS_PUBLIC_MSG_TO_TAXISNET_ERROR":             Error("Δημιουργήθηκε πρόβλημα κατά την ενημέρωση των εισερχόμενων μηνυμάτων στο MyTAXIS net."),
	"RG_WS_PUBLIC_NO_INPUT_PARAMETERS":               Error("Δε δόθηκαν υποχρεωτικές παράμετροι εισόδου για την κλήση της υπηρεσίας."),
	"RG_WS_PUBLIC_SERVICE_NOT_ACTIVE":                Error("Η υπηρεσία δεν είναι ενεργή."),
	"RG_WS_PUBLIC_TAXPAYER_NF":                       Error("O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δε βρέθηκε στους έγκυρους Α.Φ.Μ του Μητρώου TAXIS."),
	"RG_WS_PUBLIC_TOKEN_AFM_BLOCKED":                 Error("Ο χρήστης (ή ο εξουσιοδοτημένος τρίτος) που καλεί την υπηρεσία έχει προσωρινά αποκλειστεί από τη χρήση της."),
	"RG_WS_PUBLIC_TOKEN_AFM_NOT_AUTHORIZED":          Error("Ο τρέχον χρήστης δεν έχει εξουσιοδοτηθεί από τον Α.Φ.Μ για χρήση της υπηρεσίας."),
	"RG_WS_PUBLIC_TOKEN_AFM_NOT_FOUND":               Error("Ο Α.Φ.Μ. του τρέχοντος χρήστη δε βρέθηκε στους έγκυρους Α.Φ.Μ του Μητρώου TAXIS."),
	"RG_WS_PUBLIC_TOKEN_AFM_NOT_REGISTERED":          Error("Ο τρέχον χρήστης δεν έχει εγγραφεί για χρήση της υπηρεσίας."),
	"RG_WS_PUBLIC_TOKEN_USERNAME_NOT_ACTIVE":         Error("Ο κωδικός χρήστη (username) που χρησιμοποιήθηκε έχει ανακληθεί."),
	"RG_WS_PUBLIC_TOKEN_USERNAME_NOT_AUTHENTICATED":  Error("Ο συνδυασμός χρήστη/κωδικού πρόσβασης που δόθηκε δεν είναι έγκυρος."),
	"RG_WS_PUBLIC_TOKEN_USERNAME_NOT_DEFINED":        Error("Δεν ορίσθηκε ο χρήστης που καλεί την υπηρεσία."),
	"RG_WS_PUBLIC_TOKEN_USERNAME_TOO_LONG":           Error("Διαπιστώθηκε υπέρβαση του μήκους του ονόματος του χρήστη (username) της υπηρεσίας"),
	"RG_WS_PUBLIC_WRONG_AFM":                         Error("O Α.Φ.Μ. για τον οποίο ζητούνται πληροφορίες δεν είναι έγκυρος."),
}
