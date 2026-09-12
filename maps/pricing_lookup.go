package main

import "fmt"

// func main(){
// 	devSkills := map[string][]string{
// 		"Frontend" : {"React", "JavaScript", "Html"},
// 		"Backend" : {"Go", "Node", "Rust"},
// 	}

// 	fmt.Println(devSkills)
// }

// package main
// import "fmt"

func main() {
	pricing := map[string]int{
		"Basic": 499,
		"Pro":   999,
	}

	// Case 1: Jo plan exist karta hai
	price1, ok1 := pricing["Basic"]
	fmt.Println("Basic Price:", price1, "| Found:", ok1)

	// Case 2: Jo plan exist NAHI karta
	price2, ok2 := pricing["Ultra"]
	fmt.Println("Ultra Price:", price2, "| Found:", ok2)
}
