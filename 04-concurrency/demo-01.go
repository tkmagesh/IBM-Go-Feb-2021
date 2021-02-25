package main

import (
	"fmt"
	"time"
)

func say(s string, c1 chan int, c2 chan int) {

	for i := 0; i < 5; i++ {
		<-c1
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		c2 <- i
	}

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go say("Hello", ch1, ch2)
	go say("World", ch2, ch1)
	ch1 <- 0
	var input string
	fmt.Scanln(&input)
}
