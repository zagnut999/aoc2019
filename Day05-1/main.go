package main

import (
	"aoc2019/utilities"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func runProgram(program []int, input int) ([]int, []int, error) {
	index := 0
	hit99 := false
	output := []int{}
	for index < len(program) && !hit99 {
		fmt.Printf("index: %d\n", index)
		code, modeFirst, modeSecond, _ := parseCode(program[index])
		switch code {
		case 99:
			hit99 = true

		case 1:
			// add
			var first, second int
			if modeFirst == 0 {
				first = program[program[index+1]]
			} else {
				first = program[index+1]
			}
			if modeSecond == 0 {
				second = program[program[index+2]]
			} else {
				second = program[index+2]
			}
			destination := program[index+3]
			value := first + second

			program[destination] = value
			index += 4

		case 2:
			// multiply
			var first, second int
			if modeFirst == 0 {
				first = program[program[index+1]]
			} else {
				first = program[index+1]
			}
			if modeSecond == 0 {
				second = program[program[index+2]]
			} else {
				second = program[index+2]
			}
			destination := program[index+3]
			value := first * second
			program[destination] = value
			index += 4

		case 3:
			// Load input
			destination := program[index+1]
			program[destination] = input
			index += 2

		case 4:
			// Output
			var first int
			if modeFirst == 0 {
				first = program[program[index+1]]
			} else {
				first = program[index+1]
			}
			output = append(output, first)
			fmt.Printf("Output: %d\n", first)

			index += 2

		default:
			return program, nil, errors.New("Invalid Code")
		}
	}
	return program, output, nil
}

func parseCode(code int) (int, int, int, int) {
	if code < 100 {
		return code, 0, 0, 0
	}

	fullCode := fmt.Sprintf("%05d", code)
	codePart, _ := strconv.Atoi(fullCode[3:5])
	firstParamPart, _ := strconv.Atoi(fullCode[2:3])
	secondParamPart, _ := strconv.Atoi(fullCode[1:2])
	thirdParamPart, _ := strconv.Atoi(fullCode[0:1])
	return codePart, firstParamPart, secondParamPart, thirdParamPart
}

func parseProgram(program string) ([]int, error) {
	return utilities.Slice_Atoi(strings.Split(program, ","))
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func getTestData(filePath string) string {
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
