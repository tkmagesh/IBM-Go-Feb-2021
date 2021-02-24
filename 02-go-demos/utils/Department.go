package utils

import "fmt"

type Department struct {
	Identity
}

func (d *Department) Display() {
	fmt.Printf("Id=%d, Name=%s\n", d.Id, d.Name)
}

/* func (d *Department) PrintIdentity() {
	fmt.Printf("Id = %d, Name = %s\n", d.Id, d.Name)
} */
