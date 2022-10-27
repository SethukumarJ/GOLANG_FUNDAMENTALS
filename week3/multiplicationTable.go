package main

import (
	"fmt"
)

func table(){
	var number,product int

	fmt.Print("Enter a number : ")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {

       product = i*number 

		fmt.Printf("\n%d x %d = %d", i,number,product)
		
	}
}

func main(){

	table()

}