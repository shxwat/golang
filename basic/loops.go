package main

import "fmt"

// func main(){
// 	score := 80
//     //basic if-else
// 	if score >= 80 {
// 		fmt.Println("Passed!!")
// 	}else {
// 		fmt.Println("Failed!!")
// 	}

// 	if num := 10; num%2 == 0 {
// 		fmt.Println(num, "is Even")
// 	} else{
// 		fmt.Println(num, "is Odd")
// 	}
// }

func main() {
	// the classic for loop (Start, conditon, step)
	// for i :=1; i <= 3; i++{
	// 	fmt.Println("Classic:", i)
	// }

	// the while loop, in go there is no keyword while. so we direct write condition it becomes while loop
	// count := 1
	// for count <= 3{
	// 	fmt.Println("While Style:", count)
	// 	count++
	// }

	//infinite loop, in go no need to write any thing just give brackets..
	// for {
	// 	fmt.Println("infinite loop....")
	// 	break // for stopping the infinite loop we need to add break...
	// }

	//The range loop
	name := []string{"Rahul", "Shashwat", "Aakash"}
	for index, name := range name {
		fmt.Println("On index", index, "name:", name)
	}
}
