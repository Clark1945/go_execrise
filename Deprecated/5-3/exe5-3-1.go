package main

import "fmt"

func nums(id string, i ...int) { //參數不定函式

	total := 0
	for _, num := range i {
		total += num
	}
	fmt.Println(total)
}

func main() {
	nums("John", 100, 70)
	fmt.Println("---------------------------------------")
	best := []int{70, 80, 55, 77, 120}
	nums("Kelvin", best...)
	// nums()
	// nums(8763)
	mes := "Your mejesty"
	func(message string) {
		fmt.Println(message)
	}(mes)

	AGL := func() {
		fmt.Println("I wish to go Japan")
	}
	AGL()

}

//Printf = Print have format
