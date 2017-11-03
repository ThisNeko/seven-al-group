package main

import (
	"../../../src/wifi"
	"../../../src/voiture"
	"time"
	"bufio"
	"os"
	"fmt"
	"log"
	"net"
)

func main(){
	shutdownWifi := make(chan struct{})
	go wifi.StartWifi(shutdownWifi)

	mat := voiture.Materiel{
		ID: 1,
		Vitesse: 100,
		Position: voiture.Position{X: 90},
		Frein: false,
	}

	conducteur := voiture.ConducteurLog{}
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After)

	listener, err := net.Listen("tcp", "localhost:25252")
	if err != nil {
		log.Fatal(err)
	}
	connVoiture, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input +, - or nothing to tick")
	for{
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		switch text {
		case "+\n":
			mat.Vitesse += 1
		case "-\n":
			mat.Vitesse -= 1
		case "\n":
			log.Println("Ticking")
			mat.Position.X += mat.Vitesse
			time.Sleep(time.Millisecond * 51)
			fmt.Fprintf(connVoiture, "%d\n%d\n", uint64(mat.Vitesse), uint64(mat.Position.X))
		//case "q\n":
		//	log.Println("Quitting...")
		//	shutdownCar <- struct{}{}
		//	shutdownWifi <- struct{}{}
		//	time.Sleep(time.Millisecond * 350)
		}
	}
}