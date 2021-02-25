package main

import (
	"fmt"
	"time"
)

func say(s string) {
	/* for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	} */
	time.Sleep(5000 * time.Millisecond)
}

func main() {
	go say("Hello")
	fmt.Println("World")
}
