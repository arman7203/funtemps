package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/arman7203/funtemps/conv"
)

// define main variables
var (
	fahrenheit float64
	celsius    float64
	kelvin     float64
	out        string
)

// initializes variable-flags, where the parameters "point" to the mainn variables above,
// where the first string refers to the flag, default value "0.0", and finally info/help for the consumer
func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temprature in fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temprature in celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temprature in kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K - kelvin")

}

func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Println(fahrenheit, "°F is equal to", math.Round(conv.FahrenheitToCelsius(fahrenheit)*100)/100, "°C")
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to", math.Round(conv.CelsiusToFahrenheit(celsius)*100)/100, "°F")
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to", math.Round(conv.CelsiusToKelvin(celsius)*100)/100, "°K")
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to", math.Round(conv.KelvinToCelsius(kelvin)*100)/100, "°C")
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to", math.Round(conv.KelvinToFahrenheit(kelvin)*100)/100, "°F")
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
