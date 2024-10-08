package mydataInvoices

import (
	"errors"
	"time"
)

type RequestVatInfoParams struct {
	EntityVatNumber  string    // ΑΦΜ οντότητας
	DateFrom         time.Time // * Αρχή χρονικού διαστήματος αναζήτησης για την ημερομηνία έκδοσης (μορφή dd/MM/yyyy)
	DateTo           time.Time // * Λήξη χρονικού διαστήματος αναζήτησης για την ημερομηνία έκδοσης (μορφή dd/MM/yyyy)
	GroupedPerDay    bool      // Παράμετρος που δηλώνει εάν τα αποτελέσματα πρέπει να ομαδοποιηθούν ανά ημέρα. Δέχεται τιμές "true" ή "false"
	NextPartitionKey string    // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων, όταν GroupedPerDay=false
	NextRowKey       string    // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων, όταν GroupedPerDay=false
}

/*
1) Σε περίπτωση που τα αποτελέσματα αναζήτησης υπερβαίνουν το μέγιστο
επιτρεπτό όριο, ο χρήστης θα τα λάβει τμηματικά. Τα πεδία nextPartitionKey και
nextRowKey θα παρέχονται σε κάθε απόκριση και θα πρέπει να χρησιμοποιούνται
ως παράμετροι για την ανάκτηση του επόμενου συνόλου αποτελεσμάτων. Είναι
σημαντικό να σημειωθεί ότι εάν η παράμετρος GroupedPerDay έχει τιμή false, τότε
οι παράμετροι nextPartitionKey και nextRowKey δεν απαιτούνται και, αν δοθούν,
δεν θα λαμβάνονται υπόψη για την ανάκτηση των αποτελεσμάτων.
2) Σε περίπτωση που η παράμετρος entityVatNumber δεν παρέχεται, θα γίνει
αναζήτηση βάσει του ΑΦΜ του χρήστη που κάνει την κλήση.
3) Οι παράμετροι dateFrom και dateTo είναι υποχρεωτικές και πρέπει να δοθούν στη
σωστή μορφή (dd/MM/yyyy).
*/

func (p *RequestVatInfoParams) ToMap() (map[string]string, error) {
	result := make(map[string]string)
	// region required fields
	emptyTime := time.Time{}
	if p.DateFrom == emptyTime {
		return nil, errors.New("το πεδίο dateFrom είναι υποχρεωτικό")
	}
	if p.DateTo == emptyTime {
		return nil, errors.New("το πεδίο dateTo είναι υποχρεωτικό")
	}
	// endregion
	// region struct to map
	result["dateFrom"] = p.DateFrom.Format("02/01/2006")
	result["dateTo"] = p.DateTo.Format("02/01/2006")
	if p.GroupedPerDay {
		result["GroupedPerDay"] = "true"
	}
	if p.NextPartitionKey != "" && !p.GroupedPerDay {
		result["nextPartitionKey"] = p.NextPartitionKey
	}
	if p.NextRowKey != "" && !p.GroupedPerDay {
		result["nextRowKey"] = p.NextRowKey
	}
	// endregion
	return result, nil
}
