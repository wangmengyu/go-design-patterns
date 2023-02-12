package observer

type Observer interface {
	Update(string) // when item trigger updated , what update for observer , 当有变化的时候. 对于观察者要执行的操作
	GetId() string // get observer  id
}
