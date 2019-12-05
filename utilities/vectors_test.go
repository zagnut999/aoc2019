package utilities

import "testing"

func TestManhattanDistance1(t *testing.T) {
	expected := 10

	point1 := Point{0, 0}
	point2 := Point{10, 0}

	distance := ManhattanDistance(point1, point2)

	if distance != expected {
		t.Errorf("Expected distance of %d, actually %d", expected, distance)
	}
}

func TestManhattanDistance2(t *testing.T) {
	expected := 20

	point1 := Point{0, 0}
	point2 := Point{10, 10}

	distance := ManhattanDistance(point1, point2)

	if distance != expected {
		t.Errorf("Expected distance of %d, actually %d", expected, distance)
	}
}

func TestManhattanDistance3(t *testing.T) {
	expected := 20

	point1 := Point{-5, -5}
	point2 := Point{5, 5}

	distance := ManhattanDistance(point1, point2)

	if distance != expected {
		t.Errorf("Expected distance of %d, actually %d", expected, distance)
	}
}
