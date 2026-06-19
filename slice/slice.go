package main

import "fmt"

func main() {

	//1.slice from existing array
	//2. Slice from slice
	/*
		arr := [6]string{"this", "is", "a", "go", "interview", "question"}

		fmt.Println(arr)

		s := arr[1:4]
		fmt.Println(s)
		s1 := s[1:2]
		fmt.Println(s1)

		fmt.Println("length of s1 is ", len(s1),"capacity of s1 is ", cap(s1))
	*/

	//slice literal

	s := []int{1, 2, 3}

	fmt.Println("slice:", s, "len:", len(s), "capasity:", cap(s))

	// slice with make function with len

	s1 := make([]int, 3)
	s1[0] = 4
	fmt.Println("slice:", s1, "len:", len(s1), "capasity:", cap(s1))

	// make function with len capacity

	s2 := make([]int, 3, 10)

	s2[0] = 10
	s2[1] = 9
	s2[2] = 34

	fmt.Println("slice:", s2, "len:", len(s2), "capasity:", cap(s2))

	//empty slice or nil slice

	var s3 []int

	fmt.Println(s3)

}
