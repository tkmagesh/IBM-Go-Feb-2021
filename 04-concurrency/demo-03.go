package main

import "fmt"

func sum(list []int) int {
	result := 0
	for _, val := range list {
		result += val
	}
	return result
}

func main() {
	nos := []int{3, 6, 1, 7, 4, 2, 8, 9}
	result := sum(nos)
	fmt.Println(result)
}
