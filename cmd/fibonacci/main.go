package main

import (
	"flag"
	"fmt"

	"github.com/vaeu/fibonacci-go"
)

const (
	MIN = 0
	MAX = 10
)

func main() {
	min := flag.Int("min", MIN, "initial number to start iterating from")
	max := flag.Int("max", MAX, "last number to iterate through")
	flag.Parse()

	for i := *min; i <= *max; i++ {
		fmt.Println(fibonacci.Calculate(i))
	}
}
