package convert

import (
	"math"
	"testing"
)

const tolerence = 0.00001

func testConv(t *testing.T, f func(float64) float64, arg, want float64) {
	if got := f(arg); math.Abs(got-want) >= tolerence {
		t.Errorf("passed %f and got %f but wanted %f", arg, got, want)
	}
}

func TestMM(t *testing.T) {
	testConv(t, MM, 2, 78.7401574804)
	testConv(t, MM, 43, 1692.9133858286)
}

func TestThousandths(t *testing.T) {
	testConv(t, Thousandths, 2, .0508)
	testConv(t, Thousandths, 1234, 31.3436)
}

func TestFahrenheit(t *testing.T) {
	testConv(t, Fahrenheit, 32, 0)
	testConv(t, Fahrenheit, 213, 100.55555555555556)
}

func TestCelsius(t *testing.T) {
	testConv(t, Celsius, 0, 32)
	testConv(t, Celsius, 99, 210.2)
}

func TestRadiusArea(t *testing.T) {
	testConv(t, RadiusArea, 0, 0)
	testConv(t, RadiusArea, 43, 5808.80480985)
}

func TestRadiusCircumference(t *testing.T) {
	testConv(t, RadiusCircumference, 12, 75.39822360000001)
	testConv(t, RadiusCircumference, 125, 785.3981625)
}
