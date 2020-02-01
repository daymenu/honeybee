package main

import (
	"fmt"
	"log"
	"net/http"

	webscket "github.com/daymenu/honeybee/pkg/websocket"
)

func main() {
	fmt.Println("Honeybee v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ws://", r.Host)

	ws, err := webscket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	webscket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", serveWs)
}
