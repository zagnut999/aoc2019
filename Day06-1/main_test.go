package main

import (
	"testing"
)

func TestParseOrbits(t *testing.T) {
	orbits := []string{"COM)B", "B)C","C)D","D)E", "E)F","B)G","G)H","D)I","E)J","J)K","K)L"}
	expectedOrbits := 42
	orbitMap := parseOrbits(orbits)

	directCount, indirectCount := countOrbits(orbitMap)

	actualOrbits :=directCount + indirectCount
	if actualOrbits != expectedOrbits {
		t.Errorf("expected %d, got %d", expectedOrbits, actualOrbits)
	}

}

func TestFinal(t *testing.T) {
	orbits := getTestData("testdata.txt")
	expectedOrbits := 1 //1438

	orbitMap := parseOrbits(orbits)
	
	directCount, indirectCount := countOrbits(orbitMap)

	actualOrbits :=directCount + indirectCount
	if actualOrbits != expectedOrbits {
		t.Errorf("expected %d, got %d", expectedOrbits, actualOrbits)
	}
}
