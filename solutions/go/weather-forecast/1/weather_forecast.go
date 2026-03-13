// Package weather provides tools to get the weather forecast of a city.
package weather

var (
	// CurrentCondition stores the weather condition of a given location.
	CurrentCondition string

	// CurrentLocation stores a given location.
	CurrentLocation  string
)

// Forecast returns a string value representing the weather condition of a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
