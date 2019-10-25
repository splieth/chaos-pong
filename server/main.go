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

type Server struct {
	Clients []Client
}

type Client struct {
	Id string
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

func registerClient(clientId string) {
	server.Clients = append(server.Clients, Client{Id: clientId})
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

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	currentClient := randSeq(5)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		switch string(message[0]) {
		case "r":
			registerClient(currentClient)
			log.Println(server.Clients)
			c.WriteMessage(websocket.TextMessage, []byte("r "+currentClient))

		case "u":
			log.Println("moving up")
		case "d":
			log.Println("moving down")
		default:
			log.Println("defaultism")
		}
	}
	defer deregisterClient(currentClient)
}

func main() {
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(":4321", nil))
}
