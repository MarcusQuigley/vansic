package main

import (
	"fmt"
	"github/marcusquigley/firstproject/models"
)

func main() {
	u := models.User{
		Id:        1,
		Firstname: "Joe",
		Lastname:  "D",
	}

//	var users models.
	fmt.Printf("%v\n", u)
	

}
