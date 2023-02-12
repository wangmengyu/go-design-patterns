package state

type State interface {
	AddItem(int) error           // 添加商品
	RequestItem() error          // 请求商品
	InsertMoney(money int) error // 插入钱
	DispenseItem() error         // 分配商品
}
