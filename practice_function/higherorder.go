package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func addition(x int, y int) {
	c := x + y
	fmt.Println("Addition: ", c)
}

func main() {

	processOperation(5, 3, addition)
}
