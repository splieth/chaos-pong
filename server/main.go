package main

import (
	"github.com/gorilla/websocket"
	"github.com/splieth/chaos-pong/game"
	"log"
	"net/http"
)

type Session struct {
	Game    *game.Game
	Clients []Client
	Name    string
}

type Server struct {
	Addr string
}

type Client struct {
	Id string
}

var sessions []Session

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(":4321", nil))
}
