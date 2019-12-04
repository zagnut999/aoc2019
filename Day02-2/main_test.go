package main

import (
	"aoc2019/utilities"
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

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

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

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

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

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

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

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestFinal(t *testing.T) {
	program := GetTestData("testdata.txt")
	expected := "3562624,12,2,2,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,48,1,5,19,49,1,23,5,50,1,27,13,55,1,31,5,56,1,9,35,59,2,13,39,295,1,43,10,299,1,47,13,304,2,10,51,1216,1,55,5,1217,1,59,5,1218,1,63,13,1223,1,13,67,1228,1,71,10,1232,1,6,75,1234,1,6,79,1236,2,10,83,4944,1,87,5,4945,1,5,91,4946,2,95,10,19784,1,9,99,19787,1,103,13,19792,2,10,107,79168,2,13,111,395840,1,6,115,395842,1,119,10,395846,2,9,123,1187538,2,127,9,3562614,1,131,10,3562618,1,135,2,3562620,1,10,139,0,99,2,0,14,0"

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

	actual := strings.Join(utilities.Slice_Itoa(result), ",")

	if actual != expected {
		t.Errorf("Expected '%s', actually '%s'", expected, actual)
	}
}

func TestFinalDay2(t *testing.T) {
	program := GetTestData("testdata.txt")
	expected := 8298

	intcodes, err1 := ParseProgram(program)
	if err1 != nil {
		t.Errorf("Error parsing program %s", err1.Error())
	}

	foundNoun := -1
	foundVerb := -1
	found := false
	for noun := 0; noun < 100 && !found; noun++ {
		for verb := 0; verb < 100 && !found; verb++ {
			rundata := utilities.Aray_Clone(intcodes)
			rundata[1] = noun
			rundata[2] = verb

			result, err2 := RunProgram(rundata)
			if err2 != nil {
				t.Errorf("Error running program %s", err2.Error())
			}

			if result[0] == 19690720 {
				found = true
				foundNoun = noun
				foundVerb = verb
			}
		}
	}

	actual := 100*foundNoun + foundVerb

	if actual != expected {
		t.Errorf("Expected '%d', actually '%d'", expected, actual)
	}
}
