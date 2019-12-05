package main

import (
	"encoding/json"
	"testing"
)

func TestWalkPath1(t *testing.T) {
	pathString := "U8"
	expected := 8
	path := parsePath(pathString)
	visited := walkPath(path)
	if len(visited) != expected {
		t.Errorf("Expected %d nodes, actually %d", expected, len(visited))
	}

	for index, element := range visited {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("Index :%s Element: %d", string(out), element)
	}
}

func TestWalkPath2(t *testing.T) {
	pathString := "R8,U5,L5,D3"
	expected := 21
	path := parsePath(pathString)

	visited := walkPath(path)
	if len(visited) != expected {
		t.Errorf("Expected %d nodes, actually %d", expected, len(visited))
	}

	for index, element := range visited {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("Index :%s Element: %d", string(out), element)
	}
}

func TestWalkPath4(t *testing.T) {
	path1String := "R8,U5,L5,D3"
	path2String := "U7,R6,D4,L4"
	expected := 6
	path1 := parsePath(path1String)
	path2 := parsePath(path2String)

	visited1 := walkPath(path1)
	visited2 := walkPath(path2)

	distance := distanceToClosestIntersection2(visited1, visited2)
	if distance != expected {
		t.Errorf("Expected %d distance, actually %d", expected, distance)
	}

	for index, element := range visited1 {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("Index :%s Element: %d", string(out), element)
	}
}

func TestWalkPath5(t *testing.T) {
	path1String := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	path2String := "U62,R66,U55,R34,D71,R55,D58,R83"
	expected := 159
	path1 := parsePath(path1String)
	path2 := parsePath(path2String)

	visited1 := walkPath(path1)
	visited2 := walkPath(path2)

	distance := distanceToClosestIntersection2(visited1, visited2)
	if distance != expected {
		t.Errorf("Expected %d distance, actually %d", expected, distance)
	}

	for index, element := range visited1 {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("1: Index :%s Element: %d", string(out), element)
	}
	for index, element := range visited2 {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("2: Index :%s Element: %d", string(out), element)
	}
}

func TestWalkPath6(t *testing.T) {
	path1String := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	path2String := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	expected := 135
	path1 := parsePath(path1String)
	path2 := parsePath(path2String)

	visited1 := walkPath(path1)
	visited2 := walkPath(path2)

	distance := distanceToClosestIntersection2(visited1, visited2)
	if distance != expected {
		t.Errorf("Expected %d distance, actually %d", expected, distance)
	}

	for index, element := range visited1 {
		out, err := json.Marshal(index)
		if err != nil {
			panic(err)
		}
		t.Logf("Index :%s Element: %d", string(out), element)
	}
}

func TestParseTestData(t *testing.T) {
	paths := getTestData("testdata.txt")

	if len(paths) != 2 {
		t.Errorf("Expected %d lines, actually %d", 2, len(paths))
	}
}

func TestFinal(t *testing.T) {
	paths := getTestData("testdata.txt")
	expected := 1674
	path1 := parsePath(paths[0])
	path2 := parsePath(paths[1])

	visited1 := walkPath(path1)
	visited2 := walkPath(path2)

	distance := distanceToClosestIntersection2(visited1, visited2)
	if distance != expected {
		t.Errorf("Expected %d distance, actually %d", expected, distance)
	}

}
