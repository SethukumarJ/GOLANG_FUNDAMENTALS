package main

import (
	"fmt"
)


func displayWeek(){
    var  day int
	weekDays := []string{"SUNDAY","MONDAY","TUESDAY","WEDNESDAY","THURSDAY","FRIDAY","SATURDAY"}

	for i := 0; i < 8; i++ {

	fmt.Print("Enter a number from 1 to 7 to show the respective day : ")
	fmt.Scan(&day)
    
	switch(day){
	
		case 1:
	    	fmt.Println(weekDays[day-1])
		case 2:
	    	fmt.Println(weekDays[day-1])
		case 3:
	    	fmt.Println(weekDays[day-1])
		case 4:
	    	fmt.Println(weekDays[day-1])
		case 5:
	    	fmt.Println(weekDays[day-1])
		case 6:
	    	fmt.Println(weekDays[day-1])
		case 7:
	    	fmt.Println(weekDays[day-1])
		case 8:
	    	fmt.Println("invalid")
		}
	}
}
func main(){

displayWeek()
}