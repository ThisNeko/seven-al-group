package voiture

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
	"bufio"
)

type message struct{
		TypeEnum string
		Info string
}

func NewMessage(status Status) StatusVoiture {
	return StatusVoiture(status.Get())
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
	//TODO: se reconnecter en cas de perte de connexion?
	for{
		info := <- c.info
		c.broadcast(info)
	}
}

//lit les messages reçus
func receive(conn net.Conn) message {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var mat message
	json.Unmarshal([]byte(line),&mat)
	return mat
}

//goroutine du receiver
//qui lit les messages reçus et fais des choses avec
func (c connection) receiverLoop(reg *Registre){
	for{
		mess := receive(c.conn)
		//log.Println("Received:")
		//log.Println(mess)
		if mess.TypeEnum=="VOITURE"{
			var mat StatusVoiture
			json.Unmarshal([]byte(mess.Info),&mat)
			reg.UpdateVoiture(mat)
			log.Println("La voiture recoit un message d'une autre voiture")
		}
		if mess.TypeEnum=="FEU"{
			var mat Feu
			json.Unmarshal([]byte(mess.Info),&mat)
			reg.UpdateFeu(mat)
			log.Println("La voiture recoit un message d'un FEU")
		}
	}
}

//methode exportée qui met les infos à envoyer dans le channel
func (c connection) Broadcast(inf StatusVoiture){
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
func NewConnection(ip string, reg *Registre) connection{
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
	go c.receiverLoop(reg)
	return c
}
