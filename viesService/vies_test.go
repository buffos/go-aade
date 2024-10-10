package viesService

import (
	"strings"
	"testing"
)

func TestPrepareVatRequest(t *testing.T) {
	tests := []struct {
		vat      string
		includes string
		hasError bool
	}{
		{"", "", true},                       // Empty VAT number
		{"123456789", "", true},              // Invalid VAT number. No country code
		{"GR123456789", "", true},            // Invalid VAT number. Greek country code is EL.
		{"EL123456789", "123456789", false},  // Valid Greek VAT number
		{"DE123456789", "123456789", false},  // Valid German VAT number
		{"FR123456789", "123456789", false},  // Valid French VAT number
		{"INVALIDVAT", "", true},             // Invalid VAT format
		{"GR12345678901234567890", "", true}, // VAT number too long
	}

	for _, test := range tests {
		result, err := prepareVatRequest(test.vat)
		if (err != nil) != test.hasError {
			t.Errorf("prepareVatRequest(%q) error = %v, wantErr %v", test.vat, err, test.hasError)
		}
		if !strings.Contains(result, test.includes) {
			t.Errorf("prepareVatRequest(%q) = %q, want %q", test.vat, result, test.includes)
		}
	}
}

func TestGetViesVatInfo(t *testing.T) {
	testingVat := "DE814584193"
	info, err := GetViesVatInfo(testingVat)
	if err != nil {
		t.Errorf("GetViesVatInfo(%q) error = %v", testingVat, err.Error())
	}
	if info.VatNumber == "" {
		t.Fatalf("GetViesVatInfo(%q) returned nil", testingVat)
	}
	if info.CountryCode != "DE" {
		t.Errorf("GetViesVatInfo(%q) = %q, want %q", testingVat, info.CountryCode, "DE")
	}

	if info.VatNumber != "814584193" {
		t.Errorf("GetViesVatInfo(%q) = %q, want %q", testingVat, info.VatNumber, "814584193")
	}
	if info.Valid != true {
		t.Errorf("GetViesVatInfo(%q) = %t, want %t", testingVat, info.Valid, true)
	}
}

func TestGetViesVatInfoInvalid(t *testing.T) {
	testingVat := "DE123456789"
	info, err := GetViesVatInfo(testingVat)
	if err != nil {
		t.Errorf("GetViesVatInfo(%q) error = %v, wantErr %v", testingVat, err, false)
	}
	if info.Valid != false {
		t.Errorf("GetViesVatInfo(%q) = %t, want %t", testingVat, info.Valid, false)
	}
}

func TestGetViesVatInfoInvalidFormat(t *testing.T) {
	testingVat := "123456789"
	info, err := GetViesVatInfo(testingVat)
	if err == nil {
		t.Errorf("GetViesVatInfo(%q) error = %v, wantErr %v", testingVat, err, true)
	}
	if info.Valid != false {
		t.Errorf("GetViesVatInfo(%q) = %t, want %t", testingVat, info.Valid, false)
	}
}
