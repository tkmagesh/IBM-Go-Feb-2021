package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine1(wg *sync.WaitGroup) {
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("goRoutine1 completed")
	wg.Done()
}

func goRoutine2(wg *sync.WaitGroup) {
	time.Sleep(7000 * time.Millisecond)
	fmt.Println("goRoutine2 completed")
	wg.Done()
}

func goRoutine3(wg *sync.WaitGroup) {
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("goRoutine3 completed")
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	wg := &sync.WaitGroup{}
	//wg.Add(3)
	wg.Add(1)
	go goRoutine1(wg)
	wg.Add(1)
	go goRoutine2(wg)
	wg.Add(1)
	go goRoutine3(wg)

	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println("All go routines are completed")
}
