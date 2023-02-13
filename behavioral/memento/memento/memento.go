package memento

type Memento struct {
	State string // 备忘录中保留的一份状态值
}

func (m *Memento) GetSavedState() string {
	return m.State //  获取上次保存的状态值
}
