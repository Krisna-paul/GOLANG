package main

import "fmt"

type Student struct {
	Name string
	Id   int
}

func (stu Student) show() {

	fmt.Println("Name", stu.Name)
	fmt.Println("Id", stu.Id)
}

func (stu Student) show2(a int) {
	fmt.Println("Name:", stu.Name)
	fmt.Println("Id:", stu.Id)
	fmt.Println("a:", a)
}

func main() {
	var stu1 Student

	stu1 = Student{
		Name: "Krisna",
		Id:   34,
	}

	stu1.show()

	stu2 := Student{
		Name: "Dewi",
		Id:   24,
	}
	stu2.show2(10)
}
