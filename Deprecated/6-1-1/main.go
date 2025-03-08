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
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourRate
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHourWorked
	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		hoursRegular := hoursWorked - hoursOver
		return hoursRegular*hourlyRate + hoursOver*hourlyRate*2, nil
	}
	return hoursWorked * hourlyRate, nil
}
