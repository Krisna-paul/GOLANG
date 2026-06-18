package main

import "fmt"

func main() {
	arr := [6]string{"this", "is", "a", "go", "interview", "question"}

	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)
	s1 := s[1:2]
	fmt.Println(s1)

	fmt.Println("length of s1 is ", len(s1))
	fmt.Println("capacity of s1 is ", cap(s1))
}
