package main

import (
	"fmt"
)

func sumOfOdd(){

	var limit,sum int
	fmt.Print("Enter the limit : ")
	fmt.Scan(&limit)

	for i := 1; i <=limit; i++ {

		if(i%2!=0){
			sum += i
		}
	}

	fmt.Println(sum)

}

func main(){

	sumOfOdd()

}