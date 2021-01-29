package mediator

import "fmt"

type PassengerTrain struct {
	//客车, 要包含一个中介对象
	Mediator Mediator
}

func (pt *PassengerTrain) Arrive() {
	//判断是不是可以到站
	if !pt.Mediator.CanArrive(pt) {
		//可不可到站
		fmt.Println("PassengerTrain: Arrival Blocked! Waiting")
		return
	}
	//可以到站, 设置
	fmt.Println("PassengerTrain: Arrived")
}

func (pt *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	pt.Mediator.NotifyAboutDepart() // 发车, 腾出位置, 将队列的第一个放进来
}

// 允许入
func (pt *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted")
	pt.Arrive()

}
