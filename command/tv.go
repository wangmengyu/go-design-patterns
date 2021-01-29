package command

import "fmt"

type Tv struct {
	isRunning bool //是否在运行
}

func (tv *Tv) On() {
	tv.isRunning = true
	fmt.Println("tv is on")
}

func (tv *Tv) Off() {
	tv.isRunning = false
	fmt.Println("tv is off")
}
