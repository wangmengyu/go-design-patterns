package chain

import "fmt"

type Doctor struct {
	Next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		//已经看过医生了
		fmt.Println("病人已经看过医生了")
		d.Next.Execute(p)
		return
	}

	//没有看过, 去看
	fmt.Println("看医生")
	p.DoctorCheckUpDone = true
	d.Next.Execute(p)
	return
}
func (d *Doctor) SetNext(next Department) {
	d.Next = next
}
