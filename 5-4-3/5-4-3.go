package main

import "fmt"

func main() {
	max := 4

	counter := decrement(max)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
func decrement(i int) func() int {
	return func() int {
		if i > 0 {
			i--
		}
		return i
	}
}
