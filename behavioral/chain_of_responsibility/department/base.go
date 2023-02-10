package department

import "design_patterns/behavioral/chain_of_responsibility/patient"

type IDepartment interface {
	Execute(patient *patient.Patient) // 处理病人
	SetNext(department IDepartment)   // 指定转交的下一个部门
}
