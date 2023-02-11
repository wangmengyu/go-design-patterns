package command

// ICommand 【命令】所具有的方法列表
type ICommand interface {
	Execute() // 命令只要调用执行方法即可.
}
