package calc

import (
	"fmt"
)


type Calc struct {
	Num1 , Num2 int
}

func (c Calc)  CalCuLate(action string)(interface{}, error){

	if action == "add" {
		return sum(&c)
	} else if action == "sub" {
		return  subtract(&c)
	} else if action == "mul" {
		return multiply(&c)
	} else if action == "div" {
		return divide(&c)
	} else {
		return nil, fmt.Errorf("Error: `%s` is Not supported caculation", action)
	}
}



func divide(c *Calc) (float32, error)  {

	if c.Num2 == 0 {
		return float32(c.Num2), fmt.Errorf("Error: caNNot divide %d by %d", c.Num1, c.Num2)
	}
	result := float32(c.Num1) / float32(c.Num2)
	return result, nil
}


func multiply(c *Calc) (int, error)  {
	result := c.Num1 * c.Num2
	return result, nil
}


func subtract(c *Calc) (int, error)  {
	result := c.Num1 - c.Num2
	return result, nil
}


func sum(c *Calc) (int, error)  {
	result := c.Num1 + c.Num2
	return result, nil
}

