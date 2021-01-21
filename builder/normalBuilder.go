package builder

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

//创建新的NormalBuilder
func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *NormalBuilder) SetDoorType() {
	b.DoorType = "Wooden Door"

} // 设置门的类型
func (b *NormalBuilder) SetNumFloor() {
	b.Floor = 2

} // 设置楼梯类型
func (b *NormalBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
