package main

import (
	"net"
	"bufio"
	"log"
	"time"
)

type user struct {
	conn     net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func newUser(conn net.Conn) (user, error) {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	user := user{ conn, reader, writer}
	return user, nil
}

type broadcastMessage struct {
	user user
	message string
}

type userPool struct {
	add    chan user
	remove chan user
	cast   chan broadcastMessage
}

func newUserPool() *userPool {
	//create a user pool
	pool := userPool{
		add:    make(chan user),
		remove: make(chan user),
		cast:  	make(chan broadcastMessage),
	}

	go userManagement(&pool)

	return &pool
}

func userManagement(pool *userPool) {
	//goroutine that manages interactions with the user pool
	var users []user
	for {
		select {
		case u := <-pool.add:
			users = append(users, u)
		case ru := <-pool.remove:
			for i, u := range users {
				if u.conn == ru.conn {
					users = append(users[:i], users[i+1:]...)
					break
				}
			}
		case s := <-pool.cast:
			message := s.message
			user := s.user
			for _, u := range users {
				if u.conn == user.conn {
					continue
				}
				u.send(message)
			}
		}
	}
}

func (pool *userPool) addUser(user user) {
	//add a user to the pool
	pool.add <- user
	log.Println("Login of " + user.conn.RemoteAddr().String())
}

func (pool *userPool) removeUser(user user) {
	//remove a user from the pool
	pool.remove <- user
	log.Println("Logout of " + user.conn.RemoteAddr().String())
}

func (pool *userPool) broadcast(user user, message string) {
	//send message to everyone except user
	pool.cast <- broadcastMessage{user,message}
}


func (user user) send(message string) {
	//send message only to user
	user.writer.WriteString(message)
	user.writer.Flush()
}


func (user user) getMessage(message chan string, disconnect chan struct{}) {
	//get message sent by the user
	line, err := user.reader.ReadString('\n')
	if err != nil {
		disconnect <- struct{}{}
		return
	}
	log.Println(user.conn.RemoteAddr().String() + " sent : " + line)
	message <- line
}

func handleConnection(pool *userPool, conn net.Conn) {
	defer conn.Close()

	//create the user
	user, err := newUser(conn)
	if err != nil {
		log.Println(err)
		user.send("Error connecting to the server.")
		return
	}

	//add the user to the pool
	pool.addUser(user)
	defer pool.removeUser(user)

	for {
		message := make(chan string)
		disconnect := make(chan struct{})

		go user.getMessage(message,disconnect)

		select {
		case <-disconnect:
			return
		case m := <-message:
			//send the message to every other users
			pool.broadcast(user, m)
		case <-time.After(1 * time.Minute):
			//disconnect after 1 minute of inactivity
			log.Println(user.conn.RemoteAddr().String() + " seems to be out. Force disconnection.")
			return
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	pool := newUserPool()
	log.Println("Server localhost:1234 ready.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(pool, conn)
	}
}
