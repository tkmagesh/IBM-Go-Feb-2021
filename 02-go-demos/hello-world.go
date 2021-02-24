package main

import (
	"fmt"
	"go-demos/utils"
)

//package level variables

//declaring 1 variable
/*
var id int
var name string
var age int
var location string
*/

//declaring more than one variables
/*
var (
	id int
	name string
	age int
	location string
)
*/

/*
var (
	id, age int
	name, location string
)
*/

//variable declaration with initialization
/*
var (
	id int = 100
	name string = "Magesh"
	age int = 40
	location string = "Bangalore"
)
*/

//data types can be omitted when declared and initialized together
/*
var (
	id = 100
	name ="Magesh"
	age = 40
	location = "Bangalore"
)
*/

var (
	id, name, age, location = 100, "Magesh", 40, "Bangalore"
)

func Test() {
	fmt.Println(utils.Message)
	utils.Greet()
	//
	/* var pi = 3.14 */

	/* := can be used ONLY in a function */
	pi := 3.14

	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)
	fmt.Println("Hello world!")
	fmt.Println(utils.Add(1000, 2000))
	fmt.Printf("pi = %v\n", pi)
	fmt.Printf("type of pi = %T\n", pi)

	//fmt.Printf("id = %d, name = %s, age = %d\n", id, name, age)
	s := fmt.Sprintf("id = %d, name = %s, age = %d\n", id, name, age)
	fmt.Println(s)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

/* func add(x, y int) int {
	return x + y
} */
