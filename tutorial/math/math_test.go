package math

import (
	"testing"
)

// test sum
func TestSum(t *testing.T) {

	result, _ := Sum(1, 2, 3, 4, 6)
	if result != 16 {
		t.Errorf("result must be 16")
	}
}

// test divide
func TestDivide(t *testing.T) {
	result, _ := Divide(1, 2)
	if result != 0.5 {
		t.Errorf("result must be 0.5")
	}
}

// test divide
func TestDivideByZero(t *testing.T) {
	_, err := Divide(1, 0)
	if err == nil {
		t.Errorf("cannot be devided by zero")
	}
}

// test multiply
func TestMultiply(t *testing.T) {
	result, _ := Multiply(1, 2, 3, 4, 5, 6)
	if result != 720 {
		t.Errorf("result must be 720")
	}
}


// test substract
func TestSubtract(t *testing.T) {

	result, _ := Subtract(1, 2)
	if result != -1 {
		t.Errorf("result must be -1")
	}
}