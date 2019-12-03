package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	mass := 12
	fuel := CalcFuel(mass)
	if fuel != 2 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 2)
	}
}
func Test3(t *testing.T) {
	mass := 14
	fuel := CalcFuel(mass)
	if fuel != 2 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 2)
	}
}

func Test4(t *testing.T) {
	mass := 1969
	fuel := CalcFuel(mass)
	if fuel != 654 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 654)
	}
}

func Test5(t *testing.T) {
	mass := 100756
	fuel := CalcFuel(mass)
	if fuel != 33583 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", fuel, 33583)
	}
}

func CalcFuel(mass int) int {
	return mass/3 - 2
}
func Test6(t *testing.T) {
	// fmt.Println("readFileWithReadString")
	//t.Logf("readFileWithReadString")
	totalFuel := 0
	file, err := os.Open("testdata.txt")
	defer file.Close()

	if err != nil {
		t.Errorf("File not found %s", err.Error())
		return
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			//t.Errorf("Error reading line: %s", err.Error())
			break
		}
		// fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		// fmt.Println(" > > " + limitLength(line, 50))
		mass, err := strconv.Atoi(strings.TrimRight(line, "\r\n"))
		fuel := CalcFuel(mass)
		totalFuel = totalFuel + fuel
		//fmt.Printf("TotalFuel: %d, Fuel %d\n", totalFuel, fuel)
		if err != nil {
			t.Errorf("Error converting to int line: '%s' %s", line, err.Error())
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	if totalFuel != 3432671 {
		t.Errorf("Fuel was incorrect, got: %d, want: %d.", totalFuel, 3432671)
	}
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func readFileWithReadString(fn string) (err error) {
	fmt.Println("readFileWithReadString")

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		fmt.Println(" > > " + limitLength(line, 50))

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}

func limitLength(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[:length]
}
