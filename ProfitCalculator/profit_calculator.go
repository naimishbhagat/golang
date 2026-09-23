package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile = "profit.txt"
func writeBalanceToFile(ebt, profit, ratio float64){
	formattedText := fmt.Sprint("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt,profit,ratio)
	os.WriteFile(profitFile, []byte(formattedText), 0644)
}

func main(){
	revenue,err := getUserInput("Revenue: ")
	if err != nil{
		fmt.Println(err)
		return
	}
	expenses,err := getUserInput("Expenses: ")
	if err != nil{
		fmt.Println(err)
		return
	}
	tax_rate,err := getUserInput("Tax Rate: ")
	if err != nil{
		fmt.Println(err)
		return
	}
	
	ebt, profit, ratio := calcualteFinancials(revenue, expenses, tax_rate)
	writeBalanceToFile(ebt, profit, ratio)
	formattedEBT := fmt.Sprintf("Earning Before Tax: %.1f\n", ebt)
	formattedEAT := fmt.Sprintf("Earning After Tax: %.1f\n", profit)
	fmt.Print(formattedEBT,formattedEAT)
	fmt.Println("Ratio: ", ratio)
	
}

func outputText(text string){
	fmt.Print(text)
	
}

func calcualteFinancials(revenue, expenses, tax_rate float64)(float64, float64, float64){
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate /100)
	ratio := ebt / profit
	return ebt, profit,ratio
}

func getUserInput(infoText string) (float64, error){
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0{
		fmt.Println("Invalid amount. Must be greater than 0.")
		return 0, errors.New("Value must be a positive number.")
	}
	return userInput, nil
}
