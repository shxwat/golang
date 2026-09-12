package main

import "fmt"

func checkAge(age int) (string, error) {
	if age < 18 {
		return "", fmt.Errorf("entry denied: provided age is %d", age)
	}
	return "Entry allowed. Enjoy the ride.", nil
}
func main() {
	age, err := checkAge(12)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Success:", age)
}
