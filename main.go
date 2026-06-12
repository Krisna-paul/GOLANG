package main

/*
import (
	"fmt"
	"golang/mathlib"
)

func main() {
	fmt.Println("showing custom package")
	mathlib.Add(10, 30)
}
*/

import "fmt"

type User struct {
	Name string
	Age  int
}

type Customer struct {
	Name string
	Age  int
}

func main() {
	var user1 User

	user1 = User{"paul", 30}

	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)

	user2 := Customer{
		Name: "Dewi",
		Age:  24,
	}
	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)
}
