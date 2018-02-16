package main

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type message struct{
	TypeEnum string
	Info string
}

type Position struct{
	X float64
	Y float64
}

type Vitesse struct{
	X float64
	Y float64
}

type InfoVoiture struct{
	ID int
	Vitesse Vitesse
	Position Position
	Panne bool
	Timestamp int64
}

func NewMessage(id int, v Vitesse, p Position, panne bool) InfoVoiture {
	return InfoVoiture{
		id,
		v,
		p,
		panne,
		time.Now().UTC().UnixNano(),
	}
}

func NewRandomMessage() InfoVoiture{
	return NewMessage(
		rand.Int(),
		Vitesse{rand.Float64(),rand.Float64()},
		Position{rand.Float64(),rand.Float64()},
		rand.Intn(2) != 0,
	)
}

type connection struct{
	id int
	ip string
	conn net.Conn
	info chan message
}

//envoie le message sur le reseau
func (c connection) broadcast(info message){
	j, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c.conn,string(j)+"\n")
	//log.Println("Transformation []byte -> string = "+string(j))
}

//goroutine du broadcaster
//qui lit les message de son channel pour les envoyer
func (c connection) broadcastLoop(){
	for{
		info := <- c.info
		c.broadcast(info)
	}
}

//methode exportée qui met les infos à envoyer dans le channel
func (c connection) Broadcast(inf InfoVoiture){
	info, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(err)
	}

	select{
	case <- c.info:
	default:
	}
	c.info <- message{
		"VOITURE",
		string(info),
	}
}

//créé un connection
func NewConnection(ip string) connection{
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	c := connection{
		id: rand.Int(),
		ip: ip,
		conn: conn,
		info: make(chan message, 1),
	}
	return c
}


func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	conn := NewConnection("localhost:1234")

	go conn.broadcastLoop()

	for{
		m := NewRandomMessage()
		conn.Broadcast(m)
		time.After(1*time.Millisecond)
	}

}