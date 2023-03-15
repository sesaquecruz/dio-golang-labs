package converter

func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

func FahrenheitToKelvin(value float64) float64 {
	return (value + 459.67) * 5 / 9
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func KelvinToFahrenheit(value float64) float64 {
	return value*9/5 - 459.67
}
