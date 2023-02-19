package conv

func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

func KelvinToFahrenheit(value float64) float64 {
	return value*1.8 - 459.67
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}
