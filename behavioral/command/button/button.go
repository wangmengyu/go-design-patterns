package button

import "design_patterns/behavioral/command/command"

// Button 按钮， 实例化时候需要绑定命令
type Button struct {
	Command command.ICommand
}

func (b *Button) Press() {
	b.Command.Execute()
}
