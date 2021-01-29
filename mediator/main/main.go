package main

import "design_patterns/mediator"

func main() {

	st := mediator.NewStationManager()

	fr := &mediator.FreightTrain{Mediator: st}

	pa := mediator.PassengerTrain{Mediator: st}

	fr.Arrive()
	pa.Arrive()
	fr.Depart()

}
