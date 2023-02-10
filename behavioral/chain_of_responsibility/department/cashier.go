package department

import (
	"design_patterns/behavioral/chain_of_responsibility/patient"
	"fmt"
)

type Cashier struct {
	next IDepartment
}

func (d *Cashier) Execute(p *patient.Patient) {

	fmt.Println("Payment done")
	p.PaymentDone = true
	//d.next.Execute(p) 这里就没有下一步了
}

func (d *Cashier) SetNext(next IDepartment) {
	d.next = next
}
