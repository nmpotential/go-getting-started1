package main

import (
	"fmt"
	"github.com/nmpotential/webservice/10-adding-variables-to-webservice/models"
)

func main()  {
	u := models.User{
		ID: 2,
		FirstName: "Ben",
		LastName: "Ten",
	}
	fmt.Println(u)
}
