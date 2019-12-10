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

func findMaxThrust(program []int) int {
	phases := []int{0,1,2,3,4}
	maxThrust := -1
	for _, value0 := range phases {
		for _, value1 := range phases {
			if value1 == value0 {
				continue
			}
			for _, value2 := range phases {
				if value2 == value0 || value2 == value1 {
					continue
				}
				for _, value3 := range phases {
					if value3 == value0 || value3 == value1 || value3 == value2 {
						continue
					}
					for _, value4 := range phases {
						if value4 == value0 || value4 == value1 || value4 == value2 || value4 == value3 {
							continue
						}
						nextInput := 0
						phaseSequence := []int{value0, value1, value2, value3, value4}
						for _, phase := range phaseSequence {
							input := []int{phase, nextInput}
							_, output, err2 := runProgram(program, input)
							if err2 != nil {
								panic(err2)
							}
							nextInput = output[len(output)-1]
						}
						if nextInput > maxThrust {
							maxThrust = nextInput
						}
					}
				}
			}
		}
	}
	return maxThrust
}

func runProgram(program []int, input []int) ([]int, []int, error) {
	index := 0
	hit99 := false
	output := []int{}
	inputIndex := 0
	for index < len(program) && !hit99 {
		//fmt.Printf("index: %d (%v)\n", index, program)
		code, modeFirst, modeSecond, _ := parseCode(program[index])
		switch code {
		case 99:
			hit99 = true

		case 1:
			// add
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			destination := program[index+3]
			value := first + second

			program[destination] = value
			index += 4

		case 2:
			// multiply
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			destination := program[index+3]
			value := first * second
			program[destination] = value
			index += 4

		case 3:
			// Load input
			destination := program[index+1]
			program[destination] = input[inputIndex]
			inputIndex++
			index += 2

		case 4:
			// Output
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			output = append(output, first)
			//fmt.Printf("Output: %d\n", first)
			index += 2

		case 5:
			//jump-if-true
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			if first != 0 {
				index = second
			} else {
				index += 3
			}

		case 6:
			//jump-if-false
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			if first == 0 {
				index = second
			} else {
				index += 3
			}

		case 7:
			//less than
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			third := program[index+3]

			if first < second {
				program[third] = 1
			} else {
				program[third] = 0
			}
			index += 4
		case 8:
			//equals
			first := valueBasedOnParameterMode(program, index+1, modeFirst)
			second := valueBasedOnParameterMode(program, index+2, modeSecond)
			third := program[index+3]

			if first == second {
				program[third] = 1
			} else {
				program[third] = 0
			}
			index += 4

		default:
			return program, nil, errors.New("Invalid Code")
		}
	}
	return program, output, nil
}

func valueBasedOnParameterMode(program []int, index int, mode int) int {
	var value int
	if mode == 0 {
		value = program[program[index]]
	} else {
		value = program[index]
	}
	return value
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
