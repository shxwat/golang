package main

import "fmt"

type User struct {
	Name string
	Age  int
	Role string
}

func printUser(user User) {
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Role:", user.Role)
}

func main() {
	user1 := User{
		Name: "Shashwat",
		Age:  22,
		Role: "Go Developer",
	}
	printUser(user1)
}
