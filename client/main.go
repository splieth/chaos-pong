package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/splieth/chaos-pong/game"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var (
	pong game.Game
	id   string
)

func move(c *websocket.Conn) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		_ = c.WriteMessage(websocket.TextMessage, []byte("m "+id+" "+strconv.Itoa(1)))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		_ = c.WriteMessage(websocket.TextMessage, []byte("m "+id+" "+strconv.Itoa(2)))
	}
}

func main() {
	addr := "localhost:4321"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/register"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	log.Printf("connected to %s", u.String())
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err == nil && len(message) > 0 {
				fmt.Println(message)
				parts := strings.Split(string(message), " ")
				id = parts[1]
				switch parts[0] {
				case "r":
					log.Println("Ich bims, dem " + id)
				case "s":
					side := parts[1]
					log.Println("I spuil auf der Seite " + side)
					pong.StartGame(side)
				}
				if err != nil {
					log.Println("read:", err)
					return
				}
			}
		}
	}()

	register := "r"
	_ = c.WriteMessage(websocket.TextMessage, []byte(register))

	go func() {
		for {
			move(c)
		}
	}()

	ebiten.SetRunnableOnUnfocused(true) // only for local debugging in order to run 2 clients
	basePath := os.Getenv("PWD")
	pong = game.NewGame(1280, 720, basePath)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Chaos Pong!")

	if err := ebiten.RunGame(&pong); err != nil {
		log.Fatal(err)
	}
}
