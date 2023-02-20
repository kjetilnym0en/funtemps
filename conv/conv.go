package conv

import "math"

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	return round((celsius * 9.0 / 5.0) + 32.0, 2)
}

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return round((fahrenheit - 32.0) * 5.0 / 9.0, 2)
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(celsius float64) float64 {
	return round(celsius + 273.15, 2)
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(kelvin float64) float64 {
	return round(kelvin - 273.15, 2)
}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(fahrenheit float64) float64 {
	return round(CelsiusToKelvin(FahrenheitToCelsius(fahrenheit)), 2)
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(kelvin float64) float64 {
	return round(CelsiusToFahrenheit(KelvinToCelsius(kelvin)), 2)
}

// Runder av en float-verdi til n desimaler
func round(f float64, n int) float64 {
    scale := math.Pow(10, float64(n))
    return math.Round(f*scale) / scale
}

