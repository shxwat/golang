package main

import (
	"fmt"
	// "strconv" // use to convert string to int
)

func main() {
	// myMap := map[string]int {"Apple": 5, "Banana": 10}

	// value, ok := myMap["Apple"]

	// if ok {
	// fmt.Println("Avaiable in map with value:", value)
	// } else {
	// fmt.Println("Not availabe in map")
	// }

	// value, ok = myMap["cherry"]

	// if ok {
	// 	fmt.Println("Available in map with vlaue", value)
	// }else {
	// 	fmt.Println("Not available in map")
	// }

	// if num, err := strconv.Atoi("123"); err == nil{
	// 	fmt.Println("Successfully converted to num:", num)
	// }else {
	// 	fmt.Println("Conversion failed", err)
	// }
	// if num, err := strconv.Atoi("abc"); err == nil{
	// 	fmt.Println("Successfully converted to num:", num)
	// }else {
	// 	fmt.Println("Conversion failed", err)
	// }

	var i interface{} = "Hello"
	s, ok := i.(string)
	if ok {
		fmt.Println("'i' is a string:", s)
	} else {
		fmt.Println("'i' is not a string")
	}

}
