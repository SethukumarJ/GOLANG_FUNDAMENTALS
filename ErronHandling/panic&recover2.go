package main

import (
	"errors"
	"fmt"
)

func main () {
	chain(2)
	chain(4)
	chain(6)
	chain(0)
	chain(8)

}

var result = 1

func chain(n int) {
	defer func () {
		if err:=recover(); err!=nil {
			fmt.Println("panic happened")
		}
	}()
	if n ==0 {
		panic(errors.New("O can't be used!"))
	} else {
		result *= n
		fmt.Println(n)
	}
}