package utils

import "fmt"

type Displayable interface {
	Display()
}

type NewDisplayable interface {
	Display()
}

func PrintData(obj Displayable) {
	obj.Display()
}

func NewPrintData(obj NewDisplayable) {
	fmt.Println("From newPrintData")
	obj.Display()
}
