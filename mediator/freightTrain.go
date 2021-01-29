package mediator

import "fmt"

type FreightTrain struct {
	//货车, 要包含一个中介对象
	Mediator Mediator
}

func (pt *FreightTrain) Arrive() {
	//判断是不是可以到站
	if !pt.Mediator.CanArrive(pt) {
		//可不可到站
		fmt.Println("FreightTrain: Arrival Blocked! Waiting")
		return
	}
	//可以到站, 设置
	fmt.Println("FreightTrain: Arrived")
}

func (pt *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	pt.Mediator.NotifyAboutDepart() // 发车, 腾出位置, 将队列的第一个放进来
}

// 允许入
func (pt *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	pt.Arrive()

}
