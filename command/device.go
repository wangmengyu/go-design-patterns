package command

//设备接口. 所有设备的方法列表
type Device interface {
	On()
	Off()
}
