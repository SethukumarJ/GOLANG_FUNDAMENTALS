package main

import (
	"fmt"
)

func PaliondromeChecker() {

	var input string
	var flag int
	fmt.Print("Enter the string to check whether paliodrome or not : ")
	fmt.Scan(&input)
	length := len(input)

	for i := 0; i < length; i++ {

		if (input[i] != input[(length-i)-1] ){
			flag++
			break;
		} 
		

	}

	if flag >= 1 {
		fmt.Print("The entered string is not a paliondrome")
	} else {
		fmt.Print("The entered string is paliondrome")
	}

}

func main() {

	PaliondromeChecker()

}
