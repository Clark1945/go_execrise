package main

import "fmt"

type salaryFunc func(int, int) int

/*
	func main() {
		devSalary := salary(50, 2080, developerSalary)
		bossSalary := salary(150000, 205000, managerSalary)
		fmt.Printf("打雜薪資 %d", devSalary)
		fmt.Printf("經理薪資 %d", bossSalary)
	}

	func salary(x, y int, f salaryFunc) int {
		pay := f(x, y)
		return pay
	}

	func managerSalary(baseSalary, bonus int) int {
		return baseSalary + bonus
	}

	func developerSalary(hourlyRate, hoursWorked int) int {
		return hourlyRate * hoursWorked
	}
*/
func main() {
	add := calculator("+")
	substract := calculator("-")
	fmt.Println(add(5, 6))
	fmt.Println(substract(10, 5))
}
func calculator(op string) func(int, int) int {
	switch op {
	case "+":
		return func(i, j int) int {
			return i + j
		}
	case "-":
		return func(i, j int) int {
			return i - j
		}
	}
	return nil
}
