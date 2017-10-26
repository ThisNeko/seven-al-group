package feu

import "time"

func NewFeux(ip string, feu *Feu) {

	stat := NewFeu()
	conn := NewConnection(ip)

	for{
		<- time.After(50 * time.Millisecond)
		stat.Update(*feu)
		conn.Broadcast(NewMessage(stat))
	}
}
