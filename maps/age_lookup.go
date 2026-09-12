package main

import "fmt"

func main() {
	// phoneBook := make(map[string]int)

	// phoneBook["Raju"] = 987654321
	// phoneBook["Shamu"] = 88888888

	// fmt.Println("Raju number:", phoneBook["Raju"])

	//map shprthand with value itself

	ages := map[string]int{
		"Raju":  22,
		"Shamu": 25,
	}

	// fmt.Println(ages)

	check, ok := ages["Raju"]
	if ok == true {
		fmt.Println("The age of Raju is:", check)
	} else {
		fmt.Println("Raju is not in map")
	}

	delete(ages, "Raju")
}
