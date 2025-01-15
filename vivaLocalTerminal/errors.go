package vivaLocalTerminal

import (
	"errors"
)

var errorMap = map[int]error{
	1000: errors.New("Transaction canceled by user"),
	1001: errors.New("Payments SDK is busy at the moment"),
	1002: errors.New("There are stored transactions that need to be voided"),
	1003: errors.New("Terminal timed out"),
	1004: errors.New("Terminal declined transaction"),
	1006: errors.New("Transaction declined by server"),
	1007: errors.New("Declined by card"),
	1008: errors.New("Non reversible transaction"),
	1009: errors.New("Invalid amount for void"),
	1010: errors.New("Invalid card"),
	1011: errors.New("No transactions found"),
	1012: errors.New("Transaction does not exist"),
	1013: errors.New("Could not initialize an ISV Sale. Please check your input parameters"),
	1014: errors.New("Refund is disabled merchant. AllowLiquidation role missing"),
	1016: errors.New("Transaction aborted"),
	1017: errors.New("Invalid amount for capture"),
	1018: errors.New("No parameters found for this merchant"),
	1019: errors.New("Try another card"),
	1020: errors.New("Insufficient funds"),
	1021: errors.New("Non reversible transaction"),
	1099: errors.New("Generic transaction error"),
	1100: errors.New("There is already a request in progress"),
	1101: errors.New("Could not fetch parameters for parentSessionId"),
	1102: errors.New("Transaction expired"),
	2000: errors.New("Location services disabled"),
	3000: errors.New("Terminal is not connected"),
	3001: errors.New("Terminal connection error"),
	3002: errors.New("Terminal connection time out"),
	3099: errors.New("Generic terminal error"),
	4000: errors.New("Network connection error"),
	5000: errors.New("Missing 'MOTO' capability"),
	5001: errors.New("Missing 'Instalments' capability"),
	5002: errors.New("Missing 'Preauth' capability"),
	5003: errors.New("Missing 'QR' capability"),
	5004: errors.New("Missing 'Void' capability"),
	5005: errors.New("Missing 'Tipping' capability"),
	5006: errors.New("Missing 'Sale' capability"),
	6000: errors.New("Wrong request parameters: (additional descriptive message)"),
	6102: errors.New("Missing ISV Role"),
	6201: errors.New("Only regular sale is allowed in SaF mode"),
	7001: errors.New("Amount exceeded SaF transaction limit"),
	7002: errors.New("Total sales amount during SaF exceeded total SaF limit"),
	7003: errors.New("SaF mode duration limit exceeded"),
	7004: errors.New("Mastercard, Visa or AMEX card is required"),
	7005: errors.New("Transaction type unsupported in SaF mode"),
}

func GetError(code int) error {
	err, ok := errorMap[code]
	if !ok {
		return errors.New("unknown error code")
	}
	return err
}
