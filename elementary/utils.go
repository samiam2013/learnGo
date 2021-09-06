package elementary

func triangular(input int64) int64 {
	sum := int64(0)
	for input > 0 {
		sum = sum + input
		input = input - 1
	}
	return sum
}
