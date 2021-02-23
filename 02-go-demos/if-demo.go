package main

import "fmt"

func main() {
	if no, x := 6, 10; no%2 == 0 {
		fmt.Printf("%d is even & x = %d\n", no, x)
	} else {
		fmt.Printf("%d is odd\n", no)
	}
}
