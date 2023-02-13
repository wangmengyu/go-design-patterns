package memento

type Originator struct {
	State string // 元发器的状态值
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{State: e.State} // 创建新的备忘录
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.GetSavedState() // 恢复状态值
}

func (e *Originator) SetState(state string) {
	e.State = state
}

func (e *Originator) GetState() string {
	return e.State
}
