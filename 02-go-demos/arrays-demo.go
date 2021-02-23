package main

import "fmt"

func main() {
	var nos [5]int
	fmt.Println(nos)

	nos[2] = 10
	fmt.Println(nos)

	newNos := [5]int{10, 20, 30, 40, 50}
	fmt.Println(newNos)
	fmt.Println(len(newNos))
	fmt.Println(newNos[:2])
	fmt.Println(newNos[2:])
	fmt.Println(newNos[2:3])
}
