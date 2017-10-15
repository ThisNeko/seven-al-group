package voiture

import (
	"./comm"
)

type Position struct{
	X float64
	Y float64
}

type Materiel struct {
	ID int
	Vitesse float64
	Position Position
	Frein bool
}

func NewVoiture(ip string, materiel *Materiel) {
	conn := comm.NewConnection(ip)
	mess := comm.VoitureMessage{
		ID:materiel.ID,
		Vitesse:comm.Vitesse{materiel.Vitesse,0},
		Position:comm.Position{materiel.Position.X,materiel.Position.Y},
		Frein:materiel.Frein,
	}
	conn.Broadcast(mess)
	for{
		continue
	}
}
