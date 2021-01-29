package command

type OffCommand struct {
	Device Device
}

func (o *OffCommand) Execute() {
	o.Device.Off()
}
