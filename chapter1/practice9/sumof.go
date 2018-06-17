package practice9

func SumOf(a, b int) int {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}
