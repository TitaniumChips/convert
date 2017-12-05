//this package contains many useful convertors
package convert

// MM converts mm to thousandths
func MM(input float64) float64 {
	return input * 39.3700787402
}

// Thousandths converts thousandths to mm
func Thousandths(input float64) float64 {
	return input * .0254
}

// Fahrenheit converts Fahrenheit to Celsius
func Fahrenheit(input float64) float64 {
	return ((input - 32) * 5 / 9)
}

// Celsius converts Celsius to Fahrenheit
func Celsius(input float64) float64 {
	return ((input * 9 / 5) + 32)
}

// RadiusArea converts the radius of a circle into its area
func RadiusArea(input float64) float64 {
	return ((input * input) * 3.14159265)
}

// RadiusCircumference converts the radius of a circle into its circumferece
func RadiusCircumference(input float64) float64 {
	return (input * 6.28318530)
}
