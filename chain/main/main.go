package main

import "design_patterns/chain"

func main() {

	//创建一个病人
	p := chain.Patient{
		Name: "abc",
	}

	cashier := &chain.Cashier{}
	medical := &chain.Medical{}
	medical.SetNext(cashier)
	doctor := &chain.Doctor{}
	doctor.SetNext(medical)

	//去挂号处
	reception := &chain.Reception{}
	reception.SetNext(doctor)
	reception.Execute(&p)

}
