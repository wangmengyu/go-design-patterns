package subject

import "design_patterns/behavioral/observer/observer"

type Item struct {
	Name         string              // product name
	ObserverList []observer.Observer // the list of observer
	InStock      bool
}

func NewItem(name string) *Item {
	return &Item{Name: name}
}

func (i *Item) Register(observer observer.Observer) {
	// 订阅
	i.ObserverList = append(i.ObserverList, observer)
}
func (i *Item) DeRegister(observer observer.Observer) {
	// 取消订阅
	for pos, v := range i.ObserverList {
		if v.GetId() == observer.GetId() {
			i.ObserverList = append(i.ObserverList[:pos], i.ObserverList[pos+1:]...)
			break
		}
	}
}
func (i *Item) notifyAll() {
	for _, row := range i.ObserverList {
		row.Update(i.Name)
	}
}

func (i *Item) UpdateAvailability() {
	i.InStock = true
	i.notifyAll()
}
