package state

import (
	"fmt"
)

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) AddItem(num int) error {

	i.VendingMachine.IncrementItemCount(num)
	// 状态转换成有库存
	i.VendingMachine.SetState(i.VendingMachine.HasItem)
	return nil
}
func (s *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock\n") // 当前无库存
}
func (s *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock\n") // 没库存不能交易
}
func (s *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock\n") // 没库存不能移走物品
}
