package mydataInvoices

import (
	"errors"
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"time"
)

type RequestMyIncomeParams struct {
	DateFrom         time.Time
	DateTo           time.Time
	CounterVatNumber string                   // ΑΦΜ αντισυμβαλλόμενου
	EntityVatNumber  string                   // ΑΦΜ αναφοράς
	InvType          mydatavalues.InvoiceType // Τύπος Παραστατικού
	NextPartitionKey string                   // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextRowKey       string                   // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
}

func (p *RequestMyIncomeParams) ToMap() (map[string]string, error) {
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
	if p.CounterVatNumber != "" {
		result["counterVatNumber"] = p.CounterVatNumber
	}
	if p.EntityVatNumber != "" {
		result["entityVatNumber"] = p.EntityVatNumber
	}
	if p.InvType != "" {
		result["invType"] = string(p.InvType)
	}
	if p.NextPartitionKey != "" {
		result["nextPartitionKey"] = p.NextPartitionKey
	}
	if p.NextRowKey != "" {
		result["nextRowKey"] = p.NextRowKey
	}
	// endregion

	return result, nil
}

type RequestMyExpensesParams struct {
	DateFrom         time.Time
	DateTo           time.Time
	CounterVatNumber string                   // ΑΦΜ αντισυμβαλλόμενου
	EntityVatNumber  string                   // ΑΦΜ αναφοράς
	InvType          mydatavalues.InvoiceType // Τύπος Παραστατικού
	NextPartitionKey string                   // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextRowKey       string                   // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
}

func (p *RequestMyExpensesParams) ToMap() (map[string]string, error) {
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
	if p.CounterVatNumber != "" {
		result["counterVatNumber"] = p.CounterVatNumber
	}
	if p.EntityVatNumber != "" {
		result["entityVatNumber"] = p.EntityVatNumber
	}
	if p.InvType != "" {
		result["invType"] = string(p.InvType)
	}
	if p.NextPartitionKey != "" {
		result["nextPartitionKey"] = p.NextPartitionKey
	}
	if p.NextRowKey != "" {
		result["nextRowKey"] = p.NextRowKey
	}
	// endregion

	return result, nil
}
