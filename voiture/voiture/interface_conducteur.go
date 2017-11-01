package voiture

import "log"

type Conducteur interface {
	AlerteFrein()
	VitesseFeu(vitesse float64)
}

type ConducteurLog struct {}

func (c ConducteurLog) AlerteFrein() {
	log.Println("FREIN!")
}

func (c ConducteurLog) VitesseFeu(vitesse float64) {
	//log.Println("Rouler a %d pour avoir le feu vert",vitesse)
}