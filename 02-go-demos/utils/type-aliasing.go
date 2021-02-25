package utils

import "strings"

type MyStr string
type MyFloat float64

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
