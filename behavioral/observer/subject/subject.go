package subject

import "design_patterns/behavioral/observer/observer"

type Subject interface {
	Register(observer observer.Observer)
	DeRegister(observer observer.Observer)
	notifyAll()
}
