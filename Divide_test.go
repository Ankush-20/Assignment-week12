package main

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{2, 39, 78},
		{2, 0, 0},
		{2, -2, -4},
		{-2, -2, 4},
	}

	for _, tc := range testCases {
		result := Divide(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Divide(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}

}
