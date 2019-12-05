package main

import (
	"aoc2019/utilities"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

}
func distanceToClosestIntersection2(visited1 map[utilities.Point]int, visited2 map[utilities.Point]int) int {
	intersections := make(map[utilities.Point]int)
	for point, steps := range visited1 {
		if visited2[point] != 0 {
			intersections[point] = steps + visited2[point]
		}
	}

	minSteps := -1

	for _, steps := range intersections {
		if minSteps == -1 || steps < minSteps {
			minSteps = steps
		}
	}

	return minSteps
}

func walkPath(path []string) map[utilities.Point]int {
	current := utilities.Point{X: 0, Y: 0}
	visited := make(map[utilities.Point]int)
	steps := 0
	for index, move := range path {
		direction, distance, err := parseMove(move)

		if err != nil {
			fmt.Printf("paring error ('%s', %d) %s\n", move, index, err.Error())
			return visited
		}

		for i := 0; i < distance; i++ {
			steps++
			switch direction {
			case "U":
				current = utilities.Point{X: current.X, Y: current.Y + 1}
			case "D":
				current = utilities.Point{X: current.X, Y: current.Y - 1}
			case "L":
				current = utilities.Point{X: current.X - 1, Y: current.Y}
			case "R":
				current = utilities.Point{X: current.X + 1, Y: current.Y}
			}
			if visited[current] == 0 {
				visited[current] = steps
			}
		}
	}

	return visited
}

func parseMove(move string) (string, int, error) {
	direction := move[0:1]
	distance, err := strconv.Atoi(move[1:])
	if err != nil {
		return "", -1, err
	}
	return direction, distance, nil
}

func parsePath(pathString string) []string {
	return strings.Split(pathString, ",")
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func getTestData(filePath string) []string {
	testdata := []string{}
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Printf("File not found %s\n", err.Error())
		return testdata
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		testdata = append(testdata, strings.TrimRight(line, "\r\n"))
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return testdata
}
