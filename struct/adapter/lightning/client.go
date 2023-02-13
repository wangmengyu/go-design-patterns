package lightning

import "fmt"

type Client struct {
}

// 客户端向要插到电脑,  用lightning 口
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type Computer interface { // 电脑抽象接口.
	InsertIntoLightningPort() // 插入InsertIntoLightningPort
}
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() { // 苹果没问题直接插入
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() { // windows 只能插入USB
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	// 先做一层转换在插入.
	w.WindowMachine.InsertIntoUSBPort()
}
