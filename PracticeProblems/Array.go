package main

import (

	"fmt"
)

func getArray() (array []int){
   var element,size1 int
   var array1 []int

   fmt.Print("Enter the size of the array : ") 
   fmt.Scan(&size1)

   for i := 0; i < size1; i++ {
	 fmt.Scan(&element)
	  array1 = append(array1, element)
   }

   return array1
}


func displayArray(array []int){
    fmt.Print("The array is : ",array," ")
}


func main(){

displayArray(getArray())

}