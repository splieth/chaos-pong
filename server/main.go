package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type side int

const (
	LEFT  = 0
	RIGHT = 1
	NONE  = 2
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
	if len(server.Clients) == 0 {
		client.Side = LEFT
	} else if len(server.Clients) == 1 {
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
	go sendRegistrationMessage(client)
	server.Clients = append(server.Clients, client)
	go getMessages(c, client.Ingress)
	for {
		select {
		case inboundMsg := <-client.Ingress:
			if strings.HasPrefix(inboundMsg, "m") {
				parts := strings.Split(inboundMsg, " ")
				clientId := parts[1]
				moved := parts[2]
				log.Printf("%s said: %s", clientId, moved)
			} else {
				log.Println(inboundMsg)
			}
		case outboundMsg := <-client.Egress:
			_ = c.WriteMessage(websocket.TextMessage, []byte(outboundMsg))
		}
	}
	defer deregisterClient(client.Id)
}

func sendRegistrationMessage(client Client) {
	registerMessage := "r " + client.Id
	client.Egress <- registerMessage
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
			fmt.Println("Received register")
			return createClient()
		}
	}
}

func main() {
	http.HandleFunc("/register", register)

	go func() {
		for {
			if !server.Started && len(server.Clients) == 2 {
				log.Println("Clients are ready, time to start ze game!")
				server.Started = true
				for _, c := range server.Clients {
					log.Println("Sending start message to", c.Id)
					c.Egress <- "s " + c.Id + " " + strconv.Itoa(int(c.Side)) //rock solid!
				}
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":4321", nil))
}
