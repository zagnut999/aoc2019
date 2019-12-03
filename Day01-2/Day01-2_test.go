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
	if fuel != 966 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 966)
	}
}

func TestCalcFuel4(t *testing.T) {
	mass := 100756
	fuel := CalcFuel(mass)
	if fuel != 50346 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 50346)
	}
}

func TestCalcFuel6(t *testing.T) {
	filePath := "testdata.txt"
	expected := 5146132

	actual := CalcFuelCounterUpper(filePath)

	if actual != expected {
		t.Errorf("Fuel was incorrect: actual:%d expected:%d", actual, expected)
	}
}
