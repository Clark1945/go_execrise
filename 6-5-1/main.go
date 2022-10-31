package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println("main DEFER EXIST")
	test()
	fmt.Println("無印")
	msg := "Bye"
	message(msg)
	fmt.Println("無印")
	// nums := []int{2, 4, 6, 8, 10}
	// total := 0
	// for i := 0; i <= 10; i++ {
	// 	total += nums[i]
	// }

	// fmt.Println(total)
}
func message(msg string) {
	defer fmt.Println("message defer")
	if msg == "Bye" {
		panic(errors.New("出四出四了"))
	}
}
func test() {
	defer fmt.Println("test defer")
	msg := "Bye"
	message(msg)
}
