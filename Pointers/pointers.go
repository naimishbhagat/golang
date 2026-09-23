package main

import "fmt"

func main(){
	age := 31
	var agePointer *int
	agePointer = &age
	fmt.Println(agePointer) 
	fmt.Println("Age: ",age)
	fmt.Println(*agePointer) 
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) int{
	*age = *age - 18
	return *age
}