package main

import (
	"fmt"
)

func adjascent() {

	var arrayOne, arrayTwo []int
	var size, element int

	fmt.Print("Enter the size of the array : ")
	fmt.Scan(&size)

	fmt.Print("Enter the elements of the array : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&element)
		arrayOne = append(arrayOne, element)
	}

	for i := 0; i < len(arrayOne)-1; i++ {

		element = arrayOne[i] + arrayOne[i+1]

		arrayTwo = append(arrayTwo, element)

	}

	fmt.Print(arrayTwo)

}

func main() {
	adjascent()
}
