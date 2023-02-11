package command

import "design_patterns/behavioral/command/device"

type OnCommand struct {
	Device device.IDevice // 具体命令需要包含接收设备
}

func (o *OnCommand) Execute() {
	o.Device.On()
}
