//this package contains many useful convertors
package convert

func MM (input float64) float64 {
  //converts mm to thousandths
  return input * 39.3700787402
}

func Thousandths (input float64) float64 {
  //converts thousandths to mm
  return input * .0254
}

func Fahrenheit (input float64) float64 {
  //converts Fahrenheit to Celsius
  return ((input -32) * 5/9)
}

func Celsius (input float64) float64 {
  //converts Celsius to Fahrenheit
  return ((input * 9/5) + 32)
}

func RadiusArea (input float64) float64 {
  //converts the radius of a circle into its area
  return ((input * input) * 3.14159265)
}

func RadiusCircumference (input float64) float64 {
  //converts the radius of a circle into its circumferece
  return (input * 6.28318530)
}
