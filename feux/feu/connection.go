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
}

func NewMessage(feu Feu) FeuMessage{
	return FeuMessage(feu)
}

type connection struct{
	id int
	ip string
	conn net.Conn
	info chan FeuMessage
}

func (c connection) broadcast(info FeuMessage){
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
func (c connection) Broadcast(info FeuMessage){
	select{
	case <- c.info:
	default:
	}
	c.info <- info
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
		info: make(chan FeuMessage, 1),
	}
	go c.broadcastLoop()
	return &c
}




