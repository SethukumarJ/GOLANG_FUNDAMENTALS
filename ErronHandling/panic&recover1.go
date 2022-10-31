package main


import (
	"errors"
	"fmt"
)

func main() {

	const one = 2
	
	fmt.Println("corrent int")
	defer exitSafe()
	func() {
		
		if one != 1 {
			panic(errors.New("one can't be any other numbers"))
		}
	
	}()
}

func exitSafe() {
	if err := recover(); err != nil {
		fmt.Println("recovered!")
	}
}
