package main

import (
	"fmt"
)

func main() {
	x := []int{2, 2, 5, 6, 8}
	sqr := func(i ...int) int {
		total := 0
		for _, y := range i {
			total += y
		}
		return total
	}
	fmt.Printf("%d\n", sqr(x...))
	g := 0
	increment := func() int {
		g += 1
		return g
	}
	fmt.Println(increment())
	fmt.Println(increment())
	g += 100
	fmt.Println(increment())
}
