package main

import (
	"design_patterns/behavioral/mediator/mediator"
)

func main() {
	st := mediator.NewStationMediator()

	fr := &mediator.FreightTrain{
		M: st,
	}

	pe := &mediator.PassengerTrain{
		M: st,
	}

	pe.Arrive()
	fr.Arrive()
	pe.Depart()

}
