package countries

import (
	"testing"
)

func Test_AllCountriesISO3(t *testing.T) {
	if len(AllISO3Countries) != 249 {
		t.Errorf("AllISO3Countries should have 249 elements, but it has %d", len(AllISO3Countries))
	}
	if len(AllISO3Countries) != len(AllISO2Countries) {
		t.Errorf("AllISO3Countries and AllISO2Countries should have the same length, but they have %d and %d", len(AllISO3Countries), len(AllISO2Countries))
	}
	for i := 0; i < len(AllISO3Countries)-1; i++ {
		if AllISO3Countries[i] > AllISO3Countries[i+1] {
			t.Errorf("AllISO3Countries should be sorted, but %s is after %s", AllISO3Countries[i], AllISO3Countries[i+1])
		}
	}
	for i := 0; i < len(AllISO3Countries)-1; i++ {
		if AllISO3Countries[i] == AllISO3Countries[i+1] {
			t.Errorf("AllISO3Countries should have unique elements, but %s is repeated", AllISO3Countries[i])
		}
	}
}

func Test_ISO3CountryNaming(t *testing.T) {
	for _, iso3Country := range AllISO3Countries {
		value := string(iso3Country)
		if len(value) != 3 {
			t.Errorf("ISO3 country code %s should be exactly 3 characters long", value)
		}
		for _, char := range value {
			if char < 'A' || char > 'Z' {
				t.Errorf("ISO3 country code %s should only contain uppercase letters", value)
				break
			}
		}
	}
}
