package math

func Sum(nums ... float32) (float32, error)  {
	var result float32
	for _, num := range nums {
		result += num
	}
	return result, nil
}

