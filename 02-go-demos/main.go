package main

import (
	"fmt"
	"go-demos/utils"
)

func main() {
	//Test()
	//utils.SliceDemo()
	//utils.MapDemos()
	//utils.WordStats()
	//utils.SwitchDemo()
	//fmt.Println(utils.Divide(30, 8))
	//fmt.Println(utils.Sum(3, 1, 4, 2, 5))

	/* name := ""
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	age := 0
	fmt.Scanf("%d", &age) */
	/* %s (string), %d (int), %g (float), %t(bool) */
	/*
		fmt.Println(age + 10)
		evenNoGenerator := utils.EvenNumbers()
		for index := 0; index < 5; index++ {
			fmt.Println(evenNoGenerator())
		}
	*/

	emp := new(utils.Employee)
	emp.Id = 100
	emp.Name = "Magesh"
	emp.IsEmployed = true
	emp.City = "Bangalore"
	fmt.Println(emp)

	emp2 := &utils.Employee{Id: 100, Name: "Magesh", IsEmployed: true, City: "Bangalore "}
	fmt.Println(emp2)

}
