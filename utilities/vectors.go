package utilities

import "math"

type Point struct {
	X int
	Y int
}

func ManhattanDistance(first Point, second Point) int {

	return int(math.Abs(float64(first.X-second.X)) + math.Abs(float64(first.Y-second.Y)))
}
