package state

import (
	"fmt"
)

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.VendingMachine.ItemCount == 0 {
		// 无库存
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
		return fmt.Errorf("No item present\n")
	}
	// 有库存->请求物品状态
	fmt.Printf("Item requestd\n")
	i.VendingMachine.SetState(i.VendingMachine.ItemRequested) //
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.VendingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first\n") // 有库存状态不支持这个操作
}
func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item firs\nt") // 有库存状态不支持这个操作
}
