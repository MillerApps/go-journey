// Package weather represents the current conditions for the current location.
package weather
// CurrentCondition is the current weather condition.
var CurrentCondition string
// CurrentLocation is the current location in which the CurrentCondition takes place.
var CurrentLocation string

// Forecast provides a weather forecast for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
