package main

import (
	"fmt"
)

func addition(n int) {

	var number, sum int

	fmt.Print("\nEnter the number to add : ")
	fmt.Scan(&number)
	sum = number + n

	fmt.Print("The sum is ", sum)

}

func substraction(n int) {

	var number, difference int

	fmt.Print("\nEnter the number to substract : ")
	fmt.Scan(&number)
	difference = n - number

	fmt.Print("The difference is ", difference)

}

func multiplication(n int) {

	var number, product int

	fmt.Print("\nEnter the number to substract : ")
	fmt.Scan(&number)
	product = number * n

	fmt.Print("\nThe product is ", product)

}

func division(n int) {
	var number int
	var quotient float64
	fmt.Print("\nEnter the number to divide : ")
	fmt.Scan(&number)
	quotient = float64(n) / float64(number)

	fmt.Print("\nThe quotient is ", quotient)

}

func main() {

	for {
		var choice, con, element int
		fmt.Print("\nEnter the number you want to aplly the mathematical operations : ")
		fmt.Scan(&element)
		fmt.Print("\nEnter 1 for addition \n2 for substraction \n3 for multiplication \n4 for division :")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addition(element)
			break
		case 2:
			substraction(element)
			break
		case 3:
			multiplication(element)
			break
		case 4:
			division(element)
			break
		}

		fmt.Print("\nDo you want to continue? press 1 to continue\n")

		fmt.Scan(&con)
		if con != 1 {
			break
		}
	}

}
