package main

import "fmt"

type Laptop struct {
	Brand string
	RAM   int
}

func (mac *Laptop) showDetails(newRam int) {
	mac.RAM = newRam
	fmt.Println("Laptop brand:", mac.Brand, "| RAM:", mac.RAM)

}

//	func UpgradeRAM(mac *Laptop){
//		mac.RAM = 32
//		fmt.Println("The upgraded ram of my Mac is Gonna!!:", mac.RAM)
//	}
func main() {
	myLaptop := Laptop{
		Brand: "Apple",
		RAM:   16,
	}

	myLaptop.showDetails(32)

	fmt.Println("Laptop brand:", myLaptop.Brand)
	fmt.Println("Upgraded RAM:", myLaptop.RAM)
}
