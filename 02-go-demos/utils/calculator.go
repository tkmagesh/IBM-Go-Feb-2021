package utils

/* var Message string = "[Default message]" */

func Add(x, y int) int {
	return x + y
}

func Divide(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//variadic function
func Sum(nos ...int) int {
	result := 0
	for _, val := range nos {
		result += val
	}
	return result
}

func EvenNumbers() func() int {
	n := 0
	return func() int {
		n += 2
		return n
	}
}
