package main

import "fmt"

const (
	MIN = 0
	MAX = 10
)

func calculate(n int) int {
	if n > 1 {
		return calculate(n-1) + calculate(n-2)
	}
	return n
}

func main() {
	for i := MIN; i <= MAX; i++ {
		fmt.Println(calculate(i))
	}
}
