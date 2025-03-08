package main

import (
	"errors"
	"fmt"
)

func main() {
	pay := payDay(100, 25)
	fmt.Printf("周薪水%d\n", pay)
	pay = payDay(100, 200)
	fmt.Printf("周薪水%d\n", pay)
	pay = payDay(60, 25)
	fmt.Printf("周薪水%d\n", pay)
}

var (
	ErrHourRate   = errors.New("無效時薪")
	ErrHourWorked = errors.New("無效工時")
)

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourRate {
				fmt.Printf("時薪%d錯誤%v\n", hourlyRate, r) //r 是 panic裡面的訊息
			}
			if r == ErrHourWorked {
				fmt.Printf("工作時間%d錯誤%v", hoursWorked, r)
			}
		}
		fmt.Printf("計算周新的依據:公時%d/時薪%d\n", hoursWorked, hourlyRate)
	}()
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHourWorked)
	}
	if hoursWorked > 40 {
		hourover := hoursWorked - 40
		overtime := hourover * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overtime
	}
	return hoursWorked * hourlyRate
}
