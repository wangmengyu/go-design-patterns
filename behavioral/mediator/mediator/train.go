package mediator

type Train interface {
	Arrive()        // 到站
	Depart()        // 离站
	PermitArrival() // 被允许到站
}
