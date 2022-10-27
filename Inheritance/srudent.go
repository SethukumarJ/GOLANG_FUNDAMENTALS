package main

import "fmt"

type vehicle struct {
	name  string
	color string
}

type bike struct {
	vehicle
	name  string
	model int
}

func main() {

	v := vehicle{"audi", "white"}

	b := bike{v, "yamaha", 2000}

	fmt.Println(b.vehicle.name)

}
