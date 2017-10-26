package voiture

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
	"bufio"
)

type VoitureMessage struct {
	ID int
	Vitesse Vitesse
	Position Position
	Frein bool
}

func NewMessage(status Status) VoitureMessage{
	return VoitureMessage(status.status)
}

type connection struct{
	id int
	ip string
	conn net.Conn
	info chan VoitureMessage
}

//envoie le message sur le reseau
func (c connection) broadcast(info VoitureMessage){
	j, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c.conn,string(j)+"\n")
}

//goroutine du broadcaster
//qui lit les message de son channel pour les envoyer
func (c connection) broadcastLoop(){
	//TODO: se reconnecter en cas de perte de connexion?
	for{
		info := <- c.info
		c.broadcast(info)
	}
}

//lit les messages reçus
func receive(conn net.Conn) VoitureMessage{
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var mat VoitureMessage
	json.Unmarshal([]byte(line),&mat)
	return mat
}

//goroutine du receiver
//qui lit les messages reçus et fais des choses avec
func (c connection) reveiverLoop(reg *Registre){
	for{
		mess := receive(c.conn)
		//log.Println("Received:")
		//log.Println(mess)
		reg.Update(mess)
	}
}

//methode exportée qui met les infos à envoyer dans le channel
func (c connection) Broadcast(info VoitureMessage){
	select{
	case <- c.info:
	default:
	}
	c.info <- info
}

//créé un connection
func NewConnection(ip string, reg *Registre) *connection{
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	c := connection{
		id: rand.Int(),
		ip: ip,
		conn: conn,
		info: make(chan VoitureMessage, 1),
	}
	go c.broadcastLoop()
	go c.reveiverLoop(reg)
	return &c
}
