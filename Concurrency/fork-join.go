package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    

	var wg sync.WaitGroup
    now := time.Now()
	wg.Add(1)
    
	go func() {    // forkkPoint
		defer wg.Done()
		work() 
		
	}()

	wg.Wait() //Join point
	fmt.Println("The excecution time : ", time.Since(now))
    fmt.Println("Waiting done! main function exited....")

}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("something have worked!")

	
}
