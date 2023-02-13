package main

import "design_patterns/struct/adapter/lightning"

func main() {

	client := &lightning.Client{}
	mac := &lightning.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &lightning.Windows{}
	windowsMachineAdapter := &lightning.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	// 适配器 可以视作一种电脑类型 .
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
