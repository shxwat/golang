package main

import "fmt"

// func changeNum(num *int){
// 	*num = 5
// 	fmt.Println("In changeNum:", *num)
// }

// func main(){
// 	num := 1
// 	changeNum(&num)
// 	fmt.Println("After changeNum in main:", num)
// }

// func updateName (name *string){
// 	*name = "Ram"
// 	fmt.Println("In updateName:", *name)

// }
// func main(){
// 	name := "Shyam"
// 	updateName(&name)
// 	fmt.Println("After update name in main:", name)
// }

func updatedChocolateBox(chocolate *int) {
	*chocolate = 50
	fmt.Println("number of chocolate in updatedChocolateBox is:", *chocolate)
}
func main() {
	chocolate := 10
	updatedChocolateBox(&chocolate)
	fmt.Println("number of chocolate in main is:", chocolate)
}
