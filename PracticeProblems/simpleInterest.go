package main

import (
	"fmt"
)

func simpleInterest() (SI float64) {
	var principalAmount int
	var interestRate, N float64
	fmt.Print("Enter the principal amount : ")
	fmt.Scan(&principalAmount)
	fmt.Print("Enter the interest rate : ")
	fmt.Scan(&interestRate)
	fmt.Print("Enter number of years")
	fmt.Scan(&N)
    
	SI=(float64(principalAmount)*interestRate*N)/100

	return
}

func main() {

	finalResult := simpleInterest()

	fmt.Print("Simple interest for the given inputs is ", finalResult)

}
