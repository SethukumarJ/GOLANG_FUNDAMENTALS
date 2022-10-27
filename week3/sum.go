package main

import (
	"fmt"
)


func sum() (sum float64){
    var number1 int
	var number2 float64

	fmt.Print("Enter number 1: ")
	fmt.Scan(&number1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&number2)

	sum = float64(number1) + number2

	return sum


}

func main(){
	fmt.Println(sum())
}
