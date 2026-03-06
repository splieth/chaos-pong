package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// side represents which side of the play field a client is assigned to.
type side int

const (
	LEFT side = iota
	RIGHT
	NONE
)

// Server manages connected clients and tracks whether the game has started.
type Server struct {
	Clients []Client
	Started bool
}

// Client represents a connected player with message channels for
// bidirectional WebSocket communication.
type Client struct {
	ID      string
	Side    side
	Egress  chan string
	Ingress chan string
}

var server Server
var upgrader = websocket.Upgrader{}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSeq generates a random alphabetic string of length n.
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// assignSide returns the side for the next client based on how many
// clients are already connected.
func assignSide(clientCount int) side {
	switch clientCount {
	case 0:
		return LEFT
	case 1:
		return RIGHT
	default:
		return NONE
	}
}

// createClient builds a new client with a random ID and assigns it
// a side based on the current number of connected clients.
func createClient() Client {
	return Client{
		ID:      randSeq(5),
		Side:    assignSide(len(server.Clients)),
		Ingress: make(chan string),
		Egress:  make(chan string),
	}
}

// deregisterClient removes a client from the server by ID.
func deregisterClient(clientID string) {
	var remaining []Client
	for _, c := range server.Clients {
		if c.ID != clientID {
			remaining = append(remaining, c)
		}
	}
	server.Clients = remaining
}

// readMessages continuously reads from the WebSocket connection and
// forwards messages to the client's ingress channel. Returns on error.
func readMessages(conn *websocket.Conn, ch chan string) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		ch <- string(message)
	}
}

// handleMessage processes an inbound message from a client.
// Movement messages have the format "m <clientID> <direction>".
func handleMessage(msg string) {
	if !strings.HasPrefix(msg, "m") {
		log.Println(msg)
		return
	}
	parts := strings.Split(msg, " ")
	if len(parts) < 3 {
		log.Println("malformed movement message:", msg)
		return
	}
	clientID := parts[1]
	direction := parts[2]
	log.Printf("%s moved: %s", clientID, direction)
}

// register handles the /register WebSocket endpoint. It upgrades the
// HTTP connection, waits for a registration message, assigns the client
// a side, and then processes messages in a loop.
func register(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := waitForRegister(conn)
	server.Clients = append(server.Clients, client)
	defer deregisterClient(client.ID)

	go sendRegistrationMessage(client)
	go readMessages(conn, client.Ingress)

	for {
		select {
		case msg := <-client.Ingress:
			handleMessage(msg)
		case msg := <-client.Egress:
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

// sendRegistrationMessage sends the client its assigned ID.
func sendRegistrationMessage(client Client) {
	client.Egress <- "r " + client.ID
}

// waitForRegister blocks until the client sends a registration message ("r"),
// then creates and returns a new Client.
func waitForRegister(conn *websocket.Conn) Client {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			continue
		}
		if len(message) > 0 && message[0] == 'r' {
			fmt.Println("Received register")
			return createClient()
		}
	}
}

func main() {
	http.HandleFunc("/register", register)

	// Wait for two clients to connect, then broadcast the start signal.
	go func() {
		for {
			if !server.Started && len(server.Clients) == 2 {
				log.Println("Clients are ready, time to start ze game!")
				server.Started = true
				for _, c := range server.Clients {
					log.Println("Sending start message to", c.ID)
					c.Egress <- "s " + c.ID + " " + strconv.Itoa(int(c.Side))
				}
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":4321", nil))
}
