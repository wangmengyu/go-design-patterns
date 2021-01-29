package chain

import "fmt"

//挂号, 包含一个next部门
type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *Patient) {
	//检查是否已经执行过了
	if p.RegistrationDone {
		fmt.Println("病人已经挂号过了")
		r.Next.Execute(p)
		return
	}
	//还没挂号, 挂号. 执行下一步
	fmt.Println("病人挂号")
	p.RegistrationDone = true
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.Next = next

}
