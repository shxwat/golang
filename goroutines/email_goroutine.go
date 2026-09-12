package main

import "fmt"

// type User struct {
// 	Name string
// 	Age  int
// }

// func updateUser(user User) {
// 	user.Name = "Rahul"
// 	user.Age = 30
// }

// func main() {
// 	user1 := User{
// 		Name: "Shashwat",
// 		Age:  22,
// 	}

// 	updateUser(user1)

//		fmt.Println(user1.Name)
//		fmt.Println(user1.Age)
//	}
func sendEmail(name string) {
	fmt.Printf("Email sent to %s!\n", name)
}

func main() {
	done := make(chan struct{})

	go func() {
		sendEmail("Shashwat")
		close(done)
	}()

	fmt.Println("Result calculated!")
	<-done
}
