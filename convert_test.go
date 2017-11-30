package convert

import "testing"

func TestMM(t *testing.T) {
	if MM(2) != 78.7401574804 {
		t.Error("MM(9) should be 78.7401574804")
	}

	if MM(43) != 1692.9133858286 {
		t.Error("MM(3) should be 1692.9133858286")
	}
}

func TestThousandths(t *testing.T) {
	if Thousandths(2) != .0508 {
		t.Error("Thousandths(2) should be .0508")
	}

	if Thousandths(1234) != 31.3436 {
		t.Error("Thousandths(1234) should be 31.3436")
	}
}

func TestFahrenheit(t *testing.T) {
	if Fahrenheit(32) != 0{
		t.Error("Fahrenheit(32) should be 0")
	}

	if Fahrenheit(213) != 100.55555555555556 {
		t.Error("Fahrenheit(213) should be 100.55555555555556")
	}
}

func TestCelsius(t *testing.T) {
	if Celsius(0) != 32 {
		t.Error("Celsius(0) should be 32")
	}

	if Celsius(99) != 210.2 {
		t.Error("Celsius(99) should be 210.2")
	}
}

func TestRadiusArea(t *testing.T) {
	if RadiusArea(0) != 0  {
		t.Error("RadiusArea(0) should be 0")
	}

	if RadiusArea(43) != 5808.80480985  {
		t.Error("RadiusArea(43) should be 5808.80480985")
	}
}

func TestRadiusCircumference(t *testing.T) {
	if RadiusCircumference(12) != 75.39822360000001 {
		t.Error("RadiusCircumference(12) should be 75.39822360000001")
	}

	if RadiusCircumference(125) != 785.3981625 {
		t.Error("RadiusCircumference(125) should be 785.3981625")
	}
}
