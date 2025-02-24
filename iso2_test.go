package main

import (
	"testing"
)

func Test_AllCountriesISO2(t *testing.T) {
	if len(AllISO2Countries) != 249 {
		t.Errorf("AllISO2Countries should have 249 elements, but it has %d", len(AllISO2Countries))
	}
	if len(AllISO2Countries) != len(AllISO3Countries) {
		t.Errorf("AllISO2Countries and AllISO3Countries should have the same length, but they have %d and %d", len(AllISO2Countries), len(AllISO3Countries))
	}
	for i := 0; i < len(AllISO2Countries)-1; i++ {
		if AllISO2Countries[i] > AllISO2Countries[i+1] {
			t.Errorf("AllISO2Countries should be sorted, but %s is after %s", AllISO2Countries[i], AllISO2Countries[i+1])
		}
	}
	for i := 0; i < len(AllISO2Countries)-1; i++ {
		if AllISO2Countries[i] == AllISO2Countries[i+1] {
			t.Errorf("AllISO2Countries should have unique elements, but %s is repeated", AllISO2Countries[i])
		}
	}
}

func Test_ISO2CountryNaming(t *testing.T) {
	for _, iso2Country := range AllISO2Countries {
		value := string(iso2Country)
		if len(value) != 2 {
			t.Errorf("ISO2 country code %s should be exactly 2 characters long", value)
		}
		for _, char := range value {
			if char < 'A' || char > 'Z' {
				t.Errorf("ISO2 country code %s should only contain uppercase letters", value)
				break
			}
		}
	}
}
