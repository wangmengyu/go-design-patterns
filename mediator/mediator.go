package mediator

type Mediator interface { // 中介接口:
	CanArrive(train Train) bool
	NotifyAboutDepart() // 通知离站
}
