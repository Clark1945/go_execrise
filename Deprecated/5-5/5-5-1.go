package main

import "fmt"

type calc func(int, int) string

func main() {
	calcalator(add, 5, 6)
}
func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("%d+%d=%d", i, j, result) //?
}
func calcalator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
