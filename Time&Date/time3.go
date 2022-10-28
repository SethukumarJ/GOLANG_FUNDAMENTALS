package main

import (
	"fmt"
	"time"
)

func main() {

	date := time.Date(2020,10,28,11,15,0, 0, time.Now().Location())
    fmt.Println(date)

	next_date := date.AddDate(0,3,1)
	fmt.Println(next_date)
}
