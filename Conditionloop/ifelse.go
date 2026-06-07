package main

import "fmt"

func main() {
	age := 5

	if age < 18 {
		fmt.Println("you are not eligible to be married")
	} else if age == 18 {
		fmt.Println("you are eligible to be married")
	} else {
		fmt.Println("you are already married")
	}
}
