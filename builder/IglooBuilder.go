package builder

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

//创建新的NormalBuilder
func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.WindowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.DoorType = "Snow Door"

} // 设置门的类型
func (b *IglooBuilder) SetNumFloor() {
	b.Floor = 1

} // 设置楼梯类型
func (b *IglooBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
