package mediator

type StationManager struct {
	//站台是否是空的
	IsPlatformFree bool
	TrainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
	}
}

func (sm *StationManager) CanArrive(train Train) bool {
	//当前站台是空的, 可以进入
	if sm.IsPlatformFree {
		sm.IsPlatformFree = false
		return true
	}

	// 非空. 放入队列中
	sm.TrainQueue = append(sm.TrainQueue, train)
	return false
}

/**
  通知离站, 设置当前平台是空的, 并且将等待队列中的第一个元素取出, 将它设置成允许进入状态
*/
func (sm *StationManager) NotifyAboutDepart() {

	if sm.IsPlatformFree == false {
		sm.IsPlatformFree = true
	}

	if len(sm.TrainQueue) > 0 {
		firstElem := sm.TrainQueue[0]
		firstElem.PermitArrival()
		if len(sm.TrainQueue) > 1 {
			sm.TrainQueue = sm.TrainQueue[1:]
		}
	}
}
