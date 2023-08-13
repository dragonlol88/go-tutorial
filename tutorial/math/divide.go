package math

import "fmt"

func Divide(num1, num2 float32) (float32, error)  {
	if num2 == 0 {
		return num1, fmt.Errorf("Error: cannot divide %0.2f by %0.2f", num1, num2)
	}
	result := num1 / num2
	return result, nil
}
