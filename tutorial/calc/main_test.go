package calc

import (
	"testing"
)

// test sum
func TestSum(t *testing.T) {

	c := Calc{1, 2}
	result, _ := sum(&c)
	if result != 3 {
		t.Errorf("result must be 3")
	}
}

// test divide
func TestDivide(t *testing.T) {
	c := Calc{1, 2}
	result, _ := divide(&c)
	if result != 0.5 {
		t.Errorf("result must be 0.5")
	}
}

// test divide
func TestDivideByZero(t *testing.T) {
	c := Calc{1, 0}

	_, err := divide(&c)
	if err == nil {
		t.Errorf("cannot be devided by zero")
	}
}

// test multiply
func TestMultiply(t *testing.T) {
	c := Calc{1, 2}
	result, _ := multiply(&c)
	if result != 2 {
		t.Errorf("result must be 2")
	}
}


// test substract
func TestSubtract(t *testing.T) {
	c := Calc{1, 2}
	result, _ := subtract(&c)
	if result != -1 {
		t.Errorf("result must be -1")
	}
}