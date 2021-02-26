package calculator

import "testing"

func Test_Shoud_Add_2_Numbers(t *testing.T) {
	x, y, expected := 10, 20, 30

	result := Add(x, y)

	if result != expected {
		t.Errorf("Adding %d and %d results in %d but expected %d", x, y, result, expected)
	}
}
