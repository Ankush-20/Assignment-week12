package main

import (
	"testing"
)

// TestFactorial tests the Factorial function for correct functionality and edge cases.
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		hasError bool
	}{
		{5, 120, false},
		{0, 1, false},
		{1, 1, false},
		{-1, -1, true},
		{10, 3628800, false},
	}

	for _, test := range tests {
		result, err := Factorial(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Factorial(%d) expected an error, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Factorial(%d) expected no error, but got: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Factorial(%d) = %d; want %d", test.input, result, test.expected)
			}
		}
	}
}
