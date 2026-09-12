package main

import "fmt"

func calculateSalary(basicSalary, bonus int) int {
	return basicSalary + bonus
}
func main() {
	total := calculateSalary(21000, 4000)
	fmt.Println(total)
}
