package main 

import (

	"fmt"
)




func EvenInt (even chan int) {

	var num = 2

	for num <9 {
		
		even <- num

		num = num+2
	}
	close(even)
}


func OddInt (odd chan int) {

	var num = 1


	for num <9 {
		
		odd <- num

		num = num+2
	}
	close(odd)

}



func main() {
 
	even := make(chan int)
	odd := make(chan int)

	go EvenInt(even)
	go OddInt(odd)


	for {
 
		even , e := <- even
		odd , o  := <- odd

		if (e == false && o == false ) {
			break
		}

		fmt.Println(even , e, odd, o)
	}

}