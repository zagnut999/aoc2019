package main

import (
	"strings"
	"bufio"
	"fmt"
	"io"
	"os"
	
)

func main() {

}

type node struct {
	name string
	children []*node
	parent *node
}

func findRoute(start string, end string, orbitMap map[string]*node) []*node {
	path := []*node{}

	startNode := orbitMap[start].parent
	path = walk(startNode, nil, end)
	
	return path
}

func walk (current *node, from *node, goal string) []*node {
	if from != nil {
		fmt.Printf("%s -> %s\n", from.name, current.name)
	} else {
		fmt.Printf("O -> %s\n", current.name)
	}
	for _, child:= range current.children {
		if child.name == goal {
			return []*node{current}
		}
	}

	for _, child:= range current.children {
		if (child != from) {
			result := walk(child, current, goal)
			if len(result) != 0 {
				return append(result, current)
			}
		}
	}
	 
	if current.parent != from {
		result := walk(current.parent, current, goal)
		if len(result) != 0 {
			return append(result, current)
		}
	}
	return []*node{}
}

func countOrbits(orbitMap map[string]*node) (int, int) {
	directCount :=0
	indirectCount := 0

	for _, value := range orbitMap {
		if value.parent != nil {
			directCount++
			temp := value.parent
			for {
				if temp.parent == nil {
					break
				}

				indirectCount++
				temp = temp.parent
			}
		}
	}
	return directCount, indirectCount
}

func parseOrbits (orbitPairs []string) map[string]*node {
	orbits := make(map[string]*node)

	for _, pair := range orbitPairs {
		temp := strings.Split(pair, ")")
		parentName := temp[0]
		childName := temp[1]

		var child, parent *node
		if orbits[parentName] != nil {
			parent = orbits[parentName]
		} else {
			parent = &node{}
			parent.name = parentName
			parent.children = []*node{}
			orbits[parentName] = parent
		}
		if orbits[childName] != nil {
			child = orbits[childName]
		} else {
			child = &node{}
			child.name = childName
			child.children = []*node{}
			orbits[childName] = child
		}

		child.parent = parent
		parent.children = append(parent.children, child)
		
	}

	return orbits
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func getTestData(filePath string) []string {
	testdata := []string{}
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic(err)
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
