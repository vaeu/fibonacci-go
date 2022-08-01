package main

import (
	"fmt"

	"github.com/vaeu/fibonacci-go"
)

const (
	MIN = 0
	MAX = 10
)

func main() {
	for i := MIN; i <= MAX; i++ {
		fmt.Println(fibonacci.Calculate(i))
	}
}
