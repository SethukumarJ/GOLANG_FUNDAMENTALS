package main

import (
	"fmt"
)
func main() {

	M := Matrix{}
	M.displayMatrix(M.getMatrix())
	M.displayMatrix(M.getMatrix())
}

type Matrix struct {
	matrix [][]int
	size   int
}

func (m *Matrix) getMatrix() [][]int {
	var elementArray []int
	var element int
	fmt.Print("\nEnter the size of matrix : ")
	fmt.Scan(&m.size)
	fmt.Print("Enter the elements of first matrix : ")
	for i := 0; i < m.size; i++ {

		for j := 0; j < m.size; j++ {
			fmt.Scan(&element)
			elementArray = append(elementArray, element)
		}

		m.matrix = append(m.matrix, elementArray)
		elementArray = nil
	}
	return m.matrix
}

func (m Matrix) displayMatrix(matrix [][]int) {
	fmt.Print("The matris is : \n")
	for _, elementarray := range m.matrix {
		fmt.Println(elementarray)
	}
}


