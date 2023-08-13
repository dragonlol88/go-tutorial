package math

func Multiply(nums ... float32) (float32, error)  {
	var result float32 = 1

	for _, num := range nums {
		result *= num
	}
	return result, nil
}
