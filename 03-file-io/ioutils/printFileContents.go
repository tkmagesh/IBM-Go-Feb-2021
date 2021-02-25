package ioutils

import (
	"fmt"
	"io/ioutil"
)

func CopyFileContents() {
	fileContents, err := ioutil.ReadFile("data.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("new-data.dat", fileContents, 0x777)
	if err != nil {
		fmt.Println(err)
	}
}
