package ioutils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PrintSalarySum() {
	inputFile, inputError := os.Open("employee.dat")
	if inputError != nil {
		fmt.Println(inputError)
		return
	}
	defer inputFile.Close()
	inputReader := csv.NewReader(inputFile)
	sum := float64(0)
	for {
		data, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		sal, _ := strconv.ParseFloat(data[2], 64)
		sum += sal
	}
	fmt.Printf("Sum of salary = %v\n", sum)
}
