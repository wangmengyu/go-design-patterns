package mediator

import (
	"fmt"
)

type PassengerTrain struct {
	// 必须包含中间人
	M Mediator
}

func (pt *PassengerTrain) Arrive() {
	if !pt.M.CanArrive(pt) {
		// 不可以到站, 通知
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")

}
func (pt *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain : Leaving")
	pt.M.NotifyAboutDeparture()

}
func (pt *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted")
	pt.Arrive() // 被允许, 执行到站操作
}

/**
Arrive()        // 到站
	Depart()        // 离站
	PermitArrival() // 被允许到站
*/
