package main

import (
	"flag"
	"fmt"

	"github.com/kjetilnym0en/funtemps/conv"

)
//Definerer flag-variablene i hoved-"scope"
var fahr, celsius, kelvin float64
var out string

// Definerer og initialiserer flagg-variablene
func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
}

func main() {

	var inputTemp float64
	flag.Parse()

	switch {
	case isFlagPassed("F"):
		inputTemp = fahr
	case isFlagPassed("C"):
		inputTemp = celsius
	case isFlagPassed("K"):
		inputTemp = kelvin
	default:
		fmt.Println("No input temperature provided.")
		return
	}

	// Konverterer input temperatur til Celsius
	switch {
	case isFlagPassed("F"):
		celsius = conv.FahrenheitToCelsius(inputTemp)
	case isFlagPassed("K"):
		celsius = conv.KelvinToCelsius(inputTemp)
	}

	// Konverterer input temperatur til output temperatur
	switch out {
	case "C":
		fmt.Printf("%.2f°C\n", celsius)
	case "F":
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°F\n", fahr)
	case "K":
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°K\n", kelvin)
	default:
		fmt.Println("Invalid output temperature specified.")
		return
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}