package main

import (
	"fmt"

)

func main() {
	 

	var name string

	fmt.Print("enter the name : ")
	fmt.Scan(&name)
    fmt.Println(name)

	fmt.Printf("%s",name[3])

	// for i := len(name)-1; i>=0; i-- {

	// 	defer fmt.Print(name[i])
	// }

}
