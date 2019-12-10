package main

import (
	"testing"
)

func TestCombos(t *testing.T) {
	expected := 5 * 4 * 3 * 2 * 1
	input := []int{0,1,2,3,4}
	actual := 0

	for _, value0 := range input {
		for _, value1 := range input {
			if value1 == value0 {
				continue
			}
			for _, value2 := range input {
				if value2 == value0 || value2 == value1 {
					continue
				}
				for _, value3 := range input {
					if value3 == value0 || value3 == value1 || value3 == value2 {
						continue
					}
					for _, value4 := range input {
						if value4 == value0 || value4 == value1 || value4 == value2 || value4 == value3 {
							continue
						}
						actual++
					}
				}
			}
		}
	}

	if actual != expected {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}

func Test1a(t *testing.T) {
	program := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	phaseSequence := []int{4,3,2,1,0}
	input := 0
	expected := 43210

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	nextInput :=input
	for _, phase := range phaseSequence {
		input := []int{phase, nextInput}
		_, output, err2 := runProgram(intcodes, input)
		if err2 != nil {
			t.Errorf("Error running program %s", err2.Error())
		}
		nextInput = output[len(output)-1]
	}

	actual := nextInput

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}

func Test1b(t *testing.T) {
	program := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	expected := 43210

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	actual := findMaxThrust(intcodes)

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}

func Test2a(t *testing.T) {
	program := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	phaseSequence := []int{0,1,2,3,4}
	input := 0
	expected := 54321

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	nextInput :=input
	for _, phase := range phaseSequence {
		input := []int{phase, nextInput}
		_, output, err2 := runProgram(intcodes, input)
		if err2 != nil {
			t.Errorf("Error running program %s", err2.Error())
		}
		nextInput = output[len(output)-1]
	}

	actual := nextInput

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}

func Test2b(t *testing.T) {
	program := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	expected := 54321

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	actual := findMaxThrust(intcodes)

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}

func Test3a(t *testing.T) {
	program := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	phaseSequence := []int{1,0,4,3,2}
	input := 0
	expected := 65210

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	nextInput :=input
	for _, phase := range phaseSequence {
		input := []int{phase, nextInput}
		_, output, err2 := runProgram(intcodes, input)
		if err2 != nil {
			t.Errorf("Error running program %s", err2.Error())
		}
		nextInput = output[len(output)-1]
	}

	actual := nextInput

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}

func Test3b(t *testing.T) {
	program := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	expected := 65210

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	actual := findMaxThrust(intcodes)

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}


func TestFinal(t *testing.T) {
	program := getTestData("testdata.txt")
	expected := 12648139

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	actual := findMaxThrust(intcodes)

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}