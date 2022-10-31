package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(4)
	go task1(&wg)
	
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
    
	wg.Wait()
	fmt.Println("elapsed : ", time.Since(now))

}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
	
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2")
}

func task3(wg *sync.WaitGroup) {
    defer wg.Done()
	fmt.Println("task 3")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 4")
}
