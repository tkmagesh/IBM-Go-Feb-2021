package main

import "fmt"

func main() {
	if no := 5; no%2 == 0 {
		fmt.Printf("%d is even\n", no)
	} else {
		fmt.Printf("%d is odd\n", no)
	}
}
