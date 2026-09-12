package main

import "fmt"

type Device interface {
	Charge()
}
type Laptop struct {
	Brand string
}

func (l Laptop) Charge() {
	fmt.Println(l.Brand, "Laptop is charging...")
}

type Mobile struct {
	Model string
}

func (m Mobile) Charge() {
	fmt.Println(m.Model, "Mobile is charging...")
}

type SmartWatch struct {
	Watch_Model string
}

func (s SmartWatch) Charge() {
	fmt.Println(s.Watch_Model, "Smart Watch is charging...")
}
func PlugIntoCharger(gadget Device) {
	gadget.Charge()
}
func main() {
	mac := Laptop{Brand: "Apple"}
	phone := Mobile{Model: "Samsung"}
	watch := SmartWatch{Watch_Model: "Apple"}

	fmt.Println("----CHARGER START----")
	PlugIntoCharger(mac)
	PlugIntoCharger(phone)
	PlugIntoCharger(watch)
}
