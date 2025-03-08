package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourRate   = errors.New("無效時薪")
	ErrHourWorked = errors.New("無效工時")
)

func main() {
	pay := payDay(81, 50)
	fmt.Println(pay)

}

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("工時%d\n時薪%d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHourWorked)
	}
	return hoursWorked * hourlyRate
}
