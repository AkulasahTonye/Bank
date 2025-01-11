package main

import (
	"example.com/Bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR MESSAGE!!")
		fmt.Println("---------------")
		fmt.Println(err)
		//panic("sorry unable to continue!!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

		var Choice int
		fmt.Print("Enter Your Choice: ")
		fmt.Scan(&Choice)

		switch Choice {
		case 1:
			fmt.Println("Your Account Balance is:", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Money Deposited:")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount // same thing as setting AccountBalance = AccountBalance + DepositAmount
			fmt.Println("Updated Account Balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("Withdraw Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Must be Greater than 0")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("You can't withdraw more than what you have.")
				continue
			}

			accountBalance -= withdrawalAmount // same thing as setting AccountBalance = AccountBalance - WithdrawalAmount
			fmt.Println("Updated Account Balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			println("Goodbye!")
			fmt.Println("Thank you for choosing our Bank!")
			return
		}
	}
}
