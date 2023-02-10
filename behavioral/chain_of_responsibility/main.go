package main

import (
	"design_patterns/behavioral/chain_of_responsibility/department"
	"design_patterns/behavioral/chain_of_responsibility/patient"
	"fmt"
)

func main() {
	p := &patient.Patient{
		Name: "abc",
	}
	fmt.Println(p)

	cash := &department.Cashier{}
	medical := &department.Medical{}
	medical.SetNext(cash)

	doctor := &department.Doctor{}
	doctor.SetNext(medical)

	recep := &department.Reception{}
	recep.SetNext(doctor)

	recep.Execute(p)

}
