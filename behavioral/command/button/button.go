package button

import "design_patterns/behavioral/command/command"

// Button 按钮， 实例化时候需要绑定命令
type Button struct {
	Command command.ICommand // 包含了命令抽象
}

func (b *Button) Press() { // 执行命令.
	b.Command.Execute()
}
