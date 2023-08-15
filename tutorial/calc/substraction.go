package calc


func (c Calc )subtract() (int, error)  {
	result := c.Num1 - c.Num2
	return result, nil
}
