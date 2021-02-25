package main

import (
	"file-io/ioutils"
	"fmt"
)

func main() {
	/* err := ioutils.PrintFile()
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
		return
	} */

	//ioutils.CopyFileContents()
	ioutils.PrintSalarySum()
	x := test()
	r, _ := x.(int)
	fmt.Println(r + 10)
}

func test() interface{} {
	return 10
}
