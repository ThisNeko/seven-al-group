package main

import(
	"net"
	"log"
	"encoding/json"
	"fmt"
	"bufio"
	"math/rand"
	"time"
)

type position struct{
	X float64
	Y float64
}

type materiel struct {
	ID int
	Vitesse float32
	Position position
	Frein bool
}

func broadcast(conn net.Conn, mat materiel){
	j, err := json.Marshal(mat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn,string(j)+"\n")
}

func receive(conn net.Conn) materiel{
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var mat materiel
	json.Unmarshal([]byte(line),&mat)
	return mat
}

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	mat := materiel{
		ID: rand.Int(),
		Vitesse: 80,
		Position: position{0,0},
		Frein: false,
	}

	log.Println(mat)

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	broadcast(conn,mat)
	for {
		mat2 := receive(conn)
		log.Println("Received:")
		log.Println(mat2)
	}

}