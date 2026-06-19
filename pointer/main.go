package main

import "fmt"

func print(numbers *[3]int) {
	fmt.Println(numbers)

}

type User struct {
	Name string
	ID   int
	Age  int
}

func main() {

	arr := [3]int{1, 2, 3}
	print(&arr)

	object := User{
		Name: "krisna",
		ID:   2008034,
		Age:  26,
	}

	fmt.Println(object)
	fmt.Println(object.Age)

	p := &object

	fmt.Println(*p)
	fmt.Println(p.Name)
	fmt.Println(p.ID)
	fmt.Println(p.Age)

}

// pass by value  = copy of the value is passed to the function
// pass by reference = address of the value is passed to the function
