package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Now()
	fmt.Println(t)

	fmt.Println("Location : ", t.Location())


	location1, _ := time.LoadLocation("America/New_York")
	fmt.Println(location1, t.In(location1))

	location2,_ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location2)
	fmt.Printf(  "Place : |%-16s| time : |%-40s|\n ",location2, now)

}
