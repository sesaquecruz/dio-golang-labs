package main

import (
	"errors"
	"fmt"
	"lab-5/external/calculator"
)

func main() {
	for {
		fmt.Println("\n[1] Sum")
		fmt.Println("[2] Subtraction")
		fmt.Println("[3] Multiplication")
		fmt.Println("[4] Division")
		fmt.Println("[0] Exit")

		fmt.Print("\n[?] ")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if option < 0 || option > 5 {
			fmt.Println(errors.New("invalid option"))
			continue
		} else if option == 0 {
			break
		}

		fmt.Print("\nNumber 1: ")
		var number1 float64
		_, err = fmt.Scan(&number1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("Number 2: ")
		var number2 float64
		_, err = fmt.Scan(&number2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var result float64
		switch option {
		case 1:
			result = calculator.Sum(number1, number2)
		case 2:
			result = calculator.Sub(number1, number2)
		case 3:
			result = calculator.Mult(number1, number2)
		case 4:
			result, err = calculator.Div(number1, number2)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		fmt.Printf("Restult:  %f\n", result)
	}
}
