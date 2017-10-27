package voiture

import "log"

type Conducteur interface {
	AlerteFrein()
}

type ConducteurLog struct {}

func (c ConducteurLog) AlerteFrein() {
	log.Println("FREIN!")
}