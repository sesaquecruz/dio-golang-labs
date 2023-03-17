package calculator

import "errors"

func Sum(a float64, b float64) float64 {
	return a + b
}

func Sub(a float64, b float64) float64 {
	return a - b
}

func Mult(a float64, b float64) float64 {
	return a * b
}

func Div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is undefined")
	}
	return a / b, nil
}
