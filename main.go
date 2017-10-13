package main

import(
	"net"
	"log"
)

type position struct{
	x float64
	y float64
}

type materiel struct {
	vitesse float32
	position position
	frein bool
}

func main(){
	materiel := materiel{
		vitesse: 80,
		position: position{0,0},
		frein: false,
	}



}