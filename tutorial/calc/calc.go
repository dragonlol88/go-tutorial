package calc

import (
	"fmt"
)


type Calc struct {
	Num1 , Num2 int
}

func (c Calc)  CalCuLate(action string)(interface{}, error){

	if action == "add" {
		return c.sum()
	} else if action == "sub" {
		return  c.subtract()
	} else if action == "mul" {
		return c.multiply()
	} else if action == "div" {
		return c.divide()
	} else {
		return nil, fmt.Errorf("Error: `%s` is not supported caculation", action)
	}
}
