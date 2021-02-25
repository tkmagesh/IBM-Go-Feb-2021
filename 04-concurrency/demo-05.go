package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		fmt.Println("Writing 10")
		c <- 10
		fmt.Println("Writing 20")
		c <- 20
		fmt.Println("Writing 30")
		c <- 30
		fmt.Println("Writing 40")
		c <- 40
		fmt.Println("Writing 50")
		c <- 50
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
