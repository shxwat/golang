package main

import "fmt"
import "time"

func main() {
	// basic switch
	// day:= "Sunday"
	// switch day{
	// case "Saturday", "Sunday":
	// 	fmt.Println("Weekends")
	// case "Monday":
	// 	fmt.Println("WeekDay")
	// default:
	// 	fmt.Println("Normal Day")
	// }

	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("good morning")
	case hour < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")

	}
}
