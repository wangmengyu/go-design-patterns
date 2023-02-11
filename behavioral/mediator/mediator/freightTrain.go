package mediator

import (
	"fmt"
)

type FreightTrain struct {
	M Mediator
}

func (t *FreightTrain) Arrive() {
	if !t.M.CanArrive(t) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")

}
func (t *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	t.M.NotifyAboutDeparture()

}
func (t *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	t.Arrive() // 被允许, 执行到站
}

/**
Arrive()        // 到站
	Depart()        // 离站
	PermitArrival() // 被允许到站
*/
