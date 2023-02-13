package computer

import "fmt"

// Computer 电脑 ,抽象层
type Computer interface { // 抽象层, 代表电脑
	Print()             // 执行打印
	SetPrinter(Printer) // 设置打印机
}
type Mac struct {
	printer Printer // 包含一个打印机
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer // 包含一个打印机
}

func (w *Windows) Print() { //  执行打印
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
