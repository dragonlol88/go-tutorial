package calc

import (
	"testing"
)

// test sum
func TestSum(t *testing.T) {

	c := Calc{1, 2}
	result, _ := c.sum()
	if result != 3 {
		t.Errorf("result must be 3")
	}
}

// test divide
func TestDivide(t *testing.T) {
	c := Calc{1, 2}
	result, _ := c.divide()
	if result != 0.5 {
		t.Errorf("result must be 0.5")
	}
}

// test divide
func TestDivideByZero(t *testing.T) {
	c := Calc{1, 0}

	_, err := c.divide()
	if err == nil {
		t.Errorf("cannot be devided by zero")
	}
}

// test multiply
func TestMultiply(t *testing.T) {
	c := Calc{1, 2}
	result, _ := c.multiply()
	if result != 2 {
		t.Errorf("result must be 2")
	}
}


// test substract
func TestSubtract(t *testing.T) {
	c := Calc{1, 2}
	result, _ := c.subtract()
	if result != -1 {
		t.Errorf("result must be -1")
	}
}