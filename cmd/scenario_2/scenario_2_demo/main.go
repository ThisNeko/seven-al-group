package main

import (
	"../../../src/voiture"
	"time"
)
func main(){
	mat := voiture.Materiel{
		ID: 2,
		Vitesse: 100,
		Position: voiture.Position{X: 100},
		Frein: false,
	}

	conducteur := voiture.ConducteurLog{}
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After, make(chan struct{}))


}