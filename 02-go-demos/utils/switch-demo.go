package utils

import (
	"fmt"
	"time"
)

func SwitchDemo() {
	no := 2
	switch no {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend!")
	default:
		fmt.Println("Its a weekday!")
	}
}
