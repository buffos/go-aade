package mydataInvoices

import (
	"github.com/buffos/go-aade/mydata/mydatavalues"
	"golang.org/x/exp/slices"
)

func AllowedIncomeCharacterisations(
	invType *mydatavalues.InvoiceType,
	iClType *mydatavalues.IncomeClassificationTypeStringType,
	iCategory *mydatavalues.IncomeClassificationCategoryStringType) bool {

	if invType == nil {
		return false
	}

	if iCategory == nil {
		return false
	}

	if *invType == mydatavalues.InvoiceTypeSales {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return *iClType == mydatavalues.IE3_880_001
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_564,
					mydatavalues.IE3_565,
					mydatavalues.IE3_566,
					mydatavalues.IE3_567,
					mydatavalues.IE3_568,
					mydatavalues.IE3_570,
					mydatavalues.IE3_561_002,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_596,
					mydatavalues.IE3_597,
				}, *iClType) || iClType == nil
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSalesInsideEU {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return *iClType == mydatavalues.IE3_880_003
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return *iClType == mydatavalues.IE3_880_003
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSalesOutsideEU {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return *iClType == mydatavalues.IE3_880_004
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return *iClType == mydatavalues.IE3_880_004
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSalesOnBehalfOf {
		switch *iCategory {
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSalesOnBehalfOfPayment {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSalesComplementary {
		// normally, we should fetch the correlated invoice and check the type
		// and apply the same constraints as the correlated invoice
		return true
	}

	if *invType == mydatavalues.InvoiceTypeServices {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_564,
					mydatavalues.IE3_565,
					mydatavalues.IE3_566,
					mydatavalues.IE3_567,
					mydatavalues.IE3_568,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeServicesInsideEU {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeServicesOutsideEU {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeServiceComplementary {
		// normally, we should fetch the correlated invoice and check the type
		// and apply the same constraints as the correlated invoice
		return true
	}

	if *invType == mydatavalues.InvoiceTypeOwnershipTitleNoObligationIssuer {
		// no checks
		return true
	}
	if *invType == mydatavalues.InvoiceTypeOwnershipTitleRefuseByIssuer {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return *iClType == mydatavalues.IE3_880_001
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_564,
					mydatavalues.IE3_565,
					mydatavalues.IE3_566,
					mydatavalues.IE3_567,
					mydatavalues.IE3_568,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_596,
					mydatavalues.IE3_597,
				}, *iClType) || iClType == nil
		default:
			return false
		}
	}

	if *invType == mydatavalues.InvoiceTypeCreditSalesWithReference {
		// normally, we should fetch the correlated invoice and check the type
		// and apply the same constraints as the correlated invoice
		return true
	}
	if *invType == mydatavalues.InvoiceTypeCreditSalesWithoutReference {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_880_001,
					mydatavalues.IE3_880_003,
					mydatavalues.IE3_880_004,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_564,
					mydatavalues.IE3_565,
					mydatavalues.IE3_566,
					mydatavalues.IE3_567,
					mydatavalues.IE3_568,
					mydatavalues.IE3_570,
					mydatavalues.IE3_561_002,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_596,
					mydatavalues.IE3_597,
				}, *iClType) || iClType == nil
		default:
			return false
		}
	}

	if *invType == mydatavalues.InvoiceTypeSelfDelivery {
		switch *iCategory {
		case mydatavalues.ICategory1_6:
			return *iClType == mydatavalues.IE3_595
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeSelfUsage {
		switch *iCategory {
		case mydatavalues.ICategory1_6:
			return *iClType == mydatavalues.IE3_595
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}

	if *invType == mydatavalues.InvoiceTypeContractIncome {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return *iClType == mydatavalues.IE3_880_001
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}

	if *invType == mydatavalues.InvoiceTypeRentIncome {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return *iClType == mydatavalues.IE3_561_001
		case mydatavalues.ICategory1_5:
			return *iClType == mydatavalues.IE3_562
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeReceiptOfAccommodationTax {
		switch *iCategory {
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}

	if *invType == mydatavalues.RetailReceipt {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_880_002,
					mydatavalues.IE3_880_003,
					mydatavalues.IE3_880_004,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.RetailServiceReceipt {
		switch *iCategory {
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.RetailSimplifiedInvoice {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_880_002,
					mydatavalues.IE3_880_003,
					mydatavalues.IE3_880_004,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.RetailCreditReceipt {
		switch *iCategory {
		case mydatavalues.ICategory1_1:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_2:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_3:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_563,
				}, *iClType)
		case mydatavalues.ICategory1_4:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_880_002,
					mydatavalues.IE3_880_003,
					mydatavalues.IE3_880_004,
					mydatavalues.IE3_561_007,
				}, *iClType)
		case mydatavalues.ICategory1_5:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.RetailReceiptOnBehalfOf {
		switch *iCategory {
		case mydatavalues.ICategory1_7:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
				}, *iClType)
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}

	if *invType == mydatavalues.InvoiceTypeMiscIncomeRegistrationsAccountingBase {
		switch *iCategory {
		case mydatavalues.ICategory1_8:
			return true
		case mydatavalues.ICategory1_9:
			return true
		case mydatavalues.ICategory1_10:
			return slices.Contains(
				[]mydatavalues.IncomeClassificationTypeStringType{
					mydatavalues.IE3_561_001,
					mydatavalues.IE3_561_002,
					mydatavalues.IE3_561_003,
					mydatavalues.IE3_561_004,
					mydatavalues.IE3_561_005,
					mydatavalues.IE3_561_006,
					mydatavalues.IE3_561_007,
					mydatavalues.IE3_562,
					mydatavalues.IE3_563,
					mydatavalues.IE3_595,
					mydatavalues.IE3_596,
					mydatavalues.IE3_597,
					mydatavalues.IE3_880_001,
					mydatavalues.IE3_880_002,
					mydatavalues.IE3_880_003,
					mydatavalues.IE3_880_004,
					mydatavalues.IE3_881_001,
					mydatavalues.IE3_881_002,
					mydatavalues.IE3_881_003,
					mydatavalues.IE3_881_004,
					mydatavalues.IE3_564,
					mydatavalues.IE3_565,
					mydatavalues.IE3_566,
					mydatavalues.IE3_567,
					mydatavalues.IE3_568,
					mydatavalues.IE3_570,
				}, *iClType)
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}
	if *invType == mydatavalues.InvoiceTypeMiscIncomeRegistrationsTaxBase {
		switch *iCategory {
		case mydatavalues.ICategory1_10:
			return true
		case mydatavalues.ICategory1_95:
			return true
		default:
			return false
		}
	}

	return false
}

func AllowedExpenseCharacterisations(
	invType *mydatavalues.InvoiceType,
	iClType *mydatavalues.ExpensesClassificationCategoryStringType,
	iCategory *mydatavalues.ExpensesClassificationCategoryStringType) bool {
	return false
}
