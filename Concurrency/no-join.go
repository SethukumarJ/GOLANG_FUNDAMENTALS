package main

import (
	"fmt"
	"time"
)

func main() {


	go work()  // forkkPoint
	time.Sleep(100*time.Millisecond)
    fmt.Println("Waiting done! main function exited....")
	// expecting join point
    

}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("something have worked!")
}
