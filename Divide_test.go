package main

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{6, 3, 2},
		{78, 39, 2},
		{2, 1, 2},
		{2, -2, -1},
		{-4, -2, 2},
	}

	for _, tc := range testCases {
		result := Divide(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Divide(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
