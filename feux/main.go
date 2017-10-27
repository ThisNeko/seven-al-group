package main

import(
	"log"
	"math/rand"
	"time"
	"./feu"
)



func main(){


	rand.Seed(time.Now().UTC().UnixNano())
	feux := feu.Feu{
		ID: rand.Int(),
		Position: feu.Position{0,0},
		Couleur: feu.Couleur(1),
		Ticker: 0,
	}

	log.Println(feux)

	feu.NewFeux("localhost:1234",&feux)

}