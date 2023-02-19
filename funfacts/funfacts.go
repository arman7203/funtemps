package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	subject := flag.String("s", "", "the subject of the fun fact")
	unit := flag.String("t", "", "the unit of temperature")

	// Parse command line arguments
	flag.Parse()

	// Check that the subject and unit flags are set
	if *subject == "" || *unit == "" {
		fmt.Println("Please provide a subject and a temperature unit using the -s and -t flags.")
		return
	}

	// Print the fun fact based on the subject and unit flags
	if *subject == "sun" {
		if *unit == "C" {
			fmt.Println("The temperature on the outer layer of the Sun is 5506.85°C.")
			fmt.Println("The temperature in the core of the Sun is 15,000,000°C.")
		} else if *unit == "F" {
			// Convert to Fahrenheit
			outerLayerF := 5506.85*1.8 + 32
			coreF := 15000000*1.8 + 32
			fmt.Printf("The temperature on the outer layer of the Sun is %.2f°F.\n", outerLayerF)
			fmt.Printf("The temperature in the core of the Sun is %.2f°F.\n", coreF)
		} else if *unit == "K" {
			// Convert to Kelvin
			outerLayerK := 5506.85 + 273.15
			coreK := 15000000 + 273.15
			fmt.Printf("The temperature on the outer layer of the Sun is %.2fK.\n", outerLayerK)
			fmt.Printf("The temperature in the core of the Sun is %.2fK.\n", coreK)
		} else {
			fmt.Println("Invalid temperature unit. Please use C, F, or K.")
		}
	} else {
		fmt.Println("Invalid subject. Please use 'sun'.")
	}
}
