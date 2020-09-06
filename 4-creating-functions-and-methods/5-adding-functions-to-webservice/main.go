package main

import (
	"fmt"
	"github.com/nmpotential/webservice/3-working-with-collections/5-adding-variables-to-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Ben",
		LastName:  "Ten",
	}
	fmt.Println(u)
}