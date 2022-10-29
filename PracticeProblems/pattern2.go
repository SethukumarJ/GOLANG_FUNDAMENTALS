package main

import (
	"fmt"
)

func main(){
	var n int
	count :=1
	fmt.Print("Enter number of rows :")
	fmt.Scan(&n)

	for i := 1; i <=n; i++ {

		for j := 1; j<=i; j++ {

			fmt.Print(count," ")
			count++

		}
		fmt.Print("\n")
		
	}

}