package state

import (
	"fmt"
)

type VendingMachine struct {
	ItemCount int //  商品数量
	ItemPrice int // 商品价格

	CurrentState State //  当前状态

	HasItem       State //  有ITEM的状态
	ItemRequested State // 被请求商品的状态
	HasMoney      State //  有钱的状态
	NoItem        State // 没有ITEM的状态
}

func NewVendingMachine(cnt int, price int) *VendingMachine {
	m := &VendingMachine{
		ItemCount:     cnt,
		ItemPrice:     price,
		CurrentState:  nil,
		HasItem:       nil,
		ItemRequested: nil,
		HasMoney:      nil,
		NoItem:        nil,
	}
	hasItemState := &HasItemState{
		VendingMachine: m,
	}
	m.HasItem = hasItemState
	itemReqSt := &ItemRequestedState{
		VendingMachine: m,
	}
	m.ItemRequested = itemReqSt
	hasMoneySt := &HasMoneyState{
		VendingMachine: m,
	}
	m.HasMoney = hasMoneySt
	noItemSt := &NoItemState{
		VendingMachine: m,
	}
	m.NoItem = noItemSt

	m.SetState(hasItemState)

	return m
}

func (v *VendingMachine) IncrementItemCount(cnt int) {
	fmt.Printf("Adding %d items\n", cnt)
	v.ItemCount += cnt
}

func (v *VendingMachine) SetState(st State) {
	v.CurrentState = st
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}
