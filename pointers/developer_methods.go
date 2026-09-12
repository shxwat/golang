package main

import "fmt"

type Developer struct {
	Name      string
	Role      string
	BugsFixed int
}

func (d Developer) Status() {
	fmt.Printf("Hi, I am %s, a %s, and I have fixed %d bugs", d.Name, d.Role, d.BugsFixed)
}
func (d *Developer) FixBug() {
	d.BugsFixed++
	fmt.Printf("Hi, I am %s, a %s, and I have fixed %d bugs", d.Name, d.Role, d.BugsFixed)
}
func main() {
	employee := Developer{
		Name:      "Shashwat",
		Role:      "SDE-1",
		BugsFixed: 5,
	}
	employee.Status()
	employee.FixBug()
}
