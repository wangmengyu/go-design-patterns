package adapter

import "fmt"

type Windows struct {
	// 这个结构只有一个插入USB的方法
}

func (w *Windows) InsertIntoUSBPort() { //这个方法要再适配器的插入到Lightning接口去调用
	fmt.Println("USB Connector is plugged into windows machine.")
}
