package main

import "fmt"

func main() {
	scores := []int{10, 20, 30}

	scores = append(scores, 40)
	scores = append(scores, 50, 60)

	// fmt.Println(scores)

	cut1 := scores[1:4]
	fmt.Println("cut slices", cut1)

	cut2 := scores[:3]
	fmt.Println("starting slices", cut2)
}
