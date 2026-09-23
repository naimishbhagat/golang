package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)
const accountBalanceFile = "balance.txt"


func main(){
	var accountBalance,err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		//panic("Can't continue, sorry.")
	}
	//for someCondition {}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach Us 24/7",randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is : ", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount:",accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Amount to be Withdrawn: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New Amount:",accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}
		// if choice == 1 {
		// 	fmt.Println("Your balance is : ", accountBalance)
		// 	continue
		// } else if choice == 2{
		// 	fmt.Print("Your Deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		return
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New Amount:",accountBalance)
		// } else if choice == 3{
		// 	fmt.Print("Amount to be Withdrawn: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		return
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You cant withdraw more than you have.")
		// 		return
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New Amount:",accountBalance)
		// } else if choice == 4{
		// 	fmt.Println("Goodbye!")
		// 	//return
		// 	break	
		// }
	}
}

