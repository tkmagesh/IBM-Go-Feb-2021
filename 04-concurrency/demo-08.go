package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	end := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-end:
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
