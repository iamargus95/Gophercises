// Package weather contains functions to return the current weather conditions in a city.
package weather

// Variable CurrentCondition is a declaration of type string to store weather data.
var CurrentCondition string

// Variable CurrentLocation is a declaration of type string to store location data.
var CurrentLocation string

// Forecast returns a string with current weather data in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
