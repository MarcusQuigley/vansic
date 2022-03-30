package main

import (
	"fmt"
	"github/marcusquigley/firstproject/models"
)

func main() {
	u := models.User{
		Id:        1,
		Firstname: "Joe",
		Lastname:  "Donohue",
	}

	fmt.Printf("%v\n", u)

}
