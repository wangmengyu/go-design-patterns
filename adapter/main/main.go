package main

import "design_patterns/adapter"

func main() {
	//创建客户端对象
	client := &adapter.Client{}
	//创建mac对象
	mac := &adapter.Mac{}
	//调用client直接插入到mac接口

	client.InsertLightningConnectorIntoComputer(mac)

	//创建windows机器
	win := &adapter.Windows{}
	//创建windows适配器
	winAda := &adapter.WindowsAdapter{
		WindowsMachine: win,
	}
	//调用客户端插入到window适配器接口中
	client.InsertLightningConnectorIntoComputer(winAda)

}
