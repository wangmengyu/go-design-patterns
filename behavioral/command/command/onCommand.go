package command

import "design_patterns/behavioral/command/device"

type OnCommand struct {
	Device device.IDevice // 命令上又包含了需要控制的设备
}

func (o *OnCommand) Execute() {
	o.Device.On()
}
