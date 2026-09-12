package main

import "fmt"

type Storage interface {
	SaveUser(name string)
}
type MemoryDB struct {
	Users []string
}

func (m *MemoryDB) SaveUser(name string) {
	m.Users = append(m.Users, name)
	fmt.Println("User saved safely in memory:", name)
}
func RegisterNewUser(db *MemoryDB, name string) {
	db.SaveUser(name)
}
func main() {
	myDB := MemoryDB{
		Users: []string{},
	}

	RegisterNewUser(&myDB, "Shashwat")
	RegisterNewUser(&myDB, "Aman")

	fmt.Println("\nThe list of Users in Database:", myDB.Users)
}
