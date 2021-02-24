package utils

import "fmt"

func SliceDemo() {
	nos := make([]int, 3)
	fmt.Println(len(nos))
	fmt.Println(nos)
	nos[0] = 10
	nos[1] = 20
	nos[2] = 30
	fmt.Println(nos)

	nos = append(nos, 40)
	fmt.Println(nos)

	nos = append(nos, 50, 60, 70)
	fmt.Println(nos)
	fmt.Println(len(nos))

	fmt.Println(nos[2:])
	fmt.Println(nos[:2])
	fmt.Println(nos[2:5])
}
