package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	Done := make(chan struct{})
	go func() { // forkkPoint
		work()
		Done <- struct{}{}
    }()

	<-Done //Join point
	fmt.Println("The excecution time : ", time.Since(now))
	fmt.Println("Waiting done! main function exited....")

}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("something have worked!")

}
