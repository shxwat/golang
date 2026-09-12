package main

import "fmt"

// func main(){
// 	ch := make(chan string)

// 	go func(){
// 		fmt.Println("Background: I am sending data in background") //2nd

// 		ch <- "Hello Boss!!"

// 		fmt.Println("Background: this line run only when the boss take the data off.. ") //3rd

// 	}()

// 	fmt.Println("Main: I am waiting for take data from main...")  //1st

// 	msg := <-ch

// 	fmt.Println("Main: I get it ->", msg) //4th
// }

func main() {

	ch := make(chan string)

	fmt.Println("Sender: preparing to send an item.")

	ch <- "Here are the noodles."

	go func() {
		fmt.Println("Receiver: ready to receive the item.")
		msg := <-ch
		fmt.Println("Receiver: received ->", msg)
	}()

	fmt.Println("Sender: task completed.")
}
