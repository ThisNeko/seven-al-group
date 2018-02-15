package main

import (
	"../../src/voiture"
	"time"
	"bufio"
	"os"
	"fmt"
	"strconv"
)
func main(){
	mat := voiture.Materiel{
		ID: 3,
		Vitesse: 70,
		Position: voiture.Position{X: 100},
	}

	conducteur := voiture.ConducteurLog{}
	go voiture.NewVoiture("localhost:1234", &mat, conducteur, time.After)

	for{
		rdr := bufio.NewReader(os.Stdin)
		tmp, _ := rdr.ReadString('\n')
		tmp2, _ :=  strconv.Atoi(tmp[:len(tmp)-1])
		posDelta := float64(tmp2)

		mat.Position.X += posDelta

		//fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf(">>> [Speed: %d, PosX: %d] >>>\n", int64(mat.Vitesse),
			int64(mat.Position.X))
	}
}