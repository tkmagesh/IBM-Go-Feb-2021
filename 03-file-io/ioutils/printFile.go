package ioutils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintFile() (err error) {
	inputHandle, inputError := os.Open("data1.dat")

	defer func() {
		inputHandle.Close()
	}()

	if inputError != nil {
		err = inputError
		return
	}
	inputReader := bufio.NewReader(inputHandle)
	for {
		line, err := inputReader.ReadString('\n')
		fmt.Println(err)
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
	return
}
