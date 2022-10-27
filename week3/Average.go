package main

import (
	"fmt"
)


func main(){


	var written,lab,ass int
	var grade float64

	fmt.Print("\nEnter writtern exam mark : ")
	fmt.Scan(&written)
	fmt.Print("\nEnter lab exam mark : ")
	fmt.Scan(&lab)
	fmt.Print("\nEnter assignment mark : ")
	fmt.Scan(&ass)

	grade =((float64(written)*70)/100) + ((float64(lab)*20)/100) + ((float64(ass)*10)/100)


	fmt.Print("\nThe overAll grade is : ",grade,"%")


}