package main

import "design_patterns/command"

func main() {
	//创建电视对象
	tv := &command.Tv{}
	//创建开启命令, 并且绑定设备,
	onCommand := &command.OnCommand{Device: tv}
	//创建开按钮, 绑定命令
	onBtn := command.Button{Command: onCommand}

	//创建关闭命令, 并且绑定设备
	offCommand := &command.OffCommand{Device: tv}
	offBtn := command.Button{Command: offCommand}

	onBtn.Press()
	offBtn.Press()

}
