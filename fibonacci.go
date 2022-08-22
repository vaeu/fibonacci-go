package fibonacci

func Calculate(n int) int {
	back2, back1 := 0, 1

	if n == 0 {
		return 0
	}

	for i := 2; i < n; i++ {
		next := back1 + back2
		back2, back1 = back1, next
	}
	return back1 + back2
}
