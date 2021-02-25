package main

import (
	"file-io/ioutils"
	"fmt"
)

func main() {
	err := ioutils.PrintFile()
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
		return
	}
}
