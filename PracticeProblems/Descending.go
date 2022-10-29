package main

import (
	"fmt"
	"sort"
)

func descendingSort() {

	var array []int
	var size, element int

	fmt.Print("Enter the size fo the array : ")
	fmt.Scan(&size)

	fmt.Print("Enter the elements of the array : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&element)
		array = append(array, element)
	}

	sort.Ints(array)
	for i := len(array)-1; i >= 0; i--{
		fmt.Print(array[i]," ")
	}

}

func main() {
	descendingSort()

}
