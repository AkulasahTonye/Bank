package main

import "fmt"

func main() {
	AccountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var Choice int
	fmt.Print("Enter Your Choice: ")
	fmt.Scan(&Choice)
	fmt.Println("Your choice: ", Choice)

	if Choice == 1 {
		fmt.Println("Your Account Balance:", AccountBalance)
		fmt.Scan(&AccountBalance)

	} else if Choice == 2 {
		var depositAmount float64
		fmt.Print("Money Deposited:")
		fmt.Scan(&depositAmount)
		AccountBalance += depositAmount // same thing as setting AccountBalance = AccountBalance + DepositAmount
		fmt.Println("Updated Account Balance:", AccountBalance)

	} else if Choice == 3 {
		fmt.Print("How much do you want to Withdraw:")
		var MoneyWithdraw float64
		fmt.Scan(&MoneyWithdraw)
		AccountBalance -= MoneyWithdraw
		fmt.Println("Your Current Account Balance is:", AccountBalance)
	}

}
