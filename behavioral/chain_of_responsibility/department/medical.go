package department

import (
	"design_patterns/behavioral/chain_of_responsibility/patient"
	"fmt"
)

type Medical struct {
	next IDepartment // 包含下一个节点。
}

func (d *Medical) Execute(p *patient.Patient) {

	fmt.Println("medical done")
	p.MedicineDone = true
	d.next.Execute(p)
}

func (d *Medical) SetNext(next IDepartment) {
	d.next = next
}
