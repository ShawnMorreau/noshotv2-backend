package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

// func serveWs(game *noshotv2.Game, w http.ResponseWriter, r *http.Request) {
func serveWs(lobby *noshotv2.Lobby, w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit")
	conn, err := noshotv2.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	//I'd rather see the Human constructor be used here
	player := &noshotv2.Human{
		Conn:  conn,
		Lobby: lobby,
		ID:    randomdata.SillyName(),
	}
	lobby.AddPlayerToLobby <- player

	player.Read()
}
func setupRoutes() {
	lobby := noshotv2.NewLobby()
	go lobby.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(lobby, w, r)
	})
}
func main() {
	setupRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Println("Starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
