package main

import "fmt"

type Employee struct {
	name   string
	salary int
	role   string
}

type Company struct {
	name      string
	employees [3]Employee
}

func printCompany(aCompany Company) {
	fmt.Printf("Company Name = %v; Employees = %v", aCompany.name, aCompany.employees)
}

func main() {
	aEmployee := Employee{
		name:   "Usman",
		salary: 80000,
		role:   "GO Developer",
	}
	bEmployee := Employee{
		name:   "Nawfal",
		salary: 90000,
		role:   "CISO",
	}
	cEmployee := Employee{
		name:   "Faheem",
		salary: 50000,
		role:   "Manager",
	}

	employees := [3]Employee{aEmployee, bEmployee, cEmployee}

	aCompany := Company{
		name:      "Tetra",
		employees: employees,
	}

	printCompany(aCompany)
}
