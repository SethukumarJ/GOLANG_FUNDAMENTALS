package main

import (
	"fmt"
	"time"
)


func main() {


	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
    
	fmt.Println(t.Format(time.RubyDate))

    dateFromString := "2002-03-12T11:35:39+05:30"
	
	parseTimeFromString,_ := time.Parse(time.RFC3339,dateFromString)

	fmt.Println(parseTimeFromString)
	 

	
}