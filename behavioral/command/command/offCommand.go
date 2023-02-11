package command

import "design_patterns/behavioral/command/device"

type OffCommand struct {
	Device device.IDevice // 具体命令需要包含接收设备
}

func (o *OffCommand) Execute() {
	o.Device.Off()
}
