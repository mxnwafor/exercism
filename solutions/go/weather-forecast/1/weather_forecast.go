// Package weather provides tools to perform a weather forecast for a particular location and condition.
package weather

var (
    // CurrentCondition represents the current weather condition of a specific location.
	CurrentCondition string
    // CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast returns a string value which depicts the forecast of a certain city and its condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
