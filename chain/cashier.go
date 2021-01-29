package chain

import "fmt"

type Cashier struct {
	Next Department
}

func (c *Cashier) SetNext(next Department) {
	c.Next = next
}
func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("已经付过钱了")
		//最终
		return
	}

	p.PaymentDone = true
	fmt.Println("付好钱了")

}
