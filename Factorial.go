package main

import (
	"errors"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("negative input is not allowed")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
