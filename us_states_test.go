package countries

import (
	"testing"
)

func Test_AllUSStates(t *testing.T) {
	// 50 states + DC + 5 territories = 56
	expectedCount := 56
	if len(AllUSStates) != expectedCount {
		t.Errorf("AllUSStates should have %d elements, but it has %d", expectedCount, len(AllUSStates))
	}

	for i := 0; i < len(AllUSStates)-1; i++ {
		if AllUSStates[i] > AllUSStates[i+1] {
			t.Errorf("AllUSStates should be sorted, but %s is after %s", AllUSStates[i], AllUSStates[i+1])
		}
	}

	for i := 0; i < len(AllUSStates)-1; i++ {
		if AllUSStates[i] == AllUSStates[i+1] {
			t.Errorf("AllUSStates should have unique elements, but %s is repeated", AllUSStates[i])
		}
	}
}

func Test_USStateNaming(t *testing.T) {
	for _, stateCode := range AllUSStates {
		value := string(stateCode)
		if len(value) != 2 {
			t.Errorf("US state code %s should be exactly 2 characters long", value)
		}
		for _, char := range value {
			if char < 'A' || char > 'Z' {
				t.Errorf("US state code %s should only contain uppercase letters", value)
				break
			}
		}
	}
}
