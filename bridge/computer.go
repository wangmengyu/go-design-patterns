package bridge

//抽象接口, computer 准备调用打印机的
type computer interface {
	//设置打印机
	SetPrinter(Printer)
	//调用打印机
	Print()
}
