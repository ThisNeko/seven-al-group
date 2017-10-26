package voiture

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
	"bufio"
)

func NewMessage(status Status) StatusVoiture {
	return StatusVoiture(status.Get())
}

type connection struct{
	id int
	ip string
	conn net.Conn
	info chan StatusVoiture
}

//envoie le message sur le reseau
func (c connection) broadcast(info StatusVoiture){
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
func receive(conn net.Conn) StatusVoiture {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var mat StatusVoiture
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
func (c connection) Broadcast(info StatusVoiture){
	select{
	case <- c.info:
	default:
	}
	c.info <- info
}

//créé un connection
func NewConnection(ip string, reg *Registre) connection{
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	c := connection{
		id: rand.Int(),
		ip: ip,
		conn: conn,
		info: make(chan StatusVoiture, 1),
	}
	go c.broadcastLoop()
	go c.reveiverLoop(reg)
	return c
}
