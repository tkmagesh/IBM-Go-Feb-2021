package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func test(ch chan int) {
	fmt.Println("starting test fn")
	fmt.Println("Reading data from the channel")
	val := <-ch
	fmt.Println(val)
}

func main() {
	ch := make(chan int)
	fmt.Println("executing test fn concurrently")
	go test(ch)
	fmt.Println(("Writing data into the channel"))
	ch <- 10
	fmt.Println("main fn exits")
}
