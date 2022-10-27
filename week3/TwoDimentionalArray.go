package main

import (
	"fmt"
	
)


func sumOfarrays(){
    var size int
	fmt.Print("Enter the size of matrix : ")
	fmt.Scan(&size)

	var arrayOne [3][3]int
    var arrayTwo [3][3]int
	var sumArray [3][3]int

	
	fmt.Print("\nEnter the elements of matrix one : ")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			
			fmt.Scan(&arrayOne[i][j])
			}
		
	}

	fmt.Print("\nEnter the elements of matrix two : ")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			
			fmt.Scan(&arrayTwo[i][j])
			}
		
	}
	
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			

			sumArray[i][j] =arrayOne[i][j] + arrayTwo[i][j]
			
			}
		
	}
	fmt.Println("\nMatrix One : ")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			
			fmt.Print(arrayOne[i][j]," ")
			}
			fmt.Print("\n")
		
	}
	fmt.Println("\nMatrix two : ")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			
			fmt.Print(arrayTwo[i][j]," ")
			}
			fmt.Print("\n")
		
	}


fmt.Println("\nThe sum of Matrix one and Matrix two is : ")
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			
			fmt.Print(sumArray[i][j]," ")
			}
			fmt.Print("\n")
		
	}

}

func main(){
	sumOfarrays()

}