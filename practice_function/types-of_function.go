package main

import "fmt"

func main() {

	func(a int, b int) {

		c := a + b
		fmt.Println("the sum is ", c)
	}(10, 24)

	mul := func(a int, b int) int {
		return a * b
	}
	fmt.Println("the product is ", mul(10, 24))
}

func init() {
	fmt.Println("I'll be invoked before the main function")
}

func multiply(a int, b int) int {
	return a * b
}
