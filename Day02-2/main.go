package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

}

func RunProgram(program []int) ([]int, error) {
	index := 0
	hit99 := false
	for index < len(program) && !hit99 {
		fmt.Print(index)
		switch code := program[index]; code {
		case 99:
			hit99 = true

		case 1:
			// add
			first := program[program[index+1]]
			second := program[program[index+2]]
			destination := program[index+3]
			value := first + second
			//fmt.Printf("%d %d %d %d", first, second, destination, value)

			program[destination] = value
			index += 4

		case 2:
			// multiply
			first := program[program[index+1]]
			second := program[program[index+2]]
			destination := program[index+3]
			value := first * second
			//fmt.Printf("%d %d %d %d", first, second, destination, value)

			program[destination] = value
			index += 4

		default:
			return program, errors.New("Invalid Code")
		}
	}
	return program, nil
}

func ParseProgram(program string) ([]int, error) {
	return Slice_Atoi(strings.Split(program, ","))
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func GetTestData(filePath string) string {
	var testdata string
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Printf("File not found %s\n", err.Error())
		return ""
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		testdata += strings.TrimRight(line, "\r\n")
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return testdata
}
