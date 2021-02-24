package utils

import "fmt"

type Identity struct {
	Id   int
	Name string
}

func (i *Identity) PrintIdentity() {
	fmt.Printf("Id = %d, Name = %s\n", i.Id, i.Name)
}
