package department

import (
	"design_patterns/behavioral/chain_of_responsibility/patient"
	"fmt"
)

type Reception struct {
	next IDepartment // 必须要包含的下一个部门
}

func (r *Reception) Execute(p *patient.Patient) {
	p.RegistrationDone = true
	fmt.Println("Patient registration already done")
	r.next.Execute(p)
}
func (r *Reception) SetNext(d IDepartment) {
	r.next = d
}
