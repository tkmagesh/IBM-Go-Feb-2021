package utils

import "fmt"

type Any interface{}

func PrintType(myVar Any /* interface{} */) {
	switch t := myVar.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type bool %T\n", t)
	case *Department:
		fmt.Printf("Type Department %T\n", t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
