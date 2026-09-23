package main

import "fmt"

func main(){
	result := add(1,2)
	fmt.Println(result)
}

func add[T int|float64|string](a, b T) T {
	return a+b
}
// func add(a, b interface{}) interface{} {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)

// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}

// 	aString, aIsString := a.(string)
// 	bString, bIsString := b.(string)

// 	if aIsString && bIsString {
// 		return aString + bString
// 	}

// 	return nil
// }