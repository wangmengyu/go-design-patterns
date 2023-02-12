package state

import (
	"fmt"
)

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress\n")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress\n")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.VendingMachine.ItemCount = i.VendingMachine.ItemCount - 1
	if i.VendingMachine.ItemCount == 0 { // 没有库存了. 设成无库存状态
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
	} else {
		// 有库存
		i.VendingMachine.SetState(i.VendingMachine.HasItem)
	}
	return nil
}
