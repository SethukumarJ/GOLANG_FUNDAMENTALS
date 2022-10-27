// Go program to illustrate how to compare
// string using compare() function
package main

import (
	"fmt"
	"strings"
)

func main() {

	// Comparing string using Compare function
	fmt.Println(strings.Compare("gfg", "Geeks"))

	fmt.Println(strings.Compare("GeeksforGeeks",
		"GeeksforGeeks"))

	fmt.Println(strings.Compare("Geeks", " GFG"))

	fmt.Println(strings.Compare("GeeKS", "geeks"))
	fmt.Println(strings.Compare("SethU", "Sethu"))


	str := "Sethu"
    byte_arr := []byte(str)
    fmt.Println(byte_arr)

	str1 := "sethu"
    byte_arr2 := []byte(str1)
    fmt.Println(byte_arr2)

}
