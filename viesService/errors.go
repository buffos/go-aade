package viesService

import "errors"

var (
	ErrorVatTooShort        = errors.New("VAT number is too short")
	ErrorInvalidCountryCode = errors.New("invalid country code")
)
