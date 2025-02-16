package util

import (
	"errors"
)

func Sum(a int, b int) int {
	return a + b
}

func Mult(a, b int) int {
	return a * b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot Divide With Zero")
	}

	return a / b, nil
}
