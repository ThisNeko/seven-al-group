package voiture

import (
	"net"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
	"bufio"
	//"time"
	"time"
)

type message struct{
	TypeEnum string
	Info string
}

type InfoVoiture struct{
	ID int
	Vitesse Vitesse
	Position Position
	Panne bool
	Timestamp int64
}

func NewMessage(status Data) InfoVoiture {
	s := status.GetStatus()
	return InfoVoiture{
		s.ID,
		s.Vitesse,
		s.Position,
		s.Panne,
		time.Now().UTC().UnixNano(),
		}
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
func (c connection) receiverLoop(reg *Data){

	for{
		mess := receive(c.conn)
		timestamp := time.Now().UTC().UnixNano()
		//log.Println("Received:")
		//log.Println(mess)
		if mess.TypeEnum=="VOITURE"{
			var mat InfoVoiture
			err := json.Unmarshal([]byte(mess.Info),&mat)
			if err != nil {
				log.Fatal(err)
			}
			delay := timestamp - mat.Timestamp
			log.Printf("Delay : %v ms\n",delay)
			s := StatusVoiture{mat.ID,mat.Vitesse, mat.Position,mat.Panne}
			reg.UpdateVoiture(s)
			//log.Println("La voiture recoit un message d'une autre voiture")
		}
		if mess.TypeEnum=="FEU"{
			var mat Feu
			err := json.Unmarshal([]byte(mess.Info),&mat)
			if err != nil {
				log.Fatal(err)
			}
			//log.Println("Id du Feu %d",mat.ID)
			reg.UpdateFeu(mat)
			//log.Println("La voiture recoit un message d'un FEU")
		}
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

func (c connection) Start(reg *Data){
	go c.broadcastLoop()
	go c.receiverLoop(reg)
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
