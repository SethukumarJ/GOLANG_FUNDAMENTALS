package main

import (
	"enc/student"
	"fmt"
)

func main() {

	s := student.Student{Rollno: 22}
	s.SetName("Sethukumar")
	s.SetCourse("btech cse")
	fmt.Println("Student Name : ", s.GetName())
	fmt.Println("Student Course : ", s.GetCourse())
	fmt.Println("Student Rollno : ", s.GetRollno())
}
