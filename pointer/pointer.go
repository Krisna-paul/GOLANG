package main

import "fmt"

func main() {
	//pointer or memory of adress

	x := 10
	fmt.Println("Value of x:", x)

	p := &x // &= ampersand = adress of

	*p = 20 // * = dereference operator...assign value to the address

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value at address p:", *p)

}
