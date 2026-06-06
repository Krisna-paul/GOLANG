package main

import "fmt"

/*

func add(num1 int, num2 int) {

	sum := num1 + num2

	fmt.Println("the sum is ", sum)
}

func main() {

	a := 10
	b := 20

	add(a, b)
}

*/

/*
func add(num1 int, num2 int) int {

	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	diff := num1 - num2
	mul := num1 * num2

	return diff, mul
}

func PrintSomething() {
	fmt.Println("Education must be free for everyone")
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func main() {

	a := 10
	b := 20

	result := add(a, b)

	diff, mul := getNumbers(a, b)
	fmt.Println("the sum is ", result)
	fmt.Println("the difference is ", diff)
	fmt.Println("the product is ", mul)

	PrintSomething()
	sayHello("Krisna")
}

*/

//Why Function are needed?
//Functions are needed to avoid code duplication, improve readability, and make code more maintainable.

func PrintWelcomeMessage() {
	fmt.Println("Welcome to the Application,")
}

func getUserName() string {
	var name string
	fmt.Println("Please Enter Your Name- ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {

	var num1 int
	var num2 int
	fmt.Println("Enter Your First  Number -")
	fmt.Scanln(&num1)
	fmt.Println("Enter Your Second Number -")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func Display(name string, sum int) {
	fmt.Println("Hello ", name, " the sum of the two numbers is ", sum)

}

func sayGoodBye() {
	fmt.Println("Thank you for using the Application, Goodbye!")
}

func main() {

	PrintWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	Display(name, sum)
	sayGoodBye()
}
