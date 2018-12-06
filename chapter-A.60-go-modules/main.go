package main

import (
	"fmt"
	"github.com/novalagung/gubrak"
	"github.com/novalagung/myproject/models"
)

func main() {
	user := models.User{"u001", "Noval"}
	fmt.Println(user)
	fmt.Println(gubrak.RandomInt(10, 20))
}
