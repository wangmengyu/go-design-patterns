package mediator

type Mediator interface {
	CanArrive(train Train) bool // 可否到站
	NotifyAboutDeparture()      // 通知离站
}

func NewStationMediator() *StationMediator {
	return &StationMediator{
		isFree:     true,
		trainQueue: make([]Train, 0),
	}
}

type StationMediator struct {
	// 站点, 具体实现中介人
	isFree     bool    // 是否空闲
	trainQueue []Train // 维护的火车队列
}

// CanArrive 如果是可以到站的, 设置到站状态.
func (s *StationMediator) CanArrive(t Train) bool {
	if s.isFree {
		s.isFree = false // 把站点状态设置成 false
		return true      // 直接消费掉了不用加入队列
	}
	// 追加当前车辆到当前队列
	s.trainQueue = append(s.trainQueue, t)

	return false
}

func (s *StationMediator) NotifyAboutDeparture() {
	// 设置站点状态为空闲
	if s.isFree == false {
		s.isFree = true
	}
	// 取得当前队列第一个元素
	if len(s.trainQueue) == 0 {
		return
	}
	elem := s.trainQueue[0]
	s.trainQueue = s.trainQueue[1:]
	elem.PermitArrival() //  通知允许到站.
}
