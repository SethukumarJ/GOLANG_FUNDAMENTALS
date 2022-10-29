package main

import (
	"fmt"
)

func printPattern() {
	var limit int

	fmt.Print("Enter the limit : ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {

			print(j)

		}
		println(" ")
	}
}

func main() {

	printPattern()

}
