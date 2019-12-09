package main

import (
	"testing"
)

func TestParseOrbits(t *testing.T) {
	orbits := []string{"COM)B", "B)C","C)D","D)E", "E)F","B)G","G)H","D)I","E)J","J)K","K)L","K)YOU","I)SAN"}
	expected := 4
	orbitMap := parseOrbits(orbits)

	transfers := findRoute("YOU", "SAN", orbitMap)

	actual := len(transfers) - 1
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestFinal(t *testing.T) {
	orbits := getTestData("testdata.txt")
	expected := 307

	orbitMap := parseOrbits(orbits)
	
	transfers := findRoute("YOU", "SAN", orbitMap)

	actual := len(transfers) - 1
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
