package main

import (
	"design_patterns/behavioral/command/button"
	"design_patterns/behavioral/command/command"
	"design_patterns/behavioral/command/device"
)

func main() {
	tv := device.Tv{}
	// 用按钮驱动
	onCmd := command.OnCommand{
		Device: &tv,
	}
	onBtn := button.Button{Command: &onCmd}
	offCmd := command.OffCommand{Device: &tv}
	offBtn := button.Button{
		Command: &offCmd,
	}

	onBtn.Press()
	offBtn.Press()

}
