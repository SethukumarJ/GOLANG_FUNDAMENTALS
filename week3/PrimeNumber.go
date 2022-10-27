package main

import (

	"fmt"
)
func PrimeOrNot(){

	var number,flag int

	fmt.Print("Enter the number to check whether prime or not : ")
	fmt.Scan(&number)

	for i := 2; i < number/2; i++ {

		if(number%i==0){
			flag=1
			break
		} 
		
	}


	if(flag!=0){
		fmt.Print("\nNot a prime number ")
	} else {
		fmt.Print("\nIs a prime number")
	}
}

func main(){

	PrimeOrNot()

}