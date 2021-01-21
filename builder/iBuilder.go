package builder

type iBuilder interface {
	SetWindowType()  // 设置窗户类型
	SetDoorType()    // 设置门的类型
	SetNumFloor()    // 设置楼梯类型
	GetHouse() House // 获得房屋实例
}

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		// 获得正常房构建器
		return &NormalBuilder{}
	} else if builderType == "igloo" {
		//获得冰屋构建器
		return &IglooBuilder{}
	}
	return nil
}
