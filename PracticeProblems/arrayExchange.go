package main

import (
	
	"fmt"
)

func interchange(){
    var size ,temp,append1,append2 int
	fmt.Print("Enter the size of the arrays : ")
    fmt.Scan(&size)
	var arrayOne ,arrayTwo []int
	

   
   fmt.Print("Enter the elements of array 1 :")
   for i := 0; i < size; i++ {
	
	fmt.Scan(&append1)
	arrayOne = append(arrayOne, append1)

   }
   fmt.Print("Enter the elements of array 2 :")
   for i := 0; i < size; i++ {
	fmt.Scan(&append2)
	arrayTwo = append(arrayTwo, append2)

   }


   for i := 0; i < size; i++ {
    
	 temp = arrayOne[i]
	 arrayOne[i] = arrayTwo[i]
	 arrayTwo[i] = temp
}

fmt.Print("array one : ")
for i := 0; i < len(arrayOne); i++ {
 
 fmt.Print(arrayOne[i]," ")

}
fmt.Print("\narray Two : ")
for i := 0; i < len(arrayTwo); i++ {
 
 fmt.Print( arrayTwo[i]," ")

}




}


func main(){

	interchange()


}