package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR MESSAGE!!")
		fmt.Println("---------------")
		fmt.Println(err)
		//panic("sorry unable to continue!!")
	}

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
			writeBalanceToFile(accountBalance)

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
			writeBalanceToFile(accountBalance)

		default:
			println("Goodbye!")
			fmt.Println("Thank you for choosing our Bank!")
			return
		}
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		return
	}
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("failed to get data input")
	}

	balanceTicket := string(data)
	balance, err := strconv.ParseFloat(balanceTicket, 64)
	if err != nil {
		return 0, errors.New("input invalid")
	}
	return balance, err
}
