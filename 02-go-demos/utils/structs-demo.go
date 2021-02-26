package utils

import "fmt"

type Employee struct {
	Identity
	IsEmployed bool
	City       string
	Orgs       []string
}

func (e *Employee) ToggleEmploymentStatus() {
	e.IsEmployed = !e.IsEmployed
	fmt.Printf(" Inside toggle fn : %v\n", e)
}

func NewEmployee(id int, name string, isEmployed bool, city string) *Employee {
	return &Employee{Identity{id, name}, isEmployed, city, make([]string, 10)}
}

func (e *Employee) Display() {
	fmt.Printf("Id=%d, Name=%s, IsEmployed=%t, City=%s\n", e.Id, e.Name, e.IsEmployed, e.City)
}
