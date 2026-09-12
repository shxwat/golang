package main

import "fmt"

// func multipleReturn()(int, string){
//     marks := 99
// 	name := "Rohit"

// 	return marks, name
// }
// func main(){
// 	myMarks, myName := multipleReturn()
//     fmt.Println("The name of student is", myName, "who get", myMarks , "marks")
// }

func checkResult(score int) (string, bool) {
	if score >= 80 {
		return "Selected", true
	}
	return "Rejected", false
}

// func main(){
// 	status, isQualified := checkResult(55)
// 	fmt.Println(status, isQualified)

// 	onlyStatus, _ := checkResult(50)
// 	fmt.Println("Just the status:", onlyStatus)
// }

func checkStatus(marks int) (string, bool) {
	if marks >= 30 {
		return "Passed", true
	}
	return "Failed", false
}

// func main(){
// 	status, isPassed := checkStatus(50);
// 	fmt.Println(status, isPassed)

//     onlyStatus, _ := checkStatus(20)
// 	fmt.Println("Only Status", onlyStatus)
// }

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}
func main() {
	q, r := divide(10, 5)
	fmt.Println("Quotient is:", q, "and the Remainder is:", r)
}
