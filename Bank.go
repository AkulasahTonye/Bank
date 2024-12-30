package main

import "fmt"

func main() {
	AccountBalance := 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var Choice int
		fmt.Print("Enter Your Choice: ")
		fmt.Scan(&Choice)
		fmt.Println("Your choice: ", Choice)

		switch Choice {
		case 1:
			fmt.Println("Your Account Balance:", AccountBalance)
			fmt.Scan(&AccountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Money Deposited:")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			AccountBalance += depositAmount // same thing as setting AccountBalance = AccountBalance + DepositAmount
			fmt.Println("Updated Account Balance:", AccountBalance)

		case 3:
			fmt.Print("Withdraw Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Must be Greater than 0")
				continue
			}
			if withdrawalAmount > AccountBalance {
				fmt.Println("You can't withdraw more than what you have.")
				continue
			}

			AccountBalance -= withdrawalAmount // same thing as setting AccountBalance = AccountBalance - WithdrawalAmount
			fmt.Println("Updated Account Balance:", AccountBalance)

		default:
			println("Goodbye!")
			fmt.Println("Thank you for choosing our Bank!")
			return

		}
		
	}

}
