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

	// Konverterer Celsius til output temperatur
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
	var found bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

/*

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {


	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
}

func main() {
    var fromTemp float64
    var fromUnit string
    var toUnit string

    // Legg til flagg for inputverdi, inputenhet og ønsket utenheter
    flag.Float64Var(&fromTemp, "temp", 0.0, "input temperature")
    flag.StringVar(&fromUnit, "from", "C", "input unit [C, F, K]")
    flag.StringVar(&toUnit, "to", "C", "output unit [C, F, K]")
    flag.Parse()
// Velg konverteringsfunksjon basert på input- og utenheter
    var conversionFunc func(float64) float64
    switch fromUnit {
    case "C":
        switch toUnit {
        case "F":
            conversionFunc = conv.CelsiusToFahrenheit
        case "K":
            conversionFunc = conv.CelsiusToKelvin
        default:
            conversionFunc = func(x float64) float64 { return x }
        }
    case "F":
        switch toUnit {
        case "C":
            conversionFunc = conv.FahrenheitToCelsius
        case "K":
            conversionFunc = conv.FahrenheitToKelvin
        default:
            conversionFunc = func(x float64) float64 { return x }
        }
    case "K":
        switch toUnit {
        case "C":
            conversionFunc = conv.KelvinToCelsius
        case "F":
            conversionFunc = conv.KelvinToFahrenheit
        default:
            conversionFunc = func(x float64) float64 { return x }
        }
    default:
        conversionFunc = func(x float64) float64 { return x }
    }

    // Utfør konvertering
    toTemp := conversionFunc(fromTemp)

    // Skriv ut resultatet
    fmt.Printf("Input verdi %.2f %s er konvertert til %.2f %s\n", fromTemp, fromUnit, toTemp, toUnit)
}



/*
func convertTemp(temp float64, from string, to string) (float64, error) {
    switch from {
    case "F":
        switch to {
        case "C":
            return conv.FahrenheitToCelsius(temp), nil
        case "K":
            return conv.FahrenheitToKelvin(temp), nil
        default:
            return 0, fmt.Errorf("invalid to scale: %s", to)
        }
    case "C":
        switch to {
        case "F":
            return conv.CelsiusToFahrenheit(temp), nil


        case "K":
            return conv.CelsiusToKelvin(temp), nil
        default:
            return 0, fmt.Errorf("invalid to scale: %s", to)
        }
    case "K":
        switch to {
        case "F":
            return conv.KelvinToFahrenheit(temp), nil
        case "C":
            return conv.KelvinToCelsius(temp), nil
        default:
            return 0, fmt.Errorf("invalid to scale: %s", to)
        }
    default:
    fmt.Printf("Input verdi %.2f %s er konvertert til %.2f %s\n", fromTemp, fromUnit, toTemp, toUnit)}
}

// Bruk pakken conv til å vise temperaturkonverteringen som string
func formatTempConversion(temp float64, from string, to string) string {
    converted, _ := convertTemp(temp, from, to)
    return fmt.Sprintf("%g%s er %g%s", temp, from, converted, to)
}

// Bruk pakken conv til å konvertere temperaturer
// funksjonen vil konvertere temperaturen fra F, C eller K til alle de tre andre
func convertAllTempScales(temp float64, from string) (float64, float64, error) {
    switch from {
    case "F":
        return conv.FahrenheitToCelsius(temp), conv.FahrenheitToKelvin(temp), nil
    case "C":
        return conv.CelsiusToFahrenheit(temp), conv.CelsiusToKelvin(temp), nil
    case "K":
        return conv.KelvinToFahrenheit(temp), conv.KelvinToCelsius(temp), nil
    default:
        return 0, 0, fmt.Errorf("invalid from scale: %s", from)
    }
}


// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
*/
