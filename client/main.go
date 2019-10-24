package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	addr := "localhost:4321"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			log.Println(message[0])
			log.Println(message[1])
			switch string(message[0]) {
			case "r":
				id := string(message[1])
				log.Println("Ich bin jetzt", id)
			case "u":
				log.Println(string(message[1]))
				log.Println("moving up")
			case "d":
				log.Println(string(message[1]))
				log.Println("moving down")
			default:
				log.Println("defaultism")
			}
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	register := "r"
	up := "u1"
	_ = c.WriteMessage(websocket.TextMessage, []byte(register))
	_ = c.WriteMessage(websocket.TextMessage, []byte(up))

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
