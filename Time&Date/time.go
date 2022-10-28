package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	fmt.Println(t)

	fmt.Println("location :" , t.Location())

	tf := t.Format("02-06-2006")
	fmt.Println(tf)
	fmt.Println("date : " , t.Format("01-06-2006"))



	
}
