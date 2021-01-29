package chain

import "fmt"

type Medical struct {
	Next Department
}

func (m *Medical) Execute(p *Patient) {

	if p.MedicineDone {
		fmt.Println("已经配过药了")
		m.Next.Execute(p)
		return
	}
	fmt.Println("配药")
	p.MedicineDone = true
	m.Next.Execute(p)
	return

}
func (m *Medical) SetNext(next Department) {
	m.Next = next
}
