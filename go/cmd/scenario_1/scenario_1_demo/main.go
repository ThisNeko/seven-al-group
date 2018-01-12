package main

import (
	"../../../src/voiture"
	"time"
	"net"
	"bufio"
	"fmt"
	"strconv"
)
func main(){
	mat := voiture.Materiel{
		ID: 2,
		Vitesse: 100,
		Position: voiture.Position{X: 0},
		Frein: false,
	}

	conducteur := voiture.ConducteurLog{}
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After)

	conn, _ := net.Dial("tcp", "localhost:25252")

	for{
		rdr := bufio.NewReader(conn)
		tmp, _ := rdr.ReadString('\n')
		tmp2, _ :=  strconv.Atoi(tmp[:len(tmp)-1])
		vitesse := int64(tmp2)

		tmp, _ = rdr.ReadString('\n')
		tmp2, _ =  strconv.Atoi(tmp[:len(tmp)-1])

		posX := int64(tmp2)

		mat.Position.X += mat.Vitesse

		//fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf(">>> [Speed: %d, PosX: %d]   [Speed: %d, PosX: %d]  >>>\n", int64(mat.Vitesse),
			int64(mat.Position.X), vitesse, posX)

		gap := int64(mat.Position.X) - posX
		if gap < 0 {
			gap = -gap
		}

		remaining := posX + vitesse - int64(mat.Position.X) - int64(mat.Vitesse)

		fmt.Printf("Gap: %dm       Next gap: %dm\n", gap, remaining)
	}
}