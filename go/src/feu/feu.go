package feu

import (
	"math/rand"
	"time"
	"log"
)

func NewFeu(ip string) {

	rand.Seed(time.Now().UTC().UnixNano())
	feu := Feu{
		ID:       rand.Int(),
		Position: Position{40, 1},
		Couleur:  Couleur(1),
		Ticker:   0,
		Timer:    10,
	}

	log.Println(feu)



	stat := newFeu()
	conn := NewConnection(ip)
	tickerChan:=time.NewTicker(time.Second).C
	//tickChan:=time.NewTicker(time.Second).C
	b:=true


	for {
		<- time.After(50 * time.Millisecond)

		go func() {
			select {
			case <-tickerChan:
				if feu.Ticker == feu.Timer {
					feu.Ticker = 0
					if b {
						feu.Couleur = feu.Couleur + 1
						b = false
					} else {
						feu.Couleur = feu.Couleur - 1
						b = true
					}
				} else {
					feu.Ticker = feu.Ticker + 1
				}
			}
		}()

		stat.Update(feu)
		conn.Broadcast(NewMessage(stat))

	}
}
