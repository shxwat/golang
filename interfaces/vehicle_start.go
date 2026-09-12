package main

import "fmt"

type Vechile interface {
	Start()
}

type Car struct {
	Brand string
}

func (c *Car) Start() {
	fmt.Println(c.Brand, "car started successfully.")

}
func DriveNow(v Vechile) {
	v.Start()
}
func main() {
	myCar := Car{Brand: "LEXUS"}
	DriveNow(&myCar)
}
