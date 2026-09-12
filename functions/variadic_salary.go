package main

import "fmt"

func totalSalary(amounts ...int) int {
	total := 0
	for _, amount := range amounts {
		total += amount
	}
	return total
}
func main() {
	fmt.Println(totalSalary(21000, 4000, 5000))
}
