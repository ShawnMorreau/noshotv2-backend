package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

func serveWs(game *noshotv2.Game, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit")
	conn, err := noshotv2.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	player := &noshotv2.Human{
		Conn: conn,
		Game: game,
		ID:   randomdata.SillyName(),
	}
	game.Register <- player
	player.Read()
}
func setupRoutes() {
	game := noshotv2.NewGame()
	go game.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(game, w, r)
	})
}
func main() {
	setupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
