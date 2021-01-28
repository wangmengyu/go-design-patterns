package bridge

import "fmt"

type Mac struct {
	//内部必须包含一个要使用的打印机对象
	printer Printer
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile() // 调用里面的方法
}
