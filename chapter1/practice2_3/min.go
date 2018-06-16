package practice2_3

func Min(ints ...int) int {
	min := ints[0]

	for _, i := range ints[1:] {
		if i < min {
			min = i
		}
	}

	return min
}
