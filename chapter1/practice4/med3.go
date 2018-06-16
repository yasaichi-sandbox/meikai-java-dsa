package practice4

func Med3(a, b, c int) int {
	if a >= b {
		if b >= c {
			return b
		} else if a <= c {
			return a
		} else {
			return c
		}
	} else if a > c {
		return a
	} else if b > c {
		return c
	} else {
		return b
	}
}
