package main

import "fmt"

func printWelcome() {
	fmt.Println(" Welcome to ATM System Simulation App ")

}

func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func greetUser(name string) {
	fmt.Printf("\nHello, %s! Welcome to your account.\n", name)
}

func showMenu() {
	fmt.Println("ATM MENU")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Choose option: ")
}

func checkBalance(balance float64) {
	fmt.Printf("Your current balance is: %.2f\n", balance)
}

func deposit(balance float64) float64 {
	var amount float64
	fmt.Print("Enter deposit amount: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return balance
	}

	balance += amount
	fmt.Printf("Deposited successfully: %.2f\n", amount)
	return balance
}

func withdraw(balance float64) float64 {
	var amount float64
	fmt.Print("Enter withdraw amount: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return balance
	}

	if amount > balance {
		fmt.Println("Insufficient balance!")
		return balance
	}

	balance -= amount
	fmt.Printf("Withdraw successful: %.2f\n", amount)
	return balance
}

func startATM(balance float64) float64 {
	var choice int

	for {
		showMenu()
		fmt.Scanln(&choice)

		if choice == 1 {
			checkBalance(balance)

		} else if choice == 2 {
			balance = deposit(balance)

		} else if choice == 3 {
			balance = withdraw(balance)

		} else if choice == 4 {
			fmt.Println("Thank you for using ATM. Goodbye!")
			return balance

		} else {
			fmt.Println("Invalid option! Try again.")
		}
	}
}

func main() {
	printWelcome()

	name := getUserName()
	greetUser(name)

	var balance float64 = 1000

	balance = startATM(balance)
}
