package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/splieth/chaos-pong/game"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pong game.Game
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
			parts := strings.Split(string(message), " ")
			id := parts[1]
			switch parts[0] {
			case "r":
				paddle := parts[2]
				log.Println("Ich bin " + id)
				log.Println("Will control " + paddle + " paddle.")
			case "s":
				pong.StartGame()
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

	//TODO extract
	basePath := os.Getenv("PWD")
	screen, _ := ebiten.NewImage(1280, 720, ebiten.FilterDefault)
	width, height := screen.Size()
	pong = game.NewGame(screen, basePath)

	if err := ebiten.Run(pong.Tick, width, height, 1, "Chaos Pong!"); err != nil {
		log.Fatal(err)
	}

	// TODO do we ever reach this?
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

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
