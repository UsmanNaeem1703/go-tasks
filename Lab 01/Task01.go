// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Doctor struct {
	number   int
	name     string
	patients []string
}

func printDoc(aDoctor Doctor) {
	fmt.Printf("Doctor ID = %v; Doctor Name = %v; Patients = %v", aDoctor.number, aDoctor.name, aDoctor.patients)
}

func main() {
	fmt.Println("Hello World")
	aDoctor := Doctor{
		number:   6,
		name:     "Usman",
		patients: []string{"A", "B", "C", "D"},
	}
	printDoc(aDoctor)
}
