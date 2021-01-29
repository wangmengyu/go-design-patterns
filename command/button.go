package command

//按钮: 按钮中包含一个命令成员,
type Button struct {
	Command Command // 命令
}

//按钮有一个按下操作会调用命令的执行方法
func (b *Button) Press() {
	b.Command.Execute()
}
