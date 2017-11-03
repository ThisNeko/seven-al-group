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
	log.Printf("La vitesse recommandée est : %d km/h ; le couleur du feu est %s et va changer de couleur dans %d sec\n",int(vitesse),feu.Couleur.String(),10 - feu.Ticker)
}