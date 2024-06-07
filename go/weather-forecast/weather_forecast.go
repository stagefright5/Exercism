// Package weather provides API to forcast weather
// based on the current condition and location.
package weather

// CurrentCondition holds the value of the current weather condition.
var CurrentCondition string

// CurrentLocation holds the values of the location at the current time for
// which the current condition corresponds to.
var CurrentLocation string

// Forecast return the weather forecast by taking in the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
