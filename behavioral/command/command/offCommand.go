package command

import "design_patterns/behavioral/command/device"

type OffCommand struct {
	Device device.IDevice // 命令上又包含了需要控制的设备
}

func (o *OffCommand) Execute() {
	o.Device.Off()
}
