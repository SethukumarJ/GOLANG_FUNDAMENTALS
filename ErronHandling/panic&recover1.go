package main

import "fmt"

func main() {

	const one = 2
    defer exitSafe()
	if one != 1 {
		panic("one can't be 1, something bad happened ")
	}

}

func exitSafe() {
	if err := recover(); err != nil {
		fmt.Println("error recovered")
	}
}
