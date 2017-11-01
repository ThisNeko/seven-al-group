package feu

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
)

type FeuMessage struct {
	ID int
	Position Position
	Couleur Couleur
	Ticker int
	Timer int
}

type message struct{
	TypeEnum string
	Info string
}

func NewMessage(feu Feu) FeuMessage{
	return FeuMessage(feu)
}

type connection struct{
	id int
	ip string
	conn net.Conn
	info chan message
}

func (c connection) broadcast(info message){
	j, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c.conn,string(j)+"\n")
}

func (c connection) broadcastLoop(){
	//TODO: se reconnecter en cas de perte de connexion?
	for{
		info := <- c.info
		c.broadcast(info)
	}
}

//methode exportée qui met les infos à envoyer dans le channel
func (c connection) Broadcast(inf FeuMessage){
	info, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(err)
	}

	select{
	case <- c.info:
	default:
	}
	c.info <- message{
		"FEU",
		string(info),
	}
}

//créé un connection
func NewConnection(ip string) *connection{
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
	go c.broadcastLoop()
	return &c
}

