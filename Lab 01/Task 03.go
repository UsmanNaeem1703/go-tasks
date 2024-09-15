package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

// Print method to display each student's data
func (ls *StudentList) Print() {
	for _, student := range ls.list {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("Roll Number: %d\n", student.rollnumber)
		fmt.Printf("Name: %s\n", student.name)
		fmt.Printf("Address: %s\n", student.address)
	}
	fmt.Println(strings.Repeat("=", 20)) // Print the final line of equals
}

func main() {
	studentList := new(StudentList)
	studentList.CreateStudent(24, "Usman", "AAAAAA")
	studentList.CreateStudent(25, "Nawfal", "BBBBBB")
	studentList.Print() // Print all student data
}
