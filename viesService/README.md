# VIES VAT Validation

## Info

This is based on the official endpoint of EU. For more information look [here](https://ec.europa.eu/taxation_customs/vies/#/vat-validation)
The system no longer returns company names and other info. Just if the VAT is a valid VIES number.

## Usage

Usage is simple

```go
info, err := GetViesVatInfo(vatString)
```

if the vat string is not in the proper format you will get an error and the `info.Valid = false`.
If the vat string is invalid, but the correct format, you will get `info.Valid = false` but no error.
