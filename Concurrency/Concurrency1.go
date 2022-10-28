package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	task1()
	task2()
	task3()
	task4()

	fmt.Println("elapsed : ", time.Since(now))

}

func task1() {
	time.Sleep(time.Microsecond * 100)
	fmt.Println("task 1")
}

func task2() {
	time.Sleep(time.Microsecond * 200)
	fmt.Println("task 2")
}

func task3() {

	fmt.Println("task 3")
}

func task4() {
	time.Sleep(time.Microsecond * 100)
	fmt.Println("task 4")
}
