package main

import (
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
)

var (
	pong game.Game
)

func main() {
	addr := "localhost:4321"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/register"}
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
			if err == nil && len(message) > 0 {
				parts := strings.Split(string(message), " ")
				id := parts[1]
				switch parts[0] {
				case "r":
					log.Println("Ich bims, dem " + id)
				}
				if err != nil {
					log.Println("read:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}
	}()

	register := "r"
	_ = c.WriteMessage(websocket.TextMessage, []byte(register))

	//TODO extract
	basePath := os.Getenv("PWD")
	screen, _ := ebiten.NewImage(1280, 720, ebiten.FilterDefault)
	//width, height := screen.Size()
	pong = game.NewGame(screen, basePath)

	//if err := ebiten.Run(pong.Tick, width, height, 1, "Chaos Pong!"); err != nil {
	//	log.Fatal(err)
	//}
}
