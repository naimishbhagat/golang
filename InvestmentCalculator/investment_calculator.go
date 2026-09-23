package main

import (
	"fmt"
	"math"
)
const inflationRate float64 = 2.5
func main(){
	
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate,years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calculateFutureValue(investmentAmount, expectedReturnRate,years float64) (float64,float64){
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100,years)
	frv := fv /math.Pow(1+inflationRate/100, years)
	return fv, frv
}