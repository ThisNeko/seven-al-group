package main

import (
	"../../src/voiture"
	"math/rand"
	"time"
	"strconv"
	"os"
	"log"
	"bufio"
	"fmt"
)

func vitesseLoop(materiel * voiture.Materiel , vitesse chan float64, maxSpeed float64){
	for{
		v := <- vitesse
		if v > maxSpeed {
			materiel.Vitesse = maxSpeed
		} else {
			materiel.Vitesse = v
		}
	}
}

func mouvementLoop(materiel * voiture.Materiel, maxSpeed float64){
	for{
		<- time.After(10*time.Millisecond)
		v := (materiel.Vitesse/3.6)/100.0
		materiel.Position.X += v
		//log.Println(materiel)
	}
}

func main(){
	rand.Seed(time.Now().UTC().UnixNano())

	X, err := strconv.ParseFloat(os.Args[1],64)
	if err != nil {
		X = 0
	}

	Y, err := strconv.ParseFloat(os.Args[2],64)
	if err != nil {
		Y = 0
	}

	maxSpeed, err := strconv.ParseFloat(os.Args[3],64)
	if err != nil {
		maxSpeed = 80
	}


	id := rand.Int()
	if id % 2 == 1 {
		id += 1
	}

	mat := voiture.Materiel{
		ID: id,
		Vitesse: maxSpeed,
		Position: voiture.Position{X,Y},
		Panne: false,
	}

	vitesse := make(chan float64)

	log.Println(mat)
	var conducteur voiture.Conducteur = voiture.ConducteurAuto{vitesse}

	go vitesseLoop(&mat,vitesse, maxSpeed)
	go mouvementLoop(&mat, maxSpeed)
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After)

	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		switch text {
		case "p\n":
			mat.Panne = !mat.Panne
			log.Printf("Panne = %t",mat.Panne)
		case "+\n":
			mat.Position.Y += 1
		case "-\n":
			mat.Position.Y -= 1
		}
	}
}