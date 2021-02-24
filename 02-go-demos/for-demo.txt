package main

import "fmt"

func main() {
	//using 'for' as a 'while' construct
	i := 1
	for i <= 3 {
		i += 1
	}

	//traditional for loop
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	//using while(true) loop
	for {
		fmt.Println("Inside a indefinite loop")
		break
	}

}
