package services

import "testing"

func TestIsValidTrue(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"4539148803436467", true},  // Valid
		{"6011111111111117", true},  // Valid
		{"4539148803436466", false}, // Invalid
		{"4485266390581110", false}, // Invalid
	}

	for _, test := range tests {
		result := Is_valid(test.number)
		if result != test.expected {
			t.Errorf("Is_valid(%s) = %v; expected %v", test.number, result, test.expected)
		}
	}
}
