package main

import "fmt"

//here it gives us error:
// Player spawned: Shashwat. Current weapon memory address: <nil>
// Attempting to upgrade weapon damage to 100...
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x2 addr=0x10 pc=0x104d04df4]

// goroutine 1 [running]:
// main.main()
//         /Users/shashwat/Goo/pointers/Method_practice/pointer_challenge_2.go:21 +0xb4
// exit status 2
// type Weapon struct{
// 	Name string
// 	Damage int
// }
// type Player struct{
// 	Name string
// 	EquippedWeapon *Weapon
// }

// func main(){
// 	hero := Player{
// 		Name: "Shashwat",
// 	}
// 	fmt.Printf("Player spawned: %s. Current weapon memory address: %v\n", hero.Name, hero.EquippedWeapon)
//     fmt.Println("Attempting to upgrade weapon damage to 100...")

// 	hero.EquippedWeapon.Damage = 100

// }

type Weapon struct {
	Name   string
	Damage int
}
type Player struct {
	Name           string
	EquippedWeapon *Weapon
}

func main() {
	gun := Weapon{
		Name:   "m416",
		Damage: 60,
	}
	heroFix := Player{
		Name:           "shashwat",
		EquippedWeapon: &gun,
	}
	fmt.Printf("Player %s entered the battleground.\n", heroFix.Name)
	fmt.Printf("Equipped weapon: %v\n\n", heroFix.EquippedWeapon.Name)

	fmt.Println("Updating weapon damage...")
	heroFix.EquippedWeapon.Damage = 100

	fmt.Printf("Weapon updated: %s's %s now deals %d damage.\n", heroFix.Name, heroFix.EquippedWeapon.Name, heroFix.EquippedWeapon.Damage)
}
