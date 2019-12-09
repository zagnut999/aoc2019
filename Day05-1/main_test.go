package main

import (
	"aoc2019/utilities"
	"strings"
	"testing"
)

func TestParseProgram(t *testing.T) {
	program := "1,9,10,3,2,3,11,0,99,30,40,50"
	intcodes, err := parseProgram(program)
	if err != nil {
		t.Errorf("Error parsing program %s", err.Error())
	}
	if len(intcodes) != 12 {
		t.Errorf("Expected %d codes, actually %d", 12, len(intcodes))
	}
}

func TestParseTestData(t *testing.T) {
	program := getTestData("testdata.txt")
	intcodes, err := parseProgram(program)
	if err != nil {
		t.Errorf("Error parsing program %s", err.Error())
	}
	if len(intcodes) != 678 {
		t.Errorf("Expected %d codes, actually %d", 149, len(intcodes))
	}
}

func TestProgramResult1(t *testing.T) {
	program := "1,0,0,0,99"
	expected := "2,0,0,0,99"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult2(t *testing.T) {
	program := "2,3,0,3,99"
	expected := "2,3,0,6,99"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult3(t *testing.T) {
	program := "2,4,4,5,99,0"
	expected := "2,4,4,5,99,9801"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult4(t *testing.T) {
	program := "1,1,1,4,99,5,6,0,99"
	expected := "30,1,1,4,2,5,6,0,99"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestInputOutput(t *testing.T) {
	program := "3,0,4,0,99"
	expectedResult := "1,0,4,0,99"
	expectedOutput := 1
	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, output, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if output[0] != expectedOutput {
		t.Errorf("Expected Output '%d', actually '%v'", expectedOutput, output)
	}
	if actual != expectedResult {
		t.Errorf("Expected Result '%s', actually '%s'", expectedResult, actual)
	}
}

func TestParameterModes1(t *testing.T) {
	code := 1
	expectedCode := 1
	expectedFirst := 0
	expectedSecond := 0
	expectedThird := 0

	actualCode, actualFirst, actualSecond, actualThird := parseCode(code)

	if actualCode != expectedCode {
		t.Errorf("Expected code '%d', actual '%d'", expectedCode, actualCode)
	}

	if actualFirst != expectedFirst {
		t.Errorf("Expected first '%d', actual '%d'", expectedFirst, actualFirst)
	}

	if actualSecond != expectedSecond {
		t.Errorf("Expected second '%d', actual '%d'", expectedSecond, actualSecond)
	}

	if actualThird != expectedThird {
		t.Errorf("Expected third '%d', actual '%d'", expectedThird, actualThird)
	}
}

func TestParameterModes2(t *testing.T) {
	code := 99
	expectedCode := 99
	expectedFirst := 0
	expectedSecond := 0
	expectedThird := 0

	actualCode, actualFirst, actualSecond, actualThird := parseCode(code)

	if actualCode != expectedCode {
		t.Errorf("Expected code '%d', actual '%d'", expectedCode, actualCode)
	}

	if actualFirst != expectedFirst {
		t.Errorf("Expected first '%d', actual '%d'", expectedFirst, actualFirst)
	}

	if actualSecond != expectedSecond {
		t.Errorf("Expected second '%d', actual '%d'", expectedSecond, actualSecond)
	}

	if actualThird != expectedThird {
		t.Errorf("Expected third '%d', actual '%d'", expectedThird, actualThird)
	}
}

func TestParameterModes3(t *testing.T) {
	code := 1002
	expectedCode := 2
	expectedFirst := 0
	expectedSecond := 1
	expectedThird := 0

	actualCode, actualFirst, actualSecond, actualThird := parseCode(code)

	if actualCode != expectedCode {
		t.Errorf("Expected code '%d', actual '%d'", expectedCode, actualCode)
	}

	if actualFirst != expectedFirst {
		t.Errorf("Expected first '%d', actual '%d'", expectedFirst, actualFirst)
	}

	if actualSecond != expectedSecond {
		t.Errorf("Expected second '%d', actual '%d'", expectedSecond, actualSecond)
	}

	if actualThird != expectedThird {
		t.Errorf("Expected third '%d', actual '%d'", expectedThird, actualThird)
	}
}

func TestParameterModes4(t *testing.T) {
	code := 10103
	expectedCode := 3
	expectedFirst := 1
	expectedSecond := 0
	expectedThird := 1

	actualCode, actualFirst, actualSecond, actualThird := parseCode(code)

	if actualCode != expectedCode {
		t.Errorf("Expected code '%d', actual '%d'", expectedCode, actualCode)
	}

	if actualFirst != expectedFirst {
		t.Errorf("Expected first '%d', actual '%d'", expectedFirst, actualFirst)
	}

	if actualSecond != expectedSecond {
		t.Errorf("Expected second '%d', actual '%d'", expectedSecond, actualSecond)
	}

	if actualThird != expectedThird {
		t.Errorf("Expected third '%d', actual '%d'", expectedThird, actualThird)
	}
}

func TestParameterModes5(t *testing.T) {
	program := "1002,3,1,3,99"
	expected := "1002,3,1,3,99"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestParameterModes6(t *testing.T) {
	program := "1001,3,-10,3,99"
	expected := "1001,3,-10,-7,99"

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, _, err2 := runProgram(intcodes, 1)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestFinal(t *testing.T) {
	program := getTestData("testdata.txt")
	expected := 4511442

	intcodes, err1 := parseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	_, output, err2 := runProgram(intcodes, 1)
	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}
	actual := output[len(output)-1]
	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}
