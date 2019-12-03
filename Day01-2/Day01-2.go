package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func CalcFuel(mass int) int {
	return mass/3 - 2
}

//https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func CalcFuelCounterUpper(filePath string) int {
	totalFuel := 0
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Printf("File not found %s\n", err.Error())
		return -1
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		mass, err := strconv.Atoi(strings.TrimRight(line, "\r\n"))
		fuel := CalcFuel(mass)
		totalFuel = totalFuel + fuel
		//fmt.Printf("TotalFuel: %d, Fuel %d\n", totalFuel, fuel)
		if err != nil {
			fmt.Printf("Error converting to int line: '%s' %s\n", line, err.Error())
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return totalFuel
}
