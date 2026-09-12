package main

import "fmt"

type Player struct {
	Name   string
	Health int
}

func TakeDamage(p *Player, damage int) {
	p.Health = p.Health - damage
	fmt.Printf("%s took %d damage. Current health: %d\n", p.Name, damage, p.Health)

}
func main() {
	hero := Player{
		Name:   "Shashwat",
		Health: 100,
	}
	enemy := Player{
		Name:   "Rahul",
		Health: 100,
	}
	fmt.Println("The health of player before fight is:", hero.Health)
	fmt.Println("The health of enemy before fight is:", enemy.Health)

	TakeDamage(&hero, 30)
	TakeDamage(&enemy, 50)

	fmt.Printf("The health of player %s is: %d\n", hero.Name, hero.Health)
	fmt.Printf("The health of enemy %s is: %d\n", enemy.Name, enemy.Health)
}
