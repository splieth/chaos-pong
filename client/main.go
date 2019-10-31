package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
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

	//done := make(chan struct{})

	go func() {
		//defer close(done)
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
					//log.Println("Geht los für " + id)
					pong.StartGame(side)
				}
				if err != nil {
					log.Println("read:", err)
					return
				}
				//log.Printf("recv: %s", message)
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

	//TODO extract
	ebiten.SetRunnableInBackground(true) // only for local debugging in order to run 2 clients. But we'll need that for the server
	basePath := os.Getenv("PWD")
	screen, _ := ebiten.NewImage(1280, 720, ebiten.FilterDefault)
	width, height := screen.Size()
	pong = game.NewGame(screen, basePath)

	if err := ebiten.Run(pong.Tick, width, height, 1, "Chaos Pong!"); err != nil {
		log.Fatal(err)
	}
}
