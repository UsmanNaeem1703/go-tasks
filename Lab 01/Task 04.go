package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type student struct {
	rollno   int
	name     string
	address  string
	subjects []string
}

func (s student) Print(no int) {

	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), no, strings.Repeat("=", 25))

	fmt.Printf("student rollno   %d\n", s.rollno)
	fmt.Printf("student name     %s\n", s.name)
	fmt.Printf("student address  %s\n", s.address)
	fmt.Printf("student subjects %v\n", s.subjects)

	data := fmt.Sprintf("%d%s%s%v", s.rollno, s.name, s.address, s.subjects)
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("Hash: %x\n", hash)
}

func main() {

	students := []student{
		{rollno: 24, name: "Asim", address: "AAAAAA", subjects: []string{"Math", "Physics", "Chemistry"}},
		{rollno: 25, name: "Naveed", address: "BBBBBB", subjects: []string{"Biology", "History", "Geography"}},
	}

	for i, s := range students {
		s.Print(i)
	}
}
