package state

import (
	"fmt"
)

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested\n")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress\n") //  	请求状态不支持这个操作
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.ItemPrice {
		// 金额不够
		return fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.ItemPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.HasMoney) // 转成有钱状态
	return nil
}
func (i *ItemRequestedState) DispenseItem() error { // 请求状态不支持这个操作
	return fmt.Errorf("Please insert money first\n")
}
