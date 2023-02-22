package main

import (
	"flag"
	"fmt"

	"github.com/kjetilnym0en/funtemps/conv"

)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

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


	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

// Bruk pakken conv til å konvertere temperaturer
// funksjonen vil konvertere temperaturen fra F, C eller K til de to andre

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
*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
/*
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

}
*/

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
