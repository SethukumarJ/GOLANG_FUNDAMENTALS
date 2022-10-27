package main

import (
	"fmt"
)

func main(){

	var annualIncome,tax float64
	var choice int
	
   for i := 0; i < 100; i++ {
	fmt.Print("\nEnter the annual income : ")
	fmt.Scan(&annualIncome)

	if(annualIncome >250000){
		tax = annualIncome * 0.5
		fmt.Print("\nThe income tax amout  is ",tax)
	} else if (annualIncome>500000){
		 
			tax =  annualIncome * .20
			fmt.Print("\nThe income tax amout  is ",tax)
					 
	} else if(annualIncome>1000000){
		tax = annualIncome * .30
		fmt.Print("\nThe income tax amout  is ",tax)
		
	} else {
		fmt.Print("\nNo income tax to pay!")
	}
	fmt.Print("\nPress 1 to continue!")
    fmt.Scan(&choice)
	if(choice!=1){
		break
	}
	
   }

 }