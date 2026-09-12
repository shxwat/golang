package main

import "fmt"

type Player struct {
	Name   string
	Health int
}

func (p *Player) TakeDamage(damage int) {
	p.Health = p.Health - damage
	fmt.Printf("%s took %d damage. Current health: %d\n", p.Name, damage, p.Health)
}
func main() {
	hero := Player{
		Name:   "Shashwat",
		Health: 100,
	}
	fmt.Printf("The health of player %s before fight is %d\n", hero.Name, hero.Health)
	hero.TakeDamage(30)
	fmt.Printf("The health of player %s after fight is %d\n", hero.Name, hero.Health)
}
