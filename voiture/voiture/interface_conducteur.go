package voiture

import "log"

type Conducteur interface {
	AlerteFrein()
	VitesseFeu(vitesse float64,feu Feu)
}

type ConducteurLog struct {}

func (c ConducteurLog) AlerteFrein() {
	log.Println("FREIN!")
}

func (c ConducteurLog) VitesseFeu(vitesse float64, feu Feu) {
	log.Println(vitesse, feu.Ticker, feu.Couleur)
}