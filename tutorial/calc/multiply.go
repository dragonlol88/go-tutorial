package calc

func (c Calc) multiply() (int, error)  {
	result := c.Num1 * c.Num2
	return result, nil
}
