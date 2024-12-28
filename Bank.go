package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var Choice int

	fmt.Print("your choice: ")
	fmt.Scan(&Choice)

	fmt.Println("Your choice: ", Choice)

}
