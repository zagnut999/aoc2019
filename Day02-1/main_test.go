package main

import (
	"strings"
	"testing"
)

func TestParseProgram(t *testing.T) {
	program := "1,9,10,3,2,3,11,0,99,30,40,50"
	intcodes, err := ParseProgram(program)
	if err != nil {
		t.Errorf("Error parsing program %s", err.Error())
	}
	if len(intcodes) != 12 {
		t.Errorf("Expected %d codes, actually %d", 12, len(intcodes))
	}
}

func TestParseTestData(t *testing.T) {
	program := GetTestData("testdata.txt")
	intcodes, err := ParseProgram(program)
	if err != nil {
		t.Errorf("Error parsing program %s", err.Error())
	}
	if len(intcodes) != 149 {
		t.Errorf("Expected %d codes, actually %d", 149, len(intcodes))
	}
}

func TestProgramResult1(t *testing.T) {
	program := "1,0,0,0,99"
	expected := "2,0,0,0,99"

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, err2 := RunProgram(intcodes)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult2(t *testing.T) {
	program := "2,3,0,3,99"
	expected := "2,3,0,6,99"

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, err2 := RunProgram(intcodes)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult3(t *testing.T) {
	program := "2,4,4,5,99,0"
	expected := "2,4,4,5,99,9801"

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, err2 := RunProgram(intcodes)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestProgramResult4(t *testing.T) {
	program := "1,1,1,4,99,5,6,0,99"
	expected := "30,1,1,4,2,5,6,0,99"

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	result, err2 := RunProgram(intcodes)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestFinal(t *testing.T) {
	program := GetTestData("testdata.txt")
	expected := "30,1,1,4,2,5,6,0,99"

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	// Modify to 1202 program alarm
	intcodes[1] = 12
	intcodes[2] = 2

	result, err2 := RunProgram(intcodes)

	if err2 != nil {
		t.Errorf("Error running program %s", err2.Error())
	}

	actual := strings.Join(Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}
