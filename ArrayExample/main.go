package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main(){
	userNames := make([]string,2, 5)
	userNames[0] = "Julie"
	userNames[1] = "Mac"
	
	userNames = append(userNames, "Max","Manuel")
	userNames = append(userNames, "Max","Manuel")
	fmt.Println(userNames)

	courseRatings := make(floatMap,3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["node"] = 4.7 // allocate memory

	courseRatings.output()
	//fmt.Println(courseRatings)
	for index, value := range userNames{
		fmt.Println(index)
		fmt.Println(value)
	}

	for index,value := range courseRatings{
		fmt.Println(index , value)	
	}
}