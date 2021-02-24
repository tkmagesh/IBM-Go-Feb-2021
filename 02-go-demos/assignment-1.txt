/* Create an array of 100 integers
Populate them with random values (range 1 to 100)
Find out the number of odd and even values from the array and print them */

package main

import (
	"fmt"
	"math/rand"
)

var my_array [100]int

func randomInt(min, max int) int {

	return min + rand.Intn(max)
}

func main() {

	for i := 0; i < 100; i++ {
		my_array[i] = randomInt(1, 100)
	}
	even_count := 0
	odd_count := 0
	for i := 0; i < len(my_array); i++ {

		if my_array[i]%2 == 0 {
			even_count += 1
		} else {
			odd_count += 1
		}
	}
	fmt.Println(my_array)
	fmt.Println(even_count, odd_count)
}
