package main

import "fmt"

func getDeveloper() (string, string) {
	return "Shashwat", "Go Backend"
}
func main() {
	name, role := getDeveloper()

	fmt.Println(name)
	fmt.Println(role)
}
