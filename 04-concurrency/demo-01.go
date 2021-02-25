package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan string) {
	/* for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	} */
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(s)
	ch <- "done"
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)
	fmt.Println("World")
	<-ch
}
