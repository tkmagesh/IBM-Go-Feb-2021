package main

import (
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

	/* emp := new(utils.Employee)
	emp.Id = 100
	emp.Name = "Magesh"
	emp.IsEmployed = true
	emp.City = "Bangalore"
	fmt.Println(emp)

	//Composite Literal Syntax
	emp2 := &utils.Employee{Id: 100, Name: "Magesh", IsEmployed: true, City: "Bangalore "}
	fmt.Println(emp2) */

	//i := 100

	//getting the address of i
	//j := &i

	//getting the value using the address of i
	//*j++

	//fmt.Println(i, *j)

	/* addressOfi := &i

	fmt.Printf("%d, %v, %d \n ", i, addressOfi, *addressOfi) */

	/*
		emp := new(utils.Employee)
		emp.Id = 100
		emp.Name = "Magesh"
		emp.IsEmployed = true
		emp.City = "Bangalore"
	*/

	/*
		emp := &utils.Employee{IsEmployed: true, City: "Bangalore"}
		//emp.Identity = utils.Identity{Id: 100, Name: "Magesh"}
		emp.Identity = utils.Identity{100, "Magesh"}
		emp.PrintIdentity() */

	//emp := &utils.Employee{Identity: utils.Identity{Id: 100, Name: "Magesh"}, IsEmployed: true, City: "Bangalore"}
	emp := &utils.Employee{utils.Identity{100, "Magesh"}, true, "Banglaore"}
	emp.PrintIdentity()
}
