package main

import "fmt"

// func main(){
// 	candidates := []string{"Ramesh", "Suresh"}
// 	candidates = append(candidates, "Shashwat", "Rohit", "Rahul")

// 	fmt.Println(candidates)
// }

//sort top 3 using slice end

// func main(){
// 	candidates := []string{"Ramesh", "Suresh", "Mohan", "Sohan", "Rohan"}
// 	topThree := candidates[0:3]
// 	fmt.Println("The top three candidates are:", topThree)

// }

//append using this (...)

// func main(){
// 	frontendDevs := []string{"Shashwat", "Rahul", "Rohit"}
// 	backendDevs := []string{"Rohan", "Mohan", "Sohan"}
// 	allDevs := append(frontendDevs, backendDevs...)

// 	fmt.Println(allDevs)
// }

func showSkills(name string, skills ...string) {
	fmt.Printf("%s's skills: %v\n", name, skills)
}

func main() {
	showSkills("Shashwat", "React", "Go")
}
