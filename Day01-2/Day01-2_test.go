package main

import (
	"testing"
)

func TestCalcFuel1(t *testing.T) {
	mass := 12
	fuel := CalcFuel(mass)
	if fuel != 2 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 2)
	}
}
func TestCalcFuel2(t *testing.T) {
	mass := 14
	fuel := CalcFuel(mass)
	if fuel != 2 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 2)
	}
}

func TestCalcFuel3(t *testing.T) {
	mass := 1969
	fuel := CalcFuel(mass)
	if fuel != 654 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 654)
	}
}

func TestCalcFuel4(t *testing.T) {
	mass := 100756
	fuel := CalcFuel(mass)
	if fuel != 33583 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 33583)
	}
}

func TestCalcFuel6(t *testing.T) {
	filePath := "testdata.txt"
	expected := 3432671

	actual := CalcFuelCounterUpper(filePath)

	if actual != expected {
		t.Errorf("Fuel was incorrect: actual:%d expected:%d", actual, expected)
	}
}
