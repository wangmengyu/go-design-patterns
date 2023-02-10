package department

import (
	"design_patterns/behavioral/chain_of_responsibility/patient"
	"fmt"
)

type Doctor struct {
	next IDepartment
}

func (d *Doctor) Execute(p *patient.Patient) {

	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next IDepartment) {
	d.next = next
}
