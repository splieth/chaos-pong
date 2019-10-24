package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	Addr    string
	Clients []Client
}

type Client struct {
	Id string
}

var server Server
var counter int = 0
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		//s := string(message)
		switch string(message[0]) {
		case "r":
			server.Clients = append(server.Clients, Client{Id: string(counter)})
			c.WriteMessage(websocket.TextMessage, []byte("r"+string(counter)))
			counter++
		case "u":
			log.Println(string(message[1]))
			log.Println("moving up")
		case "d":
			log.Println(string(message[1]))
			log.Println("moving down")
		default:
			log.Println("defaultism")
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(":4321", nil))
}
