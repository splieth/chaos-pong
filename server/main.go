package main

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type side int

const (
	LEFT = iota
	RIGHT
	NONE
)

type Server struct {
	Clients []Client
	Started bool
}

type Client struct {
	Id      string
	Side    side
	Egress  chan string
	Ingress chan string
}

var server Server
var upgrader = websocket.Upgrader{} // use default options
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createClient() Client {
	client := Client{
		Id:      randSeq(5),
		Side:    NONE,
		Ingress: make(chan string),
		Egress:  make(chan string),
	}
	if len(server.Clients) == 1 {
		client.Side = LEFT
	} else if len(server.Clients) == 2 {
		client.Side = RIGHT
	}
	return client
}

func deregisterClient(clientId string) {
	var newClients []Client
	for _, c := range server.Clients {
		if c.Id != clientId {
			newClients = append(newClients, c)
		}
	}
	server.Clients = newClients
}

func getMessages(conn *websocket.Conn, ch chan string) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		ch <- string(message)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	client := waitForRegister(c)
	log.Println("Socket iz da")
	client.Egress <- "r " + client.Id
	defer deregisterClient(client.Id)
	server.Clients = append(server.Clients, client)
	go getMessages(c, client.Ingress)
	for {
		select {
		case inboudMsg := <-client.Ingress:
			log.Println(inboudMsg)
		case outboundMsg := <-client.Egress:
			log.Println(outboundMsg)
			_ = c.WriteMessage(websocket.TextMessage, []byte(outboundMsg))
		}
	}
}

func waitForRegister(c *websocket.Conn) Client {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		msgType := string(message[0])
		switch msgType {
		case "r":
			return createClient()
		}
	}
}

func main() {
	http.HandleFunc("/register", register)

	go func() {
		for {
			if !server.Started || len(server.Clients) == 2 {
				for _, c := range server.Clients {
					c.Egress <- "s"
				}
				server.Started = true
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":4321", nil))
}
