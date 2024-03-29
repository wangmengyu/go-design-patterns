package main

import (
	"design_patterns/struct/bridge/computer"
	"fmt"
)

func main() {
	/**
	桥接是一种结构型设计模式， 可将业务逻辑或一个大类拆分为不同的层次结构， 从而能独立地进行开发。
	层次结构中的第一层 （通常称为抽象部分） 将包含对第二层 （实现部分） 对象的引用。
	抽象部分将能将一些 （有时是绝大部分）
	对自己的调用委派给实现部分的对象。
	所有的实现部分都有一个通用接口， 因此它们能在抽象部分内部相互替换。
	*/
	hpPrinter := &computer.Hp{}
	epsonPrinter := &computer.Epson{}

	macComputer := &computer.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &computer.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
