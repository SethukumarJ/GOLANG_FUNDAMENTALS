package main

import (
	"fmt"
)

type Shape interface {
	area() float32
	
}

type Circle struct {
	radius, pi float32
}
type Rectangle struct {
	length, breadth float32
}
type Square struct {
	side float32
}
type triangle struct {
	base, height float32
}

// Area funtions

func (c Circle) area() float32 {
	c.pi = 3.14

	
	fmt.Print("Enter the radius of the circle : ")
	fmt.Scan(&c.radius)
	return c.radius * c.radius * c.pi
}

func (r Rectangle) area() float32 {
	fmt.Print("\nEnter the length : ")
	fmt.Scan(&r.length)
	fmt.Print("\nEnter the breadth :")
	fmt.Scan(&r.breadth)
	return r.length * r.breadth
}
func (s Square) area() float32 {
	fmt.Print("Enter the lenght of side of square : ")
	fmt.Scan(&s.side)
	return s.side * s.side
}
func (t triangle) area() float32 {
	fmt.Print("Enter the baselength :")
	fmt.Scan(&t.base)
	fmt.Print("Enter the height :")
	fmt.Scan(&t.height)
	return (t.base * t.height) * .5
}

func main() {

	c := Circle{}
	r := Rectangle{}
	s := Square{}
	t := triangle{}

	var choice,contine int
    contine =5


	for contine==5 {
	fmt.Print("Enter 1 for circle\n 2 for rectangle\n 3 for square\n 4 for triangle\n")
	fmt.Scan(&choice)
	if(choice>5){
		fmt.Print("Enter valid option")
	}

	switch choice {

	case 1:
		Shapes := []Shape{c}

		for _, Shapes := range Shapes {
			fmt.Print("The area is : ", Shapes.area())
		}
	case 2:
		Shapes := []Shape{r}

		for _, Shapes := range Shapes {
			fmt.Print("The area is : ", Shapes.area())
		}

	case 3:
		Shapes := []Shape{s}

		for _, Shapes := range Shapes {
			fmt.Print("The area is : ", Shapes.area())
		}
	case 4:
		Shapes := []Shape{t}

		for _, Shapes := range Shapes {
			fmt.Print("The area is : ", Shapes.area())
		}
		

		
     
	}

	fmt.Print("\nEnter 5 to contine : ")
	fmt.Scan(&contine)

	if(contine!=5){
		break
	}
	 }

}
