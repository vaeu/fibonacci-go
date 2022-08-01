package fibonacci

func Calculate(n int) int {
	if n > 1 {
		return Calculate(n-1) + Calculate(n-2)
	}
	return n
}
