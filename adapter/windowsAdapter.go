package adapter

import "fmt"

type WindowsAdapter struct {
	//包含一个要适配的对象
	WindowsMachine *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {

	//先转换型号 从Lighting 到 USB
	fmt.Println("Adapter convert Lighting signal to USB")
	// 在插入到Windows的USB里面去 // 适配方法基本就在这地方写
	wa.WindowsMachine.InsertIntoUSBPort()
}
