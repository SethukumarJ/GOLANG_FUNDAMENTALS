package main

import (
	"fmt"
)

func getArray() ([][]int, int) {

	var size int
	var elementSlice []int
	var matrix [][]int

	fmt.Print("Enter the size of matrix :")
	fmt.Scan(&size)
	fmt.Print("Enter the elements of the matrix :")
	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {

			var element int
			fmt.Scan(&element)

			elementSlice = append(elementSlice, element)
		}

		matrix = append(matrix, elementSlice)
		elementSlice = nil
	}

	return matrix, size
}

func addArray(slice1 [][]int, size1 int, slice2 [][]int, size int) {
	var elementSlice []int
	var matrix [][]int
	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {

			elementSlice = append(elementSlice, slice1[i][j]+slice2[i][j])

		}

		matrix = append(matrix, elementSlice)
		elementSlice = nil

	}

	fmt.Println(matrix)

}

func main() {

	g1, s := getArray()
	g2, s := getArray()

	addArray(g1, s, g2, s)

}
