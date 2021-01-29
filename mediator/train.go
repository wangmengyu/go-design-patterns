package mediator

type Train interface { // 车辆的基类,
	Arrive()        // 到站
	Depart()        // 离站
	PermitArrival() // 允许进站
}
