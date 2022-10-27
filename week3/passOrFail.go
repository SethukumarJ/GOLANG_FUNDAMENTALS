package main

import (
	"fmt"
)

func PassOrFail() {

	var subject, result string
	var mark float64

	fmt.Print("Enter the subject : ")
	fmt.Scan(&subject)
	fmt.Print("Emter the marks : ")
	fmt.Scan(&mark)

	if mark < 50 {
		result = "\nyou are failed in " + subject
		fmt.Println(result)
	} else if mark >= 50 && mark <= 100 {
		result = "you are passed in " + subject
		fmt.Println(result)
	} else {
		fmt.Println("Invalid mark entered")
	}
}

func main() {
	var count int
	fmt.Print("\nEnter number of subjects : ")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {

		PassOrFail()

	}

}
