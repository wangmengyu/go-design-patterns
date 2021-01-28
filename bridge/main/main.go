package main

import "design_patterns/bridge"

func main() {
	//创建好HP和EPSON的打印机
	hp := &bridge.Hp{}
	epson := &bridge.Epson{}

	//创建好MAC
	mac := &bridge.Mac{}
	// 设置HP为MAC的打印机, 并且执行打印
	mac.SetPrinter(hp)
	mac.Print()

	//设置EPSON为MAC的打印机 , 并且执行打印
	mac.SetPrinter(epson)
	mac.Print()

	//创建好Windows
	windows := &bridge.Windows{}

	//设置HP为WINDOWS的打印机, 并且执行打印
	windows.SetPrinter(hp)
	windows.Print()

	//设置EPSON为WINDOWS的打印机, 并且执行打印
	windows.SetPrinter(epson)
	windows.Print()

}
