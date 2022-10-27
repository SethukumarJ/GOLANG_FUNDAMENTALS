package main

import (
	"fmt"
	
)


func gradeFinder(){

	var total float64

	fmt.Print("Enter the total percentage obtained : ")
	fmt.Scan(&total)

	if(total>= 90){
		fmt.Println("A grade")
	} else if(total >= 80){
		fmt.Println("B grade")
	} else if(total >= 70){
		fmt.Println("C grade")
	} else if(total >= 60){
		fmt.Println("D grade")
	} else if(total >= 50){
		fmt.Println("E grade")
	} else if(total > 100){
		fmt.Println("Invalid input")
	} else{
		fmt.Println("failed")
	}


}

func main () {
     var count int
	fmt.Print("Enter the number of marks you want to check the grade : ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		gradeFinder()
	}
}