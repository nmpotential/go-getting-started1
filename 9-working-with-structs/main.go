package main

import "fmt"

func main() {
	//Define the struct
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	//Initialise object using struct
	var u user
	u.ID = 1
	u.FirstName = "Ben"
	u.LastName = "Ten"
	fmt.Println(u)

	u2 := user{ ID: 1,
			FirstName: "Ben",
			LastName: "Ten",
	}
	fmt.Println(u2)
}
