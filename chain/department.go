package chain

//部门抽象接口, 每个部门有 要执行的操作和设置下一个操作
type Department interface {
	Execute(*Patient)
	SetNext(department Department)
}
