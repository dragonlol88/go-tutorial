package calc

import "fmt"

func (c Calc)divide() (float32, error)  {

	if c.Num2 == 0 {
		return float32(c.Num2), fmt.Errorf("Error: cannot divide %d by %d", c.Num1, c.Num2)
	}
	result := float32(c.Num1) / float32(c.Num2)
	return result, nil
}
