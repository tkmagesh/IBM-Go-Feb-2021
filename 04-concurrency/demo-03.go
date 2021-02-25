package main

import "fmt"

func sum(list []int, ch chan int) {
	result := 0
	for _, val := range list {
		result += val
	}
	ch <- result
}

func main() {
	nos := []int{3, 6, 1, 7, 4, 2, 8, 9}
	first := nos[:len(nos)/2]
	second := nos[len(nos)/2:]
	/* firstCh := make(chan int)
	secondCh := make(chan int) */
	ch := make(chan int)
	go sum(first, ch)
	go sum(second, ch)
	total := <-ch + <-ch
	fmt.Println(total)
}
