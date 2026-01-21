package main

import "fmt"

func main() {
	var accountBalance = 1000.00
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do ? ")
	fmt.Println("1 Check balance")
	fmt.Println("2 Deposit money")
	fmt.Println("3 Withdraw money")
	fmt.Println("4 Exit ")

	var choice int
	fmt.Print("Your choice ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your balance is ", accountBalance)

	case 2:
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Add greater than 0")
			return
		}
		accountBalance += depositAmount

		fmt.Println("Balance update ! New amount: ", accountBalance)
	case 3:
		fmt.Print("Withdrawl amount: ")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)
		if withdrawlAmount > accountBalance {
			fmt.Println("Not possible ")
			return
		}
		accountBalance -= withdrawlAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	default:

		fmt.Println("TATA")
		return
	}

}
