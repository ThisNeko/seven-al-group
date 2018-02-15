package voiture

import "log"

type Conducteur interface {
	AlerteFrein()
	VitesseFeu(vitesse float64,feu Feu)
	PanneLead()
	Vitesse(vitesse float64)
}

type ConducteurLog struct {}

func (c ConducteurLog) AlerteFrein() {
	log.Println("FREIN!")
}

func (c ConducteurLog) VitesseFeu(vitesse float64, feu Feu) {
	log.Printf("La vitesse recommandée est : %d km/h ; le couleur du feu est %s et va changer de couleur dans %d sec\n",int(vitesse),feu.Couleur.String(),10 - feu.Ticker)
}

func (c ConducteurLog) PanneLead() {
	log.Println("Le véhicule d'en face est en panne, veuillez changer de voie.")
}

func (c ConducteurLog) Vitesse(vitesse float64) {
	log.Printf("La vitesse recommandée est : %d km/h\n",int(vitesse))
}



type ConducteurAuto struct {
	VitesseChan chan float64
}

func (c ConducteurAuto) AlerteFrein() {
	log.Println("FREIN!")
}

func (c ConducteurAuto) VitesseFeu(vitesse float64, feu Feu) {
	log.Printf("La vitesse recommandée est : %d km/h ; le couleur du feu est %s et va changer de couleur dans %d sec\n",int(vitesse),feu.Couleur.String(),10 - feu.Ticker)
}

func (c ConducteurAuto) PanneLead() {
	log.Println("Le véhicule d'en face est en panne, veuillez changer de voie.")
}

func (c ConducteurAuto) Vitesse(vitesse float64) {
	log.Printf("La vitesse recommandée est : %d km/h\n",int(vitesse))
	c.VitesseChan <- vitesse
}
