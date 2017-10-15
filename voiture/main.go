package main

import(
	"log"
	"math/rand"
	"time"
	"./voiture"
)



func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	mat := voiture.Materiel{
		ID: rand.Int(),
		Vitesse: 80,
		Position: voiture.Position{0,0},
		Frein: false,
	}

	log.Println(mat)

	voiture.NewVoiture("localhost:1234",&mat)

}