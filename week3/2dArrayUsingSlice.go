package main

import (
	"fmt"
)

func getArray() (slice [][]int, slice2 [][]int, s int) {

	var size int
	var elementArray []int
	var array [][]int
	var array2 [][]int

	fmt.Print("Enter the size of the matrix 1 : ")
	fmt.Scan(&size)
	fmt.Print("Enter the elements : ")

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {
			var element int
			fmt.Scan(&element)
			elementArray = append(elementArray, element)
		}

		array = append(array, elementArray)

		elementArray = nil

	}

	fmt.Print("\nEnter the elements of matrix 2 : ")

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {
			var element int
			fmt.Scan(&element)
			elementArray = append(elementArray, element)
		}

		array2 = append(array2, elementArray)

		elementArray = nil

	}

	return array, array2, size

}

func addArray(slice1 [][]int, slice2 [][]int, s int) (sumArray [][]int) {
	size := s
	var elementArray []int
	var sArray [][]int

	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {
			elementArray = append(elementArray, slice1[i][j]+slice2[i][j])

		}

		sArray = append(sArray, elementArray)
		elementArray = nil

	}

	return sArray

}

func displayArray(sumArray [][]int) {
	fmt.Print("The sum of matrix is : \n")
	fmt.Println(sumArray)
}

func main() {

	displayArray(addArray(getArray()))
}
