package main

import (
	"errors"
	"fmt"
	converter "lab-1/pkg"
	"log"
)

func menu() (*uint8, *float64, error) {
	fmt.Println("\n[1] Celsius -> Fahrenheit")
	fmt.Println("[2] Celsius -> Kelvin")
	fmt.Println("[3] Fahrenheit -> Celsius")
	fmt.Println("[4] Fahrenheit -> Kelvin")
	fmt.Println("[5] Kelvin -> Celsius")
	fmt.Printf("[6] Kelvin -> Fahrenheit\n\n")

	fmt.Print("Option: ")
	var option uint8
	_, err := fmt.Scan(&option)
	if err != nil {
		return nil, nil, err
	}

	fmt.Print("Value:  ")
	var value float64
	_, err = fmt.Scan(&value)
	if err != nil {
		return nil, nil, err
	}

	return &option, &value, nil
}

func result(option *uint8, value *float64) (*float64, error) {
	var result float64

	switch *option {
	case 1:
		result = converter.CelsiusToFahrenheit(*value)
	case 2:
		result = converter.CelsiusToKelvin(*value)
	case 3:
		result = converter.FahrenheitToCelsius(*value)
	case 4:
		result = converter.FahrenheitToKelvin(*value)
	case 5:
		result = converter.KelvinToCelsius(*value)
	case 6:
		result = converter.KelvinToFahrenheit(*value)
	default:
		return nil, errors.New("invalid option")
	}

	return &result, nil
}

func main() {
	option, value, err := menu()
	if err != nil {
		log.Fatal(err)
	}

	result, err := result(option, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %g\n\n", *result)
}
