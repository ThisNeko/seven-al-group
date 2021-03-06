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

	mat := voiture.Materiel{
		ID: rand.Int(),
		Vitesse: 80,
		Position: voiture.Position{X,Y},
		Panne: false,
	}

	log.Println(mat)
	conducteur := voiture.ConducteurLog{}
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After)

	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		switch text {
		case "f\n":
			mat.Vitesse = 10
			log.Println("Vitesse = 10")
		case "a\n":
			mat.Vitesse = 80
			log.Println("Vitesse = 80")
		case "p\n":
			mat.Panne = !mat.Panne
			log.Printf("Panne = %t",mat.Panne)
		}
	}
}