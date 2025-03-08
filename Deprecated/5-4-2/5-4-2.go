package main

import "fmt"

func main() {
	increment := incrementor()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
