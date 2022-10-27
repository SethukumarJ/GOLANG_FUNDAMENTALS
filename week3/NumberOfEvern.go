package main

import (
	"fmt"
)

func evenCount(){

	package main

	import (
		"fmt"
	)
	
	func evenCount(){
	
	  var array []int
	  var size,element,count int
	
	  fmt.Print("Enter the size fo the array : ")
	  fmt.Scan(&size)
	
	  fmt.Print("Enter the elements of the array : ")
	  for i := 0; i < size; i++ {
		fmt.Scan(&element)
		if(element%2==0){
			count++
		}
		array = append(array,element)
	}
	
	fmt.Print("Number of even numbers is ",count)
	
	
	}
	
	func main(){
		evenCount()
	
	}

fmt.Print("Number of even numbers is ",count)


}

func main(){
	evenCount()

}