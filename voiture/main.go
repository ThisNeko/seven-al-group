package main

import(
	"log"
	"math/rand"
	"time"
	"./voiture"
	"bufio"
	"os"
	"fmt"
	//"strconv"
)



func main(){
	rand.Seed(time.Now().UTC().UnixNano())

	//X, err := strconv.ParseFloat(os.Args[1],64)
	//if err != nil {
		X := float64(0)
	//}

	//Y, err := strconv.ParseFloat(os.Args[2],64)
	//if err != nil {
		Y := float64(0)
	//}

	mat := voiture.Materiel{
		ID: rand.Int(),
		Vitesse: 80,
		Position: voiture.Position{X,Y},
		Frein: false,
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
		}
	}
}