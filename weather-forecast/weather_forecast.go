// Package weather provides a function to get the current weather CurrentCondition
// for a given city CurrentLocation.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast returns the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
