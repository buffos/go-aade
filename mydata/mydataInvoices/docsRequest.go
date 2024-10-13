package mydataInvoices

import (
	"errors"
	"time"

	"github.com/buffos/go-aade/mydata/mydatavalues"
)

type RequestDocsParams struct {
	Mark              string                    // * Μοναδικός αριθμός καταχώρησης
	EntityVatNumber   string                    // ΑΦΜ αναφοράς
	DateFrom          time.Time                 // Αρχή χρονικού διαστήματος αναζήτησης για την ημερομηνία έκδοσης
	DateTo            time.Time                 // Τέλος χρονικού διαστήματος αναζήτησης για την ημερομηνία έκδοσης
	ReceiverVatNumber string                    // ΑΦΜ παραλήπτη
	InvType           *mydatavalues.InvoiceType // Τύπος Παραστατικού
	MaxMark           string                    // Μέγιστος Αριθμός ΜΑΡΚ
	NextPartitionKey  string                    // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextRowKey        string                    // Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
}

func (p *RequestDocsParams) SetNextPartitionData(partitionKey, rowKey string) {
	p.NextPartitionKey = partitionKey
	p.NextRowKey = rowKey
}

func (p *RequestDocsParams) ToMap() (map[string]string, error) {
	result := make(map[string]string)
	// region required fields
	if p.Mark == "" {
		return nil, errors.New("το πεδίο mark είναι υποχρεωτικό")
	}
	// endregion
	// region struct to map
	result["mark"] = p.Mark
	if p.EntityVatNumber != "" {
		result["entityVatNumber"] = p.EntityVatNumber
	}
	emptyTime := time.Time{}
	if p.DateFrom != emptyTime {
		result["dateFrom"] = p.DateFrom.Format("02/01/2006")
	}
	if p.DateTo != emptyTime {
		result["dateTo"] = p.DateTo.Format("02/01/2006")
	}
	if p.ReceiverVatNumber != "" {
		result["receiverVatNumber"] = p.ReceiverVatNumber
	}
	if p.InvType != nil && *p.InvType != "" {
		result["invType"] = string(*p.InvType)
	}
	if p.MaxMark != "" {
		result["maxMark"] = p.MaxMark
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

type CancelInvoiceParams struct {
	Mark            string // * Μοναδικός αριθμός καταχώρησης
	EntityVatNumber string // ΑΦΜ αναφοράς
}

func (p *CancelInvoiceParams) ToMap() (map[string]string, error) {
	result := make(map[string]string)
	// region required fields
	if p.Mark == "" {
		return nil, errors.New("το πεδίο mark είναι υποχρεωτικό")
	}
	// endregion
	// region struct to map
	result["mark"] = p.Mark
	if p.EntityVatNumber != "" {
		result["entityVatNumber"] = p.EntityVatNumber
	}
	// endregion
	return result, nil
}
