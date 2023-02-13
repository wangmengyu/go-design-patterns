package memento

type Caretaker struct {
	MementoArray []*Memento // 备忘录列表,  负责管理做过的备忘录.
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m) // 从尾部追加
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.MementoArray[index] // 取得指定版本
}
