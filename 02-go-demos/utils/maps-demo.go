package utils

import "fmt"

func MapDemos() {
	m := make(map[string]int)
	m["foo"] = 10
	m["bar"] = 20

	fmt.Println(m)
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

	colors := map[string]int{"red": 1, "green": 2, "blue": 3}
	colors["yellow"] = 4
	fmt.Println(colors)
	fmt.Println(len(colors))

	/* _, found := colors["purple"] //=> val, bool
	fmt.Println(found) */

	colors["purple"] = 5
	if _, found := colors["purple"]; found {
		fmt.Println("purple exits")
	} else {
		fmt.Println("purple doesnot exist")
	}
}
